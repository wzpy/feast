// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/types/Value.proto

package types // import "github.com/gojektech/feast/go-feast-proto/feast/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type ValueType_Enum int32

const (
	ValueType_UNKNOWN   ValueType_Enum = 0
	ValueType_BYTES     ValueType_Enum = 1
	ValueType_STRING    ValueType_Enum = 2
	ValueType_INT32     ValueType_Enum = 3
	ValueType_INT64     ValueType_Enum = 4
	ValueType_DOUBLE    ValueType_Enum = 5
	ValueType_FLOAT     ValueType_Enum = 6
	ValueType_BOOL      ValueType_Enum = 7
	ValueType_TIMESTAMP ValueType_Enum = 8
)

var ValueType_Enum_name = map[int32]string{
	0: "UNKNOWN",
	1: "BYTES",
	2: "STRING",
	3: "INT32",
	4: "INT64",
	5: "DOUBLE",
	6: "FLOAT",
	7: "BOOL",
	8: "TIMESTAMP",
}
var ValueType_Enum_value = map[string]int32{
	"UNKNOWN":   0,
	"BYTES":     1,
	"STRING":    2,
	"INT32":     3,
	"INT64":     4,
	"DOUBLE":    5,
	"FLOAT":     6,
	"BOOL":      7,
	"TIMESTAMP": 8,
}

func (x ValueType_Enum) String() string {
	return proto.EnumName(ValueType_Enum_name, int32(x))
}
func (ValueType_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{0, 0}
}

type ValueType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueType) Reset()         { *m = ValueType{} }
func (m *ValueType) String() string { return proto.CompactTextString(m) }
func (*ValueType) ProtoMessage()    {}
func (*ValueType) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{0}
}
func (m *ValueType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueType.Unmarshal(m, b)
}
func (m *ValueType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueType.Marshal(b, m, deterministic)
}
func (dst *ValueType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueType.Merge(dst, src)
}
func (m *ValueType) XXX_Size() int {
	return xxx_messageInfo_ValueType.Size(m)
}
func (m *ValueType) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueType.DiscardUnknown(m)
}

var xxx_messageInfo_ValueType proto.InternalMessageInfo

type Value struct {
	// Types that are valid to be assigned to Val:
	//	*Value_BytesVal
	//	*Value_StringVal
	//	*Value_Int32Val
	//	*Value_Int64Val
	//	*Value_DoubleVal
	//	*Value_FloatVal
	//	*Value_BoolVal
	//	*Value_TimestampVal
	Val                  isValue_Val `protobuf_oneof:"val"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{1}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Val interface {
	isValue_Val()
}

type Value_BytesVal struct {
	BytesVal []byte `protobuf:"bytes,1,opt,name=bytesVal,proto3,oneof"`
}

type Value_StringVal struct {
	StringVal string `protobuf:"bytes,2,opt,name=stringVal,proto3,oneof"`
}

type Value_Int32Val struct {
	Int32Val int32 `protobuf:"varint,3,opt,name=int32Val,proto3,oneof"`
}

type Value_Int64Val struct {
	Int64Val int64 `protobuf:"varint,4,opt,name=int64Val,proto3,oneof"`
}

type Value_DoubleVal struct {
	DoubleVal float64 `protobuf:"fixed64,5,opt,name=doubleVal,proto3,oneof"`
}

type Value_FloatVal struct {
	FloatVal float32 `protobuf:"fixed32,6,opt,name=floatVal,proto3,oneof"`
}

type Value_BoolVal struct {
	BoolVal bool `protobuf:"varint,7,opt,name=boolVal,proto3,oneof"`
}

type Value_TimestampVal struct {
	TimestampVal *timestamp.Timestamp `protobuf:"bytes,8,opt,name=timestampVal,proto3,oneof"`
}

func (*Value_BytesVal) isValue_Val() {}

func (*Value_StringVal) isValue_Val() {}

func (*Value_Int32Val) isValue_Val() {}

func (*Value_Int64Val) isValue_Val() {}

func (*Value_DoubleVal) isValue_Val() {}

func (*Value_FloatVal) isValue_Val() {}

func (*Value_BoolVal) isValue_Val() {}

func (*Value_TimestampVal) isValue_Val() {}

func (m *Value) GetVal() isValue_Val {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *Value) GetBytesVal() []byte {
	if x, ok := m.GetVal().(*Value_BytesVal); ok {
		return x.BytesVal
	}
	return nil
}

func (m *Value) GetStringVal() string {
	if x, ok := m.GetVal().(*Value_StringVal); ok {
		return x.StringVal
	}
	return ""
}

func (m *Value) GetInt32Val() int32 {
	if x, ok := m.GetVal().(*Value_Int32Val); ok {
		return x.Int32Val
	}
	return 0
}

func (m *Value) GetInt64Val() int64 {
	if x, ok := m.GetVal().(*Value_Int64Val); ok {
		return x.Int64Val
	}
	return 0
}

func (m *Value) GetDoubleVal() float64 {
	if x, ok := m.GetVal().(*Value_DoubleVal); ok {
		return x.DoubleVal
	}
	return 0
}

func (m *Value) GetFloatVal() float32 {
	if x, ok := m.GetVal().(*Value_FloatVal); ok {
		return x.FloatVal
	}
	return 0
}

func (m *Value) GetBoolVal() bool {
	if x, ok := m.GetVal().(*Value_BoolVal); ok {
		return x.BoolVal
	}
	return false
}

func (m *Value) GetTimestampVal() *timestamp.Timestamp {
	if x, ok := m.GetVal().(*Value_TimestampVal); ok {
		return x.TimestampVal
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Value) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Value_OneofMarshaler, _Value_OneofUnmarshaler, _Value_OneofSizer, []interface{}{
		(*Value_BytesVal)(nil),
		(*Value_StringVal)(nil),
		(*Value_Int32Val)(nil),
		(*Value_Int64Val)(nil),
		(*Value_DoubleVal)(nil),
		(*Value_FloatVal)(nil),
		(*Value_BoolVal)(nil),
		(*Value_TimestampVal)(nil),
	}
}

func _Value_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Value)
	// val
	switch x := m.Val.(type) {
	case *Value_BytesVal:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.BytesVal)
	case *Value_StringVal:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.StringVal)
	case *Value_Int32Val:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int32Val))
	case *Value_Int64Val:
		b.EncodeVarint(4<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Int64Val))
	case *Value_DoubleVal:
		b.EncodeVarint(5<<3 | proto.WireFixed64)
		b.EncodeFixed64(math.Float64bits(x.DoubleVal))
	case *Value_FloatVal:
		b.EncodeVarint(6<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.FloatVal)))
	case *Value_BoolVal:
		t := uint64(0)
		if x.BoolVal {
			t = 1
		}
		b.EncodeVarint(7<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *Value_TimestampVal:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampVal); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Value.Val has unexpected type %T", x)
	}
	return nil
}

func _Value_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Value)
	switch tag {
	case 1: // val.bytesVal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Val = &Value_BytesVal{x}
		return true, err
	case 2: // val.stringVal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Val = &Value_StringVal{x}
		return true, err
	case 3: // val.int32Val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Val = &Value_Int32Val{int32(x)}
		return true, err
	case 4: // val.int64Val
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Val = &Value_Int64Val{int64(x)}
		return true, err
	case 5: // val.doubleVal
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Val = &Value_DoubleVal{math.Float64frombits(x)}
		return true, err
	case 6: // val.floatVal
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Val = &Value_FloatVal{math.Float32frombits(uint32(x))}
		return true, err
	case 7: // val.boolVal
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Val = &Value_BoolVal{x != 0}
		return true, err
	case 8: // val.timestampVal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(timestamp.Timestamp)
		err := b.DecodeMessage(msg)
		m.Val = &Value_TimestampVal{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Value_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Value)
	// val
	switch x := m.Val.(type) {
	case *Value_BytesVal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.BytesVal)))
		n += len(x.BytesVal)
	case *Value_StringVal:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.StringVal)))
		n += len(x.StringVal)
	case *Value_Int32Val:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int32Val))
	case *Value_Int64Val:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Int64Val))
	case *Value_DoubleVal:
		n += 1 // tag and wire
		n += 8
	case *Value_FloatVal:
		n += 1 // tag and wire
		n += 4
	case *Value_BoolVal:
		n += 1 // tag and wire
		n += 1
	case *Value_TimestampVal:
		s := proto.Size(x.TimestampVal)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ValueList struct {
	// Types that are valid to be assigned to ValueList:
	//	*ValueList_BytesList
	//	*ValueList_StringList
	//	*ValueList_Int32List
	//	*ValueList_Int64List
	//	*ValueList_DoubleList
	//	*ValueList_FloatList
	//	*ValueList_BoolList
	//	*ValueList_TimestampList
	ValueList            isValueList_ValueList `protobuf_oneof:"valueList"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ValueList) Reset()         { *m = ValueList{} }
func (m *ValueList) String() string { return proto.CompactTextString(m) }
func (*ValueList) ProtoMessage()    {}
func (*ValueList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{2}
}
func (m *ValueList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueList.Unmarshal(m, b)
}
func (m *ValueList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueList.Marshal(b, m, deterministic)
}
func (dst *ValueList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueList.Merge(dst, src)
}
func (m *ValueList) XXX_Size() int {
	return xxx_messageInfo_ValueList.Size(m)
}
func (m *ValueList) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueList.DiscardUnknown(m)
}

var xxx_messageInfo_ValueList proto.InternalMessageInfo

type isValueList_ValueList interface {
	isValueList_ValueList()
}

type ValueList_BytesList struct {
	BytesList *BytesList `protobuf:"bytes,1,opt,name=bytesList,proto3,oneof"`
}

type ValueList_StringList struct {
	StringList *StringList `protobuf:"bytes,2,opt,name=stringList,proto3,oneof"`
}

type ValueList_Int32List struct {
	Int32List *Int32List `protobuf:"bytes,3,opt,name=int32List,proto3,oneof"`
}

type ValueList_Int64List struct {
	Int64List *Int64List `protobuf:"bytes,4,opt,name=int64List,proto3,oneof"`
}

type ValueList_DoubleList struct {
	DoubleList *DoubleList `protobuf:"bytes,5,opt,name=doubleList,proto3,oneof"`
}

type ValueList_FloatList struct {
	FloatList *FloatList `protobuf:"bytes,6,opt,name=floatList,proto3,oneof"`
}

type ValueList_BoolList struct {
	BoolList *BoolList `protobuf:"bytes,7,opt,name=boolList,proto3,oneof"`
}

type ValueList_TimestampList struct {
	TimestampList *TimestampList `protobuf:"bytes,8,opt,name=timestampList,proto3,oneof"`
}

func (*ValueList_BytesList) isValueList_ValueList() {}

func (*ValueList_StringList) isValueList_ValueList() {}

func (*ValueList_Int32List) isValueList_ValueList() {}

func (*ValueList_Int64List) isValueList_ValueList() {}

func (*ValueList_DoubleList) isValueList_ValueList() {}

func (*ValueList_FloatList) isValueList_ValueList() {}

func (*ValueList_BoolList) isValueList_ValueList() {}

func (*ValueList_TimestampList) isValueList_ValueList() {}

func (m *ValueList) GetValueList() isValueList_ValueList {
	if m != nil {
		return m.ValueList
	}
	return nil
}

func (m *ValueList) GetBytesList() *BytesList {
	if x, ok := m.GetValueList().(*ValueList_BytesList); ok {
		return x.BytesList
	}
	return nil
}

func (m *ValueList) GetStringList() *StringList {
	if x, ok := m.GetValueList().(*ValueList_StringList); ok {
		return x.StringList
	}
	return nil
}

func (m *ValueList) GetInt32List() *Int32List {
	if x, ok := m.GetValueList().(*ValueList_Int32List); ok {
		return x.Int32List
	}
	return nil
}

func (m *ValueList) GetInt64List() *Int64List {
	if x, ok := m.GetValueList().(*ValueList_Int64List); ok {
		return x.Int64List
	}
	return nil
}

func (m *ValueList) GetDoubleList() *DoubleList {
	if x, ok := m.GetValueList().(*ValueList_DoubleList); ok {
		return x.DoubleList
	}
	return nil
}

func (m *ValueList) GetFloatList() *FloatList {
	if x, ok := m.GetValueList().(*ValueList_FloatList); ok {
		return x.FloatList
	}
	return nil
}

func (m *ValueList) GetBoolList() *BoolList {
	if x, ok := m.GetValueList().(*ValueList_BoolList); ok {
		return x.BoolList
	}
	return nil
}

func (m *ValueList) GetTimestampList() *TimestampList {
	if x, ok := m.GetValueList().(*ValueList_TimestampList); ok {
		return x.TimestampList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ValueList) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ValueList_OneofMarshaler, _ValueList_OneofUnmarshaler, _ValueList_OneofSizer, []interface{}{
		(*ValueList_BytesList)(nil),
		(*ValueList_StringList)(nil),
		(*ValueList_Int32List)(nil),
		(*ValueList_Int64List)(nil),
		(*ValueList_DoubleList)(nil),
		(*ValueList_FloatList)(nil),
		(*ValueList_BoolList)(nil),
		(*ValueList_TimestampList)(nil),
	}
}

func _ValueList_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ValueList)
	// valueList
	switch x := m.ValueList.(type) {
	case *ValueList_BytesList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BytesList); err != nil {
			return err
		}
	case *ValueList_StringList:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringList); err != nil {
			return err
		}
	case *ValueList_Int32List:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int32List); err != nil {
			return err
		}
	case *ValueList_Int64List:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int64List); err != nil {
			return err
		}
	case *ValueList_DoubleList:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DoubleList); err != nil {
			return err
		}
	case *ValueList_FloatList:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FloatList); err != nil {
			return err
		}
	case *ValueList_BoolList:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BoolList); err != nil {
			return err
		}
	case *ValueList_TimestampList:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TimestampList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ValueList.ValueList has unexpected type %T", x)
	}
	return nil
}

func _ValueList_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ValueList)
	switch tag {
	case 1: // valueList.bytesList
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BytesList)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_BytesList{msg}
		return true, err
	case 2: // valueList.stringList
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StringList)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_StringList{msg}
		return true, err
	case 3: // valueList.int32List
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Int32List)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_Int32List{msg}
		return true, err
	case 4: // valueList.int64List
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Int64List)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_Int64List{msg}
		return true, err
	case 5: // valueList.doubleList
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DoubleList)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_DoubleList{msg}
		return true, err
	case 6: // valueList.floatList
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FloatList)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_FloatList{msg}
		return true, err
	case 7: // valueList.boolList
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(BoolList)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_BoolList{msg}
		return true, err
	case 8: // valueList.timestampList
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TimestampList)
		err := b.DecodeMessage(msg)
		m.ValueList = &ValueList_TimestampList{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ValueList_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ValueList)
	// valueList
	switch x := m.ValueList.(type) {
	case *ValueList_BytesList:
		s := proto.Size(x.BytesList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_StringList:
		s := proto.Size(x.StringList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_Int32List:
		s := proto.Size(x.Int32List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_Int64List:
		s := proto.Size(x.Int64List)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_DoubleList:
		s := proto.Size(x.DoubleList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_FloatList:
		s := proto.Size(x.FloatList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_BoolList:
		s := proto.Size(x.BoolList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ValueList_TimestampList:
		s := proto.Size(x.TimestampList)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BytesList struct {
	Val                  [][]byte `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BytesList) Reset()         { *m = BytesList{} }
func (m *BytesList) String() string { return proto.CompactTextString(m) }
func (*BytesList) ProtoMessage()    {}
func (*BytesList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{3}
}
func (m *BytesList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BytesList.Unmarshal(m, b)
}
func (m *BytesList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BytesList.Marshal(b, m, deterministic)
}
func (dst *BytesList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesList.Merge(dst, src)
}
func (m *BytesList) XXX_Size() int {
	return xxx_messageInfo_BytesList.Size(m)
}
func (m *BytesList) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesList.DiscardUnknown(m)
}

var xxx_messageInfo_BytesList proto.InternalMessageInfo

func (m *BytesList) GetVal() [][]byte {
	if m != nil {
		return m.Val
	}
	return nil
}

type StringList struct {
	Val                  []string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringList) Reset()         { *m = StringList{} }
func (m *StringList) String() string { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()    {}
func (*StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{4}
}
func (m *StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringList.Unmarshal(m, b)
}
func (m *StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringList.Marshal(b, m, deterministic)
}
func (dst *StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringList.Merge(dst, src)
}
func (m *StringList) XXX_Size() int {
	return xxx_messageInfo_StringList.Size(m)
}
func (m *StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_StringList proto.InternalMessageInfo

func (m *StringList) GetVal() []string {
	if m != nil {
		return m.Val
	}
	return nil
}

type Int32List struct {
	Val                  []int32  `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32List) Reset()         { *m = Int32List{} }
func (m *Int32List) String() string { return proto.CompactTextString(m) }
func (*Int32List) ProtoMessage()    {}
func (*Int32List) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{5}
}
func (m *Int32List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32List.Unmarshal(m, b)
}
func (m *Int32List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32List.Marshal(b, m, deterministic)
}
func (dst *Int32List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32List.Merge(dst, src)
}
func (m *Int32List) XXX_Size() int {
	return xxx_messageInfo_Int32List.Size(m)
}
func (m *Int32List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32List.DiscardUnknown(m)
}

var xxx_messageInfo_Int32List proto.InternalMessageInfo

func (m *Int32List) GetVal() []int32 {
	if m != nil {
		return m.Val
	}
	return nil
}

type Int64List struct {
	Val                  []int64  `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64List) Reset()         { *m = Int64List{} }
func (m *Int64List) String() string { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()    {}
func (*Int64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{6}
}
func (m *Int64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64List.Unmarshal(m, b)
}
func (m *Int64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64List.Marshal(b, m, deterministic)
}
func (dst *Int64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64List.Merge(dst, src)
}
func (m *Int64List) XXX_Size() int {
	return xxx_messageInfo_Int64List.Size(m)
}
func (m *Int64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64List.DiscardUnknown(m)
}

var xxx_messageInfo_Int64List proto.InternalMessageInfo

func (m *Int64List) GetVal() []int64 {
	if m != nil {
		return m.Val
	}
	return nil
}

type DoubleList struct {
	Val                  []float64 `protobuf:"fixed64,1,rep,packed,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DoubleList) Reset()         { *m = DoubleList{} }
func (m *DoubleList) String() string { return proto.CompactTextString(m) }
func (*DoubleList) ProtoMessage()    {}
func (*DoubleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{7}
}
func (m *DoubleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleList.Unmarshal(m, b)
}
func (m *DoubleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleList.Marshal(b, m, deterministic)
}
func (dst *DoubleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleList.Merge(dst, src)
}
func (m *DoubleList) XXX_Size() int {
	return xxx_messageInfo_DoubleList.Size(m)
}
func (m *DoubleList) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleList.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleList proto.InternalMessageInfo

func (m *DoubleList) GetVal() []float64 {
	if m != nil {
		return m.Val
	}
	return nil
}

type FloatList struct {
	Val                  []float32 `protobuf:"fixed32,1,rep,packed,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FloatList) Reset()         { *m = FloatList{} }
func (m *FloatList) String() string { return proto.CompactTextString(m) }
func (*FloatList) ProtoMessage()    {}
func (*FloatList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{8}
}
func (m *FloatList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FloatList.Unmarshal(m, b)
}
func (m *FloatList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FloatList.Marshal(b, m, deterministic)
}
func (dst *FloatList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FloatList.Merge(dst, src)
}
func (m *FloatList) XXX_Size() int {
	return xxx_messageInfo_FloatList.Size(m)
}
func (m *FloatList) XXX_DiscardUnknown() {
	xxx_messageInfo_FloatList.DiscardUnknown(m)
}

var xxx_messageInfo_FloatList proto.InternalMessageInfo

func (m *FloatList) GetVal() []float32 {
	if m != nil {
		return m.Val
	}
	return nil
}

type BoolList struct {
	Val                  []bool   `protobuf:"varint,1,rep,packed,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolList) Reset()         { *m = BoolList{} }
func (m *BoolList) String() string { return proto.CompactTextString(m) }
func (*BoolList) ProtoMessage()    {}
func (*BoolList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{9}
}
func (m *BoolList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolList.Unmarshal(m, b)
}
func (m *BoolList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolList.Marshal(b, m, deterministic)
}
func (dst *BoolList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolList.Merge(dst, src)
}
func (m *BoolList) XXX_Size() int {
	return xxx_messageInfo_BoolList.Size(m)
}
func (m *BoolList) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolList.DiscardUnknown(m)
}

var xxx_messageInfo_BoolList proto.InternalMessageInfo

func (m *BoolList) GetVal() []bool {
	if m != nil {
		return m.Val
	}
	return nil
}

type TimestampList struct {
	Val                  []*timestamp.Timestamp `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TimestampList) Reset()         { *m = TimestampList{} }
func (m *TimestampList) String() string { return proto.CompactTextString(m) }
func (*TimestampList) ProtoMessage()    {}
func (*TimestampList) Descriptor() ([]byte, []int) {
	return fileDescriptor_Value_d050960009ca8e98, []int{10}
}
func (m *TimestampList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimestampList.Unmarshal(m, b)
}
func (m *TimestampList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimestampList.Marshal(b, m, deterministic)
}
func (dst *TimestampList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimestampList.Merge(dst, src)
}
func (m *TimestampList) XXX_Size() int {
	return xxx_messageInfo_TimestampList.Size(m)
}
func (m *TimestampList) XXX_DiscardUnknown() {
	xxx_messageInfo_TimestampList.DiscardUnknown(m)
}

var xxx_messageInfo_TimestampList proto.InternalMessageInfo

func (m *TimestampList) GetVal() []*timestamp.Timestamp {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*ValueType)(nil), "feast.types.ValueType")
	proto.RegisterType((*Value)(nil), "feast.types.Value")
	proto.RegisterType((*ValueList)(nil), "feast.types.ValueList")
	proto.RegisterType((*BytesList)(nil), "feast.types.BytesList")
	proto.RegisterType((*StringList)(nil), "feast.types.StringList")
	proto.RegisterType((*Int32List)(nil), "feast.types.Int32List")
	proto.RegisterType((*Int64List)(nil), "feast.types.Int64List")
	proto.RegisterType((*DoubleList)(nil), "feast.types.DoubleList")
	proto.RegisterType((*FloatList)(nil), "feast.types.FloatList")
	proto.RegisterType((*BoolList)(nil), "feast.types.BoolList")
	proto.RegisterType((*TimestampList)(nil), "feast.types.TimestampList")
	proto.RegisterEnum("feast.types.ValueType_Enum", ValueType_Enum_name, ValueType_Enum_value)
}

func init() { proto.RegisterFile("feast/types/Value.proto", fileDescriptor_Value_d050960009ca8e98) }

var fileDescriptor_Value_d050960009ca8e98 = []byte{
	// 626 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xc1, 0x6e, 0xda, 0x4e,
	0x10, 0xc6, 0xbd, 0x18, 0x83, 0x3d, 0x24, 0x92, 0xb5, 0xd2, 0xff, 0x9f, 0x28, 0x4a, 0x52, 0x8b,
	0x93, 0x0f, 0x8d, 0x2d, 0x41, 0x8a, 0xd4, 0x43, 0xa5, 0xc6, 0x0a, 0x29, 0xa8, 0x04, 0x22, 0xe3,
	0xa4, 0x6a, 0x6f, 0x90, 0x3a, 0x0e, 0xad, 0x89, 0x51, 0xbc, 0x44, 0xca, 0xa1, 0x6f, 0xd3, 0x77,
	0xeb, 0x6b, 0x54, 0x33, 0xb6, 0x97, 0x45, 0x42, 0xbd, 0xad, 0xe7, 0xf7, 0x7d, 0x3b, 0xeb, 0xf9,
	0xec, 0x85, 0x83, 0x87, 0x78, 0x96, 0x0b, 0x5f, 0xbc, 0xae, 0xe2, 0xdc, 0xbf, 0x9b, 0xa5, 0xeb,
	0xd8, 0x5b, 0x3d, 0x67, 0x22, 0xe3, 0x2d, 0x02, 0x1e, 0x81, 0xa3, 0x37, 0x49, 0x96, 0x25, 0x69,
	0xec, 0x13, 0x9a, 0xaf, 0x1f, 0x7c, 0xb1, 0x58, 0xc6, 0xb9, 0x98, 0x2d, 0x57, 0x85, 0xba, 0xfd,
	0x0b, 0x2c, 0x32, 0x47, 0xaf, 0xab, 0xb8, 0xbd, 0x82, 0x7a, 0xff, 0x69, 0xbd, 0xe4, 0x2d, 0x68,
	0xde, 0x8e, 0x3f, 0x8f, 0x27, 0x5f, 0xc6, 0xb6, 0xc6, 0x2d, 0x30, 0x82, 0xaf, 0x51, 0x7f, 0x6a,
	0x33, 0x0e, 0xd0, 0x98, 0x46, 0xe1, 0x70, 0xfc, 0xc9, 0xae, 0x61, 0x79, 0x38, 0x8e, 0xba, 0x1d,
	0x5b, 0x2f, 0x97, 0xbd, 0x73, 0xbb, 0x8e, 0x8a, 0xcb, 0xc9, 0x6d, 0x30, 0xea, 0xdb, 0x06, 0x96,
	0xaf, 0x46, 0x93, 0x8b, 0xc8, 0x6e, 0x70, 0x13, 0xea, 0xc1, 0x64, 0x32, 0xb2, 0x9b, 0x7c, 0x1f,
	0xac, 0x68, 0x78, 0xdd, 0x9f, 0x46, 0x17, 0xd7, 0x37, 0xb6, 0xd9, 0xfe, 0x5d, 0x03, 0x83, 0xfa,
	0xf3, 0x63, 0x30, 0xe7, 0xaf, 0x22, 0xce, 0xef, 0x66, 0xe9, 0x21, 0x73, 0x98, 0xbb, 0x37, 0xd0,
	0x42, 0x59, 0xe1, 0xa7, 0x60, 0xe5, 0xe2, 0x79, 0xf1, 0x94, 0x20, 0xae, 0x39, 0xcc, 0xb5, 0x06,
	0x5a, 0xb8, 0x29, 0xa1, 0x7b, 0xf1, 0x24, 0xba, 0x1d, 0xc4, 0xba, 0xc3, 0x5c, 0x03, 0xdd, 0x55,
	0xa5, 0xa4, 0xbd, 0x73, 0xa4, 0x75, 0x87, 0xb9, 0x7a, 0x49, 0xa9, 0x82, 0x7b, 0x7f, 0xcf, 0xd6,
	0xf3, 0x34, 0x46, 0x6c, 0x38, 0xcc, 0x65, 0xb8, 0xb7, 0x2c, 0xa1, 0xfb, 0x21, 0xcd, 0x66, 0x02,
	0x71, 0xc3, 0x61, 0x6e, 0x0d, 0xdd, 0x55, 0x85, 0x1f, 0x41, 0x73, 0x9e, 0x65, 0x29, 0xc2, 0xa6,
	0xc3, 0x5c, 0x73, 0xa0, 0x85, 0x55, 0x81, 0x7f, 0x84, 0x3d, 0x39, 0x6f, 0x14, 0x98, 0x0e, 0x73,
	0x5b, 0x9d, 0x23, 0xaf, 0x08, 0xc5, 0xab, 0x42, 0xf1, 0xa2, 0x4a, 0x34, 0xd0, 0xc2, 0x2d, 0x47,
	0x60, 0x80, 0xfe, 0x32, 0x4b, 0xdb, 0x7f, 0xf4, 0x32, 0xa6, 0xd1, 0x22, 0x17, 0xbc, 0x07, 0x16,
	0x0d, 0x06, 0x1f, 0x68, 0x56, 0xad, 0xce, 0xff, 0x9e, 0x92, 0xba, 0x17, 0x54, 0x14, 0x5f, 0x44,
	0x4a, 0xf9, 0x7b, 0x80, 0x62, 0x62, 0x64, 0xac, 0x91, 0xf1, 0x60, 0xcb, 0x38, 0x95, 0x78, 0xa0,
	0x85, 0x8a, 0x18, 0x5b, 0xd2, 0x34, 0xc9, 0xa9, 0xef, 0x68, 0x39, 0xac, 0x28, 0xb6, 0x94, 0xd2,
	0xd2, 0xd7, 0x3b, 0x27, 0x5f, 0x7d, 0xb7, 0xaf, 0xa0, 0xa5, 0xaf, 0x78, 0xc0, 0xa3, 0x16, 0x01,
	0x90, 0xd1, 0xd8, 0x71, 0xd4, 0x4b, 0x89, 0xf1, 0xa8, 0x1b, 0x31, 0xb6, 0xa4, 0x70, 0xc8, 0xd9,
	0xd8, 0xd1, 0xf2, 0xaa, 0xa2, 0xd8, 0x52, 0x4a, 0x79, 0x17, 0x4c, 0xcc, 0x8d, 0x6c, 0x4d, 0xb2,
	0xfd, 0xb7, 0x3d, 0xd4, 0x12, 0xd2, 0x77, 0x59, 0xae, 0x79, 0x00, 0xfb, 0x32, 0x2f, 0x72, 0x56,
	0x11, 0xab, 0xce, 0x48, 0x55, 0x0c, 0xb4, 0x70, 0xdb, 0x12, 0xb4, 0xc0, 0x7a, 0xa9, 0xb2, 0x6d,
	0x9f, 0x80, 0x25, 0xd3, 0xe3, 0x36, 0xa5, 0x7f, 0xc8, 0x1c, 0xdd, 0xdd, 0x0b, 0xe9, 0x43, 0x38,
	0x05, 0xd8, 0x64, 0xa4, 0x72, 0xab, 0xe0, 0x27, 0x60, 0xc9, 0x24, 0x54, 0x6c, 0xa8, 0xb8, 0x9c,
	0xb1, 0x82, 0x75, 0xb9, 0xfb, 0x66, 0xac, 0x2a, 0x67, 0xd2, 0x2e, 0x87, 0xa7, 0xe2, 0x5a, 0x81,
	0x8f, 0xc1, 0xac, 0x86, 0xa4, 0x52, 0xb3, 0xa0, 0x1f, 0x60, 0x7f, 0x6b, 0x10, 0xfc, 0xed, 0x46,
	0xf2, 0xcf, 0x9f, 0x82, 0xec, 0xc1, 0x0d, 0xa8, 0x17, 0x5b, 0x00, 0xf4, 0x3b, 0xdc, 0xa0, 0xf8,
	0xdb, 0xbb, 0x64, 0x21, 0x1e, 0xd7, 0x73, 0xef, 0x3e, 0x5b, 0xfa, 0x49, 0xf6, 0x23, 0xfe, 0x29,
	0xe2, 0xfb, 0x47, 0xbf, 0xb8, 0x1f, 0x93, 0xec, 0x8c, 0x16, 0x67, 0xb4, 0xaf, 0xaf, 0x5c, 0x9a,
	0xf3, 0x06, 0x95, 0xba, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x66, 0xda, 0x0e, 0x5f, 0x4a, 0x05,
	0x00, 0x00,
}
