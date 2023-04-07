// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smrpc.proto

package smrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type FilesystemTrimRequest struct {
	EncryptedDevice      bool     `protobuf:"varint,1,opt,name=encrypted_device,json=encryptedDevice,proto3" json:"encrypted_device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilesystemTrimRequest) Reset()         { *m = FilesystemTrimRequest{} }
func (m *FilesystemTrimRequest) String() string { return proto.CompactTextString(m) }
func (*FilesystemTrimRequest) ProtoMessage()    {}
func (*FilesystemTrimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_754d325a94992d5d, []int{0}
}

func (m *FilesystemTrimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilesystemTrimRequest.Unmarshal(m, b)
}
func (m *FilesystemTrimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilesystemTrimRequest.Marshal(b, m, deterministic)
}
func (m *FilesystemTrimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilesystemTrimRequest.Merge(m, src)
}
func (m *FilesystemTrimRequest) XXX_Size() int {
	return xxx_messageInfo_FilesystemTrimRequest.Size(m)
}
func (m *FilesystemTrimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilesystemTrimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilesystemTrimRequest proto.InternalMessageInfo

func (m *FilesystemTrimRequest) GetEncryptedDevice() bool {
	if m != nil {
		return m.EncryptedDevice
	}
	return false
}

func init() {
	proto.RegisterType((*FilesystemTrimRequest)(nil), "FilesystemTrimRequest")
}

func init() { proto.RegisterFile("smrpc.proto", fileDescriptor_754d325a94992d5d) }

var fileDescriptor_754d325a94992d5d = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xce, 0x2d, 0x2a,
	0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5,
	0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0x92, 0x4a, 0x4e, 0x5c,
	0xa2, 0x6e, 0x99, 0x39, 0xa9, 0xc5, 0x95, 0xc5, 0x25, 0xa9, 0xb9, 0x21, 0x45, 0x99, 0xb9, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x9a, 0x5c, 0x02, 0xa9, 0x79, 0xc9, 0x45, 0x95, 0x05,
	0x25, 0xa9, 0x29, 0xf1, 0x29, 0xa9, 0x65, 0x99, 0xc9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c,
	0x41, 0xfc, 0x70, 0x71, 0x17, 0xb0, 0xb0, 0x51, 0x24, 0x97, 0x70, 0x70, 0x46, 0x62, 0x51, 0xaa,
	0x6f, 0x62, 0x5e, 0x62, 0x7a, 0x6a, 0x51, 0x70, 0x6a, 0x11, 0x48, 0x58, 0xc8, 0x89, 0x8b, 0x0f,
	0xd5, 0x68, 0x21, 0x31, 0x3d, 0xac, 0x76, 0x49, 0x89, 0xe9, 0x41, 0x9c, 0xa8, 0x07, 0x73, 0xa2,
	0x9e, 0x2b, 0xc8, 0x89, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x6f, 0xdd, 0x3d, 0x37, 0xd1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShareManagerServiceClient is the client API for ShareManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShareManagerServiceClient interface {
	FilesystemTrim(ctx context.Context, in *FilesystemTrimRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type shareManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewShareManagerServiceClient(cc *grpc.ClientConn) ShareManagerServiceClient {
	return &shareManagerServiceClient{cc}
}

func (c *shareManagerServiceClient) FilesystemTrim(ctx context.Context, in *FilesystemTrimRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ShareManagerService/FilesystemTrim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShareManagerServiceServer is the server API for ShareManagerService service.
type ShareManagerServiceServer interface {
	FilesystemTrim(context.Context, *FilesystemTrimRequest) (*empty.Empty, error)
}

// UnimplementedShareManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedShareManagerServiceServer struct {
}

func (*UnimplementedShareManagerServiceServer) FilesystemTrim(ctx context.Context, req *FilesystemTrimRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilesystemTrim not implemented")
}

func RegisterShareManagerServiceServer(s *grpc.Server, srv ShareManagerServiceServer) {
	s.RegisterService(&_ShareManagerService_serviceDesc, srv)
}

func _ShareManagerService_FilesystemTrim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesystemTrimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShareManagerServiceServer).FilesystemTrim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShareManagerService/FilesystemTrim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShareManagerServiceServer).FilesystemTrim(ctx, req.(*FilesystemTrimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShareManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShareManagerService",
	HandlerType: (*ShareManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FilesystemTrim",
			Handler:    _ShareManagerService_FilesystemTrim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smrpc.proto",
}
