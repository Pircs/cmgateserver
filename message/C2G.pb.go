// Code generated by protoc-gen-go. DO NOT EDIT.
// source: C2G.proto

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

// 玩家登录游戏服务器
type C2G_UserLogin struct {
	GameServerId         int32    `protobuf:"varint,1,opt,name=gameServerId,proto3" json:"gameServerId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2G_UserLogin) Reset()         { *m = C2G_UserLogin{} }
func (m *C2G_UserLogin) String() string { return proto.CompactTextString(m) }
func (*C2G_UserLogin) ProtoMessage()    {}
func (*C2G_UserLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2G_cfab3fffa04a3a1a, []int{0}
}
func (m *C2G_UserLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2G_UserLogin.Unmarshal(m, b)
}
func (m *C2G_UserLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2G_UserLogin.Marshal(b, m, deterministic)
}
func (dst *C2G_UserLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2G_UserLogin.Merge(dst, src)
}
func (m *C2G_UserLogin) XXX_Size() int {
	return xxx_messageInfo_C2G_UserLogin.Size(m)
}
func (m *C2G_UserLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_C2G_UserLogin.DiscardUnknown(m)
}

var xxx_messageInfo_C2G_UserLogin proto.InternalMessageInfo

func (m *C2G_UserLogin) GetGameServerId() int32 {
	if m != nil {
		return m.GameServerId
	}
	return 0
}

func (m *C2G_UserLogin) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *C2G_UserLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 玩家选择该角色登录
type C2G_SelectCharacter struct {
	RoleId               int64    `protobuf:"varint,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2G_SelectCharacter) Reset()         { *m = C2G_SelectCharacter{} }
func (m *C2G_SelectCharacter) String() string { return proto.CompactTextString(m) }
func (*C2G_SelectCharacter) ProtoMessage()    {}
func (*C2G_SelectCharacter) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2G_cfab3fffa04a3a1a, []int{1}
}
func (m *C2G_SelectCharacter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2G_SelectCharacter.Unmarshal(m, b)
}
func (m *C2G_SelectCharacter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2G_SelectCharacter.Marshal(b, m, deterministic)
}
func (dst *C2G_SelectCharacter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2G_SelectCharacter.Merge(dst, src)
}
func (m *C2G_SelectCharacter) XXX_Size() int {
	return xxx_messageInfo_C2G_SelectCharacter.Size(m)
}
func (m *C2G_SelectCharacter) XXX_DiscardUnknown() {
	xxx_messageInfo_C2G_SelectCharacter.DiscardUnknown(m)
}

var xxx_messageInfo_C2G_SelectCharacter proto.InternalMessageInfo

func (m *C2G_SelectCharacter) GetRoleId() int64 {
	if m != nil {
		return m.RoleId
	}
	return 0
}

// 玩家修改昵称
type C2G_ChangeNickName struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2G_ChangeNickName) Reset()         { *m = C2G_ChangeNickName{} }
func (m *C2G_ChangeNickName) String() string { return proto.CompactTextString(m) }
func (*C2G_ChangeNickName) ProtoMessage()    {}
func (*C2G_ChangeNickName) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2G_cfab3fffa04a3a1a, []int{2}
}
func (m *C2G_ChangeNickName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2G_ChangeNickName.Unmarshal(m, b)
}
func (m *C2G_ChangeNickName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2G_ChangeNickName.Marshal(b, m, deterministic)
}
func (dst *C2G_ChangeNickName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2G_ChangeNickName.Merge(dst, src)
}
func (m *C2G_ChangeNickName) XXX_Size() int {
	return xxx_messageInfo_C2G_ChangeNickName.Size(m)
}
func (m *C2G_ChangeNickName) XXX_DiscardUnknown() {
	xxx_messageInfo_C2G_ChangeNickName.DiscardUnknown(m)
}

var xxx_messageInfo_C2G_ChangeNickName proto.InternalMessageInfo

func (m *C2G_ChangeNickName) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

// 玩家修改头像id
type C2G_ChangeAvatarIcon struct {
	AvatarId             int32    `protobuf:"varint,1,opt,name=avatarId,proto3" json:"avatarId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2G_ChangeAvatarIcon) Reset()         { *m = C2G_ChangeAvatarIcon{} }
func (m *C2G_ChangeAvatarIcon) String() string { return proto.CompactTextString(m) }
func (*C2G_ChangeAvatarIcon) ProtoMessage()    {}
func (*C2G_ChangeAvatarIcon) Descriptor() ([]byte, []int) {
	return fileDescriptor_C2G_cfab3fffa04a3a1a, []int{3}
}
func (m *C2G_ChangeAvatarIcon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2G_ChangeAvatarIcon.Unmarshal(m, b)
}
func (m *C2G_ChangeAvatarIcon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2G_ChangeAvatarIcon.Marshal(b, m, deterministic)
}
func (dst *C2G_ChangeAvatarIcon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2G_ChangeAvatarIcon.Merge(dst, src)
}
func (m *C2G_ChangeAvatarIcon) XXX_Size() int {
	return xxx_messageInfo_C2G_ChangeAvatarIcon.Size(m)
}
func (m *C2G_ChangeAvatarIcon) XXX_DiscardUnknown() {
	xxx_messageInfo_C2G_ChangeAvatarIcon.DiscardUnknown(m)
}

var xxx_messageInfo_C2G_ChangeAvatarIcon proto.InternalMessageInfo

func (m *C2G_ChangeAvatarIcon) GetAvatarId() int32 {
	if m != nil {
		return m.AvatarId
	}
	return 0
}

func init() {
	proto.RegisterType((*C2G_UserLogin)(nil), "message.C2G_UserLogin")
	proto.RegisterType((*C2G_SelectCharacter)(nil), "message.C2G_SelectCharacter")
	proto.RegisterType((*C2G_ChangeNickName)(nil), "message.C2G_ChangeNickName")
	proto.RegisterType((*C2G_ChangeAvatarIcon)(nil), "message.C2G_ChangeAvatarIcon")
}

func init() { proto.RegisterFile("C2G.proto", fileDescriptor_C2G_cfab3fffa04a3a1a) }

var fileDescriptor_C2G_cfab3fffa04a3a1a = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xa9, 0x8b, 0xab, 0x1d, 0xf4, 0x12, 0x45, 0x8a, 0xa7, 0x25, 0xa7, 0xbd, 0x28, 0x52,
	0x3f, 0x81, 0xe4, 0xb0, 0x2c, 0xc8, 0x1e, 0xb2, 0x78, 0x96, 0x31, 0x3b, 0x64, 0x4b, 0xdb, 0xa4,
	0x4c, 0x62, 0xfd, 0xfa, 0x92, 0xfe, 0x13, 0x6f, 0xf3, 0x9b, 0xf7, 0x1e, 0x8f, 0x19, 0xc8, 0x55,
	0xb9, 0x7b, 0xee, 0xd8, 0x47, 0x2f, 0xae, 0x5a, 0x0a, 0x01, 0x2d, 0xc9, 0x1a, 0x6e, 0x55, 0xb9,
	0xfb, 0xfc, 0x08, 0xc4, 0xef, 0xde, 0x56, 0x4e, 0x48, 0xb8, 0xb1, 0xd8, 0xd2, 0x91, 0xb8, 0x27,
	0xde, 0x9f, 0x8a, 0x6c, 0x93, 0x6d, 0x2f, 0xf5, 0xbf, 0x9d, 0x78, 0x84, 0xeb, 0xef, 0x40, 0x7c,
	0xc0, 0x96, 0x8a, 0x8b, 0x4d, 0xb6, 0xcd, 0xf5, 0xc2, 0x49, 0xeb, 0x30, 0x84, 0x1f, 0xcf, 0xa7,
	0x62, 0x35, 0x6a, 0x33, 0xcb, 0x27, 0xb8, 0x4b, 0x65, 0x47, 0x6a, 0xc8, 0x44, 0x75, 0x46, 0x46,
	0x13, 0x89, 0xc5, 0x03, 0xac, 0xd9, 0x37, 0x34, 0x95, 0xad, 0xf4, 0x44, 0xf2, 0x05, 0x44, 0xb2,
	0xab, 0x33, 0x3a, 0x4b, 0x87, 0xca, 0xd4, 0x73, 0x81, 0x9b, 0xe6, 0xc1, 0x9f, 0xeb, 0x85, 0x65,
	0x09, 0xf7, 0x7f, 0x89, 0xb7, 0x1e, 0x23, 0xf2, 0xde, 0x78, 0x97, 0x32, 0x38, 0xd2, 0x7c, 0xd0,
	0xc2, 0x5f, 0xeb, 0xe1, 0x23, 0xaf, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xf7, 0xc6, 0x14,
	0x1e, 0x01, 0x00, 0x00,
}
