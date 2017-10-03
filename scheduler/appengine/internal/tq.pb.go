// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/tq.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/internal/tq.proto
	go.chromium.org/luci/scheduler/appengine/internal/triggers.proto
	go.chromium.org/luci/scheduler/appengine/internal/types.proto

It has these top-level messages:
	ReadProjectConfigTask
	LaunchInvocationTask
	LaunchInvocationsBatchTask
	TriageJobStateTask
	InvocationFinishedTask
	TriggerPayload
	GitilesTrigger
	Trigger
*/
package internal

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

// ReadProjectConfigTask is used to import jobs of some project.
//
// Queue: "read-project-config".
type ReadProjectConfigTask struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId" json:"project_id,omitempty"`
}

func (m *ReadProjectConfigTask) Reset()                    { *m = ReadProjectConfigTask{} }
func (m *ReadProjectConfigTask) String() string            { return proto.CompactTextString(m) }
func (*ReadProjectConfigTask) ProtoMessage()               {}
func (*ReadProjectConfigTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReadProjectConfigTask) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

// LaunchInvocationTask is used to start running (or retry a lunch of) a single
// invocation.
//
// It is enqueued non-transactionally, but with the deduplication key.
//
// Queue: "launches".
type LaunchInvocationTask struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	InvId int64  `protobuf:"varint,2,opt,name=inv_id,json=invId" json:"inv_id,omitempty"`
}

func (m *LaunchInvocationTask) Reset()                    { *m = LaunchInvocationTask{} }
func (m *LaunchInvocationTask) String() string            { return proto.CompactTextString(m) }
func (*LaunchInvocationTask) ProtoMessage()               {}
func (*LaunchInvocationTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LaunchInvocationTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *LaunchInvocationTask) GetInvId() int64 {
	if m != nil {
		return m.InvId
	}
	return 0
}

// LaunchInvocationsBatchTask is used to kick off several invocations at once.
//
// It is enqueued transactionally. It fans out into many LaunchInvocationTask.
//
// Queue: "batches".
type LaunchInvocationsBatchTask struct {
	Tasks []*LaunchInvocationTask `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *LaunchInvocationsBatchTask) Reset()                    { *m = LaunchInvocationsBatchTask{} }
func (m *LaunchInvocationsBatchTask) String() string            { return proto.CompactTextString(m) }
func (*LaunchInvocationsBatchTask) ProtoMessage()               {}
func (*LaunchInvocationsBatchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LaunchInvocationsBatchTask) GetTasks() []*LaunchInvocationTask {
	if m != nil {
		return m.Tasks
	}
	return nil
}

// TriageJobStateTask looks at the state of the job and decided what to do next.
//
// Enqueued non-transactionally. It is throttled to run approximately once per
// second. It looks at pending triggers and recently finished invocations and
// launches new invocations (or schedules timers to do it later).
//
// Queue: "triages".
type TriageJobStateTask struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
}

func (m *TriageJobStateTask) Reset()                    { *m = TriageJobStateTask{} }
func (m *TriageJobStateTask) String() string            { return proto.CompactTextString(m) }
func (*TriageJobStateTask) ProtoMessage()               {}
func (*TriageJobStateTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TriageJobStateTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

// InvocationFinishedTask is emitted by the invocation when it finishes.
//
// It is enqueued transactionally.
//
// Queue: "completions".
type InvocationFinishedTask struct {
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	InvId int64  `protobuf:"varint,2,opt,name=inv_id,json=invId" json:"inv_id,omitempty"`
}

func (m *InvocationFinishedTask) Reset()                    { *m = InvocationFinishedTask{} }
func (m *InvocationFinishedTask) String() string            { return proto.CompactTextString(m) }
func (*InvocationFinishedTask) ProtoMessage()               {}
func (*InvocationFinishedTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InvocationFinishedTask) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *InvocationFinishedTask) GetInvId() int64 {
	if m != nil {
		return m.InvId
	}
	return 0
}

func init() {
	proto.RegisterType((*ReadProjectConfigTask)(nil), "internal.tq.ReadProjectConfigTask")
	proto.RegisterType((*LaunchInvocationTask)(nil), "internal.tq.LaunchInvocationTask")
	proto.RegisterType((*LaunchInvocationsBatchTask)(nil), "internal.tq.LaunchInvocationsBatchTask")
	proto.RegisterType((*TriageJobStateTask)(nil), "internal.tq.TriageJobStateTask")
	proto.RegisterType((*InvocationFinishedTask)(nil), "internal.tq.InvocationFinishedTask")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/tq.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0xa5, 0x85, 0x6e, 0x6f, 0xc1, 0x4a, 0x11, 0x84, 0x98, 0x53, 0x40, 0xd8, 0x80,
	0x82, 0x82, 0x47, 0x95, 0x42, 0xc4, 0x83, 0xc4, 0x7a, 0x96, 0xcd, 0xee, 0x9a, 0x4c, 0x9a, 0xce,
	0xa4, 0xbb, 0x93, 0xfc, 0x7e, 0x69, 0xaa, 0x54, 0x44, 0x04, 0xaf, 0x1f, 0xf3, 0xbe, 0x79, 0x33,
	0xe2, 0xb6, 0x24, 0xa9, 0x2b, 0x47, 0x1b, 0xe8, 0x36, 0x92, 0x5c, 0x99, 0x36, 0x9d, 0x86, 0xd4,
	0xeb, 0xca, 0x9a, 0xae, 0xb1, 0x2e, 0x55, 0x6d, 0x6b, 0xb1, 0x04, 0xb4, 0x29, 0x20, 0x5b, 0x87,
	0xaa, 0x49, 0x79, 0x2b, 0x5b, 0x47, 0x4c, 0xe1, 0xec, 0x0b, 0x49, 0xde, 0xc6, 0xd7, 0x62, 0x9e,
	0x5b, 0x65, 0x9e, 0x1d, 0xd5, 0x56, 0xf3, 0x3d, 0xe1, 0x3b, 0x94, 0x2b, 0xe5, 0xd7, 0xe1, 0x99,
	0x10, 0xed, 0x1e, 0xbe, 0x81, 0x59, 0x04, 0x51, 0x90, 0x4c, 0xf3, 0xe9, 0x27, 0xc9, 0x4c, 0xfc,
	0x20, 0x8e, 0x9f, 0x54, 0x87, 0xba, 0xca, 0xb0, 0x27, 0xad, 0x18, 0x08, 0x87, 0xd8, 0x5c, 0x4c,
	0x6a, 0x2a, 0x0e, 0x91, 0x71, 0x4d, 0x45, 0x66, 0x76, 0x18, 0xb0, 0xdf, 0xe1, 0xa3, 0x28, 0x48,
	0x46, 0xf9, 0x18, 0xb0, 0xcf, 0x4c, 0xfc, 0x2a, 0x4e, 0x7f, 0x5a, 0xfc, 0x9d, 0x62, 0x5d, 0x0d,
	0xae, 0x1b, 0x31, 0x66, 0xe5, 0xd7, 0x7e, 0x11, 0x44, 0xa3, 0x64, 0x76, 0x79, 0x2e, 0xbf, 0x15,
	0x97, 0xbf, 0x6d, 0xcf, 0xf7, 0xf3, 0xf1, 0x85, 0x08, 0x57, 0x0e, 0x54, 0x69, 0x1f, 0xa9, 0x78,
	0x61, 0xc5, 0xf6, 0x8f, 0x6a, 0xf1, 0x52, 0x9c, 0x1c, 0x2c, 0x4b, 0x40, 0xf0, 0x95, 0x35, 0xff,
	0xbf, 0xa5, 0x98, 0x0c, 0xdf, 0xbd, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xef, 0x5b, 0xfe,
	0x9b, 0x01, 0x00, 0x00,
}
