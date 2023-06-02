// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/longhorn/longhorn-instance-manager/pkg/imrpc/common.proto

package imrpc

import (
	fmt "fmt"
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

type BackendStoreDriver int32

const (
	BackendStoreDriver_longhorn BackendStoreDriver = 0
	BackendStoreDriver_spdk     BackendStoreDriver = 1
)

var BackendStoreDriver_name = map[int32]string{
	0: "longhorn",
	1: "spdk",
}

var BackendStoreDriver_value = map[string]int32{
	"longhorn": 0,
	"spdk":     1,
}

func (x BackendStoreDriver) String() string {
	return proto.EnumName(BackendStoreDriver_name, int32(x))
}

func (BackendStoreDriver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8198dc9351c39e1c, []int{0}
}

func init() {
	proto.RegisterEnum("imrpc.BackendStoreDriver", BackendStoreDriver_name, BackendStoreDriver_value)
}

func init() {
	proto.RegisterFile("github.com/longhorn/longhorn-instance-manager/pkg/imrpc/common.proto", fileDescriptor_8198dc9351c39e1c)
}

var fileDescriptor_8198dc9351c39e1c = []byte{
	// 134 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8c, 0x51, 0x0a, 0xc2, 0x30,
	0x0c, 0x40, 0x15, 0x54, 0x46, 0xf1, 0x63, 0xf4, 0x18, 0xe2, 0xd6, 0x0f, 0x6f, 0x20, 0xbb, 0x81,
	0x27, 0xe8, 0xb2, 0xd0, 0x95, 0x9a, 0xa4, 0x64, 0xd1, 0xf3, 0x0b, 0x05, 0xf7, 0xf7, 0x78, 0xf0,
	0x9e, 0x9b, 0x52, 0xb6, 0xf5, 0x33, 0x8f, 0x20, 0x14, 0xde, 0xc2, 0x69, 0x15, 0xe5, 0x1d, 0x86,
	0xcc, 0x9b, 0x45, 0x06, 0x1c, 0x28, 0x72, 0x4c, 0xa8, 0xa1, 0x96, 0x14, 0x32, 0x69, 0x85, 0x00,
	0x42, 0x24, 0x3c, 0x56, 0x15, 0x13, 0x7f, 0x6e, 0xee, 0x76, 0x77, 0xfe, 0x19, 0xa1, 0x20, 0x2f,
	0x2f, 0x13, 0xc5, 0x49, 0xf3, 0x17, 0xd5, 0x5f, 0x5d, 0xf7, 0x1f, 0xf6, 0x07, 0xdf, 0xb9, 0xd3,
	0x56, 0x97, 0xd2, 0x1f, 0xe7, 0x4b, 0x6b, 0x1f, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xb9,
	0x38, 0x19, 0x83, 0x00, 0x00, 0x00,
}
