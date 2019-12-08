# Copyright 2019 The Feast Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import logging
import os
import shutil
import sys
import time
import uuid
from collections import OrderedDict
from typing import Dict, Union
from typing import List

import grpc
import pandas as pd
import pyarrow as pa
import pyarrow.parquet as pq
from feast.core.CoreService_pb2 import (
    GetFeastCoreVersionRequest,
    ListFeatureSetsResponse,
    ApplyFeatureSetRequest,
    ListFeatureSetsRequest,
    ApplyFeatureSetResponse,
    GetFeatureSetRequest,
    GetFeatureSetResponse,
)
from feast.core.CoreService_pb2_grpc import CoreServiceStub
from feast.exceptions import format_grpc_exception
from feast.feature_set import FeatureSet, Entity
from feast.job import Job
from feast.loaders.file import export_dataframe_to_staging_location
from feast.loaders.ingest import KAFKA_CHUNK_PRODUCTION_TIMEOUT
from feast.loaders.ingest import get_feature_row_chunks, ingest_table_to_kafka
from feast.serving.ServingService_pb2 import GetFeastServingInfoResponse
from feast.serving.ServingService_pb2 import (
    GetOnlineFeaturesRequest,
    GetBatchFeaturesRequest,
    GetFeastServingInfoRequest,
    GetOnlineFeaturesResponse,
    DatasetSource,
    DataFormat,
    FeatureSetRequest,
    FeastServingType,
)
from feast.serving.ServingService_pb2_grpc import ServingServiceStub
from confluent_kafka import Producer
from tqdm import tqdm

_logger = logging.getLogger(__name__)

GRPC_CONNECTION_TIMEOUT_DEFAULT = 3  # type: int
GRPC_CONNECTION_TIMEOUT_APPLY = 600  # type: int
FEAST_SERVING_URL_ENV_KEY = "FEAST_SERVING_URL"  # type: str
FEAST_CORE_URL_ENV_KEY = "FEAST_CORE_URL"  # type: str
BATCH_FEATURE_REQUEST_WAIT_TIME_SECONDS = 300
CPU_COUNT = os.cpu_count()  # type: int


class Client:
    def __init__(
        self, core_url: str = None, serving_url: str = None, verbose: bool = False
    ):
        self._core_url = core_url
        self._serving_url = serving_url
        self._verbose = verbose
        self.__core_channel: grpc.Channel = None
        self.__serving_channel: grpc.Channel = None
        self._core_service_stub: CoreServiceStub = None
        self._serving_service_stub: ServingServiceStub = None

    @property
    def core_url(self) -> str:
        if self._core_url is not None:
            return self._core_url
        if os.getenv(FEAST_CORE_URL_ENV_KEY) is not None:
            return os.getenv(FEAST_CORE_URL_ENV_KEY)
        return ""

    @core_url.setter
    def core_url(self, value: str):
        self._core_url = value

    @property
    def serving_url(self) -> str:
        if self._serving_url is not None:
            return self._serving_url
        if os.getenv(FEAST_SERVING_URL_ENV_KEY) is not None:
            return os.getenv(FEAST_SERVING_URL_ENV_KEY)
        return ""

    @serving_url.setter
    def serving_url(self, value: str):
        self._serving_url = value

    def version(self):
        """
        Returns version information from Feast Core and Feast Serving
        :return: Dictionary containing Core and Serving versions and status
        """

        self._connect_core()
        self._connect_serving()

        core_version = ""
        serving_version = ""
        core_status = "not connected"
        serving_status = "not connected"

        try:
            core_version = self._core_service_stub.GetFeastCoreVersion(
                GetFeastCoreVersionRequest(), timeout=GRPC_CONNECTION_TIMEOUT_DEFAULT
            ).version
            core_status = "connected"
        except grpc.RpcError as e:
            print(format_grpc_exception("GetFeastCoreVersion", e.code(), e.details()))

        try:
            serving_version = self._serving_service_stub.GetFeastServingInfo(
                GetFeastServingInfoRequest(), timeout=GRPC_CONNECTION_TIMEOUT_DEFAULT
            ).version
            serving_status = "connected"
        except grpc.RpcError as e:
            print(format_grpc_exception("GetFeastServingInfo", e.code(), e.details()))

        return {
            "core": {
                "url": self.core_url,
                "version": core_version,
                "status": core_status,
            },
            "serving": {
                "url": self.serving_url,
                "version": serving_version,
                "status": serving_status,
            },
        }

    def _connect_core(self, skip_if_connected=True):
        """
        Connect to Core API
        """
        if skip_if_connected and self._core_service_stub:
            return

        if not self.core_url:
            raise ValueError("Please set Feast Core URL.")

        if self.__core_channel is None:
            self.__core_channel = grpc.insecure_channel(self.core_url)

        try:
            grpc.channel_ready_future(self.__core_channel).result(
                timeout=GRPC_CONNECTION_TIMEOUT_DEFAULT
            )
        except grpc.FutureTimeoutError:
            print(
                f"Connection timed out while attempting to connect to Feast Core gRPC server {self.core_url}"
            )
            sys.exit(1)
        else:
            self._core_service_stub = CoreServiceStub(self.__core_channel)

    def _connect_serving(self, skip_if_connected=True):
        """
        Connect to Serving API
        """

        if skip_if_connected and self._serving_service_stub:
            return

        if not self.serving_url:
            raise ValueError("Please set Feast Serving URL.")

        if self.__serving_channel is None:
            self.__serving_channel = grpc.insecure_channel(self.serving_url)

        try:
            grpc.channel_ready_future(self.__serving_channel).result(
                timeout=GRPC_CONNECTION_TIMEOUT_DEFAULT
            )
        except grpc.FutureTimeoutError:
            print(
                f"Connection timed out while attempting to connect to Feast Serving gRPC server {self.serving_url} "
            )
            sys.exit(1)
        else:
            self._serving_service_stub = ServingServiceStub(self.__serving_channel)

    def apply(self, feature_sets: Union[List[FeatureSet], FeatureSet]):
        """
        Idempotently registers feature set(s) with Feast Core. Either a single feature set or a list can be provided.
        :param feature_sets: Union[List[FeatureSet], FeatureSet]
        """
        if not isinstance(feature_sets, list):
            feature_sets = [feature_sets]
        for feature_set in feature_sets:
            if isinstance(feature_set, FeatureSet):
                self._apply_feature_set(feature_set)
                continue
            raise ValueError(
                f"Could not determine feature set type to apply {feature_set}"
            )

    def _apply_feature_set(self, feature_set: FeatureSet):
        self._connect_core()
        feature_set._client = self

        valid, message = feature_set.is_valid()
        if not valid:
            raise Exception(message)
        try:
            apply_fs_response = self._core_service_stub.ApplyFeatureSet(
                ApplyFeatureSetRequest(feature_set=feature_set.to_proto()),
                timeout=GRPC_CONNECTION_TIMEOUT_APPLY,
            )  # type: ApplyFeatureSetResponse
            applied_fs = FeatureSet.from_proto(apply_fs_response.feature_set)

            if apply_fs_response.status == ApplyFeatureSetResponse.Status.CREATED:
                print(
                    f'Feature set updated/created: "{applied_fs.name}:{applied_fs.version}".'
                )
                feature_set._update_from_feature_set(applied_fs, is_dirty=False)
                return
            if apply_fs_response.status == ApplyFeatureSetResponse.Status.NO_CHANGE:
                print(f"No change detected in feature set {feature_set.name}")
                return
        except grpc.RpcError as e:
            print(format_grpc_exception("ApplyFeatureSet", e.code(), e.details()))

    def list_feature_sets(self) -> List[FeatureSet]:
        """
        Retrieve a list of feature sets from Feast Core
        :return: Returns a list of feature sets
        """
        self._connect_core()

        try:
            # Get latest feature sets from Feast Core
            feature_set_protos = self._core_service_stub.ListFeatureSets(
                ListFeatureSetsRequest()
            )  # type: ListFeatureSetsResponse
        except grpc.RpcError as e:
            raise Exception(
                format_grpc_exception("ListFeatureSets", e.code(), e.details())
            )

        # Store list of feature sets
        feature_sets = []
        for feature_set_proto in feature_set_protos.feature_sets:
            feature_set = FeatureSet.from_proto(feature_set_proto)
            feature_set._client = self
            feature_sets.append(feature_set)
        return feature_sets

    def get_feature_set(
        self, name: str, version: int = None, fail_if_missing: bool = False
    ) -> Union[FeatureSet, None]:
        """
        Retrieve a single feature set from Feast Core
        :param name: (str) Name of feature set
        :param version: (int) Version of feature set
        :param fail_if_missing: (bool) Throws an exception if the feature set is not
         found
        :return: Returns a single feature set

        """
        self._connect_core()
        try:
            name = name.strip()
            if version is None:
                version = 0
            get_feature_set_response = self._core_service_stub.GetFeatureSet(
                GetFeatureSetRequest(name=name, version=version)
            )  # type: GetFeatureSetResponse
            feature_set = get_feature_set_response.feature_set
        except grpc.RpcError as e:
            print(format_grpc_exception("GetFeatureSet", e.code(), e.details()))
        else:
            if feature_set is not None:
                return FeatureSet.from_proto(feature_set)

            if fail_if_missing:
                raise Exception(
                    f'Could not find feature set with name "{name}" and '
                    f'version "{version}"'
                )

    def list_entities(self) -> Dict[str, Entity]:
        """
        Returns a dictionary of entities across all feature sets
        :return: Dictionary of entity name to Entity
        """
        entities_dict = OrderedDict()
        for fs in self.list_feature_sets():
            for entity in fs.entities:
                entities_dict[entity.name] = entity
        return entities_dict

    def get_batch_features(
        self, feature_ids: List[str], entity_rows: pd.DataFrame
    ) -> Job:
        """
        Retrieves historical features from a Feast Serving deployment.

        Args:
            feature_ids: List of feature ids that will be returned for each entity.
            Each feature id should have the following format "feature_set_name:version:feature_name".

            entity_rows: Pandas dataframe containing entities and a 'datetime' column. Each entity in
            a feature set must be present as a column in this dataframe. The datetime column must
            contain timestamps in datetime64 format

        Returns:
            Feast batch retrieval job: feast.job.Job

        Example usage:
        ============================================================
        >>> from feast import Client
        >>> from datetime import datetime
        >>>
        >>> feast_client = Client(core_url="localhost:6565", serving_url="localhost:6566")
        >>> feature_ids = ["customer:1:bookings_7d"]
        >>> entity_rows = pd.DataFrame(
        >>>         {
        >>>            "datetime": [pd.datetime.now() for _ in range(3)],
        >>>            "customer": [1001, 1002, 1003],
        >>>         }
        >>>     )
        >>> feature_retrieval_job = feast_client.get_batch_features(feature_ids, entity_rows)
        >>> df = feature_retrieval_job.to_dataframe()
        >>> print(df)
        """

        self._connect_serving()

        try:
            fs_request = _build_feature_set_request(feature_ids)

            # Validate entity rows based on entities in Feast Core
            self._validate_entity_rows_for_batch_retrieval(entity_rows, fs_request)

            # We want the timestamp column naming to be consistent with the
            # rest of Feast
            entity_rows.columns = [
                "event_timestamp" if col == "datetime" else col
                for col in entity_rows.columns
            ]

            # Remove timezone from datetime column
            if isinstance(
                entity_rows["event_timestamp"].dtype,
                pd.core.dtypes.dtypes.DatetimeTZDtype,
            ):
                entity_rows["event_timestamp"] = pd.DatetimeIndex(
                    entity_rows["event_timestamp"]
                ).tz_localize(None)

            # Retrieve serving information to determine store type and staging location
            serving_info = self._serving_service_stub.GetFeastServingInfo(
                GetFeastServingInfoRequest(), timeout=GRPC_CONNECTION_TIMEOUT_DEFAULT
            )  # type: GetFeastServingInfoResponse

            if serving_info.type != FeastServingType.FEAST_SERVING_TYPE_BATCH:
                raise Exception(
                    f'You are connected to a store "{self._serving_url}" which does not support batch retrieval'
                )

            # Export and upload entity row dataframe to staging location provided by Feast
            staged_file = export_dataframe_to_staging_location(
                entity_rows, serving_info.job_staging_location
            )  # type: str

            request = GetBatchFeaturesRequest(
                feature_sets=fs_request,
                dataset_source=DatasetSource(
                    file_source=DatasetSource.FileSource(
                        file_uris=[staged_file], data_format=DataFormat.DATA_FORMAT_AVRO
                    )
                ),
            )

            # Retrieve Feast Job object to manage life cycle of retrieval
            response = self._serving_service_stub.GetBatchFeatures(request)
            return Job(response.job, self._serving_service_stub)

        except grpc.RpcError as e:
            print(format_grpc_exception("GetBatchFeatures", e.code(), e.details()))

    def _validate_entity_rows_for_batch_retrieval(
        self, entity_rows, feature_sets_request
    ):
        """
        Validate whether an entity_row dataframe contains the correct information for batch retrieval
        :param entity_rows: Pandas dataframe containing entities and datetime column. Each entity in a feature set
        must be present as a column in this dataframe.
        :param feature_sets_request: Feature sets that will
        """

        # Ensure datetime column exists
        if "datetime" not in entity_rows.columns:
            raise ValueError(
                f'Entity rows does not contain "datetime" column in columns {entity_rows.columns}'
            )

        # Validate dataframe columns based on feature set entities
        for feature_set in feature_sets_request:
            fs = self.get_feature_set(
                name=feature_set.name, version=feature_set.version
            )
            if fs is None:
                raise ValueError(
                    f'Feature set "{feature_set.name}:{feature_set.version}" could not be found'
                )
            for entity_type in fs.entities:
                if entity_type.name not in entity_rows.columns:
                    raise ValueError(
                        f'Dataframe does not contain entity "{entity_type.name}" column in columns "{entity_rows.columns}"'
                    )

    def get_online_features(
        self,
        feature_ids: List[str],
        entity_rows: List[GetOnlineFeaturesRequest.EntityRow],
    ) -> GetOnlineFeaturesResponse:
        """
        Retrieves the latest online feature data from Feast Serving
        :param feature_ids: List of feature Ids in the following format
                            [feature_set_name]:[version]:[feature_name]
                            example: ["feature_set_1:6:my_feature_1",
                                     "feature_set_1:6:my_feature_2",]

        :param entity_rows: List of GetFeaturesRequest.EntityRow where each row
                            contains entities. Timestamp should not be set for
                            online retrieval. All entity types within a feature
                            set must be provided for each entity key.
        :return: Returns a list of maps where each item in the list contains
                 the latest feature values for the provided entities
        """
        self._connect_serving()

        try:
            response = self._serving_service_stub.GetOnlineFeatures(
                GetOnlineFeaturesRequest(
                    feature_sets=_build_feature_set_request(feature_ids),
                    entity_rows=entity_rows,
                )
            )  # type: GetOnlineFeaturesResponse
        except grpc.RpcError as e:
            print(format_grpc_exception("GetOnlineFeatures", e.code(), e.details()))
        else:
            return response

    def ingest(
            self,
            feature_set: Union[str, FeatureSet],
            source: Union[pd.DataFrame, str],
            chunk_size: int,
            version: int = None,
            force_update: bool = False,
            max_workers: int = CPU_COUNT - 1,
            disable_progress_bar: bool = False,
    ) -> None:
        """
        Loads data into Feast for a specific feature set.

        Args:
            feature_set (typing.Union[str, FeatureSet]):
                Feature set object or the string name of the feature set
                (without a version).

            source (typing.Union[pd.DataFrame, str]):
                Either a file path or Pandas Dataframe to ingest into Feast
                Files that are currently supported:
                    * parquet
                    * csv
                    * json

            chunk_size (int):
                Amount of rows to load and ingest at a time.

            version (int):
                Feature set version.

            force_update (bool):
                Automatically update feature set based on.

            max_workers (int):
                Number of worker processes to use to encode values.

            disable_progress_bar (bool):
                Disable printing of progress statistics.

        Returns:
            None:
                None
        """

        if isinstance(feature_set, FeatureSet):
            name = feature_set.name
            if version is None:
                version = feature_set.version
        elif isinstance(feature_set, str):
            name = feature_set
        else:
            raise Exception(f"Feature set name must be provided")

        # Read table and get row count
        table = _read_table_from_source(source)
        row_count = table.num_rows

        time_start = time.time()

        # Update the feature set based on PyArrow table schema
        if force_update:
            feature_set.infer_fields_from_pa(
                table=table,
                discard_unused_fields=True,
                replace_existing_features=True
            )
            self.apply(feature_set)

        feature_set = self.get_feature_set(name, version, fail_if_missing=True)

        # Split table into smaller parquet files
        batches = table.to_batches(
            max_chunksize=min(int(table.num_rows / max_workers), chunk_size))
        tables = [pa.lib.Table.from_batches([batch]) for batch in batches]

        print(f"Splitting parquet file into {len(tables)} chunks.")

        temp_dir = f"_feast_{time.time_ns()}/"

        # Create a temporary work dir
        try:
            os.mkdir(temp_dir)
        except FileExistsError as e:
            print("Temporary directory exists")

        for tbl in tables:
            pa.parquet.write_table(tbl, temp_dir + str(uuid.uuid4()))

        time_end = time.time()
        exec_time = round(time_end - time_start, 2)
        print(f"Took {exec_time}s to validate and split parquet file.")

        # Remove PyArrow from memory
        del table

        # Progress bar will always display average rate
        pbar = tqdm(total=row_count,
                    unit="rows",
                    smoothing=0,
                    disable=disable_progress_bar)

        # Get a list of all parquet files created
        files = [temp_dir + x
                 for x in os.listdir(temp_dir)
                 if not x.startswith(".")]  # DO not read in .ds_store

        # Kafka configs
        brokers = feature_set.get_kafka_source_brokers()
        topic = feature_set.get_kafka_source_topic()
        producer = Producer({"bootstrap.servers": brokers})

        # Code optimizations
        send = producer.produce
        flush = producer.flush
        update = pbar.update

        # Define tracker context
        ctx = {"success_count": 0, "error_count": 0, "last_exception": ""}

        if feature_set.source.source_type == "Kafka":
            for chunk in get_feature_row_chunks(
                    files=files, fs=feature_set, max_workers=max_workers):

                # Push FeatureRow in chunk to kafka
                for row in chunk:
                    try:
                        send(topic, row.SerializeToString())
                        update(1)
                    except Exception as e:
                        # Save last exception
                        ctx["last_exception"] = e

                        # Increment error count
                        if "error_count" in ctx:
                            ctx["error_count"] += 1
                        else:
                            ctx["error_count"] = 1

                # Force a flush
                flush(timeout=KAFKA_CHUNK_PRODUCTION_TIMEOUT)

                # Remove chunk from memory
                del chunk

        else:
            raise Exception(
                f"Could not determine source type for feature set "
                f'"{feature_set.name}" with source type '
                f'"{feature_set.source.source_type}"'
            )

        # Refresh and close tqdm progress bar
        pbar.refresh()

        # Using progress bar as counter is much faster than incrementing dict
        ctx["success_count"] = pbar.n

        pbar.close()
        print("Ingestion complete!")

        failed_message = (
            ""
            if ctx["error_count"] == 0
            else f"\nFail: {ctx['error_count']}/{row_count}"
        )

        last_exception_message = (
            ""
            if ctx["last_exception"] == ""
            else f"\nLast exception:\n{ctx['last_exception']}"
        )
        print(
            f"\nIngestion statistics:"
            f"\nSuccess: {ctx['success_count']}/{row_count}"
            f"{failed_message}"
            f"{last_exception_message}"
        )

        # Cleanup, remove smaller parquet file(s) that were created
        print("Removing temporary files...")
        shutil.rmtree(temp_dir)

        return None

    def ingest_old(
        self,
        feature_set: Union[str, FeatureSet],
        source: Union[pd.DataFrame, str],
        version: int = None,
        force_update: bool = False,
        max_workers: int = CPU_COUNT,
        disable_progress_bar: bool = False,
        chunk_size: int = 5000,
        timeout: int = None,
    ):
        """
        Loads data into Feast for a specific feature set.

        :param feature_set: (str, FeatureSet) Feature set object or the
        string name of the feature set (without a version)
        :param source:
        Either a file path or Pandas Dataframe to ingest into Feast
        Files that are currently supported:
            * parquet
            * csv
            * json

        :param version: Feature set version
        :param force_update: (bool) Automatically update feature set based on
        source data prior to ingesting. This will also register changes to Feast
        :param max_workers: Number of worker processes to use to encode values
        :param disable_progress_bar: Disable printing of progress statistics
        :param timeout: Time in seconds before ingestion times out
        :param chunk_size: Amount of rows to load and ingest at a time

        """

        if isinstance(feature_set, FeatureSet):
            name = feature_set.name
            if version is None:
                version = feature_set.version
        elif isinstance(feature_set, str):
            name = feature_set
        else:
            raise Exception(f"Feature set name must be provided")

        table = _read_table_from_source(source)

        # Update the feature set based on DataFrame schema
        if force_update:
            # Use a small as reference DataFrame to infer fields
            ref_df = table.to_batches(max_chunksize=20)[0].to_pandas()

            feature_set.infer_fields_from_df(
                ref_df, discard_unused_fields=True, replace_existing_features=True
            )
            self.apply(feature_set)

        feature_set = self.get_feature_set(name, version, fail_if_missing=True)

        if feature_set.source.source_type == "Kafka":
            ingest_table_to_kafka(
                feature_set=feature_set,
                table=table,
                max_workers=max_workers,
                disable_pbar=disable_progress_bar,
                chunk_size=chunk_size,
                timeout=timeout,
            )
        else:
            raise Exception(
                f"Could not determine source type for feature set "
                f'"{feature_set.name}" with source type '
                f'"{feature_set.source.source_type}"'
            )


def _build_feature_set_request(feature_ids: List[str]) -> List[FeatureSetRequest]:
    """
    Builds a list of FeatureSet objects from feature set ids in order to retrieve feature data from Feast Serving
    """
    feature_set_request = dict()  # type: Dict[str, FeatureSetRequest]
    for feature_id in feature_ids:
        fid_parts = feature_id.split(":")
        if len(fid_parts) == 3:
            feature_set, version, feature = fid_parts
        else:
            raise ValueError(
                f"Could not parse feature id ${feature_id}, needs 3 colons"
            )

        if feature_set not in feature_set_request:
            feature_set_request[feature_set] = FeatureSetRequest(
                name=feature_set, version=int(version)
            )
        feature_set_request[feature_set].feature_names.append(feature)
    return list(feature_set_request.values())


def _read_table_from_source(source: Union[pd.DataFrame, str]) -> pa.lib.Table:
    """
    Infers a data source type (path or Pandas Dataframe) and reads it in as
    a PyArrow Table.

    :param source: Either a string path or Pandas dataframe
    :return: PyArrow table
    """

    # Pandas dataframe detected
    if isinstance(source, pd.DataFrame):
        table = pa.Table.from_pandas(df=source)

    # Inferring a string path
    elif isinstance(source, str):
        file_path = source
        filename, file_ext = os.path.splitext(file_path)

        if ".csv" in file_ext:
            from pyarrow import csv

            table = csv.read_csv(filename)
        elif ".json" in file_ext:
            from pyarrow import json

            table = json.read_json(filename)
        else:
            table = pq.read_table(file_path)
    else:
        raise ValueError(f"Unknown data source provided for ingestion: {source}")

    # Ensure that PyArrow table is initialised
    assert isinstance(table, pa.lib.Table)

    return table
