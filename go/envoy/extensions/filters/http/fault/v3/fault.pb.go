// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/fault/v3/fault.proto

package envoy_extensions_filters_http_fault_v3

import (
	fmt "fmt"
	v32 "github.com/cilium/proxy/go/envoy/config/route/v3"
	v31 "github.com/cilium/proxy/go/envoy/extensions/filters/common/fault/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type FaultAbort struct {
	// Types that are valid to be assigned to ErrorType:
	//	*FaultAbort_HttpStatus
	ErrorType isFaultAbort_ErrorType `protobuf_oneof:"error_type"`
	// The percentage of requests/operations/connections that will be aborted with the error code
	// provided.
	Percentage           *v3.FractionalPercent `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FaultAbort) Reset()         { *m = FaultAbort{} }
func (m *FaultAbort) String() string { return proto.CompactTextString(m) }
func (*FaultAbort) ProtoMessage()    {}
func (*FaultAbort) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{0}
}

func (m *FaultAbort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultAbort.Unmarshal(m, b)
}
func (m *FaultAbort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultAbort.Marshal(b, m, deterministic)
}
func (m *FaultAbort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultAbort.Merge(m, src)
}
func (m *FaultAbort) XXX_Size() int {
	return xxx_messageInfo_FaultAbort.Size(m)
}
func (m *FaultAbort) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultAbort.DiscardUnknown(m)
}

var xxx_messageInfo_FaultAbort proto.InternalMessageInfo

type isFaultAbort_ErrorType interface {
	isFaultAbort_ErrorType()
}

type FaultAbort_HttpStatus struct {
	HttpStatus uint32 `protobuf:"varint,2,opt,name=http_status,json=httpStatus,proto3,oneof"`
}

func (*FaultAbort_HttpStatus) isFaultAbort_ErrorType() {}

func (m *FaultAbort) GetErrorType() isFaultAbort_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return nil
}

func (m *FaultAbort) GetHttpStatus() uint32 {
	if x, ok := m.GetErrorType().(*FaultAbort_HttpStatus); ok {
		return x.HttpStatus
	}
	return 0
}

func (m *FaultAbort) GetPercentage() *v3.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultAbort) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultAbort_HttpStatus)(nil),
	}
}

// [#next-free-field: 14]
type HTTPFault struct {
	// If specified, the filter will inject delays based on the values in the
	// object.
	Delay *v31.FaultDelay `protobuf:"bytes,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// If specified, the filter will abort requests based on the values in
	// the object. At least *abort* or *delay* must be specified.
	Abort *FaultAbort `protobuf:"bytes,2,opt,name=abort,proto3" json:"abort,omitempty"`
	// Specifies the name of the (destination) upstream cluster that the
	// filter should match on. Fault injection will be restricted to requests
	// bound to the specific upstream cluster.
	UpstreamCluster string `protobuf:"bytes,3,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Specifies a set of headers that the filter should match on. The fault
	// injection filter can be applied selectively to requests that match a set of
	// headers specified in the fault filter config. The chances of actual fault
	// injection further depend on the value of the :ref:`percentage
	// <envoy_api_field_extensions.filters.http.fault.v3.FaultAbort.percentage>` field.
	// The filter will check the request's headers against all the specified
	// headers in the filter config. A match will happen if all the headers in the
	// config are present in the request with the same values (or based on
	// presence if the *value* field is not in the config).
	Headers []*v32.HeaderMatcher `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	// Faults are injected for the specified list of downstream hosts. If this
	// setting is not set, faults are injected for all downstream nodes.
	// Downstream node name is taken from :ref:`the HTTP
	// x-envoy-downstream-service-node
	// <config_http_conn_man_headers_downstream-service-node>` header and compared
	// against downstream_nodes list.
	DownstreamNodes []string `protobuf:"bytes,5,rep,name=downstream_nodes,json=downstreamNodes,proto3" json:"downstream_nodes,omitempty"`
	// The maximum number of faults that can be active at a single time via the configured fault
	// filter. Note that because this setting can be overridden at the route level, it's possible
	// for the number of active faults to be greater than this value (if injected via a different
	// route). If not specified, defaults to unlimited. This setting can be overridden via
	// `runtime <config_http_filters_fault_injection_runtime>` and any faults that are not injected
	// due to overflow will be indicated via the `faults_overflow
	// <config_http_filters_fault_injection_stats>` stat.
	//
	// .. attention::
	//   Like other :ref:`circuit breakers <arch_overview_circuit_break>` in Envoy, this is a fuzzy
	//   limit. It's possible for the number of active faults to rise slightly above the configured
	//   amount due to the implementation details.
	MaxActiveFaults *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=max_active_faults,json=maxActiveFaults,proto3" json:"max_active_faults,omitempty"`
	// The response rate limit to be applied to the response body of the stream. When configured,
	// the percentage can be overridden by the :ref:`fault.http.rate_limit.response_percent
	// <config_http_filters_fault_injection_runtime>` runtime key.
	//
	// .. attention::
	//  This is a per-stream limit versus a connection level limit. This means that concurrent streams
	//  will each get an independent limit.
	ResponseRateLimit *v31.FaultRateLimit `protobuf:"bytes,7,opt,name=response_rate_limit,json=responseRateLimit,proto3" json:"response_rate_limit,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_delay_percent
	DelayPercentRuntime string `protobuf:"bytes,8,opt,name=delay_percent_runtime,json=delayPercentRuntime,proto3" json:"delay_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.abort_percent
	AbortPercentRuntime string `protobuf:"bytes,9,opt,name=abort_percent_runtime,json=abortPercentRuntime,proto3" json:"abort_percent_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.delay.fixed_duration_ms
	DelayDurationRuntime string `protobuf:"bytes,10,opt,name=delay_duration_runtime,json=delayDurationRuntime,proto3" json:"delay_duration_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.abort.http_status
	AbortHttpStatusRuntime string `protobuf:"bytes,11,opt,name=abort_http_status_runtime,json=abortHttpStatusRuntime,proto3" json:"abort_http_status_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.max_active_faults
	MaxActiveFaultsRuntime string `protobuf:"bytes,12,opt,name=max_active_faults_runtime,json=maxActiveFaultsRuntime,proto3" json:"max_active_faults_runtime,omitempty"`
	// The runtime key to override the :ref:`default <config_http_filters_fault_injection_runtime>`
	// runtime. The default is: fault.http.rate_limit.response_percent
	ResponseRateLimitPercentRuntime string   `protobuf:"bytes,13,opt,name=response_rate_limit_percent_runtime,json=responseRateLimitPercentRuntime,proto3" json:"response_rate_limit_percent_runtime,omitempty"`
	XXX_NoUnkeyedLiteral            struct{} `json:"-"`
	XXX_unrecognized                []byte   `json:"-"`
	XXX_sizecache                   int32    `json:"-"`
}

func (m *HTTPFault) Reset()         { *m = HTTPFault{} }
func (m *HTTPFault) String() string { return proto.CompactTextString(m) }
func (*HTTPFault) ProtoMessage()    {}
func (*HTTPFault) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c8de848a52219f2, []int{1}
}

func (m *HTTPFault) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HTTPFault.Unmarshal(m, b)
}
func (m *HTTPFault) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HTTPFault.Marshal(b, m, deterministic)
}
func (m *HTTPFault) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HTTPFault.Merge(m, src)
}
func (m *HTTPFault) XXX_Size() int {
	return xxx_messageInfo_HTTPFault.Size(m)
}
func (m *HTTPFault) XXX_DiscardUnknown() {
	xxx_messageInfo_HTTPFault.DiscardUnknown(m)
}

var xxx_messageInfo_HTTPFault proto.InternalMessageInfo

func (m *HTTPFault) GetDelay() *v31.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *HTTPFault) GetAbort() *FaultAbort {
	if m != nil {
		return m.Abort
	}
	return nil
}

func (m *HTTPFault) GetUpstreamCluster() string {
	if m != nil {
		return m.UpstreamCluster
	}
	return ""
}

func (m *HTTPFault) GetHeaders() []*v32.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPFault) GetDownstreamNodes() []string {
	if m != nil {
		return m.DownstreamNodes
	}
	return nil
}

func (m *HTTPFault) GetMaxActiveFaults() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxActiveFaults
	}
	return nil
}

func (m *HTTPFault) GetResponseRateLimit() *v31.FaultRateLimit {
	if m != nil {
		return m.ResponseRateLimit
	}
	return nil
}

func (m *HTTPFault) GetDelayPercentRuntime() string {
	if m != nil {
		return m.DelayPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortPercentRuntime() string {
	if m != nil {
		return m.AbortPercentRuntime
	}
	return ""
}

func (m *HTTPFault) GetDelayDurationRuntime() string {
	if m != nil {
		return m.DelayDurationRuntime
	}
	return ""
}

func (m *HTTPFault) GetAbortHttpStatusRuntime() string {
	if m != nil {
		return m.AbortHttpStatusRuntime
	}
	return ""
}

func (m *HTTPFault) GetMaxActiveFaultsRuntime() string {
	if m != nil {
		return m.MaxActiveFaultsRuntime
	}
	return ""
}

func (m *HTTPFault) GetResponseRateLimitPercentRuntime() string {
	if m != nil {
		return m.ResponseRateLimitPercentRuntime
	}
	return ""
}

func init() {
	proto.RegisterType((*FaultAbort)(nil), "envoy.extensions.filters.http.fault.v3.FaultAbort")
	proto.RegisterType((*HTTPFault)(nil), "envoy.extensions.filters.http.fault.v3.HTTPFault")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/fault/v3/fault.proto", fileDescriptor_1c8de848a52219f2)
}

var fileDescriptor_1c8de848a52219f2 = []byte{
	// 703 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xaf, 0x9b, 0x7e, 0x4e, 0x6e, 0xd5, 0xd6, 0xbd, 0xb7, 0xd7, 0xb7, 0xa0, 0x12, 0x0a,
	0x42, 0xe1, 0x6b, 0x2c, 0x39, 0x59, 0x40, 0x17, 0x88, 0x86, 0xaa, 0x0a, 0x55, 0x41, 0x91, 0x29,
	0x6c, 0xad, 0x69, 0x3c, 0x49, 0x2c, 0xd9, 0x33, 0xd6, 0xcc, 0xd8, 0x4d, 0x96, 0xec, 0x78, 0x06,
	0xde, 0x07, 0x89, 0x05, 0x0f, 0xc0, 0xa3, 0x20, 0x56, 0x68, 0xce, 0xd8, 0x6e, 0xda, 0xb4, 0xa2,
	0xec, 0x9c, 0x39, 0xe7, 0x77, 0xbe, 0xfe, 0xe7, 0x04, 0x79, 0x94, 0xe5, 0x7c, 0xe2, 0xd2, 0xb1,
	0xa2, 0x4c, 0x46, 0x9c, 0x49, 0x77, 0x10, 0xc5, 0x8a, 0x0a, 0xe9, 0x8e, 0x94, 0x4a, 0xdd, 0x01,
	0xc9, 0x62, 0xe5, 0xe6, 0x2d, 0xf3, 0x81, 0x53, 0xc1, 0x15, 0xb7, 0x1f, 0x00, 0x83, 0xcf, 0x19,
	0x5c, 0x30, 0x58, 0x33, 0xd8, 0xb8, 0xe6, 0xad, 0xed, 0x27, 0x26, 0x76, 0x9f, 0xb3, 0x41, 0x34,
	0x74, 0x05, 0xcf, 0x14, 0xd5, 0xa1, 0xe0, 0x23, 0xe8, 0xf3, 0x24, 0xe5, 0x8c, 0x32, 0x25, 0x4d,
	0xd4, 0xed, 0xf6, 0xb5, 0x95, 0xf4, 0x79, 0x92, 0x70, 0x76, 0x65, 0x2d, 0xdb, 0xb7, 0x0c, 0xa5,
	0x26, 0x29, 0xc4, 0x4e, 0xa9, 0xe8, 0x53, 0x56, 0x1a, 0x77, 0x86, 0x9c, 0x0f, 0x63, 0xea, 0xc2,
	0xaf, 0xd3, 0x6c, 0xe0, 0x9e, 0x09, 0x92, 0xa6, 0xba, 0x50, 0x63, 0xbf, 0x9b, 0x85, 0x29, 0x71,
	0x09, 0x63, 0x5c, 0x11, 0x05, 0x29, 0x73, 0x2a, 0x74, 0xee, 0x88, 0x0d, 0x0b, 0x97, 0xff, 0x72,
	0x12, 0x47, 0x21, 0xd1, 0x95, 0x17, 0x1f, 0xc6, 0xb0, 0xfb, 0xcd, 0x42, 0xe8, 0x50, 0x17, 0xb2,
	0x7f, 0xca, 0x85, 0xb2, 0x31, 0xaa, 0xeb, 0xe6, 0x03, 0xa9, 0x88, 0xca, 0xa4, 0x33, 0xd7, 0xb0,
	0x9a, 0xab, 0x9d, 0xfa, 0xcf, 0xce, 0xf2, 0xa3, 0xc5, 0xf5, 0xef, 0xf3, 0xcd, 0xaf, 0x56, 0xf7,
	0x2f, 0x1f, 0x69, 0x8f, 0x77, 0xe0, 0x60, 0xbf, 0x44, 0xa8, 0xa8, 0x95, 0x0c, 0xa9, 0x53, 0x6b,
	0x58, 0xcd, 0xba, 0xd7, 0xc0, 0x66, 0xb0, 0xba, 0x19, 0x9c, 0xb7, 0xf0, 0xa1, 0x20, 0x7d, 0x5d,
	0x16, 0x89, 0x7b, 0xc6, 0xd5, 0x9f, 0x62, 0xf6, 0x5a, 0x9f, 0xbf, 0x7c, 0xda, 0xc1, 0xc8, 0x0c,
	0x19, 0x9b, 0x21, 0x17, 0x42, 0x5c, 0xd0, 0xc1, 0xc3, 0xe7, 0x65, 0x76, 0x36, 0x10, 0xa2, 0x42,
	0x70, 0x11, 0xe8, 0x1c, 0x76, 0xed, 0x47, 0xc7, 0x3a, 0x9a, 0x5f, 0xb6, 0xd6, 0xe7, 0x76, 0x3f,
	0x2e, 0xa1, 0x95, 0xee, 0xc9, 0x49, 0x0f, 0x7c, 0xed, 0x23, 0xb4, 0x10, 0xd2, 0x98, 0x4c, 0x1c,
	0x0b, 0x0a, 0x6b, 0xe3, 0x6b, 0x15, 0x37, 0xda, 0x54, 0x9a, 0x9b, 0x5c, 0x07, 0x9a, 0xf5, 0x4d,
	0x08, 0xbb, 0x8b, 0x16, 0x88, 0xce, 0x0d, 0x33, 0xa9, 0x7b, 0x1e, 0xbe, 0xd9, 0xf6, 0x4c, 0x55,
	0xed, 0x9b, 0x00, 0xf6, 0x43, 0xb4, 0x9e, 0xa5, 0x52, 0x09, 0x4a, 0x92, 0xa0, 0x1f, 0x67, 0x52,
	0x51, 0x01, 0x93, 0x5b, 0xf1, 0xd7, 0xca, 0xf7, 0x57, 0xe6, 0xd9, 0x7e, 0x81, 0x96, 0x46, 0x94,
	0x84, 0x54, 0x48, 0x67, 0xbe, 0x51, 0x6b, 0xd6, 0xbd, 0xfb, 0xf8, 0xc2, 0x9c, 0x60, 0x07, 0x75,
	0x96, 0x2e, 0x78, 0xbd, 0x21, 0xaa, 0x3f, 0xa2, 0xc2, 0x2f, 0x21, 0x9d, 0x2a, 0xe4, 0x67, 0xac,
	0x48, 0xc6, 0x78, 0x48, 0xa5, 0xb3, 0xd0, 0xa8, 0xe9, 0x54, 0xe7, 0xef, 0x6f, 0xf5, 0xb3, 0xdd,
	0x45, 0x1b, 0x09, 0x19, 0x07, 0x5a, 0xaa, 0x9c, 0x06, 0x50, 0xbe, 0x74, 0x16, 0xa1, 0xd7, 0xdb,
	0xd8, 0x2c, 0x20, 0x2e, 0x17, 0x10, 0xbf, 0x7f, 0xcd, 0x54, 0xcb, 0xfb, 0x40, 0xe2, 0x8c, 0xfa,
	0x6b, 0x09, 0x19, 0xef, 0x03, 0x05, 0xad, 0x4a, 0x7b, 0x84, 0x36, 0x05, 0x95, 0x29, 0x67, 0x92,
	0x06, 0x82, 0x28, 0x1a, 0xc4, 0x51, 0x12, 0x29, 0x67, 0x09, 0x62, 0x3d, 0xfb, 0x43, 0x0d, 0x7c,
	0xa2, 0xe8, 0xb1, 0xe6, 0xfd, 0x8d, 0x32, 0x68, 0xf5, 0x64, 0x7b, 0xe8, 0x5f, 0x10, 0x27, 0x28,
	0xf6, 0x29, 0x10, 0x19, 0x53, 0x51, 0x42, 0x9d, 0x65, 0x18, 0xe7, 0x26, 0x18, 0xcb, 0xa5, 0x33,
	0x26, 0xcd, 0x80, 0x0c, 0x33, 0xcc, 0x8a, 0x61, 0xc0, 0x78, 0x89, 0x69, 0xa3, 0x2d, 0x93, 0x27,
	0xcc, 0x04, 0x5c, 0x58, 0x05, 0x21, 0x80, 0xfe, 0x01, 0xeb, 0x41, 0x61, 0x2c, 0xa9, 0xe7, 0xe8,
	0x7f, 0x93, 0x69, 0xea, 0xa2, 0x2a, 0xb0, 0x0e, 0xe0, 0x16, 0x38, 0x74, 0xab, 0x7b, 0x9a, 0x42,
	0x67, 0xc4, 0xa8, 0xd0, 0xbf, 0x0d, 0x7a, 0x69, 0xec, 0x25, 0x7a, 0x8c, 0xee, 0x5d, 0x31, 0xfd,
	0x99, 0x6e, 0x57, 0x21, 0xc8, 0x9d, 0x99, 0x99, 0x5e, 0xec, 0x7c, 0xcf, 0xd3, 0xd7, 0xf9, 0x14,
	0x3d, 0xfe, 0xfd, 0x75, 0x56, 0x57, 0xd7, 0x39, 0x40, 0xed, 0x88, 0x1b, 0x99, 0x53, 0xc1, 0xc7,
	0x93, 0x1b, 0x5e, 0x4a, 0xc7, 0xfc, 0x0f, 0xf5, 0xf4, 0x8e, 0xf5, 0xac, 0xd3, 0x45, 0x58, 0xb6,
	0xd6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x48, 0xbe, 0x34, 0xa5, 0xda, 0x05, 0x00, 0x00,
}
