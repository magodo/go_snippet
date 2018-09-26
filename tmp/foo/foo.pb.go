// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

/*
Package foo is a generated protocol buffer package.

It is generated from these files:
	foo.proto

It has these top-level messages:
	Foo
*/
package foo

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

type Foo struct {
	I                *int32 `protobuf:"varint,1,req,name=i" json:"i,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Foo) Reset()                    { *m = Foo{} }
func (m *Foo) String() string            { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()               {}
func (*Foo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Foo) GetI() int32 {
	if m != nil && m.I != nil {
		return *m.I
	}
	return 0
}

func init() {
	proto.RegisterType((*Foo)(nil), "foo.Foo")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 61 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcb, 0xcf, 0x57, 0x12, 0xe6, 0x62, 0x76, 0xcb,
	0xcf, 0x17, 0xe2, 0xe1, 0x62, 0xcc, 0x94, 0x60, 0x54, 0x60, 0xd2, 0x60, 0x0d, 0x62, 0xcc, 0x04,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x8c, 0xbc, 0xf0, 0x25, 0x00, 0x00, 0x00,
}
