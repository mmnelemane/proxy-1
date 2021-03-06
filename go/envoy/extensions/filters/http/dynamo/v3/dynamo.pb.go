// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/dynamo/v3/dynamo.proto

package envoy_extensions_filters_http_dynamo_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Dynamo filter config.
type Dynamo struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dynamo) Reset()         { *m = Dynamo{} }
func (m *Dynamo) String() string { return proto.CompactTextString(m) }
func (*Dynamo) ProtoMessage()    {}
func (*Dynamo) Descriptor() ([]byte, []int) {
	return fileDescriptor_79057240c5b18ac4, []int{0}
}

func (m *Dynamo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dynamo.Unmarshal(m, b)
}
func (m *Dynamo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dynamo.Marshal(b, m, deterministic)
}
func (m *Dynamo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dynamo.Merge(m, src)
}
func (m *Dynamo) XXX_Size() int {
	return xxx_messageInfo_Dynamo.Size(m)
}
func (m *Dynamo) XXX_DiscardUnknown() {
	xxx_messageInfo_Dynamo.DiscardUnknown(m)
}

var xxx_messageInfo_Dynamo proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Dynamo)(nil), "envoy.extensions.filters.http.dynamo.v3.Dynamo")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/dynamo/v3/dynamo.proto", fileDescriptor_79057240c5b18ac4)
}

var fileDescriptor_79057240c5b18ac4 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd,
	0xd7, 0x2f, 0x33, 0x86, 0xb2, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd4, 0xc1, 0xba, 0xf4,
	0x10, 0xba, 0xf4, 0xa0, 0xba, 0xf4, 0x40, 0xba, 0xf4, 0xa0, 0x6a, 0xcb, 0x8c, 0xa5, 0x14, 0x4b,
	0x53, 0x0a, 0x12, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0xc0, 0xc6, 0x97, 0xa5, 0x16,
	0x81, 0x74, 0x64, 0xe6, 0xa5, 0x43, 0xcc, 0x52, 0xb2, 0xe2, 0x62, 0x73, 0x01, 0xab, 0xb7, 0x32,
	0x98, 0x75, 0xb4, 0x43, 0x4e, 0x9b, 0x4b, 0x13, 0x62, 0x78, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0x3a,
	0xd4, 0x60, 0x54, 0x73, 0x8d, 0xf4, 0x20, 0x3a, 0x9c, 0xdc, 0xb8, 0x4c, 0x33, 0xf3, 0xf5, 0xc0,
	0xea, 0x0b, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x88, 0x74, 0x97, 0x13, 0x37, 0xc4, 0x80, 0x00, 0x90,
	0x0b, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x4e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x70, 0x32,
	0xe9, 0xae, 0x0e, 0x01, 0x00, 0x00,
}
