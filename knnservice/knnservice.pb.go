// Code generated by protoc-gen-go. DO NOT EDIT.
// source: knnservice.proto

/*
Package knnservice is a generated protocol buffer package.

It is generated from these files:
	knnservice.proto

It has these top-level messages:
	KnnRequest
	Feature
	KnnResponse
	InsertionRequest
	InsertionResponse
	ServiceMessage
	Peer
	PeerMessage
	JoinRequest
	JoinResponse
*/
package knnservice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Request message for creating a new customer
type KnnRequest struct {
	Id        string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Timestamp int64     `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Timeout   int64     `protobuf:"varint,3,opt,name=timeout" json:"timeout,omitempty"`
	K         int32     `protobuf:"varint,4,opt,name=k" json:"k,omitempty"`
	Feature   []float64 `protobuf:"fixed64,5,rep,packed,name=feature" json:"feature,omitempty"`
}

func (m *KnnRequest) Reset()                    { *m = KnnRequest{} }
func (m *KnnRequest) String() string            { return proto.CompactTextString(m) }
func (*KnnRequest) ProtoMessage()               {}
func (*KnnRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KnnRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KnnRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *KnnRequest) GetTimeout() int64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *KnnRequest) GetK() int32 {
	if m != nil {
		return m.K
	}
	return 0
}

func (m *KnnRequest) GetFeature() []float64 {
	if m != nil {
		return m.Feature
	}
	return nil
}

type Feature struct {
	Feature    []float64 `protobuf:"fixed64,1,rep,packed,name=feature" json:"feature,omitempty"`
	Timestamp  int64     `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Label      string    `protobuf:"bytes,3,opt,name=label" json:"label,omitempty"`
	Grouplabel string    `protobuf:"bytes,4,opt,name=grouplabel" json:"grouplabel,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Feature) GetFeature() []float64 {
	if m != nil {
		return m.Feature
	}
	return nil
}

func (m *Feature) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Feature) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Feature) GetGrouplabel() string {
	if m != nil {
		return m.Grouplabel
	}
	return ""
}

type KnnResponse struct {
	Id       string     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Features []*Feature `protobuf:"bytes,2,rep,name=features" json:"features,omitempty"`
}

func (m *KnnResponse) Reset()                    { *m = KnnResponse{} }
func (m *KnnResponse) String() string            { return proto.CompactTextString(m) }
func (*KnnResponse) ProtoMessage()               {}
func (*KnnResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KnnResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *KnnResponse) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type InsertionRequest struct {
	Timestamp  int64     `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Label      string    `protobuf:"bytes,3,opt,name=label" json:"label,omitempty"`
	Grouplabel string    `protobuf:"bytes,4,opt,name=grouplabel" json:"grouplabel,omitempty"`
	Feature    []float64 `protobuf:"fixed64,5,rep,packed,name=feature" json:"feature,omitempty"`
}

func (m *InsertionRequest) Reset()                    { *m = InsertionRequest{} }
func (m *InsertionRequest) String() string            { return proto.CompactTextString(m) }
func (*InsertionRequest) ProtoMessage()               {}
func (*InsertionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InsertionRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *InsertionRequest) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *InsertionRequest) GetGrouplabel() string {
	if m != nil {
		return m.Grouplabel
	}
	return ""
}

func (m *InsertionRequest) GetFeature() []float64 {
	if m != nil {
		return m.Feature
	}
	return nil
}

type InsertionResponse struct {
	Code int32 `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
}

func (m *InsertionResponse) Reset()                    { *m = InsertionResponse{} }
func (m *InsertionResponse) String() string            { return proto.CompactTextString(m) }
func (*InsertionResponse) ProtoMessage()               {}
func (*InsertionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InsertionResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ServiceMessage struct {
	Services []string `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ServiceMessage) Reset()                    { *m = ServiceMessage{} }
func (m *ServiceMessage) String() string            { return proto.CompactTextString(m) }
func (*ServiceMessage) ProtoMessage()               {}
func (*ServiceMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ServiceMessage) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

type Peer struct {
	Address   string    `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Timestamp int64     `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Version   string    `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	Avg       []float64 `protobuf:"fixed64,4,rep,packed,name=avg" json:"avg,omitempty"`
	Hist      []float64 `protobuf:"fixed64,5,rep,packed,name=hist" json:"hist,omitempty"`
	N         int64     `protobuf:"varint,6,opt,name=n" json:"n,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Peer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Peer) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Peer) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Peer) GetAvg() []float64 {
	if m != nil {
		return m.Avg
	}
	return nil
}

func (m *Peer) GetHist() []float64 {
	if m != nil {
		return m.Hist
	}
	return nil
}

func (m *Peer) GetN() int64 {
	if m != nil {
		return m.N
	}
	return 0
}

type PeerMessage struct {
	Peers []*Peer `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
}

func (m *PeerMessage) Reset()                    { *m = PeerMessage{} }
func (m *PeerMessage) String() string            { return proto.CompactTextString(m) }
func (*PeerMessage) ProtoMessage()               {}
func (*PeerMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PeerMessage) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

type JoinRequest struct {
	Address   string    `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port      int32     `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	Avg       []float64 `protobuf:"fixed64,3,rep,packed,name=avg" json:"avg,omitempty"`
	Version   string    `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	Hist      []float64 `protobuf:"fixed64,5,rep,packed,name=hist" json:"hist,omitempty"`
	N         int64     `protobuf:"varint,6,opt,name=n" json:"n,omitempty"`
	Timestamp int64     `protobuf:"varint,7,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *JoinRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *JoinRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *JoinRequest) GetAvg() []float64 {
	if m != nil {
		return m.Avg
	}
	return nil
}

func (m *JoinRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *JoinRequest) GetHist() []float64 {
	if m != nil {
		return m.Hist
	}
	return nil
}

func (m *JoinRequest) GetN() int64 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *JoinRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type JoinResponse struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *JoinResponse) Reset()                    { *m = JoinResponse{} }
func (m *JoinResponse) String() string            { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()               {}
func (*JoinResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *JoinResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*KnnRequest)(nil), "knnservice.KnnRequest")
	proto.RegisterType((*Feature)(nil), "knnservice.Feature")
	proto.RegisterType((*KnnResponse)(nil), "knnservice.KnnResponse")
	proto.RegisterType((*InsertionRequest)(nil), "knnservice.InsertionRequest")
	proto.RegisterType((*InsertionResponse)(nil), "knnservice.InsertionResponse")
	proto.RegisterType((*ServiceMessage)(nil), "knnservice.ServiceMessage")
	proto.RegisterType((*Peer)(nil), "knnservice.Peer")
	proto.RegisterType((*PeerMessage)(nil), "knnservice.PeerMessage")
	proto.RegisterType((*JoinRequest)(nil), "knnservice.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "knnservice.JoinResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KnnService service

type KnnServiceClient interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	// rpc GetCustomers(CustomerFilter) returns (stream CustomerRequest) {}
	// Create a new Customer - A simple RPC
	GetKnn(ctx context.Context, in *KnnRequest, opts ...grpc.CallOption) (*KnnResponse, error)
	Insert(ctx context.Context, in *InsertionRequest, opts ...grpc.CallOption) (*InsertionResponse, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	ExchangeServices(ctx context.Context, in *ServiceMessage, opts ...grpc.CallOption) (*ServiceMessage, error)
	ExchangePeers(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerMessage, error)
}

type knnServiceClient struct {
	cc *grpc.ClientConn
}

func NewKnnServiceClient(cc *grpc.ClientConn) KnnServiceClient {
	return &knnServiceClient{cc}
}

func (c *knnServiceClient) GetKnn(ctx context.Context, in *KnnRequest, opts ...grpc.CallOption) (*KnnResponse, error) {
	out := new(KnnResponse)
	err := grpc.Invoke(ctx, "/knnservice.KnnService/GetKnn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knnServiceClient) Insert(ctx context.Context, in *InsertionRequest, opts ...grpc.CallOption) (*InsertionResponse, error) {
	out := new(InsertionResponse)
	err := grpc.Invoke(ctx, "/knnservice.KnnService/Insert", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knnServiceClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := grpc.Invoke(ctx, "/knnservice.KnnService/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knnServiceClient) ExchangeServices(ctx context.Context, in *ServiceMessage, opts ...grpc.CallOption) (*ServiceMessage, error) {
	out := new(ServiceMessage)
	err := grpc.Invoke(ctx, "/knnservice.KnnService/ExchangeServices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knnServiceClient) ExchangePeers(ctx context.Context, in *PeerMessage, opts ...grpc.CallOption) (*PeerMessage, error) {
	out := new(PeerMessage)
	err := grpc.Invoke(ctx, "/knnservice.KnnService/ExchangePeers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KnnService service

type KnnServiceServer interface {
	// Get all Customers with filter - A server-to-client streaming RPC.
	// rpc GetCustomers(CustomerFilter) returns (stream CustomerRequest) {}
	// Create a new Customer - A simple RPC
	GetKnn(context.Context, *KnnRequest) (*KnnResponse, error)
	Insert(context.Context, *InsertionRequest) (*InsertionResponse, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	ExchangeServices(context.Context, *ServiceMessage) (*ServiceMessage, error)
	ExchangePeers(context.Context, *PeerMessage) (*PeerMessage, error)
}

func RegisterKnnServiceServer(s *grpc.Server, srv KnnServiceServer) {
	s.RegisterService(&_KnnService_serviceDesc, srv)
}

func _KnnService_GetKnn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KnnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnnServiceServer).GetKnn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knnservice.KnnService/GetKnn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnnServiceServer).GetKnn(ctx, req.(*KnnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnnService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnnServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knnservice.KnnService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnnServiceServer).Insert(ctx, req.(*InsertionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnnService_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnnServiceServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knnservice.KnnService/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnnServiceServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnnService_ExchangeServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnnServiceServer).ExchangeServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knnservice.KnnService/ExchangeServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnnServiceServer).ExchangeServices(ctx, req.(*ServiceMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnnService_ExchangePeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnnServiceServer).ExchangePeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/knnservice.KnnService/ExchangePeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnnServiceServer).ExchangePeers(ctx, req.(*PeerMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _KnnService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "knnservice.KnnService",
	HandlerType: (*KnnServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetKnn",
			Handler:    _KnnService_GetKnn_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _KnnService_Insert_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _KnnService_Join_Handler,
		},
		{
			MethodName: "ExchangeServices",
			Handler:    _KnnService_ExchangeServices_Handler,
		},
		{
			MethodName: "ExchangePeers",
			Handler:    _KnnService_ExchangePeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "knnservice.proto",
}

func init() { proto.RegisterFile("knnservice.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x8d, 0xb3, 0x9b, 0xa4, 0x99, 0x94, 0x2a, 0x18, 0x04, 0xd6, 0xaa, 0xa0, 0xc8, 0x07, 0xd8,
	0x03, 0x2a, 0x52, 0x10, 0xa7, 0x1e, 0x11, 0x54, 0x50, 0x40, 0xc8, 0x7c, 0xc1, 0x36, 0x19, 0xd2,
	0x55, 0x5a, 0x7b, 0x59, 0x3b, 0x81, 0x03, 0x07, 0x3e, 0x80, 0xff, 0xe0, 0xbf, 0xf8, 0x12, 0x64,
	0xaf, 0x77, 0xe3, 0x6d, 0x94, 0xe6, 0xd4, 0xdb, 0xcc, 0xbc, 0x89, 0xe7, 0xbd, 0x37, 0xb3, 0x81,
	0xf1, 0x52, 0x4a, 0x8d, 0xe5, 0x3a, 0x9f, 0xe1, 0x49, 0x51, 0x2a, 0xa3, 0x28, 0x6c, 0x2a, 0xfc,
	0x17, 0xc0, 0xb9, 0x94, 0x02, 0xbf, 0xaf, 0x50, 0x1b, 0x7a, 0x04, 0xdd, 0x7c, 0xce, 0xc8, 0x84,
	0xa4, 0x43, 0xd1, 0xcd, 0xe7, 0xf4, 0x18, 0x86, 0x26, 0xbf, 0x46, 0x6d, 0xb2, 0xeb, 0x82, 0x75,
	0x27, 0x24, 0x8d, 0xc4, 0xa6, 0x40, 0x19, 0x0c, 0x6c, 0xa2, 0x56, 0x86, 0x45, 0x0e, 0xab, 0x53,
	0x7a, 0x08, 0x64, 0xc9, 0xe2, 0x09, 0x49, 0x7b, 0x82, 0x2c, 0x6d, 0xdf, 0x37, 0xcc, 0xcc, 0xaa,
	0x44, 0xd6, 0x9b, 0x44, 0x29, 0x11, 0x75, 0xca, 0x7f, 0xc0, 0xe0, 0x5d, 0x15, 0x86, 0x4d, 0xa4,
	0xd5, 0xb4, 0x87, 0xc4, 0x43, 0xe8, 0x5d, 0x65, 0x17, 0x78, 0xe5, 0x28, 0x0c, 0x45, 0x95, 0xd0,
	0xa7, 0x00, 0x8b, 0x52, 0xad, 0x8a, 0x0a, 0x8a, 0x1d, 0x14, 0x54, 0xf8, 0x67, 0x18, 0x39, 0xd9,
	0xba, 0x50, 0x52, 0xe3, 0x96, 0xee, 0x97, 0x70, 0xe0, 0xa7, 0x6b, 0xd6, 0x9d, 0x44, 0xe9, 0x68,
	0xfa, 0xe0, 0x24, 0xb0, 0xd1, 0x73, 0x16, 0x4d, 0x13, 0xff, 0x4d, 0x60, 0xfc, 0xde, 0xe2, 0x26,
	0x57, 0x8d, 0x9b, 0x77, 0x40, 0xfc, 0x16, 0x2f, 0x9f, 0xc3, 0xfd, 0x80, 0x81, 0x17, 0x46, 0x21,
	0x9e, 0xa9, 0x39, 0xba, 0xe9, 0x3d, 0xe1, 0x62, 0xfe, 0x02, 0x8e, 0xbe, 0x56, 0x42, 0x3e, 0xa1,
	0xd6, 0xd9, 0x02, 0x69, 0x02, 0x07, 0x5e, 0x9a, 0x76, 0xe6, 0x0f, 0x45, 0x93, 0xf3, 0x3f, 0x04,
	0xe2, 0x2f, 0x88, 0xa5, 0x9d, 0x9c, 0xcd, 0xe7, 0x25, 0x6a, 0xed, 0x8d, 0xaa, 0xd3, 0xfd, 0x57,
	0xb2, 0xc6, 0x52, 0xe7, 0x4a, 0x7a, 0xa5, 0x75, 0x4a, 0xc7, 0x10, 0x65, 0xeb, 0x05, 0x8b, 0x9d,
	0x0e, 0x1b, 0x5a, 0xba, 0x97, 0xb9, 0x36, 0x5e, 0x9a, 0x8b, 0xed, 0x2d, 0x49, 0xd6, 0x77, 0xaf,
	0x12, 0xc9, 0x5f, 0xc3, 0xc8, 0xb2, 0xa9, 0x99, 0x3f, 0x83, 0x5e, 0x81, 0x58, 0x56, 0xb4, 0x47,
	0xd3, 0x71, 0xb8, 0x25, 0xdb, 0x27, 0x2a, 0x98, 0xff, 0x25, 0x30, 0xfa, 0xa0, 0xf2, 0x66, 0x35,
	0xbb, 0xc5, 0x50, 0x88, 0x0b, 0x55, 0x9a, 0xda, 0x31, 0x1b, 0xd7, 0x44, 0xa3, 0x0d, 0xd1, 0x40,
	0x54, 0xdc, 0x16, 0xb5, 0x57, 0x42, 0xdb, 0xae, 0xc1, 0x0d, 0xbb, 0x78, 0x0a, 0x87, 0x15, 0x51,
	0xbf, 0xc1, 0x9d, 0x4c, 0xa7, 0xff, 0xba, 0xee, 0xdb, 0xf5, 0xbb, 0xa4, 0xa7, 0xd0, 0x3f, 0x43,
	0x73, 0x2e, 0x25, 0x7d, 0x14, 0xba, 0xb0, 0xf9, 0xba, 0x93, 0xc7, 0x5b, 0xf5, 0x6a, 0x06, 0xef,
	0xd0, 0x33, 0xe8, 0x57, 0xc7, 0x43, 0x8f, 0xc3, 0xa6, 0x9b, 0x27, 0x9d, 0x3c, 0xd9, 0x81, 0x36,
	0x0f, 0x9d, 0x42, 0x6c, 0xe9, 0xd3, 0xd6, 0xac, 0xc0, 0xf9, 0x84, 0x6d, 0x03, 0xcd, 0x8f, 0x3f,
	0xc2, 0xf8, 0xed, 0xcf, 0xd9, 0x65, 0x26, 0x17, 0xe8, 0x55, 0x69, 0x9a, 0x84, 0xfd, 0xed, 0xbb,
	0x4d, 0x6e, 0xc1, 0x78, 0x87, 0xbe, 0x81, 0x7b, 0xf5, 0x6b, 0xf6, 0x14, 0x74, 0x9b, 0x53, 0x70,
	0x45, 0xc9, 0x2e, 0x80, 0x77, 0x2e, 0xfa, 0xee, 0x2f, 0xf3, 0xd5, 0xff, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xe3, 0x24, 0x90, 0xe1, 0x46, 0x05, 0x00, 0x00,
}
