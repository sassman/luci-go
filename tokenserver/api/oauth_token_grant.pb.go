// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/tokenserver/api/oauth_token_grant.proto

package tokenserver

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OAuthTokenGrantBody contains the internal guts of an oauth token grant.
//
// It gets serialized, signed and stuffed into OAuthTokenGrantEnvelope, which
// then also gets serialized to get the final blob with the grant. This blob is
// then base64-encoded and returned to the caller of MintOAuthTokenGrant.
type OAuthTokenGrantBody struct {
	// Identifier of this token as generated by the token server.
	//
	// Used for logging and tracking purposes.
	//
	// TODO(vadimsh): It may later be used for revocation purposes.
	TokenId int64 `protobuf:"varint,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	// Service account email the end user wants to act as.
	ServiceAccount string `protobuf:"bytes,2,opt,name=service_account,json=serviceAccount" json:"service_account,omitempty"`
	// Who can pass this token to MintOAuthTokenViaGrant to get an OAuth token.
	//
	// A string of the form "user:<email>". On Swarming, this is Swarming's own
	// service account name.
	Proxy string `protobuf:"bytes,3,opt,name=proxy" json:"proxy,omitempty"`
	// An end user that wants to act as the service account (perhaps indirectly).
	//
	// A string of the form "user:<email>". On Swarming, this is an identity of
	// a user that posted the task.
	//
	// Used by MintOAuthTokenViaGrant to recheck that the access is still allowed.
	EndUser string `protobuf:"bytes,4,opt,name=end_user,json=endUser" json:"end_user,omitempty"`
	// When the token was generated (and when it becomes valid).
	IssuedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=issued_at,json=issuedAt" json:"issued_at,omitempty"`
	// How long the token is considered valid (in seconds).
	//
	// It may become invalid sooner if the token server policy changes and the
	// new policy doesn't allow this token.
	ValidityDuration     int64    `protobuf:"varint,6,opt,name=validity_duration,json=validityDuration" json:"validity_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthTokenGrantBody) Reset()         { *m = OAuthTokenGrantBody{} }
func (m *OAuthTokenGrantBody) String() string { return proto.CompactTextString(m) }
func (*OAuthTokenGrantBody) ProtoMessage()    {}
func (*OAuthTokenGrantBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_token_grant_480c9516e875a2c8, []int{0}
}
func (m *OAuthTokenGrantBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthTokenGrantBody.Unmarshal(m, b)
}
func (m *OAuthTokenGrantBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthTokenGrantBody.Marshal(b, m, deterministic)
}
func (dst *OAuthTokenGrantBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthTokenGrantBody.Merge(dst, src)
}
func (m *OAuthTokenGrantBody) XXX_Size() int {
	return xxx_messageInfo_OAuthTokenGrantBody.Size(m)
}
func (m *OAuthTokenGrantBody) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthTokenGrantBody.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthTokenGrantBody proto.InternalMessageInfo

func (m *OAuthTokenGrantBody) GetTokenId() int64 {
	if m != nil {
		return m.TokenId
	}
	return 0
}

func (m *OAuthTokenGrantBody) GetServiceAccount() string {
	if m != nil {
		return m.ServiceAccount
	}
	return ""
}

func (m *OAuthTokenGrantBody) GetProxy() string {
	if m != nil {
		return m.Proxy
	}
	return ""
}

func (m *OAuthTokenGrantBody) GetEndUser() string {
	if m != nil {
		return m.EndUser
	}
	return ""
}

func (m *OAuthTokenGrantBody) GetIssuedAt() *timestamp.Timestamp {
	if m != nil {
		return m.IssuedAt
	}
	return nil
}

func (m *OAuthTokenGrantBody) GetValidityDuration() int64 {
	if m != nil {
		return m.ValidityDuration
	}
	return 0
}

// OAuthTokenGrantEnvelope is what is actually being serialized and send to
// the callers of MintOAuthTokenGrant (after being encoded using base64 standard
// raw encoding).
type OAuthTokenGrantEnvelope struct {
	TokenBody            []byte   `protobuf:"bytes,1,opt,name=token_body,json=tokenBody,proto3" json:"token_body,omitempty"`
	KeyId                string   `protobuf:"bytes,2,opt,name=key_id,json=keyId" json:"key_id,omitempty"`
	Pkcs1Sha256Sig       []byte   `protobuf:"bytes,3,opt,name=pkcs1_sha256_sig,json=pkcs1Sha256Sig,proto3" json:"pkcs1_sha256_sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthTokenGrantEnvelope) Reset()         { *m = OAuthTokenGrantEnvelope{} }
func (m *OAuthTokenGrantEnvelope) String() string { return proto.CompactTextString(m) }
func (*OAuthTokenGrantEnvelope) ProtoMessage()    {}
func (*OAuthTokenGrantEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_oauth_token_grant_480c9516e875a2c8, []int{1}
}
func (m *OAuthTokenGrantEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthTokenGrantEnvelope.Unmarshal(m, b)
}
func (m *OAuthTokenGrantEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthTokenGrantEnvelope.Marshal(b, m, deterministic)
}
func (dst *OAuthTokenGrantEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthTokenGrantEnvelope.Merge(dst, src)
}
func (m *OAuthTokenGrantEnvelope) XXX_Size() int {
	return xxx_messageInfo_OAuthTokenGrantEnvelope.Size(m)
}
func (m *OAuthTokenGrantEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthTokenGrantEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthTokenGrantEnvelope proto.InternalMessageInfo

func (m *OAuthTokenGrantEnvelope) GetTokenBody() []byte {
	if m != nil {
		return m.TokenBody
	}
	return nil
}

func (m *OAuthTokenGrantEnvelope) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *OAuthTokenGrantEnvelope) GetPkcs1Sha256Sig() []byte {
	if m != nil {
		return m.Pkcs1Sha256Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*OAuthTokenGrantBody)(nil), "tokenserver.OAuthTokenGrantBody")
	proto.RegisterType((*OAuthTokenGrantEnvelope)(nil), "tokenserver.OAuthTokenGrantEnvelope")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/tokenserver/api/oauth_token_grant.proto", fileDescriptor_oauth_token_grant_480c9516e875a2c8)
}

var fileDescriptor_oauth_token_grant_480c9516e875a2c8 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0xc9, 0xed, 0xed, 0xd7, 0xb4, 0xf4, 0xf6, 0xce, 0xbd, 0x62, 0x2c, 0x88, 0xa5, 0x1b,
	0x03, 0x42, 0x82, 0x15, 0x75, 0xe3, 0xa6, 0xa2, 0x48, 0x57, 0x42, 0x5a, 0xd7, 0xc3, 0x34, 0x33,
	0x26, 0x43, 0x3e, 0x4e, 0x98, 0x8f, 0x62, 0xfe, 0xb5, 0x3f, 0x41, 0x32, 0xd3, 0x82, 0xb8, 0x3c,
	0xcf, 0x7b, 0xe6, 0x65, 0x9e, 0x83, 0x1e, 0x52, 0x08, 0x93, 0x4c, 0x42, 0x29, 0x4c, 0x19, 0x82,
	0x4c, 0xa3, 0xc2, 0x24, 0x22, 0xd2, 0x90, 0xf3, 0x4a, 0x71, 0xb9, 0xe7, 0x32, 0xa2, 0xb5, 0x88,
	0x80, 0x1a, 0x9d, 0x11, 0x4b, 0x49, 0x2a, 0x69, 0xa5, 0xc3, 0x5a, 0x82, 0x06, 0x3c, 0xfa, 0xb6,
	0x38, 0xbb, 0x48, 0x01, 0xd2, 0x82, 0x47, 0x36, 0xda, 0x99, 0xf7, 0x48, 0x8b, 0x92, 0x2b, 0x4d,
	0xcb, 0xda, 0x6d, 0x2f, 0x3e, 0x3d, 0xf4, 0xef, 0x75, 0x65, 0x74, 0xb6, 0x6d, 0x5f, 0xbd, 0xb4,
	0x3d, 0x8f, 0xc0, 0x1a, 0x7c, 0x86, 0x06, 0xae, 0x5a, 0x30, 0xdf, 0x9b, 0x7b, 0x41, 0x27, 0xee,
	0xdb, 0x79, 0xcd, 0xf0, 0x25, 0xfa, 0xd3, 0xb6, 0x8b, 0x84, 0x13, 0x9a, 0x24, 0x60, 0x2a, 0xed,
	0xff, 0x9a, 0x7b, 0xc1, 0x30, 0x9e, 0x1c, 0xf0, 0xca, 0x51, 0xfc, 0x1f, 0x75, 0x6b, 0x09, 0x1f,
	0x8d, 0xdf, 0xb1, 0xb1, 0x1b, 0xda, 0x66, 0x5e, 0x31, 0x62, 0x14, 0x97, 0xfe, 0x6f, 0x1b, 0xf4,
	0x79, 0xc5, 0xde, 0x14, 0x97, 0xf8, 0x1e, 0x0d, 0x85, 0x52, 0x86, 0x33, 0x42, 0xb5, 0xdf, 0x9d,
	0x7b, 0xc1, 0x68, 0x39, 0x0b, 0x9d, 0x41, 0x78, 0x34, 0x08, 0xb7, 0x47, 0x83, 0x78, 0xe0, 0x96,
	0x57, 0x1a, 0x5f, 0xa1, 0xbf, 0x7b, 0x5a, 0x08, 0x26, 0x74, 0x43, 0x98, 0x91, 0x54, 0x0b, 0xa8,
	0xfc, 0x9e, 0xfd, 0xf6, 0xf4, 0x18, 0x3c, 0x1d, 0xf8, 0xa2, 0x41, 0xa7, 0x3f, 0x8c, 0x9f, 0xab,
	0x3d, 0x2f, 0xa0, 0xe6, 0xf8, 0x1c, 0x21, 0x67, 0xbd, 0x03, 0xd6, 0x58, 0xef, 0x71, 0x3c, 0xb4,
	0xc4, 0x1e, 0xe5, 0x04, 0xf5, 0x72, 0xde, 0xb4, 0x27, 0x71, 0xc2, 0xdd, 0x9c, 0x37, 0x6b, 0x86,
	0x03, 0x34, 0xad, 0xf3, 0x44, 0x5d, 0x13, 0x95, 0xd1, 0xe5, 0xed, 0x1d, 0x51, 0x22, 0xb5, 0xca,
	0xe3, 0x78, 0x62, 0xf9, 0xc6, 0xe2, 0x8d, 0x48, 0x77, 0x3d, 0x6b, 0x71, 0xf3, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xa8, 0x34, 0x4e, 0x89, 0xe2, 0x01, 0x00, 0x00,
}
