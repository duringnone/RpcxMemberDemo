// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member.proto

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

type Args struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Args) Reset()         { *m = Args{} }
func (m *Args) String() string { return proto.CompactTextString(m) }
func (*Args) ProtoMessage()    {}
func (*Args) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{0}
}

func (m *Args) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Args.Unmarshal(m, b)
}
func (m *Args) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Args.Marshal(b, m, deterministic)
}
func (m *Args) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Args.Merge(m, src)
}
func (m *Args) XXX_Size() int {
	return xxx_messageInfo_Args.Size(m)
}
func (m *Args) XXX_DiscardUnknown() {
	xxx_messageInfo_Args.DiscardUnknown(m)
}

var xxx_messageInfo_Args proto.InternalMessageInfo

func (m *Args) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Reply struct {
	UID                  int64    `protobuf:"varint,1,opt,name=UID,proto3" json:"UID,omitempty"`
	AddTime              int64    `protobuf:"varint,2,opt,name=AddTime,proto3" json:"AddTime,omitempty"`
	UserType             int32    `protobuf:"varint,3,opt,name=UserType,proto3" json:"UserType,omitempty"`
	Uface                string   `protobuf:"bytes,4,opt,name=Uface,proto3" json:"Uface,omitempty"`
	UserName             string   `protobuf:"bytes,5,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{1}
}

func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (m *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(m, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetUID() int64 {
	if m != nil {
		return m.UID
	}
	return 0
}

func (m *Reply) GetAddTime() int64 {
	if m != nil {
		return m.AddTime
	}
	return 0
}

func (m *Reply) GetUserType() int32 {
	if m != nil {
		return m.UserType
	}
	return 0
}

func (m *Reply) GetUface() string {
	if m != nil {
		return m.Uface
	}
	return ""
}

func (m *Reply) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type ProtoArgs struct {
	A                    int32    `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoArgs) Reset()         { *m = ProtoArgs{} }
func (m *ProtoArgs) String() string { return proto.CompactTextString(m) }
func (*ProtoArgs) ProtoMessage()    {}
func (*ProtoArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{2}
}

func (m *ProtoArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoArgs.Unmarshal(m, b)
}
func (m *ProtoArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoArgs.Marshal(b, m, deterministic)
}
func (m *ProtoArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoArgs.Merge(m, src)
}
func (m *ProtoArgs) XXX_Size() int {
	return xxx_messageInfo_ProtoArgs.Size(m)
}
func (m *ProtoArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoArgs.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoArgs proto.InternalMessageInfo

func (m *ProtoArgs) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *ProtoArgs) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type ProtoReply struct {
	C                    int32    `protobuf:"varint,1,opt,name=C,proto3" json:"C,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoReply) Reset()         { *m = ProtoReply{} }
func (m *ProtoReply) String() string { return proto.CompactTextString(m) }
func (*ProtoReply) ProtoMessage()    {}
func (*ProtoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{3}
}

func (m *ProtoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoReply.Unmarshal(m, b)
}
func (m *ProtoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoReply.Marshal(b, m, deterministic)
}
func (m *ProtoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoReply.Merge(m, src)
}
func (m *ProtoReply) XXX_Size() int {
	return xxx_messageInfo_ProtoReply.Size(m)
}
func (m *ProtoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoReply.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoReply proto.InternalMessageInfo

func (m *ProtoReply) GetC() int32 {
	if m != nil {
		return m.C
	}
	return 0
}

func init() {
	proto.RegisterType((*Args)(nil), "pb.Args")
	proto.RegisterType((*Reply)(nil), "pb.Reply")
	proto.RegisterType((*ProtoArgs)(nil), "pb.ProtoArgs")
	proto.RegisterType((*ProtoReply)(nil), "pb.ProtoReply")
}

func init() { proto.RegisterFile("member.proto", fileDescriptor_9b9836b7e13de206) }

var fileDescriptor_9b9836b7e13de206 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x6a, 0x84, 0x30,
	0x10, 0x86, 0x99, 0x68, 0xda, 0x3a, 0x48, 0x29, 0xa1, 0x94, 0xe0, 0x49, 0x72, 0xa9, 0xa7, 0x5e,
	0xfa, 0x04, 0xd1, 0x5e, 0xbc, 0x94, 0x12, 0xf4, 0x01, 0xb4, 0x99, 0x96, 0x85, 0x0d, 0x86, 0xb8,
	0x17, 0x8f, 0xfb, 0xe6, 0x8b, 0x71, 0xf5, 0x96, 0x2f, 0xff, 0x3f, 0xcc, 0x37, 0x98, 0x3b, 0x72,
	0x23, 0x85, 0x0f, 0x1f, 0xa6, 0xcb, 0x24, 0x98, 0x1f, 0xd5, 0x1b, 0xa6, 0x3a, 0xfc, 0xcf, 0xe2,
	0x19, 0x59, 0x6b, 0x25, 0x94, 0x50, 0x25, 0x86, 0xb5, 0x56, 0x5d, 0x01, 0xb9, 0x21, 0x7f, 0x5e,
	0xc4, 0x0b, 0x26, 0x7d, 0xfb, 0x75, 0x8f, 0xd6, 0xa7, 0x90, 0xf8, 0xa8, 0xad, 0xed, 0x4e, 0x8e,
	0x24, 0x8b, 0xbf, 0x3b, 0x8a, 0x02, 0x9f, 0xfa, 0x99, 0x42, 0xb7, 0x78, 0x92, 0x49, 0x09, 0x15,
	0x37, 0x07, 0x8b, 0x57, 0xe4, 0xfd, 0xdf, 0xf0, 0x4b, 0x32, 0x2d, 0xa1, 0xca, 0xcc, 0x06, 0xfb,
	0xc4, 0xf7, 0xe0, 0x48, 0xf2, 0x18, 0x1c, 0xac, 0xde, 0x31, 0xfb, 0x59, 0x45, 0xa3, 0x60, 0x8e,
	0xa0, 0xa3, 0x04, 0x37, 0xa0, 0x57, 0xaa, 0xe3, 0x72, 0x6e, 0xa0, 0x56, 0x05, 0x62, 0x2c, 0x6e,
	0xc2, 0x39, 0x42, 0xb3, 0x37, 0x9b, 0xf1, 0x21, 0xde, 0xfa, 0x79, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0x4b, 0xeb, 0xb2, 0xfb, 0x00, 0x00, 0x00,
}
