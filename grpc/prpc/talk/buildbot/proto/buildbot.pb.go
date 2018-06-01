// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/grpc/prpc/talk/buildbot/proto/buildbot.proto

package buildbot

import prpc "go.chromium.org/luci/grpc/prpc"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildState int32

const (
	BuildState_UNSET     BuildState = 0
	BuildState_PENDING   BuildState = 1
	BuildState_RUNNING   BuildState = 2
	BuildState_SUCCESS   BuildState = 3
	BuildState_FAILURE   BuildState = 4
	BuildState_EXCEPTION BuildState = 5
)

var BuildState_name = map[int32]string{
	0: "UNSET",
	1: "PENDING",
	2: "RUNNING",
	3: "SUCCESS",
	4: "FAILURE",
	5: "EXCEPTION",
}
var BuildState_value = map[string]int32{
	"UNSET":     0,
	"PENDING":   1,
	"RUNNING":   2,
	"SUCCESS":   3,
	"FAILURE":   4,
	"EXCEPTION": 5,
}

func (x BuildState) String() string {
	return proto.EnumName(BuildState_name, int32(x))
}
func (BuildState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{0}
}

// SearchReqeust specifies a search criteria.
type SearchRequest struct {
	// Master filters by master name, e.g. "master.XXX".
	Master string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	// State filters by build state.
	State BuildState `protobuf:"varint,2,opt,name=state,enum=buildbot.BuildState" json:"state,omitempty"`
	// Builder filters by builder name.
	Builder              string   `protobuf:"bytes,3,opt,name=builder" json:"builder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{0}
}
func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (dst *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(dst, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *SearchRequest) GetState() BuildState {
	if m != nil {
		return m.State
	}
	return BuildState_UNSET
}

func (m *SearchRequest) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

type SearchResponse struct {
	Builds               []*Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{1}
}
func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (dst *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(dst, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type Build struct {
	Master               string     `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Builder              string     `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	Number               int32      `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	State                BuildState `protobuf:"varint,4,opt,name=state,enum=buildbot.BuildState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{2}
}
func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (dst *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(dst, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *Build) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *Build) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Build) GetState() BuildState {
	if m != nil {
		return m.State
	}
	return BuildState_UNSET
}

// ScheduleRequest defines builds to schedule.
type ScheduleRequest struct {
	// Master is a "master.XXX" string that defines where to schedule builds.
	Master string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	// Builds is a list of builds to schedule.
	Builds               []*ScheduleRequest_BuildDef `protobuf:"bytes,2,rep,name=builds" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ScheduleRequest) Reset()         { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()    {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{3}
}
func (m *ScheduleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleRequest.Unmarshal(m, b)
}
func (m *ScheduleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleRequest.Marshal(b, m, deterministic)
}
func (dst *ScheduleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleRequest.Merge(dst, src)
}
func (m *ScheduleRequest) XXX_Size() int {
	return xxx_messageInfo_ScheduleRequest.Size(m)
}
func (m *ScheduleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleRequest proto.InternalMessageInfo

func (m *ScheduleRequest) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ScheduleRequest) GetBuilds() []*ScheduleRequest_BuildDef {
	if m != nil {
		return m.Builds
	}
	return nil
}

// Build is a build to schedule.
type ScheduleRequest_BuildDef struct {
	// Builder defines the build script.
	Builder string `protobuf:"bytes,1,opt,name=builder" json:"builder,omitempty"`
	// Branch defines what to fetch.
	Branch string `protobuf:"bytes,2,opt,name=branch" json:"branch,omitempty"`
	// Revision is a commit hash to checkout
	Revision string `protobuf:"bytes,3,opt,name=revision" json:"revision,omitempty"`
	// Properties are "key:value" pairs.
	Properties []string `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
	// Blamelist is a list of user email addressed to blame if this build
	// fails.
	Blamelist            []string `protobuf:"bytes,5,rep,name=blamelist" json:"blamelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleRequest_BuildDef) Reset()         { *m = ScheduleRequest_BuildDef{} }
func (m *ScheduleRequest_BuildDef) String() string { return proto.CompactTextString(m) }
func (*ScheduleRequest_BuildDef) ProtoMessage()    {}
func (*ScheduleRequest_BuildDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{3, 0}
}
func (m *ScheduleRequest_BuildDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleRequest_BuildDef.Unmarshal(m, b)
}
func (m *ScheduleRequest_BuildDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleRequest_BuildDef.Marshal(b, m, deterministic)
}
func (dst *ScheduleRequest_BuildDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleRequest_BuildDef.Merge(dst, src)
}
func (m *ScheduleRequest_BuildDef) XXX_Size() int {
	return xxx_messageInfo_ScheduleRequest_BuildDef.Size(m)
}
func (m *ScheduleRequest_BuildDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleRequest_BuildDef.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleRequest_BuildDef proto.InternalMessageInfo

func (m *ScheduleRequest_BuildDef) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *ScheduleRequest_BuildDef) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *ScheduleRequest_BuildDef) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ScheduleRequest_BuildDef) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *ScheduleRequest_BuildDef) GetBlamelist() []string {
	if m != nil {
		return m.Blamelist
	}
	return nil
}

// HelloReply contains a greeting.
type ScheduleResponse struct {
	Builds               []*Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleResponse) Reset()         { *m = ScheduleResponse{} }
func (m *ScheduleResponse) String() string { return proto.CompactTextString(m) }
func (*ScheduleResponse) ProtoMessage()    {}
func (*ScheduleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_buildbot_a1cbc2dc36e5eb0e, []int{4}
}
func (m *ScheduleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleResponse.Unmarshal(m, b)
}
func (m *ScheduleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleResponse.Marshal(b, m, deterministic)
}
func (dst *ScheduleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleResponse.Merge(dst, src)
}
func (m *ScheduleResponse) XXX_Size() int {
	return xxx_messageInfo_ScheduleResponse.Size(m)
}
func (m *ScheduleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleResponse proto.InternalMessageInfo

func (m *ScheduleResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

func init() {
	proto.RegisterType((*SearchRequest)(nil), "buildbot.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "buildbot.SearchResponse")
	proto.RegisterType((*Build)(nil), "buildbot.Build")
	proto.RegisterType((*ScheduleRequest)(nil), "buildbot.ScheduleRequest")
	proto.RegisterType((*ScheduleRequest_BuildDef)(nil), "buildbot.ScheduleRequest.BuildDef")
	proto.RegisterType((*ScheduleResponse)(nil), "buildbot.ScheduleResponse")
	proto.RegisterEnum("buildbot.BuildState", BuildState_name, BuildState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildbotClient is the client API for Buildbot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildbotClient interface {
	// Search returns builds matching a criteria.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Schedule puts new builds to a queue.
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
}
type buildbotPRPCClient struct {
	client *prpc.Client
}

func NewBuildbotPRPCClient(client *prpc.Client) BuildbotClient {
	return &buildbotPRPCClient{client}
}

func (c *buildbotPRPCClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.client.Call(ctx, "buildbot.Buildbot", "Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotPRPCClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.client.Call(ctx, "buildbot.Buildbot", "Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildbotClient struct {
	cc *grpc.ClientConn
}

func NewBuildbotClient(cc *grpc.ClientConn) BuildbotClient {
	return &buildbotClient{cc}
}

func (c *buildbotClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/buildbot.Buildbot/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := c.cc.Invoke(ctx, "/buildbot.Buildbot/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildbotServer is the server API for Buildbot service.
type BuildbotServer interface {
	// Search returns builds matching a criteria.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Schedule puts new builds to a queue.
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
}

func RegisterBuildbotServer(s prpc.Registrar, srv BuildbotServer) {
	s.RegisterService(&_Buildbot_serviceDesc, srv)
}

func _Buildbot_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbot.Buildbot/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildbot_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbot.Buildbot/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).Schedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buildbot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbot.Buildbot",
	HandlerType: (*BuildbotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _Buildbot_Search_Handler,
		},
		{
			MethodName: "Schedule",
			Handler:    _Buildbot_Schedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/grpc/prpc/talk/buildbot/proto/buildbot.proto",
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/grpc/prpc/talk/buildbot/proto/buildbot.proto", fileDescriptor_buildbot_a1cbc2dc36e5eb0e)
}

var fileDescriptor_buildbot_a1cbc2dc36e5eb0e = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xb6, 0xe9, 0x9a, 0x37, 0x6d, 0x8b, 0x2c, 0x04, 0x26, 0x42, 0xa8, 0xca, 0x85,
	0x6a, 0x87, 0x46, 0x2a, 0x27, 0x40, 0x1c, 0x46, 0x17, 0x50, 0x25, 0x14, 0x26, 0x67, 0x95, 0x10,
	0xb7, 0x24, 0x33, 0x6d, 0x44, 0x52, 0x07, 0xdb, 0xe1, 0xc6, 0xbf, 0xc0, 0x91, 0x3f, 0x17, 0x21,
	0xbb, 0x71, 0x12, 0x36, 0xa6, 0x69, 0x97, 0xaa, 0xbf, 0xf7, 0x25, 0xdf, 0x7b, 0xdf, 0xb3, 0x03,
	0xe7, 0x1b, 0x36, 0xcf, 0xb6, 0x9c, 0x95, 0x79, 0x5d, 0xce, 0x19, 0xdf, 0x04, 0x45, 0x9d, 0xe5,
	0xc1, 0x86, 0x57, 0x59, 0x50, 0xa9, 0x1f, 0x99, 0x14, 0xdf, 0x82, 0xb4, 0xce, 0x8b, 0xeb, 0x94,
	0xc9, 0xa0, 0xe2, 0x4c, 0xb2, 0x16, 0xe7, 0x1a, 0xd1, 0xc4, 0xb0, 0x5f, 0xc2, 0x71, 0x4c, 0x13,
	0x9e, 0x6d, 0x09, 0xfd, 0x5e, 0x53, 0x21, 0xd1, 0x63, 0x18, 0x97, 0x89, 0x90, 0x94, 0x63, 0x6b,
	0x6a, 0xcd, 0x1c, 0xd2, 0x10, 0x3a, 0x03, 0x5b, 0xc8, 0x44, 0x52, 0x3c, 0x98, 0x5a, 0xb3, 0x93,
	0xc5, 0xa3, 0x79, 0x6b, 0xf9, 0x4e, 0xfd, 0x89, 0x95, 0x46, 0xf6, 0x8f, 0x20, 0x0c, 0x87, 0x5a,
	0xa5, 0x1c, 0x0f, 0xb5, 0x89, 0x41, 0xff, 0x15, 0x9c, 0x98, 0x76, 0xa2, 0x62, 0x3b, 0x41, 0xd1,
	0x0b, 0x18, 0x6b, 0x51, 0x60, 0x6b, 0x3a, 0x9c, 0x1d, 0x2d, 0x4e, 0x6f, 0x18, 0x93, 0x46, 0xf6,
	0x7f, 0x82, 0xad, 0x0b, 0x77, 0x4e, 0xd8, 0xeb, 0x3a, 0xf8, 0xa7, 0xab, 0x7a, 0x63, 0x57, 0x97,
	0x69, 0x33, 0x8e, 0x4d, 0x1a, 0xea, 0x32, 0x8d, 0xee, 0xcd, 0xe4, 0xff, 0xb1, 0xe0, 0x34, 0xce,
	0xb6, 0xf4, 0xba, 0x2e, 0xe8, 0x7d, 0xbb, 0x7a, 0xdd, 0x66, 0x1a, 0xe8, 0x4c, 0x7e, 0x67, 0x7c,
	0xc3, 0x62, 0xdf, 0xe8, 0x82, 0x7e, 0x35, 0x31, 0xbd, 0xdf, 0x16, 0x4c, 0x4c, 0xb1, 0x1f, 0xc9,
	0xba, 0x15, 0x29, 0xe5, 0xc9, 0x2e, 0xdb, 0x36, 0x59, 0x1b, 0x42, 0x1e, 0x4c, 0x38, 0xfd, 0x91,
	0x8b, 0x9c, 0xed, 0x9a, 0xdd, 0xb7, 0x8c, 0x9e, 0x03, 0x54, 0x9c, 0x55, 0x94, 0xcb, 0x9c, 0x0a,
	0x3c, 0x9a, 0x0e, 0x67, 0x0e, 0xe9, 0x55, 0xd0, 0x33, 0x70, 0xd2, 0x22, 0x29, 0x69, 0x91, 0x0b,
	0x89, 0x6d, 0x2d, 0x77, 0x05, 0xff, 0x0d, 0xb8, 0xdd, 0xf0, 0x0f, 0x3c, 0xbc, 0xb3, 0x2f, 0x00,
	0xdd, 0x4a, 0x91, 0x03, 0xf6, 0x3a, 0x8a, 0xc3, 0x2b, 0xf7, 0x00, 0x1d, 0xc1, 0xe1, 0x65, 0x18,
	0x5d, 0xac, 0xa2, 0x0f, 0xae, 0xa5, 0x80, 0xac, 0xa3, 0x48, 0xc1, 0x40, 0x41, 0xbc, 0x5e, 0x2e,
	0xc3, 0x38, 0x76, 0x87, 0x0a, 0xde, 0x9f, 0xaf, 0x3e, 0xae, 0x49, 0xe8, 0x8e, 0xd0, 0x31, 0x38,
	0xe1, 0xe7, 0x65, 0x78, 0x79, 0xb5, 0xfa, 0x14, 0xb9, 0xf6, 0xe2, 0x97, 0xd9, 0x58, 0xca, 0x24,
	0x7a, 0x0b, 0xe3, 0xfd, 0x05, 0x43, 0x4f, 0x7a, 0x4b, 0xef, 0xdf, 0x70, 0x0f, 0xdf, 0x16, 0xf6,
	0x71, 0xfc, 0x03, 0xb4, 0x84, 0x89, 0x09, 0x89, 0x9e, 0xde, 0x79, 0x6a, 0x9e, 0xf7, 0x3f, 0xc9,
	0x98, 0xa4, 0x63, 0xfd, 0x91, 0xbd, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x4f, 0xe1, 0x41,
	0xa9, 0x03, 0x00, 0x00,
}
