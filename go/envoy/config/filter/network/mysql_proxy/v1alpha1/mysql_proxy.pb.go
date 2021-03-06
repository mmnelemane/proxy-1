// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/mysql_proxy/v1alpha1/mysql_proxy.proto

package envoy_config_filter_network_mysql_proxy_v1alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type MySQLProxy struct {
	// The human readable prefix to use when emitting :ref:`statistics
	// <config_network_filters_mysql_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// [#not-implemented-hide:] The optional path to use for writing MySQL access logs.
	// If the access log field is empty, access logs will not be written.
	AccessLog            string   `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLProxy) Reset()         { *m = MySQLProxy{} }
func (m *MySQLProxy) String() string { return proto.CompactTextString(m) }
func (*MySQLProxy) ProtoMessage()    {}
func (*MySQLProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4bac5cccef760ed, []int{0}
}

func (m *MySQLProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLProxy.Unmarshal(m, b)
}
func (m *MySQLProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLProxy.Marshal(b, m, deterministic)
}
func (m *MySQLProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLProxy.Merge(m, src)
}
func (m *MySQLProxy) XXX_Size() int {
	return xxx_messageInfo_MySQLProxy.Size(m)
}
func (m *MySQLProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLProxy proto.InternalMessageInfo

func (m *MySQLProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MySQLProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func init() {
	proto.RegisterType((*MySQLProxy)(nil), "envoy.config.filter.network.mysql_proxy.v1alpha1.MySQLProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mysql_proxy/v1alpha1/mysql_proxy.proto", fileDescriptor_c4bac5cccef760ed)
}

var fileDescriptor_c4bac5cccef760ed = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x95, 0x8a, 0x1f, 0xd5, 0x0c, 0x48, 0x59, 0xa8, 0x2a, 0x81, 0x2a, 0xa6, 0x4e, 0x36,
	0x51, 0x77, 0x86, 0xcc, 0xad, 0x14, 0x8a, 0x98, 0xa3, 0x4b, 0xea, 0x04, 0x8b, 0xd4, 0xd7, 0xd8,
	0x97, 0x90, 0xbc, 0x02, 0x0b, 0x2b, 0xcf, 0xc9, 0xc8, 0x80, 0x90, 0xed, 0x56, 0xea, 0xc0, 0xc2,
	0x66, 0x9d, 0x73, 0xfc, 0xc9, 0x9f, 0x59, 0x2e, 0x75, 0x87, 0x83, 0xa8, 0x50, 0xd7, 0xaa, 0x11,
	0xb5, 0x6a, 0x49, 0x5a, 0xa1, 0x25, 0xbd, 0xa1, 0x7d, 0x16, 0xdb, 0xc1, 0xbd, 0xb4, 0xa5, 0xb1,
	0xd8, 0x0f, 0xa2, 0xcb, 0xa0, 0x35, 0x4f, 0x90, 0x1d, 0x86, 0xdc, 0x58, 0x24, 0x4c, 0x6f, 0x02,
	0x83, 0x47, 0x06, 0x8f, 0x0c, 0xbe, 0x63, 0xf0, 0xc3, 0xf9, 0x9e, 0x31, 0xbd, 0x7a, 0xdd, 0x18,
	0x10, 0xa0, 0x35, 0x12, 0x90, 0x42, 0xed, 0xc4, 0x56, 0x35, 0x16, 0x48, 0x46, 0xe2, 0xf4, 0xa2,
	0x83, 0x56, 0x6d, 0x80, 0xa4, 0xd8, 0x1f, 0x62, 0x71, 0xfd, 0xc0, 0xd8, 0x6a, 0xb8, 0xbf, 0x5b,
	0x16, 0x9e, 0x97, 0xce, 0xd9, 0x99, 0x23, 0xa0, 0xd2, 0x58, 0x59, 0xab, 0x7e, 0x92, 0xcc, 0x92,
	0xf9, 0x38, 0x3f, 0xfd, 0xce, 0x8f, 0xec, 0x68, 0x96, 0xac, 0x99, 0xef, 0x8a, 0x50, 0xa5, 0x97,
	0x8c, 0x41, 0x55, 0x49, 0xe7, 0xca, 0x16, 0x9b, 0xc9, 0xc8, 0x0f, 0xd7, 0xe3, 0x98, 0x2c, 0xb1,
	0xc9, 0xdf, 0x93, 0xaf, 0xcf, 0x9f, 0x8f, 0xe3, 0x2c, 0x15, 0x51, 0x45, 0xf6, 0x24, 0xb5, 0xf3,
	0x0f, 0xdb, 0xe9, 0xb8, 0xbf, 0x7d, 0x16, 0xec, 0x56, 0x21, 0x0f, 0x77, 0x62, 0xf2, 0xdf, 0x9f,
	0xc8, 0xcf, 0x57, 0x3e, 0x0d, 0x3a, 0x85, 0x37, 0x2c, 0x92, 0xc7, 0x93, 0xa0, 0xba, 0xf8, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xc9, 0x50, 0x23, 0x91, 0x9b, 0x01, 0x00, 0x00,
}
