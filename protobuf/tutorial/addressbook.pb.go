// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package tutorial is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Person
	AddressBook
*/
package tutorial

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

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x Person_PhoneType) Enum() *Person_PhoneType {
	p := new(Person_PhoneType)
	*p = x
	return p
}
func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (x *Person_PhoneType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Person_PhoneType_value, data, "Person_PhoneType")
	if err != nil {
		return err
	}
	*x = Person_PhoneType(value)
	return nil
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Person struct {
	Name             *string               `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Id               *int32                `protobuf:"varint,2,req,name=id" json:"id,omitempty"`
	Email            *string               `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones           []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Person) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number           *string           `protobuf:"bytes,1,req,name=number" json:"number,omitempty"`
	Type             *Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=tutorial.Person_PhoneType,def=1" json:"type,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

const Default_Person_PhoneNumber_Type Person_PhoneType = Person_HOME

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_Person_PhoneNumber_Type
}

type AddressBook struct {
	People           []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "tutorial.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "tutorial.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "tutorial.AddressBook")
	proto.RegisterEnum("tutorial.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0xdb, 0x74, 0x69, 0x26, 0x50, 0xe2, 0x20, 0xb2, 0x14, 0x0f, 0x4b, 0x4e, 0x0b, 0x42,
	0x0e, 0xa5, 0x20, 0x78, 0xb3, 0x50, 0x50, 0xb4, 0xa6, 0x2c, 0x82, 0x07, 0x4f, 0x29, 0x59, 0x30,
	0x34, 0xc9, 0x2c, 0x9b, 0xf4, 0xd0, 0xab, 0xbf, 0x5c, 0xf2, 0xa1, 0x88, 0xf4, 0xf6, 0xde, 0xbc,
	0xc7, 0x9b, 0x99, 0x07, 0x97, 0x59, 0x9e, 0x3b, 0xd3, 0x34, 0x7b, 0xa2, 0x43, 0x62, 0x1d, 0xb5,
	0x84, 0xb3, 0xf6, 0xd8, 0x92, 0x2b, 0xb2, 0x32, 0xfe, 0x62, 0xc0, 0x77, 0xc6, 0x35, 0x54, 0x23,
	0x82, 0x5f, 0x67, 0x95, 0x11, 0x9e, 0x64, 0x2a, 0xd0, 0x3d, 0xc6, 0x39, 0xb0, 0x22, 0x17, 0x4c,
	0x32, 0x35, 0xd5, 0xac, 0xc8, 0xf1, 0x0a, 0xa6, 0xa6, 0xca, 0x8a, 0x52, 0x4c, 0xa4, 0xa7, 0x02,
	0x3d, 0x10, 0x5c, 0x01, 0xb7, 0x9f, 0x54, 0x9b, 0x46, 0xf8, 0x72, 0xa2, 0xc2, 0xe5, 0x4d, 0xf2,
	0x93, 0x9f, 0x0c, 0xd9, 0xc9, 0xae, 0x93, 0x5f, 0x8f, 0xd5, 0xde, 0x38, 0x3d, 0x7a, 0x17, 0x1f,
	0x10, 0xfe, 0x19, 0xe3, 0x35, 0xf0, 0xba, 0x47, 0xe3, 0x01, 0x23, 0xc3, 0x15, 0xf8, 0xed, 0xc9,
	0x1a, 0xc1, 0xa4, 0xa7, 0xe6, 0xcb, 0xc5, 0xf9, 0xe8, 0xb7, 0x93, 0x35, 0xf7, 0xfe, 0x63, 0xba,
	0xdd, 0xe8, 0xde, 0x1d, 0xdf, 0x42, 0xf0, 0x2b, 0x20, 0x00, 0xdf, 0xa6, 0xeb, 0xa7, 0x97, 0x4d,
	0x74, 0x81, 0x33, 0xe8, 0x6d, 0x91, 0xd7, 0xa1, 0xf7, 0x54, 0x3f, 0x47, 0x2c, 0xbe, 0x83, 0xf0,
	0x61, 0xe8, 0x68, 0x4d, 0x74, 0x40, 0x05, 0xdc, 0x1a, 0xb2, 0x65, 0x57, 0x45, 0xf7, 0x4e, 0xf4,
	0x7f, 0xa7, 0x1e, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x98, 0x8a, 0x28, 0x5b, 0x01,
	0x00, 0x00,
}
