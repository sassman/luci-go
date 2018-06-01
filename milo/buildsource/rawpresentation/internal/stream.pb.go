// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/buildsource/rawpresentation/internal/stream.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import milo "go.chromium.org/luci/common/proto/milo"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Item is a full annotation protobuf state item. It is used to serialize and
// deserialize Step and associated fetch data into memcache.
type CachedStep struct {
	// Step is the root annotation step.
	Step *milo.Step `protobuf:"bytes,1,opt,name=step" json:"step,omitempty"`
	// Finished is true if this is the last annotation protobuf in the stream.
	Finished             bool     `protobuf:"varint,2,opt,name=finished" json:"finished,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CachedStep) Reset()         { *m = CachedStep{} }
func (m *CachedStep) String() string { return proto.CompactTextString(m) }
func (*CachedStep) ProtoMessage()    {}
func (*CachedStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_stream_48dd3ce49f63ae00, []int{0}
}
func (m *CachedStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CachedStep.Unmarshal(m, b)
}
func (m *CachedStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CachedStep.Marshal(b, m, deterministic)
}
func (dst *CachedStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CachedStep.Merge(dst, src)
}
func (m *CachedStep) XXX_Size() int {
	return xxx_messageInfo_CachedStep.Size(m)
}
func (m *CachedStep) XXX_DiscardUnknown() {
	xxx_messageInfo_CachedStep.DiscardUnknown(m)
}

var xxx_messageInfo_CachedStep proto.InternalMessageInfo

func (m *CachedStep) GetStep() *milo.Step {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *CachedStep) GetFinished() bool {
	if m != nil {
		return m.Finished
	}
	return false
}

func init() {
	proto.RegisterType((*CachedStep)(nil), "internal.CachedStep")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/milo/buildsource/rawpresentation/internal/stream.proto", fileDescriptor_stream_48dd3ce49f63ae00)
}

var fileDescriptor_stream_48dd3ce49f63ae00 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcd, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0x5c, 0x4a, 0x31, 0xea, 0xe6, 0xc9, 0x78, 0x28, 0xa6, 0x93, 0x27, 0x09, 0xda,
	0xa5, 0x7b, 0x97, 0x40, 0x36, 0xe7, 0x09, 0x64, 0xf9, 0x62, 0x1f, 0x48, 0x77, 0xe2, 0x24, 0x91,
	0xd7, 0x0f, 0xb6, 0x49, 0xa6, 0x8c, 0xf7, 0xdf, 0xdd, 0xf7, 0xab, 0xf3, 0xc2, 0xda, 0xad, 0xc2,
	0x01, 0x4b, 0xd0, 0x2c, 0x8b, 0xf1, 0xc5, 0xa1, 0x09, 0xe8, 0xd9, 0x4c, 0x05, 0xfd, 0x9c, 0xb8,
	0x88, 0x03, 0x23, 0xf6, 0x16, 0x05, 0x12, 0x50, 0xb6, 0x19, 0x99, 0x0c, 0x52, 0x06, 0x21, 0xeb,
	0x4d, 0xca, 0x02, 0x36, 0xe8, 0x28, 0x9c, 0xb9, 0xa9, 0x1f, 0x71, 0xf7, 0xf7, 0x92, 0x75, 0x1c,
	0x02, 0x93, 0xd9, 0x8f, 0x8f, 0x0e, 0x4b, 0xc4, 0x07, 0x99, 0x0e, 0xe3, 0xfb, 0xa4, 0xd4, 0xbf,
	0x75, 0x2b, 0xcc, 0x97, 0x0c, 0xb1, 0xf9, 0x52, 0xef, 0x29, 0x43, 0x6c, 0xab, 0xbe, 0x1a, 0x3e,
	0x7f, 0x94, 0xde, 0x9e, 0xf4, 0xb6, 0x19, 0xf7, 0xbc, 0xe9, 0x54, 0x7d, 0x45, 0xc2, 0xb4, 0xc2,
	0xdc, 0xbe, 0xf5, 0xd5, 0x50, 0x8f, 0xcf, 0x79, 0xfa, 0xd8, 0xc1, 0xdf, 0x7b, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0x3c, 0x43, 0x43, 0xe3, 0x00, 0x00, 0x00,
}
