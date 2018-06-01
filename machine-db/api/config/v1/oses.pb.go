// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/config/v1/oses.proto

package config

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

// An operating system to store in the database.
type OS struct {
	// The name of this operating system. Must be unique.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this operating system.
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OS) Reset()         { *m = OS{} }
func (m *OS) String() string { return proto.CompactTextString(m) }
func (*OS) ProtoMessage()    {}
func (*OS) Descriptor() ([]byte, []int) {
	return fileDescriptor_oses_f5f81609645a8d7a, []int{0}
}
func (m *OS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OS.Unmarshal(m, b)
}
func (m *OS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OS.Marshal(b, m, deterministic)
}
func (dst *OS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OS.Merge(dst, src)
}
func (m *OS) XXX_Size() int {
	return xxx_messageInfo_OS.Size(m)
}
func (m *OS) XXX_DiscardUnknown() {
	xxx_messageInfo_OS.DiscardUnknown(m)
}

var xxx_messageInfo_OS proto.InternalMessageInfo

func (m *OS) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OS) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A list of operating systems.
type OSes struct {
	// A list of operating systems.
	OperatingSystem      []*OS    `protobuf:"bytes,1,rep,name=operating_system,json=operatingSystem" json:"operating_system,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OSes) Reset()         { *m = OSes{} }
func (m *OSes) String() string { return proto.CompactTextString(m) }
func (*OSes) ProtoMessage()    {}
func (*OSes) Descriptor() ([]byte, []int) {
	return fileDescriptor_oses_f5f81609645a8d7a, []int{1}
}
func (m *OSes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OSes.Unmarshal(m, b)
}
func (m *OSes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OSes.Marshal(b, m, deterministic)
}
func (dst *OSes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OSes.Merge(dst, src)
}
func (m *OSes) XXX_Size() int {
	return xxx_messageInfo_OSes.Size(m)
}
func (m *OSes) XXX_DiscardUnknown() {
	xxx_messageInfo_OSes.DiscardUnknown(m)
}

var xxx_messageInfo_OSes proto.InternalMessageInfo

func (m *OSes) GetOperatingSystem() []*OS {
	if m != nil {
		return m.OperatingSystem
	}
	return nil
}

func init() {
	proto.RegisterType((*OS)(nil), "config.OS")
	proto.RegisterType((*OSes)(nil), "config.OSes")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/config/v1/oses.proto", fileDescriptor_oses_f5f81609645a8d7a)
}

var fileDescriptor_oses_f5f81609645a8d7a = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x8a, 0x83, 0x30,
	0x18, 0x84, 0xd1, 0x15, 0x61, 0xe3, 0x61, 0x97, 0x9c, 0x3c, 0x8a, 0x27, 0x2f, 0x9b, 0xb0, 0xbb,
	0x14, 0x4a, 0xa1, 0xcf, 0x20, 0x98, 0x07, 0x28, 0x31, 0xa6, 0xf1, 0x87, 0x26, 0x7f, 0x48, 0x62,
	0xa1, 0x6f, 0x5f, 0x88, 0x50, 0x7a, 0x1b, 0xbe, 0xf9, 0x60, 0x86, 0x1c, 0x0d, 0x32, 0xb5, 0x06,
	0xb4, 0xb0, 0x59, 0x86, 0xc1, 0xf0, 0xdb, 0xa6, 0x80, 0x5b, 0xa9, 0x56, 0x70, 0xfa, 0x67, 0x99,
	0xb9, 0xf4, 0xc0, 0x15, 0xba, 0x2b, 0x18, 0x7e, 0xff, 0xe5, 0x18, 0x75, 0x64, 0x3e, 0x60, 0x42,
	0x5a, 0xef, 0xb4, 0x3f, 0x91, 0x72, 0x14, 0x94, 0x92, 0xca, 0x49, 0xab, 0xdb, 0xa2, 0x2b, 0x86,
	0xcf, 0x29, 0x67, 0xda, 0x91, 0x66, 0xd1, 0x51, 0x05, 0xf0, 0x09, 0xd0, 0xb5, 0x65, 0xae, 0xde,
	0x51, 0x7f, 0x26, 0xd5, 0x28, 0x74, 0xa4, 0x07, 0xf2, 0x8d, 0x5e, 0x07, 0x99, 0xc0, 0x99, 0x4b,
	0x7c, 0xc4, 0xa4, 0x6d, 0x5b, 0x74, 0x1f, 0x43, 0xf3, 0x47, 0xd8, 0x3e, 0xc3, 0x46, 0x31, 0x7d,
	0xbd, 0x1c, 0x91, 0x95, 0xb9, 0xce, 0x4f, 0xfe, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x10,
	0xb8, 0xf3, 0xc5, 0x00, 0x00, 0x00,
}
