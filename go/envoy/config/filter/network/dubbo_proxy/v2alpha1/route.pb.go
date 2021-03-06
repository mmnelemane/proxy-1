// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto

package envoy_config_filter_network_dubbo_proxy_v2alpha1

import (
	fmt "fmt"
	route "github.com/cilium/proxy/go/envoy/api/v2/route"
	_type "github.com/cilium/proxy/go/envoy/type"
	matcher "github.com/cilium/proxy/go/envoy/type/matcher"
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

// [#next-free-field: 6]
type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The interface name of the service.
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	// Which group does the interface belong to.
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// The version number of the interface.
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	// Route matching parameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Method level routing matching.
	Method *MethodMatch `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	// Specifies a set of headers that the route should match on. The router will check the request’s
	// headers against all the specified headers in the route config. A match will happen if all the
	// headers in the route are present in the request with the same values (or based on presence if
	// the value field is not in the config).
	Headers              []*route.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *route.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *route.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	// The name of the method.
	Name *matcher.StringMatcher `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Method parameter definition.
	// The key is the parameter index, starting from 0.
	// The value is the parameter matching type.
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *matcher.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

// The parameter matching type.
type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_74a572433a3292e0, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *_type.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *_type.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.config.filter.network.dubbo_proxy.v2alpha1.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto", fileDescriptor_74a572433a3292e0)
}

var fileDescriptor_74a572433a3292e0 = []byte{
	// 700 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xee, 0x24, 0x75, 0x7b, 0x7b, 0x7c, 0xaf, 0x94, 0x8e, 0xae, 0x5a, 0xdf, 0x5c, 0x04, 0x6d,
	0xd8, 0x94, 0x8d, 0x4d, 0x53, 0xfe, 0x29, 0x48, 0xa4, 0x42, 0x2a, 0x8b, 0x42, 0xe5, 0xaa, 0x62,
	0x03, 0x8a, 0xa6, 0xce, 0x24, 0x19, 0x35, 0x99, 0x31, 0xe3, 0x49, 0xda, 0x3c, 0x00, 0x1b, 0x36,
	0x6c, 0x40, 0xb0, 0xe4, 0x61, 0x90, 0x78, 0x13, 0x1e, 0x80, 0x15, 0x62, 0x81, 0x90, 0xcf, 0x8c,
	0x89, 0x81, 0x6e, 0x5a, 0xd8, 0x58, 0xe3, 0x33, 0xe7, 0xfb, 0xce, 0xf9, 0x3e, 0x9d, 0x33, 0xb0,
	0xc9, 0xe5, 0x58, 0x4d, 0xa2, 0x44, 0xc9, 0xae, 0xe8, 0x45, 0x5d, 0x31, 0x30, 0x5c, 0x47, 0x92,
	0x9b, 0x23, 0xa5, 0x0f, 0xa3, 0xce, 0xe8, 0xe0, 0x40, 0xb5, 0x53, 0xad, 0x8e, 0x27, 0xd1, 0xb8,
	0xc9, 0x06, 0x69, 0x9f, 0xad, 0x47, 0x5a, 0x8d, 0x0c, 0x0f, 0x53, 0xad, 0x8c, 0xa2, 0x97, 0x11,
	0x1d, 0x5a, 0x74, 0x68, 0xd1, 0xa1, 0x43, 0x87, 0x25, 0x74, 0x58, 0xa0, 0xeb, 0x97, 0x6c, 0x3d,
	0x96, 0x8a, 0x68, 0xdc, 0xb4, 0x5c, 0xf6, 0xdb, 0x4e, 0xd4, 0x30, 0x55, 0x92, 0x4b, 0x93, 0x59,
	0xf2, 0xfa, 0x05, 0x9b, 0x6a, 0x26, 0x29, 0x8f, 0x86, 0xcc, 0x24, 0x7d, 0xae, 0xa3, 0xcc, 0x68,
	0x21, 0x7b, 0x2e, 0x61, 0xa9, 0x94, 0xa0, 0x99, 0xec, 0xb9, 0xae, 0xea, 0xe7, 0x47, 0x9d, 0x94,
	0x45, 0x4c, 0x4a, 0x65, 0x98, 0x11, 0x4a, 0x66, 0xd1, 0x50, 0xf4, 0x34, 0x2b, 0xba, 0xae, 0x2f,
	0x8f, 0xd9, 0x40, 0x74, 0x98, 0xe1, 0x51, 0x71, 0xb0, 0x17, 0x8d, 0x0f, 0x04, 0x68, 0x9c, 0x37,
	0xb3, 0x85, 0x82, 0x46, 0x1a, 0xe1, 0x94, 0xc2, 0xac, 0x64, 0x43, 0x1e, 0x90, 0x15, 0xb2, 0xb6,
	0x10, 0xe3, 0x99, 0x9e, 0x83, 0x05, 0x21, 0x0d, 0xd7, 0x5d, 0x96, 0xf0, 0xa0, 0x82, 0x17, 0xd3,
	0x00, 0xfd, 0x17, 0xbc, 0x9e, 0x56, 0xa3, 0x34, 0xa8, 0xe2, 0x8d, 0xfd, 0xa1, 0x01, 0xcc, 0x8f,
	0xb9, 0xce, 0x84, 0x92, 0xc1, 0x2c, 0xc6, 0x8b, 0x5f, 0xfa, 0x08, 0xe6, 0xd0, 0x84, 0x2c, 0xf0,
	0x56, 0xaa, 0x6b, 0x7e, 0xf3, 0x7a, 0x78, 0x5a, 0x63, 0x43, 0xec, 0x3b, 0x76, 0x34, 0x8d, 0xf7,
	0x04, 0x3c, 0x8c, 0xd0, 0x27, 0xe0, 0xa1, 0x79, 0xd8, 0xbd, 0xdf, 0xdc, 0x3c, 0x23, 0xf3, 0x4e,
	0xce, 0xd1, 0xfa, 0xeb, 0x4b, 0xcb, 0x7b, 0x41, 0x2a, 0x35, 0x12, 0x5b, 0x52, 0xfa, 0x14, 0x3c,
	0xac, 0x88, 0x16, 0xf8, 0xcd, 0x3b, 0x67, 0x64, 0xbf, 0x97, 0xe4, 0x46, 0x97, 0xe9, 0x91, 0xb5,
	0xf1, 0x8e, 0x00, 0x4c, 0xcb, 0xd3, 0x7d, 0x98, 0x1b, 0x72, 0xd3, 0x57, 0x1d, 0x27, 0xe6, 0x0c,
	0xe5, 0x76, 0x10, 0x8f, 0x74, 0xb1, 0x23, 0xa3, 0xb7, 0x61, 0xbe, 0xcf, 0x59, 0x87, 0xeb, 0x2c,
	0xa8, 0xa0, 0xfd, 0xab, 0x8e, 0x97, 0xa5, 0x22, 0x1c, 0x37, 0x43, 0x3b, 0xf1, 0xdb, 0x98, 0xb2,
	0x63, 0x07, 0x31, 0x2e, 0x10, 0x8d, 0x37, 0x04, 0xfc, 0x92, 0x06, 0x5a, 0x87, 0xf9, 0x64, 0x30,
	0xca, 0x0c, 0xd7, 0x76, 0x5e, 0xb6, 0x67, 0xe2, 0x22, 0x40, 0x63, 0x58, 0x3c, 0xe2, 0xa2, 0xd7,
	0x37, 0xbc, 0xd3, 0x76, 0xb1, 0xcc, 0x39, 0x77, 0xf1, 0xa4, 0x92, 0x8f, 0x5d, 0xf2, 0x96, 0xcd,
	0xdd, 0x9e, 0x89, 0x6b, 0x47, 0x3f, 0x86, 0xb2, 0x56, 0x00, 0x8b, 0x8e, 0xaa, 0x9d, 0xa5, 0x3c,
	0x11, 0x5d, 0xc1, 0x35, 0xad, 0x7e, 0x6e, 0x91, 0xc6, 0xc7, 0x2a, 0xf8, 0x25, 0xb9, 0xf4, 0x6a,
	0x69, 0x8c, 0xa7, 0x1a, 0xf3, 0xed, 0x09, 0xdd, 0x7a, 0x85, 0x7b, 0xb8, 0x5e, 0x85, 0x46, 0x3b,
	0xe9, 0xcf, 0xe0, 0xef, 0x94, 0x69, 0x36, 0xcc, 0xda, 0x76, 0x8e, 0xac, 0x45, 0x0f, 0x7f, 0xcb,
	0xfa, 0x70, 0x17, 0x19, 0xf1, 0x7c, 0x5f, 0x1a, 0x3d, 0x89, 0xfd, 0x74, 0x1a, 0xa9, 0xbf, 0x22,
	0xb0, 0x8c, 0x19, 0xdc, 0x38, 0xc7, 0xf7, 0xbe, 0x4b, 0x5b, 0x05, 0x9f, 0x1f, 0xb3, 0xc4, 0xb8,
	0x6e, 0xaa, 0xce, 0x63, 0xc0, 0xa0, 0x15, 0x7a, 0x13, 0x7c, 0x7c, 0x0e, 0x5c, 0xca, 0x2c, 0xea,
	0x5d, 0x2a, 0xeb, 0x7d, 0x20, 0xcd, 0xb5, 0x2b, 0x71, 0x9e, 0x93, 0x43, 0x31, 0xd9, 0x0e, 0xf8,
	0xff, 0xf0, 0x5f, 0x5a, 0x14, 0xb6, 0xf0, 0xa9, 0xab, 0xf5, 0xd7, 0x04, 0x6a, 0x3f, 0x37, 0x4e,
	0x6b, 0x50, 0x3d, 0xe4, 0x13, 0x34, 0xf5, 0x9f, 0x38, 0x3f, 0xd2, 0x43, 0xf0, 0xc6, 0x6c, 0x30,
	0x2a, 0x76, 0x62, 0xff, 0x0f, 0x38, 0xf5, 0xab, 0x0f, 0xb1, 0xad, 0x71, 0xab, 0x72, 0x83, 0xb4,
	0x9e, 0x93, 0x4f, 0x6f, 0xbf, 0xbe, 0xf4, 0xd6, 0x69, 0x64, 0x2b, 0xf1, 0x63, 0xc3, 0x65, 0xfe,
	0xae, 0x64, 0xae, 0x5a, 0x76, 0x72, 0xb9, 0x0d, 0xb8, 0x2b, 0x94, 0xed, 0xce, 0x46, 0x4e, 0xdb,
	0x68, 0xcb, 0x2e, 0xe7, 0x6e, 0xfe, 0x78, 0xee, 0x92, 0x83, 0x39, 0x7c, 0x45, 0x37, 0xbe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0x99, 0xad, 0x07, 0x54, 0x06, 0x00, 0x00,
}
