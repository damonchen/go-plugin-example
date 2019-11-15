// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EchoMesssage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMesssage) Reset()         { *m = EchoMesssage{} }
func (m *EchoMesssage) String() string { return proto.CompactTextString(m) }
func (*EchoMesssage) ProtoMessage()    {}
func (*EchoMesssage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *EchoMesssage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMesssage.Unmarshal(m, b)
}
func (m *EchoMesssage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMesssage.Marshal(b, m, deterministic)
}
func (m *EchoMesssage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMesssage.Merge(m, src)
}
func (m *EchoMesssage) XXX_Size() int {
	return xxx_messageInfo_EchoMesssage.Size(m)
}
func (m *EchoMesssage) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMesssage.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMesssage proto.InternalMessageInfo

type EchoMesssage_Req struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMesssage_Req) Reset()         { *m = EchoMesssage_Req{} }
func (m *EchoMesssage_Req) String() string { return proto.CompactTextString(m) }
func (*EchoMesssage_Req) ProtoMessage()    {}
func (*EchoMesssage_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 0}
}

func (m *EchoMesssage_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMesssage_Req.Unmarshal(m, b)
}
func (m *EchoMesssage_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMesssage_Req.Marshal(b, m, deterministic)
}
func (m *EchoMesssage_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMesssage_Req.Merge(m, src)
}
func (m *EchoMesssage_Req) XXX_Size() int {
	return xxx_messageInfo_EchoMesssage_Req.Size(m)
}
func (m *EchoMesssage_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMesssage_Req.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMesssage_Req proto.InternalMessageInfo

func (m *EchoMesssage_Req) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EchoMesssage_Resp struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoMesssage_Resp) Reset()         { *m = EchoMesssage_Resp{} }
func (m *EchoMesssage_Resp) String() string { return proto.CompactTextString(m) }
func (*EchoMesssage_Resp) ProtoMessage()    {}
func (*EchoMesssage_Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0, 1}
}

func (m *EchoMesssage_Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoMesssage_Resp.Unmarshal(m, b)
}
func (m *EchoMesssage_Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoMesssage_Resp.Marshal(b, m, deterministic)
}
func (m *EchoMesssage_Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoMesssage_Resp.Merge(m, src)
}
func (m *EchoMesssage_Resp) XXX_Size() int {
	return xxx_messageInfo_EchoMesssage_Resp.Size(m)
}
func (m *EchoMesssage_Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoMesssage_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_EchoMesssage_Resp proto.InternalMessageInfo

func (m *EchoMesssage_Resp) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*EchoMesssage)(nil), "EchoMesssage")
	proto.RegisterType((*EchoMesssage_Req)(nil), "EchoMesssage.Req")
	proto.RegisterType((*EchoMesssage_Resp)(nil), "EchoMesssage.Resp")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x72, 0xe5, 0xe2, 0x71, 0x4d, 0xce,
	0xc8, 0xf7, 0x4d, 0x2d, 0x2e, 0x2e, 0x4e, 0x4c, 0x4f, 0x95, 0x92, 0xe4, 0x62, 0x0e, 0x4a, 0x2d,
	0x14, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02,
	0xb3, 0xa5, 0xa4, 0xb8, 0x58, 0x82, 0x52, 0x8b, 0x0b, 0xb0, 0xc9, 0x19, 0x99, 0x73, 0xb1, 0x80,
	0x8c, 0x11, 0xd2, 0xe7, 0x62, 0x49, 0x05, 0xd1, 0x82, 0x7a, 0xc8, 0xa6, 0xea, 0x05, 0xa5, 0x16,
	0x4a, 0x09, 0xa1, 0x0b, 0x15, 0x17, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x9d, 0x61, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x18, 0x66, 0x71, 0xc1, 0x97, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Echo(ctx context.Context, in *EchoMesssage_Req, opts ...grpc.CallOption) (*EchoMesssage_Resp, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *EchoMesssage_Req, opts ...grpc.CallOption) (*EchoMesssage_Resp, error) {
	out := new(EchoMesssage_Resp)
	err := c.cc.Invoke(ctx, "/Echo/echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Echo(context.Context, *EchoMesssage_Req) (*EchoMesssage_Resp, error)
}

// UnimplementedEchoServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServer struct {
}

func (*UnimplementedEchoServer) Echo(ctx context.Context, req *EchoMesssage_Req) (*EchoMesssage_Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMesssage_Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoMesssage_Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
