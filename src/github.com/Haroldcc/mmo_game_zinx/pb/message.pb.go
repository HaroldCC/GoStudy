// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

// 同步玩家ID
type SyncPlayerID struct {
	PlayerID             int32    `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPlayerID) Reset()         { *m = SyncPlayerID{} }
func (m *SyncPlayerID) String() string { return proto.CompactTextString(m) }
func (*SyncPlayerID) ProtoMessage()    {}
func (*SyncPlayerID) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *SyncPlayerID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPlayerID.Unmarshal(m, b)
}
func (m *SyncPlayerID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPlayerID.Marshal(b, m, deterministic)
}
func (m *SyncPlayerID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPlayerID.Merge(m, src)
}
func (m *SyncPlayerID) XXX_Size() int {
	return xxx_messageInfo_SyncPlayerID.Size(m)
}
func (m *SyncPlayerID) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPlayerID.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPlayerID proto.InternalMessageInfo

func (m *SyncPlayerID) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

// 位置信息
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=v,proto3" json:"v,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

// 广播消息
type BroadCast struct {
	PlayerID int32 `protobuf:"varint,1,opt,name=playerID,proto3" json:"playerID,omitempty"`
	Tp       int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_Pos
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPlayerID() int32 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_Pos struct {
	Pos *Position `protobuf:"bytes,4,opt,name=Pos,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_Pos) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetPos() *Position {
	if x, ok := m.GetData().(*BroadCast_Pos); ok {
		return x.Pos
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BroadCast) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_Pos)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

func init() {
	proto.RegisterType((*SyncPlayerID)(nil), "pb.SyncPlayerID")
	proto.RegisterType((*Position)(nil), "pb.Position")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0xb7, 0x2e, 0xed, 0x11, 0x18, 0x3c, 0x59, 0x9d, 0xa2, 0x4c, 0x88, 0x21, 0x03,
	0x3c, 0x01, 0x69, 0x87, 0xb0, 0x59, 0xa6, 0x2f, 0x60, 0x17, 0x0b, 0x55, 0x02, 0xdb, 0x8a, 0xad,
	0x28, 0xce, 0xcb, 0xb0, 0xf3, 0x94, 0xc8, 0x0e, 0x41, 0x4c, 0x4c, 0x77, 0xdf, 0x7f, 0xa7, 0xff,
	0x3f, 0x1d, 0xdc, 0x7c, 0x68, 0xef, 0xe5, 0x9b, 0x6e, 0x5c, 0x6f, 0x83, 0xa5, 0xd8, 0xa9, 0xfa,
	0x1e, 0xca, 0x97, 0x68, 0xce, 0xfc, 0x5d, 0x46, 0xdd, 0x3f, 0x1f, 0xe9, 0x1e, 0xb6, 0xee, 0xa7,
	0x67, 0xa8, 0x42, 0x77, 0x44, 0xfc, 0x72, 0xdd, 0xc2, 0x96, 0x5b, 0x7f, 0x09, 0x17, 0x6b, 0x68,
	0x09, 0x68, 0xcc, 0x0b, 0x58, 0xa0, 0x31, 0x51, 0x64, 0x78, 0xa6, 0x98, 0x68, 0x62, 0xab, 0x99,
	0xa6, 0x44, 0x03, 0x5b, 0xcf, 0x34, 0xd4, 0x9f, 0x08, 0x76, 0x6d, 0x6f, 0xe5, 0xeb, 0x41, 0xfa,
	0xf0, 0x5f, 0x1a, 0xbd, 0x05, 0x7c, 0x72, 0xd9, 0x94, 0x08, 0x7c, 0x72, 0x74, 0x0f, 0x57, 0x07,
	0x6b, 0x82, 0x36, 0x21, 0x7b, 0xef, 0xba, 0x42, 0x2c, 0x02, 0xad, 0x60, 0xc5, 0xad, 0xcf, 0x29,
	0xd7, 0x0f, 0x65, 0xe3, 0x54, 0xb3, 0x1c, 0xda, 0x15, 0x22, 0x8d, 0x68, 0x05, 0xf0, 0x74, 0x4e,
	0xc2, 0x51, 0x06, 0xc9, 0x48, 0x72, 0xed, 0x0a, 0xf1, 0x47, 0x6b, 0x37, 0xb0, 0xce, 0x95, 0x7c,
	0x61, 0xcc, 0x95, 0xda, 0xe4, 0x1f, 0x3d, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x17, 0x13, 0xd4,
	0x64, 0x34, 0x01, 0x00, 0x00,
}