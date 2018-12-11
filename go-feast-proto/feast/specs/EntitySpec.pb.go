// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/specs/EntitySpec.proto

package specs // import "github.com/gojektech/feast/go-feast-proto/feast/specs"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntitySpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntitySpec) Reset()         { *m = EntitySpec{} }
func (m *EntitySpec) String() string { return proto.CompactTextString(m) }
func (*EntitySpec) ProtoMessage()    {}
func (*EntitySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_EntitySpec_acaf20f0bc303278, []int{0}
}
func (m *EntitySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntitySpec.Unmarshal(m, b)
}
func (m *EntitySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntitySpec.Marshal(b, m, deterministic)
}
func (dst *EntitySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntitySpec.Merge(dst, src)
}
func (m *EntitySpec) XXX_Size() int {
	return xxx_messageInfo_EntitySpec.Size(m)
}
func (m *EntitySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EntitySpec.DiscardUnknown(m)
}

var xxx_messageInfo_EntitySpec proto.InternalMessageInfo

func (m *EntitySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntitySpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EntitySpec) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*EntitySpec)(nil), "feast.specs.EntitySpec")
}

func init() {
	proto.RegisterFile("feast/specs/EntitySpec.proto", fileDescriptor_EntitySpec_acaf20f0bc303278)
}

var fileDescriptor_EntitySpec_acaf20f0bc303278 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4b, 0x4d, 0x2c,
	0x2e, 0xd1, 0x2f, 0x2e, 0x48, 0x4d, 0x2e, 0xd6, 0x77, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9, 0x0c, 0x2e,
	0x48, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0xcb, 0xea, 0x81, 0x65, 0x95,
	0xc2, 0xb8, 0xb8, 0x10, 0x0a, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x05, 0x2e, 0xee, 0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc,
	0x82, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x26, 0xb0, 0x14, 0xb2, 0x10, 0x48, 0x57, 0x49, 0x62, 0x7a,
	0xb1, 0x04, 0xb3, 0x02, 0x33, 0x48, 0x17, 0x88, 0xed, 0x14, 0xca, 0x85, 0x6c, 0x8d, 0x13, 0x3f,
	0xc2, 0x92, 0x00, 0x90, 0x23, 0xa2, 0x4c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0xd3, 0xf3, 0xb3, 0x52, 0xb3, 0x4b, 0x52, 0x93, 0x33, 0xf4, 0x21, 0xee, 0x4e, 0xcf,
	0xd7, 0x05, 0x33, 0x74, 0xc1, 0xee, 0xd5, 0x47, 0xf2, 0x4c, 0x12, 0x1b, 0x58, 0xc8, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0xa8, 0xdf, 0xe2, 0xe2, 0x00, 0x00, 0x00,
}
