// Code generated by protoc-gen-go.
// source: buildbot.proto
// DO NOT EDIT!

/*
Package buildbot is a generated protocol buffer package.

It is generated from these files:
	buildbot.proto

It has these top-level messages:
	SearchRequest
	SearchResponse
	Build
	ScheduleRequest
	ScheduleResponse
*/
package buildbot

import prpccommon "github.com/luci/luci-go/common/prpc"
import prpc "github.com/luci/luci-go/server/prpc"

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
func (BuildState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// SearchReqeust specifies a search criteria.
type SearchRequest struct {
	// Master filters by master name, e.g. "master.XXX".
	Master string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	// State filters by build state.
	State BuildState `protobuf:"varint,2,opt,name=state,enum=buildbot.BuildState" json:"state,omitempty"`
	// Builder filters by builder name.
	Builder string `protobuf:"bytes,3,opt,name=builder" json:"builder,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SearchResponse struct {
	Builds []*Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type Build struct {
	Master  string     `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Builder string     `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	Number  int32      `protobuf:"varint,3,opt,name=number" json:"number,omitempty"`
	State   BuildState `protobuf:"varint,4,opt,name=state,enum=buildbot.BuildState" json:"state,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// ScheduleRequest defines builds to schedule.
type ScheduleRequest struct {
	// Master is a "master.XXX" string that defines where to schedule builds.
	Master string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	// Builds is a list of builds to schedule.
	Builds []*ScheduleRequest_BuildDef `protobuf:"bytes,2,rep,name=builds" json:"builds,omitempty"`
}

func (m *ScheduleRequest) Reset()                    { *m = ScheduleRequest{} }
func (m *ScheduleRequest) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRequest) ProtoMessage()               {}
func (*ScheduleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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
	Blamelist []string `protobuf:"bytes,5,rep,name=blamelist" json:"blamelist,omitempty"`
}

func (m *ScheduleRequest_BuildDef) Reset()                    { *m = ScheduleRequest_BuildDef{} }
func (m *ScheduleRequest_BuildDef) String() string            { return proto.CompactTextString(m) }
func (*ScheduleRequest_BuildDef) ProtoMessage()               {}
func (*ScheduleRequest_BuildDef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// HelloReply contains a greeting.
type ScheduleResponse struct {
	Builds []*Build `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty"`
}

func (m *ScheduleResponse) Reset()                    { *m = ScheduleResponse{} }
func (m *ScheduleResponse) String() string            { return proto.CompactTextString(m) }
func (*ScheduleResponse) ProtoMessage()               {}
func (*ScheduleResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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

// Client API for Buildbot service

type BuildbotClient interface {
	// Search returns builds matching a criteria.
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// Schedule puts new builds to a queue.
	Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error)
}
type buildbotPRPCClient struct {
	client *prpccommon.Client
}

func NewBuildbotPRPCClient(client *prpccommon.Client) BuildbotClient {
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
	err := grpc.Invoke(ctx, "/buildbot.Buildbot/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) Schedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleResponse, error) {
	out := new(ScheduleResponse)
	err := grpc.Invoke(ctx, "/buildbot.Buildbot/Schedule", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Buildbot service

type BuildbotServer interface {
	// Search returns builds matching a criteria.
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// Schedule puts new builds to a queue.
	Schedule(context.Context, *ScheduleRequest) (*ScheduleResponse, error)
}

func RegisterBuildbotServer(s prpc.Registrar, srv BuildbotServer) {
	s.RegisterService(&_Buildbot_serviceDesc, srv)
}

func _Buildbot_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BuildbotServer).Search(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Buildbot_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BuildbotServer).Schedule(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0x51, 0x6f, 0xa2, 0x40,
	0x10, 0x16, 0x10, 0x84, 0x31, 0x22, 0xb7, 0xb9, 0xe4, 0x38, 0x5e, 0xee, 0xc2, 0xbd, 0x98, 0x7b,
	0x30, 0x39, 0x7c, 0xbe, 0x87, 0x16, 0x69, 0x63, 0xd2, 0x50, 0x23, 0x92, 0x34, 0x4d, 0xfa, 0x00,
	0xba, 0x8d, 0x24, 0x08, 0x96, 0x5d, 0xfa, 0x33, 0xfa, 0x87, 0xfa, 0xe7, 0xba, 0x0b, 0xa2, 0x44,
	0x6b, 0xd2, 0xb7, 0xf9, 0x66, 0x76, 0xbe, 0xfd, 0xbe, 0x99, 0x01, 0x3d, 0x2e, 0x93, 0x74, 0x1d,
	0xe7, 0x74, 0xbc, 0x2b, 0x72, 0x9a, 0x23, 0xb5, 0xc1, 0x76, 0x08, 0x83, 0x00, 0x47, 0xc5, 0x6a,
	0xb3, 0xc0, 0x2f, 0x25, 0x26, 0x14, 0xe9, 0xa0, 0x6c, 0x23, 0x42, 0x71, 0x61, 0x0a, 0xbf, 0x85,
	0x91, 0x86, 0xfe, 0x80, 0x4c, 0x68, 0x44, 0xb1, 0x29, 0x32, 0xa8, 0x3b, 0xdf, 0xc7, 0x07, 0xaa,
	0x6b, 0x1e, 0x04, 0xbc, 0x86, 0x86, 0xd0, 0xab, 0xd2, 0xac, 0x4b, 0xe2, 0x5d, 0xf6, 0x3f, 0xd0,
	0x1b, 0x5a, 0xb2, 0xcb, 0x33, 0x82, 0xd1, 0x2f, 0x50, 0xaa, 0x27, 0x84, 0xf1, 0x4a, 0xa3, 0xbe,
	0x33, 0x3c, 0x21, 0xb2, 0x9f, 0x40, 0xae, 0x82, 0x33, 0x05, 0x2d, 0x72, 0xb1, 0x4a, 0xb0, 0x07,
	0x59, 0xb9, 0x8d, 0xf7, 0x9f, 0xc9, 0x47, 0x89, 0xdd, 0xcb, 0x12, 0xed, 0x77, 0x01, 0x86, 0xc1,
	0x6a, 0x83, 0xd7, 0x65, 0x8a, 0x2f, 0x79, 0x75, 0x0e, 0x1a, 0xc5, 0x4a, 0xa3, 0x7d, 0x64, 0x3a,
	0x69, 0xad, 0x99, 0xa7, 0xf8, 0xd9, 0x5a, 0x83, 0xda, 0xc4, 0x6d, 0xa5, 0x42, 0xa3, 0x34, 0x2e,
	0xa2, 0x6c, 0xb5, 0xd9, 0x2b, 0x37, 0x40, 0x2d, 0xf0, 0x6b, 0x42, 0x92, 0x3c, 0xab, 0x07, 0x85,
	0x10, 0x00, 0x5b, 0xc9, 0x0e, 0x17, 0x34, 0xc1, 0x84, 0x19, 0x90, 0x58, 0xee, 0x1b, 0x68, 0x71,
	0x1a, 0x6d, 0x71, 0x9a, 0x10, 0x6a, 0xca, 0x3c, 0x65, 0x4f, 0xc0, 0x38, 0x2a, 0xf8, 0xe2, 0x44,
	0xff, 0x3e, 0x02, 0xb4, 0x76, 0xa4, 0x81, 0x1c, 0xfa, 0x81, 0xb7, 0x34, 0x3a, 0xa8, 0x0f, 0xbd,
	0xb9, 0xe7, 0x4f, 0x67, 0xfe, 0xad, 0x21, 0x70, 0xb0, 0x08, 0x7d, 0x9f, 0x03, 0x91, 0x83, 0x20,
	0x74, 0x5d, 0x2f, 0x08, 0x0c, 0x89, 0x83, 0x9b, 0xab, 0xd9, 0x5d, 0xb8, 0xf0, 0x8c, 0x2e, 0x1a,
	0x80, 0xe6, 0x3d, 0xb8, 0xde, 0x7c, 0x39, 0xbb, 0xf7, 0x0d, 0xd9, 0x79, 0x13, 0xf6, 0xbe, 0xd9,
	0x77, 0xe8, 0x3f, 0x28, 0xf5, 0xb6, 0xd1, 0x8f, 0xd6, 0xc4, 0xda, 0x67, 0x65, 0x99, 0xe7, 0x85,
	0xda, 0x86, 0xdd, 0x41, 0x2e, 0xa8, 0x8d, 0x39, 0xf4, 0xf3, 0xe2, 0xc8, 0x2d, 0xeb, 0xb3, 0x52,
	0x43, 0x12, 0x2b, 0xd5, 0x65, 0x4f, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x4a, 0xf3, 0x8a,
	0xeb, 0x02, 0x00, 0x00,
}
