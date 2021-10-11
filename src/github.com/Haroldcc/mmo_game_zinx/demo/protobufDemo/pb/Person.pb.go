// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

package pb

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

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}

func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0, 0}
}

type Person struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email                string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones               []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
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
	Number               string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type                 Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=pb.Person_PhoneType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Person_PhoneNumber) Reset()         { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()    {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0, 0}
}

func (m *Person_PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person_PhoneNumber.Unmarshal(m, b)
}
func (m *Person_PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *Person_PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person_PhoneNumber.Merge(m, src)
}
func (m *Person_PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_Person_PhoneNumber.Size(m)
}
func (m *Person_PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_Person_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_Person_PhoneNumber proto.InternalMessageInfo

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

// Our address book file is just one of these.
type AddressBook struct {
	People               []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AddressBook) Reset()         { *m = AddressBook{} }
func (m *AddressBook) String() string { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()    {}
func (*AddressBook) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *AddressBook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressBook.Unmarshal(m, b)
}
func (m *AddressBook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressBook.Marshal(b, m, deterministic)
}
func (m *AddressBook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressBook.Merge(m, src)
}
func (m *AddressBook) XXX_Size() int {
	return xxx_messageInfo_AddressBook.Size(m)
}
func (m *AddressBook) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressBook.DiscardUnknown(m)
}

var xxx_messageInfo_AddressBook proto.InternalMessageInfo

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterEnum("pb.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "pb.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "pb.AddressBook")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x4d, 0xd6, 0x05, 0xf7, 0x54, 0x46, 0x79, 0x19, 0xa3, 0x78, 0x2a, 0x3d, 0x15, 0x84,
	0x82, 0xf3, 0x13, 0x38, 0x18, 0x28, 0x3a, 0x3b, 0x82, 0xe0, 0x79, 0xa5, 0x2f, 0x58, 0x5c, 0x9b,
	0xd0, 0xce, 0xc3, 0xbe, 0xb5, 0x1f, 0x41, 0x92, 0x05, 0x15, 0x6f, 0xcf, 0x3f, 0xf2, 0x4b, 0x82,
	0xab, 0x1d, 0x0f, 0xa3, 0xe9, 0x4b, 0x3b, 0x98, 0xa3, 0x21, 0x69, 0xeb, 0xfc, 0x4b, 0x40, 0x9d,
	0x43, 0x22, 0x44, 0xfd, 0xbe, 0xe3, 0x54, 0x64, 0xa2, 0x98, 0x69, 0xaf, 0x69, 0x0e, 0xd9, 0x36,
	0xa9, 0xcc, 0x44, 0x31, 0xd5, 0xb2, 0x6d, 0x68, 0x81, 0x29, 0x77, 0xfb, 0xf6, 0x90, 0x4e, 0xfc,
	0xe8, 0x6c, 0xa8, 0x84, 0xb2, 0xef, 0xa6, 0xe7, 0x31, 0x8d, 0xb2, 0x49, 0x11, 0xaf, 0x96, 0xa5,
	0xad, 0xcb, 0x80, 0xda, 0xb9, 0xe2, 0xe5, 0xb3, 0xab, 0x79, 0xd0, 0x61, 0x75, 0x5d, 0x21, 0xfe,
	0x13, 0xd3, 0x12, 0xaa, 0xf7, 0x2a, 0xa0, 0x83, 0xa3, 0x02, 0xd1, 0xf1, 0x64, 0xd9, 0xe3, 0xe7,
	0xab, 0xc5, 0xff, 0x43, 0x5f, 0x4f, 0x96, 0xb5, 0x5f, 0xe4, 0x37, 0x98, 0xfd, 0x44, 0x04, 0xa8,
	0x6d, 0xb5, 0x7e, 0x7c, 0xde, 0x24, 0x17, 0x74, 0x89, 0xe8, 0xa1, 0xda, 0x6e, 0x12, 0xe1, 0xd4,
	0x5b, 0xa5, 0x9f, 0x12, 0x99, 0xdf, 0x22, 0xbe, 0x6f, 0x9a, 0x81, 0xc7, 0x71, 0x6d, 0xcc, 0x07,
	0xe5, 0x50, 0x96, 0x8d, 0x3d, 0xb8, 0x87, 0xbb, 0xcb, 0xe3, 0x97, 0xa3, 0x43, 0x53, 0x2b, 0xff,
	0x61, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0x6e, 0xda, 0xc3, 0x40, 0x01, 0x00, 0x00,
}