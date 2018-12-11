// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/storage/Redis.proto

package storage // import "github.com/gojektech/feast/go-feast-proto/feast/storage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/gojektech/feast/go-feast-proto/feast/types"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RedisBucketKey struct {
	// Entity key from the FeatureRow
	EntityKey string `protobuf:"bytes,2,opt,name=entityKey,proto3" json:"entityKey,omitempty"`
	// *
	// This should be the first 7 characters of a sha1 of the featureId
	// This is just to save storage space as it's kept in memory.
	FeatureIdSha1Prefix string `protobuf:"bytes,3,opt,name=featureIdSha1Prefix,proto3" json:"featureIdSha1Prefix,omitempty"`
	// *
	// This groups a feature's values (for different eventTimestamps),
	// into buckets so many can be retrieved together.
	//
	// See FeatureRowToRedisMutationDoFn.
	// bucketId = roundedToGranularity(eventTimestamp).seconds / bucketSize.seconds
	BucketId             uint64   `protobuf:"fixed64,4,opt,name=bucketId,proto3" json:"bucketId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisBucketKey) Reset()         { *m = RedisBucketKey{} }
func (m *RedisBucketKey) String() string { return proto.CompactTextString(m) }
func (*RedisBucketKey) ProtoMessage()    {}
func (*RedisBucketKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_Redis_ad00784f17632936, []int{0}
}
func (m *RedisBucketKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisBucketKey.Unmarshal(m, b)
}
func (m *RedisBucketKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisBucketKey.Marshal(b, m, deterministic)
}
func (dst *RedisBucketKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisBucketKey.Merge(dst, src)
}
func (m *RedisBucketKey) XXX_Size() int {
	return xxx_messageInfo_RedisBucketKey.Size(m)
}
func (m *RedisBucketKey) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisBucketKey.DiscardUnknown(m)
}

var xxx_messageInfo_RedisBucketKey proto.InternalMessageInfo

func (m *RedisBucketKey) GetEntityKey() string {
	if m != nil {
		return m.EntityKey
	}
	return ""
}

func (m *RedisBucketKey) GetFeatureIdSha1Prefix() string {
	if m != nil {
		return m.FeatureIdSha1Prefix
	}
	return ""
}

func (m *RedisBucketKey) GetBucketId() uint64 {
	if m != nil {
		return m.BucketId
	}
	return 0
}

// *
// Because in redis features are stored as a key per feature not per
// feature row, we need the event timestamp in the value.
type RedisBucketValue struct {
	Value                *types.Value         `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	EventTimestamp       *timestamp.Timestamp `protobuf:"bytes,2,opt,name=eventTimestamp,proto3" json:"eventTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RedisBucketValue) Reset()         { *m = RedisBucketValue{} }
func (m *RedisBucketValue) String() string { return proto.CompactTextString(m) }
func (*RedisBucketValue) ProtoMessage()    {}
func (*RedisBucketValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_Redis_ad00784f17632936, []int{1}
}
func (m *RedisBucketValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisBucketValue.Unmarshal(m, b)
}
func (m *RedisBucketValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisBucketValue.Marshal(b, m, deterministic)
}
func (dst *RedisBucketValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisBucketValue.Merge(dst, src)
}
func (m *RedisBucketValue) XXX_Size() int {
	return xxx_messageInfo_RedisBucketValue.Size(m)
}
func (m *RedisBucketValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisBucketValue.DiscardUnknown(m)
}

var xxx_messageInfo_RedisBucketValue proto.InternalMessageInfo

func (m *RedisBucketValue) GetValue() *types.Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RedisBucketValue) GetEventTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.EventTimestamp
	}
	return nil
}

// *
// This allows us to group multiple bucket values together in a
// single list to make it easier to keep sets together
type RedisBucketValueList struct {
	Values               []*RedisBucketValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RedisBucketValueList) Reset()         { *m = RedisBucketValueList{} }
func (m *RedisBucketValueList) String() string { return proto.CompactTextString(m) }
func (*RedisBucketValueList) ProtoMessage()    {}
func (*RedisBucketValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Redis_ad00784f17632936, []int{2}
}
func (m *RedisBucketValueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisBucketValueList.Unmarshal(m, b)
}
func (m *RedisBucketValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisBucketValueList.Marshal(b, m, deterministic)
}
func (dst *RedisBucketValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisBucketValueList.Merge(dst, src)
}
func (m *RedisBucketValueList) XXX_Size() int {
	return xxx_messageInfo_RedisBucketValueList.Size(m)
}
func (m *RedisBucketValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisBucketValueList.DiscardUnknown(m)
}

var xxx_messageInfo_RedisBucketValueList proto.InternalMessageInfo

func (m *RedisBucketValueList) GetValues() []*RedisBucketValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*RedisBucketKey)(nil), "feast.storage.RedisBucketKey")
	proto.RegisterType((*RedisBucketValue)(nil), "feast.storage.RedisBucketValue")
	proto.RegisterType((*RedisBucketValueList)(nil), "feast.storage.RedisBucketValueList")
}

func init() { proto.RegisterFile("feast/storage/Redis.proto", fileDescriptor_Redis_ad00784f17632936) }

var fileDescriptor_Redis_ad00784f17632936 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xd3, 0x97, 0x57, 0x22, 0x43, 0x24, 0x66, 0x35, 0xb1, 0x36, 0x26, 0x34, 0x9c, 0x7a,
	0x61, 0x57, 0xf1, 0xc0, 0xbd, 0x37, 0x82, 0x89, 0xa4, 0x12, 0x0f, 0xde, 0x5a, 0x98, 0x96, 0xca,
	0x9f, 0x25, 0xdd, 0x29, 0xb1, 0x89, 0x07, 0x3f, 0xba, 0x61, 0x16, 0x50, 0x88, 0xb7, 0xdd, 0x79,
	0x7e, 0x33, 0xf3, 0xcc, 0x0c, 0xdc, 0xa6, 0x18, 0x1b, 0x52, 0x86, 0x74, 0x11, 0x67, 0xa8, 0x22,
	0x9c, 0xe6, 0x46, 0xae, 0x0b, 0x4d, 0x5a, 0x5c, 0xb0, 0x24, 0x77, 0x92, 0xd7, 0xce, 0xb4, 0xce,
	0x16, 0xa8, 0x58, 0x4c, 0xca, 0x54, 0x51, 0xbe, 0x44, 0x43, 0xf1, 0x72, 0x6d, 0x79, 0xef, 0xc6,
	0x96, 0xa2, 0x6a, 0x8d, 0x46, 0xbd, 0xc6, 0x8b, 0x12, 0xad, 0xd0, 0xf9, 0x84, 0x16, 0xd7, 0x0d,
	0xcb, 0xc9, 0x1c, 0x69, 0x88, 0x95, 0xb8, 0x83, 0x06, 0xae, 0x28, 0xa7, 0x6a, 0x88, 0x95, 0xfb,
	0xcf, 0x77, 0x82, 0x46, 0xf4, 0x13, 0x10, 0xf7, 0x70, 0x95, 0x62, 0x4c, 0x65, 0x81, 0x83, 0xe9,
	0xcb, 0x2c, 0x7e, 0x18, 0x15, 0x98, 0xe6, 0x1f, 0x6e, 0x8d, 0xb9, 0xbf, 0x24, 0xe1, 0xc1, 0x79,
	0xc2, 0xc5, 0x07, 0x53, 0xf7, 0xbf, 0xef, 0x04, 0xf5, 0xe8, 0xf0, 0xef, 0x7c, 0x39, 0x70, 0xf9,
	0xab, 0x3d, 0x1b, 0x13, 0x01, 0x9c, 0x6d, 0xb6, 0x0f, 0xd7, 0xf1, 0x9d, 0xa0, 0xd9, 0x13, 0xd2,
	0xce, 0xca, 0xde, 0x25, 0x23, 0x91, 0x05, 0x44, 0x08, 0x2d, 0xdc, 0xe0, 0x8a, 0xc6, 0xfb, 0x69,
	0xd9, 0x6f, 0xb3, 0xe7, 0x49, 0xbb, 0x0f, 0xb9, 0xdf, 0x87, 0x3c, 0x10, 0xd1, 0x49, 0x46, 0xe7,
	0x19, 0xae, 0x4f, 0x1d, 0x3c, 0xe5, 0x86, 0x44, 0x1f, 0xea, 0xdc, 0xc4, 0xb8, 0x8e, 0x5f, 0x0b,
	0x9a, 0xbd, 0xb6, 0x3c, 0x5a, 0xb9, 0x3c, 0x4d, 0x8a, 0x76, 0x78, 0x38, 0x86, 0xe3, 0xe3, 0x84,
	0xc0, 0xe8, 0x68, 0x6b, 0xe5, 0xad, 0x9f, 0xe5, 0x34, 0x2b, 0x13, 0x39, 0xd1, 0x4b, 0x95, 0xe9,
	0x77, 0x9c, 0x13, 0x4e, 0x66, 0xca, 0x1e, 0x27, 0xd3, 0x5d, 0x7e, 0x74, 0xd9, 0xb5, 0x3a, 0x3a,
	0x7e, 0x52, 0xe7, 0xe0, 0xe3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x9d, 0x61, 0xa8, 0x14,
	0x02, 0x00, 0x00,
}
