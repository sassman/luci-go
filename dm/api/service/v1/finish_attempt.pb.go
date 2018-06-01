// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/dm/api/service/v1/finish_attempt.proto

package dm

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

// FinishAttemptReq sets the final result of an Attempt.
type FinishAttemptReq struct {
	// required
	Auth *Execution_Auth `protobuf:"bytes,1,opt,name=auth" json:"auth,omitempty"`
	// The result data for this Attempt. The `size` field is recalculated after
	// the data field is normalized, and may be omitted.
	Data                 *JsonResult `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FinishAttemptReq) Reset()         { *m = FinishAttemptReq{} }
func (m *FinishAttemptReq) String() string { return proto.CompactTextString(m) }
func (*FinishAttemptReq) ProtoMessage()    {}
func (*FinishAttemptReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_finish_attempt_e19e6442529b68e6, []int{0}
}
func (m *FinishAttemptReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishAttemptReq.Unmarshal(m, b)
}
func (m *FinishAttemptReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishAttemptReq.Marshal(b, m, deterministic)
}
func (dst *FinishAttemptReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishAttemptReq.Merge(dst, src)
}
func (m *FinishAttemptReq) XXX_Size() int {
	return xxx_messageInfo_FinishAttemptReq.Size(m)
}
func (m *FinishAttemptReq) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishAttemptReq.DiscardUnknown(m)
}

var xxx_messageInfo_FinishAttemptReq proto.InternalMessageInfo

func (m *FinishAttemptReq) GetAuth() *Execution_Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *FinishAttemptReq) GetData() *JsonResult {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*FinishAttemptReq)(nil), "dm.FinishAttemptReq")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/dm/api/service/v1/finish_attempt.proto", fileDescriptor_finish_attempt_e19e6442529b68e6)
}

var fileDescriptor_finish_attempt_e19e6442529b68e6 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcd, 0xb1, 0xae, 0xc2, 0x20,
	0x14, 0xc6, 0xf1, 0xb4, 0x69, 0xee, 0xc0, 0x4d, 0x8c, 0x61, 0x6a, 0x9c, 0x4c, 0x07, 0xe3, 0x04,
	0x51, 0x07, 0x07, 0xa7, 0x0e, 0x3a, 0x38, 0xf2, 0x00, 0x36, 0x08, 0x58, 0x48, 0x4a, 0x41, 0x38,
	0x34, 0x3e, 0xbe, 0x29, 0x7d, 0x01, 0xe7, 0xef, 0x7f, 0x7e, 0x07, 0x5d, 0x7a, 0x47, 0x84, 0x0e,
	0xce, 0x9a, 0x64, 0x89, 0x0b, 0x3d, 0x1d, 0x92, 0x30, 0x54, 0x5a, 0xca, 0xbd, 0xa1, 0x51, 0x85,
	0xc9, 0x08, 0x45, 0xa7, 0x03, 0x7d, 0x99, 0xd1, 0x44, 0xdd, 0x71, 0x00, 0x65, 0x3d, 0x10, 0x1f,
	0x1c, 0x38, 0x5c, 0x4a, 0xbb, 0x39, 0xff, 0x08, 0xf4, 0x81, 0x7b, 0xdd, 0x49, 0x0e, 0x7c, 0x39,
	0x6e, 0x1e, 0x68, 0x7d, 0xcb, 0x68, 0xbb, 0x98, 0x4c, 0xbd, 0xf1, 0x0e, 0x55, 0x3c, 0x81, 0xae,
	0x8b, 0x6d, 0xb1, 0xff, 0x3f, 0x62, 0x22, 0x2d, 0xb9, 0x7e, 0x94, 0x48, 0x60, 0xdc, 0x48, 0xda,
	0x04, 0x9a, 0xe5, 0x1d, 0x37, 0xa8, 0x9a, 0xa5, 0xba, 0xcc, 0xdd, 0x6a, 0xee, 0xee, 0xd1, 0x8d,
	0x4c, 0xc5, 0x34, 0x00, 0xcb, 0xdb, 0xf3, 0x2f, 0xbf, 0x39, 0x7d, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x20, 0x42, 0xb2, 0x57, 0xe2, 0x00, 0x00, 0x00,
}
