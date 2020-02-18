// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/dubbo_proxy/v3/dubbo_proxy.proto

package envoy_extensions_filters_network_dubbo_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Dubbo Protocol types supported by Envoy.
type ProtocolType int32

const (
	// the default protocol.
	ProtocolType_Dubbo ProtocolType = 0
)

var ProtocolType_name = map[int32]string{
	0: "Dubbo",
}

var ProtocolType_value = map[string]int32{
	"Dubbo": 0,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0aa79721abd555e, []int{0}
}

// Dubbo Serialization types supported by Envoy.
type SerializationType int32

const (
	// the default serialization protocol.
	SerializationType_Hessian2 SerializationType = 0
)

var SerializationType_name = map[int32]string{
	0: "Hessian2",
}

var SerializationType_value = map[string]int32{
	"Hessian2": 0,
}

func (x SerializationType) String() string {
	return proto.EnumName(SerializationType_name, int32(x))
}

func (SerializationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0aa79721abd555e, []int{1}
}

// [#next-free-field: 6]
type DubboProxy struct {
	// The human readable prefix to use when emitting statistics.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Configure the protocol used.
	ProtocolType ProtocolType `protobuf:"varint,2,opt,name=protocol_type,json=protocolType,proto3,enum=envoy.extensions.filters.network.dubbo_proxy.v3.ProtocolType" json:"protocol_type,omitempty"`
	// Configure the serialization protocol used.
	SerializationType SerializationType `protobuf:"varint,3,opt,name=serialization_type,json=serializationType,proto3,enum=envoy.extensions.filters.network.dubbo_proxy.v3.SerializationType" json:"serialization_type,omitempty"`
	// The route table for the connection manager is static and is specified in this property.
	RouteConfig []*RouteConfiguration `protobuf:"bytes,4,rep,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	// A list of individual Dubbo filters that make up the filter chain for requests made to the
	// Dubbo proxy. Order matters as the filters are processed sequentially. For backwards
	// compatibility, if no dubbo_filters are specified, a default Dubbo router filter
	// (`envoy.filters.dubbo.router`) is used.
	DubboFilters         []*DubboFilter `protobuf:"bytes,5,rep,name=dubbo_filters,json=dubboFilters,proto3" json:"dubbo_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DubboProxy) Reset()         { *m = DubboProxy{} }
func (m *DubboProxy) String() string { return proto.CompactTextString(m) }
func (*DubboProxy) ProtoMessage()    {}
func (*DubboProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0aa79721abd555e, []int{0}
}

func (m *DubboProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DubboProxy.Unmarshal(m, b)
}
func (m *DubboProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DubboProxy.Marshal(b, m, deterministic)
}
func (m *DubboProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DubboProxy.Merge(m, src)
}
func (m *DubboProxy) XXX_Size() int {
	return xxx_messageInfo_DubboProxy.Size(m)
}
func (m *DubboProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_DubboProxy.DiscardUnknown(m)
}

var xxx_messageInfo_DubboProxy proto.InternalMessageInfo

func (m *DubboProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *DubboProxy) GetProtocolType() ProtocolType {
	if m != nil {
		return m.ProtocolType
	}
	return ProtocolType_Dubbo
}

func (m *DubboProxy) GetSerializationType() SerializationType {
	if m != nil {
		return m.SerializationType
	}
	return SerializationType_Hessian2
}

func (m *DubboProxy) GetRouteConfig() []*RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *DubboProxy) GetDubboFilters() []*DubboFilter {
	if m != nil {
		return m.DubboFilters
	}
	return nil
}

// DubboFilter configures a Dubbo filter.
type DubboFilter struct {
	// The name of the filter to instantiate. The name must match a supported
	// filter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Filter specific configuration which depends on the filter being
	// instantiated. See the supported filters for further documentation.
	Config               *any.Any `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DubboFilter) Reset()         { *m = DubboFilter{} }
func (m *DubboFilter) String() string { return proto.CompactTextString(m) }
func (*DubboFilter) ProtoMessage()    {}
func (*DubboFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0aa79721abd555e, []int{1}
}

func (m *DubboFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DubboFilter.Unmarshal(m, b)
}
func (m *DubboFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DubboFilter.Marshal(b, m, deterministic)
}
func (m *DubboFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DubboFilter.Merge(m, src)
}
func (m *DubboFilter) XXX_Size() int {
	return xxx_messageInfo_DubboFilter.Size(m)
}
func (m *DubboFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_DubboFilter.DiscardUnknown(m)
}

var xxx_messageInfo_DubboFilter proto.InternalMessageInfo

func (m *DubboFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DubboFilter) GetConfig() *any.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.dubbo_proxy.v3.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterEnum("envoy.extensions.filters.network.dubbo_proxy.v3.SerializationType", SerializationType_name, SerializationType_value)
	proto.RegisterType((*DubboProxy)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.DubboProxy")
	proto.RegisterType((*DubboFilter)(nil), "envoy.extensions.filters.network.dubbo_proxy.v3.DubboFilter")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/dubbo_proxy/v3/dubbo_proxy.proto", fileDescriptor_c0aa79721abd555e)
}

var fileDescriptor_c0aa79721abd555e = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x97, 0xad, 0x1d, 0xdb, 0x97, 0x0c, 0x3a, 0x0b, 0x89, 0xae, 0x48, 0xa8, 0xdb, 0xa9,
	0x9a, 0x90, 0x2d, 0xd2, 0xdb, 0x58, 0x91, 0x96, 0x22, 0xc4, 0xb1, 0x0a, 0xdc, 0x2b, 0x77, 0x75,
	0x8a, 0x45, 0xb0, 0x23, 0xc7, 0x09, 0x0d, 0xdc, 0x38, 0xf1, 0x0c, 0x48, 0xbc, 0x03, 0x2f, 0xc1,
	0x4b, 0xed, 0x84, 0x62, 0x7b, 0x4a, 0x54, 0xb8, 0x84, 0x5b, 0xec, 0xcf, 0xf9, 0xfd, 0x3e, 0x7f,
	0xfe, 0xc3, 0x0d, 0x13, 0xa5, 0xac, 0x08, 0xdb, 0x6a, 0x26, 0x72, 0x2e, 0x45, 0x4e, 0x12, 0x9e,
	0x6a, 0xa6, 0x72, 0x22, 0x98, 0xfe, 0x2c, 0xd5, 0x47, 0xb2, 0x2e, 0x56, 0x2b, 0xb9, 0xcc, 0x94,
	0xdc, 0x56, 0xa4, 0x9c, 0xb6, 0x97, 0x38, 0x53, 0x52, 0x4b, 0x44, 0x0c, 0x02, 0x37, 0x08, 0xec,
	0x10, 0xd8, 0x21, 0x70, 0xfb, 0x9f, 0x72, 0x3a, 0x7a, 0xd9, 0xd5, 0xa9, 0x64, 0xa1, 0x99, 0xb5,
	0x8d, 0xce, 0x36, 0x52, 0x6e, 0x52, 0x46, 0xcc, 0x6a, 0x55, 0x24, 0x84, 0x0a, 0xd7, 0xc8, 0xe8,
	0xbc, 0x58, 0x67, 0x94, 0x50, 0x21, 0xa4, 0xa6, 0xda, 0x70, 0x4b, 0xa6, 0x6a, 0x01, 0x17, 0x1b,
	0x77, 0xe4, 0x49, 0x49, 0x53, 0xbe, 0xa6, 0x9a, 0x91, 0xfb, 0x0f, 0x5b, 0xb8, 0xf8, 0xd5, 0x03,
	0x78, 0x5d, 0x5b, 0x17, 0xb5, 0x14, 0x4d, 0xc0, 0xcf, 0x35, 0xd5, 0xcb, 0x4c, 0xb1, 0x84, 0x6f,
	0x87, 0xde, 0xd8, 0x9b, 0x1c, 0x47, 0x0f, 0xee, 0xa2, 0x9e, 0xda, 0x1f, 0x7b, 0x31, 0xd4, 0xb5,
	0x85, 0x29, 0xa1, 0x14, 0x4e, 0x0c, 0xe1, 0x56, 0xa6, 0x4b, 0x5d, 0x65, 0x6c, 0xb8, 0x3f, 0xf6,
	0x26, 0x0f, 0xc3, 0x19, 0xee, 0x38, 0x15, 0xbc, 0x70, 0x94, 0xf7, 0x55, 0xc6, 0xa2, 0xa3, 0xbb,
	0xa8, 0xff, 0xcd, 0xdb, 0x1f, 0x78, 0x71, 0x90, 0xb5, 0xf6, 0xd1, 0x57, 0x40, 0x39, 0x53, 0x9c,
	0xa6, 0xfc, 0x8b, 0xb9, 0xa2, 0x55, 0x1e, 0x18, 0x65, 0xd4, 0x59, 0xf9, 0xae, 0x8d, 0xda, 0xf1,
	0x9e, 0xe6, 0xbb, 0x45, 0x94, 0x40, 0x60, 0x5e, 0x62, 0x79, 0x2b, 0x45, 0xc2, 0x37, 0xc3, 0xde,
	0xf8, 0x60, 0xe2, 0x87, 0xf3, 0xce, 0xda, 0xb8, 0x86, 0xcc, 0x0d, 0xa3, 0x50, 0x06, 0x1f, 0xfb,
	0xaa, 0xd9, 0x43, 0x14, 0x4e, 0xec, 0x1f, 0x8e, 0x33, 0xec, 0x1b, 0xd1, 0x75, 0x67, 0x91, 0x79,
	0xd0, 0x37, 0xe6, 0x4c, 0x1c, 0xac, 0x9b, 0x45, 0x7e, 0x15, 0xfd, 0xf8, 0xfd, 0xfd, 0xd9, 0x0c,
	0x6c, 0x12, 0xb1, 0xbd, 0x8f, 0xa3, 0xfd, 0x1b, 0x16, 0xd2, 0x34, 0xfb, 0x40, 0x5f, 0xe0, 0x26,
	0x23, 0x17, 0x3f, 0x3d, 0xf0, 0x5b, 0x06, 0xf4, 0x14, 0x7a, 0x82, 0x7e, 0x62, 0xbb, 0x61, 0x31,
	0x9b, 0xe8, 0x39, 0x1c, 0xba, 0xa9, 0xd5, 0xf9, 0xf0, 0xc3, 0xc7, 0xd8, 0xe6, 0x18, 0xdf, 0xe7,
	0x18, 0xdf, 0x88, 0x2a, 0x76, 0x67, 0xae, 0xe6, 0x75, 0x7b, 0xaf, 0xe0, 0xfa, 0xff, 0xda, 0xb3,
	0xfd, 0x5c, 0x9e, 0x41, 0xd0, 0xce, 0x14, 0x3a, 0x86, 0xbe, 0x29, 0x0f, 0xf6, 0x2e, 0xcf, 0xe1,
	0xf4, 0xaf, 0xb7, 0x47, 0x01, 0x1c, 0xbd, 0x65, 0x79, 0xce, 0xa9, 0x08, 0x07, 0x7b, 0x51, 0x0c,
	0x33, 0x2e, 0xed, 0xc4, 0xad, 0xa3, 0xe3, 0xf0, 0xa3, 0x47, 0xcd, 0xa8, 0x4c, 0x1b, 0x0b, 0x6f,
	0x75, 0x68, 0x2e, 0x3b, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xab, 0xfc, 0xf4, 0x75, 0x04,
	0x00, 0x00,
}