// Code generated by protoc-gen-go. DO NOT EDIT.
// source: M2C.proto

package message

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

type M2C_EnterLobby struct {
	IsInBattle           bool     `protobuf:"varint,1,opt,name=isInBattle,proto3" json:"isInBattle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_EnterLobby) Reset()         { *m = M2C_EnterLobby{} }
func (m *M2C_EnterLobby) String() string { return proto.CompactTextString(m) }
func (*M2C_EnterLobby) ProtoMessage()    {}
func (*M2C_EnterLobby) Descriptor() ([]byte, []int) {
	return fileDescriptor_M2C_ec78165a0c0b8568, []int{0}
}
func (m *M2C_EnterLobby) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_EnterLobby.Unmarshal(m, b)
}
func (m *M2C_EnterLobby) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_EnterLobby.Marshal(b, m, deterministic)
}
func (dst *M2C_EnterLobby) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_EnterLobby.Merge(dst, src)
}
func (m *M2C_EnterLobby) XXX_Size() int {
	return xxx_messageInfo_M2C_EnterLobby.Size(m)
}
func (m *M2C_EnterLobby) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_EnterLobby.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_EnterLobby proto.InternalMessageInfo

func (m *M2C_EnterLobby) GetIsInBattle() bool {
	if m != nil {
		return m.IsInBattle
	}
	return false
}

func init() {
	proto.RegisterType((*M2C_EnterLobby)(nil), "message.M2C_EnterLobby")
}

func init() { proto.RegisterFile("M2C.proto", fileDescriptor_M2C_ec78165a0c0b8568) }

var fileDescriptor_M2C_ec78165a0c0b8568 = []byte{
	// 90 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xf4, 0x35, 0x72, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0x32,
	0xe0, 0xe2, 0xf3, 0x35, 0x72, 0x8e, 0x77, 0xcd, 0x2b, 0x49, 0x2d, 0xf2, 0xc9, 0x4f, 0x4a, 0xaa,
	0x14, 0x92, 0xe3, 0xe2, 0xca, 0x2c, 0xf6, 0xcc, 0x73, 0x4a, 0x2c, 0x29, 0xc9, 0x49, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x08, 0x42, 0x12, 0x49, 0x62, 0x03, 0x9b, 0x60, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0xd6, 0xba, 0x9d, 0x4e, 0x00, 0x00, 0x00,
}
