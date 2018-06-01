// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/triggers.proto

package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import v1 "go.chromium.org/luci/scheduler/api/scheduler/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Trigger can be emitted by the engine itself (e.g. on a schedule) or by
// triggering tasks (such as Gitiles tasks).
//
// One or multiple triggers are consumed to initiate a new invocation which has
// access to the properties of consumed triggers. For example, Buildbucket task
// knows about triggers produced by Gitiles tasks.
//
// This message is an internal representation of the trigger, as stored in
// the datastore. See also triggers.Trigger for public representation used in
// API calls.
type Trigger struct {
	// Unique in time identifier of the trigger.
	//
	// It is used to deduplicate and hence provide idempotency for adding
	// a trigger. Must be provided by whoever emits the trigger.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// ID of a job that emitted this trigger or "" if emitted by the engine.
	//
	// Set by the engine, can't be overridden.
	JobId string `protobuf:"bytes,2,opt,name=job_id,json=jobId" json:"job_id,omitempty"`
	// ID of an invocation that emitted this trigger or 0 if emitted by the
	// engine.
	//
	// Set by the engine, can't be overridden.
	InvocationId int64 `protobuf:"varint,3,opt,name=invocation_id,json=invocationId" json:"invocation_id,omitempty"`
	// Timestamp when the trigger was created.
	//
	// Can be set by whoever emits the trigger if the trigger is based on some
	// external event. If not provided, the engine will set it to the current
	// time.
	//
	// Together with 'order_in_batch' used for weak ordering of triggers that
	// aren't directly comparable (e.g. git commits from different repositories).
	// This ordering shouldn't be considered reliable.
	Created *timestamp.Timestamp `protobuf:"bytes,4,opt,name=created" json:"created,omitempty"`
	// If a bunch of triggers were emitted at the same moment in time (for example
	// through a single RPC or by a single invocation in a tight loop), a trigger
	// with smaller 'order_in_batch' is considered to be older. Value of
	// 'order_in_batch' for triggers with different 'created' timestamps are not
	// comparable.
	//
	// Should be set by whoever emits the trigger if 'created' timestamp was
	// supplied explicitly. Otherwise will be set by the engine based on the order
	// of EmitTrigger calls done by the invocation.
	//
	// Together with 'order_in_batch' used for weak ordering of triggers that
	// aren't directly comparable (e.g. git commits from different repositories).
	// This ordering shouldn't be considered reliable.
	OrderInBatch int64 `protobuf:"varint,7,opt,name=order_in_batch,json=orderInBatch" json:"order_in_batch,omitempty"`
	// User friendly name for this trigger that shows up in UI.
	//
	// Can be provided by whoever emits the trigger. Doesn't have to be unique.
	Title string `protobuf:"bytes,5,opt,name=title" json:"title,omitempty"`
	// Optional HTTP link to display in UI.
	//
	// Can be provided by whoever emits the trigger. Doesn't have to be unique.
	Url string `protobuf:"bytes,6,opt,name=url" json:"url,omitempty"`
	// For triggers emitted through public API contains identity of a user who
	// posted this trigger. Empty for triggers emitted by invocations.
	EmittedByUser string `protobuf:"bytes,8,opt,name=emitted_by_user,json=emittedByUser" json:"emitted_by_user,omitempty"`
	// Actual trigger data that depends on type of the trigger.
	//
	// Types that are valid to be assigned to Payload:
	//	*Trigger_Cron
	//	*Trigger_Noop
	//	*Trigger_Gitiles
	//	*Trigger_Buildbucket
	Payload              isTrigger_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_triggers_d5c70b81999e9321, []int{0}
}
func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (dst *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(dst, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

type isTrigger_Payload interface {
	isTrigger_Payload()
}

type Trigger_Cron struct {
	Cron *v1.CronTrigger `protobuf:"bytes,40,opt,name=cron,oneof"`
}
type Trigger_Noop struct {
	Noop *v1.NoopTrigger `protobuf:"bytes,50,opt,name=noop,oneof"`
}
type Trigger_Gitiles struct {
	Gitiles *v1.GitilesTrigger `protobuf:"bytes,51,opt,name=gitiles,oneof"`
}
type Trigger_Buildbucket struct {
	Buildbucket *v1.BuildbucketTrigger `protobuf:"bytes,52,opt,name=buildbucket,oneof"`
}

func (*Trigger_Cron) isTrigger_Payload()        {}
func (*Trigger_Noop) isTrigger_Payload()        {}
func (*Trigger_Gitiles) isTrigger_Payload()     {}
func (*Trigger_Buildbucket) isTrigger_Payload() {}

func (m *Trigger) GetPayload() isTrigger_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Trigger) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trigger) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *Trigger) GetInvocationId() int64 {
	if m != nil {
		return m.InvocationId
	}
	return 0
}

func (m *Trigger) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Trigger) GetOrderInBatch() int64 {
	if m != nil {
		return m.OrderInBatch
	}
	return 0
}

func (m *Trigger) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Trigger) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Trigger) GetEmittedByUser() string {
	if m != nil {
		return m.EmittedByUser
	}
	return ""
}

func (m *Trigger) GetCron() *v1.CronTrigger {
	if x, ok := m.GetPayload().(*Trigger_Cron); ok {
		return x.Cron
	}
	return nil
}

func (m *Trigger) GetNoop() *v1.NoopTrigger {
	if x, ok := m.GetPayload().(*Trigger_Noop); ok {
		return x.Noop
	}
	return nil
}

func (m *Trigger) GetGitiles() *v1.GitilesTrigger {
	if x, ok := m.GetPayload().(*Trigger_Gitiles); ok {
		return x.Gitiles
	}
	return nil
}

func (m *Trigger) GetBuildbucket() *v1.BuildbucketTrigger {
	if x, ok := m.GetPayload().(*Trigger_Buildbucket); ok {
		return x.Buildbucket
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Trigger) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Trigger_OneofMarshaler, _Trigger_OneofUnmarshaler, _Trigger_OneofSizer, []interface{}{
		(*Trigger_Cron)(nil),
		(*Trigger_Noop)(nil),
		(*Trigger_Gitiles)(nil),
		(*Trigger_Buildbucket)(nil),
	}
}

func _Trigger_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Trigger)
	// payload
	switch x := m.Payload.(type) {
	case *Trigger_Cron:
		b.EncodeVarint(40<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cron); err != nil {
			return err
		}
	case *Trigger_Noop:
		b.EncodeVarint(50<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Noop); err != nil {
			return err
		}
	case *Trigger_Gitiles:
		b.EncodeVarint(51<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Gitiles); err != nil {
			return err
		}
	case *Trigger_Buildbucket:
		b.EncodeVarint(52<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Buildbucket); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Trigger.Payload has unexpected type %T", x)
	}
	return nil
}

func _Trigger_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Trigger)
	switch tag {
	case 40: // payload.cron
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1.CronTrigger)
		err := b.DecodeMessage(msg)
		m.Payload = &Trigger_Cron{msg}
		return true, err
	case 50: // payload.noop
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1.NoopTrigger)
		err := b.DecodeMessage(msg)
		m.Payload = &Trigger_Noop{msg}
		return true, err
	case 51: // payload.gitiles
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1.GitilesTrigger)
		err := b.DecodeMessage(msg)
		m.Payload = &Trigger_Gitiles{msg}
		return true, err
	case 52: // payload.buildbucket
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(v1.BuildbucketTrigger)
		err := b.DecodeMessage(msg)
		m.Payload = &Trigger_Buildbucket{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Trigger_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Trigger)
	// payload
	switch x := m.Payload.(type) {
	case *Trigger_Cron:
		s := proto.Size(x.Cron)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Trigger_Noop:
		s := proto.Size(x.Noop)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Trigger_Gitiles:
		s := proto.Size(x.Gitiles)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Trigger_Buildbucket:
		s := proto.Size(x.Buildbucket)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// TriggerList is what we store in datastore entities.
type TriggerList struct {
	Triggers             []*Trigger `protobuf:"bytes,1,rep,name=triggers" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TriggerList) Reset()         { *m = TriggerList{} }
func (m *TriggerList) String() string { return proto.CompactTextString(m) }
func (*TriggerList) ProtoMessage()    {}
func (*TriggerList) Descriptor() ([]byte, []int) {
	return fileDescriptor_triggers_d5c70b81999e9321, []int{1}
}
func (m *TriggerList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerList.Unmarshal(m, b)
}
func (m *TriggerList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerList.Marshal(b, m, deterministic)
}
func (dst *TriggerList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerList.Merge(dst, src)
}
func (m *TriggerList) XXX_Size() int {
	return xxx_messageInfo_TriggerList.Size(m)
}
func (m *TriggerList) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerList.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerList proto.InternalMessageInfo

func (m *TriggerList) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

func init() {
	proto.RegisterType((*Trigger)(nil), "internal.triggers.Trigger")
	proto.RegisterType((*TriggerList)(nil), "internal.triggers.TriggerList")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/triggers.proto", fileDescriptor_triggers_d5c70b81999e9321)
}

var fileDescriptor_triggers_d5c70b81999e9321 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x86, 0x3f, 0xc7, 0x4d, 0x9c, 0x4e, 0xbe, 0x16, 0x18, 0x01, 0x1a, 0x22, 0x21, 0xa2, 0x82,
	0x50, 0x16, 0x68, 0x2c, 0xda, 0xc2, 0x12, 0x81, 0x11, 0x82, 0x48, 0x88, 0x85, 0x55, 0x36, 0x6c,
	0x2c, 0xdb, 0x73, 0x70, 0x4e, 0xb1, 0xe7, 0x58, 0xe3, 0x71, 0xa5, 0xdc, 0x2e, 0x57, 0x82, 0x3c,
	0x8e, 0x1b, 0x8b, 0x2e, 0xd8, 0x79, 0xde, 0xf7, 0x79, 0x3c, 0xbf, 0xec, 0x7d, 0x41, 0x32, 0xdf,
	0x1a, 0xaa, 0xb0, 0xad, 0x24, 0x99, 0x22, 0x2c, 0xdb, 0x1c, 0xc3, 0x26, 0xdf, 0x82, 0x6a, 0x4b,
	0x30, 0x61, 0x5a, 0xd7, 0xa0, 0x0b, 0xd4, 0x10, 0xa2, 0xb6, 0x60, 0x74, 0x5a, 0x86, 0xd6, 0x60,
	0x51, 0x80, 0x69, 0x64, 0x6d, 0xc8, 0x12, 0x7f, 0x30, 0x14, 0x72, 0x28, 0x96, 0xcf, 0x0a, 0xa2,
	0xa2, 0x84, 0xd0, 0x01, 0x59, 0xfb, 0x33, 0xb4, 0x58, 0x41, 0x63, 0xd3, 0xaa, 0xee, 0x9d, 0xe5,
	0xbb, 0x7f, 0xce, 0x3a, 0x1e, 0xdd, 0xbc, 0xfe, 0x6b, 0xce, 0xb3, 0xdf, 0x3e, 0x0b, 0xae, 0xfa,
	0x88, 0x9f, 0xb2, 0x09, 0x2a, 0xe1, 0xad, 0xbc, 0xf5, 0x71, 0x3c, 0x41, 0xc5, 0x1f, 0xb1, 0xd9,
	0x35, 0x65, 0x09, 0x2a, 0x31, 0x71, 0xd9, 0xf4, 0x9a, 0xb2, 0x8d, 0xe2, 0xcf, 0xd9, 0x09, 0xea,
	0x1b, 0xca, 0x53, 0x8b, 0xa4, 0xbb, 0xd6, 0x5f, 0x79, 0x6b, 0x3f, 0xfe, 0xff, 0x10, 0x6e, 0x14,
	0xbf, 0x64, 0x41, 0x6e, 0x20, 0xb5, 0xa0, 0xc4, 0xd1, 0xca, 0x5b, 0x2f, 0xce, 0x97, 0xb2, 0xdf,
	0x8a, 0x1c, 0xb6, 0x22, 0xaf, 0x86, 0xad, 0xc4, 0x03, 0xca, 0x5f, 0xb0, 0x53, 0x32, 0x0a, 0x4c,
	0x82, 0x3a, 0xc9, 0x52, 0x9b, 0x6f, 0x45, 0xd0, 0xff, 0xdb, 0xa5, 0x1b, 0x1d, 0x75, 0x19, 0x7f,
	0xc8, 0xa6, 0x16, 0x6d, 0x09, 0x62, 0xda, 0x2f, 0xcb, 0x0d, 0xf8, 0x7d, 0xe6, 0xb7, 0xa6, 0x14,
	0x33, 0x97, 0x75, 0x9f, 0xfc, 0x25, 0xbb, 0x07, 0x15, 0x5a, 0x0b, 0x2a, 0xc9, 0x76, 0x49, 0xdb,
	0x80, 0x11, 0x73, 0xd7, 0x9e, 0xec, 0xe3, 0x68, 0xf7, 0xbd, 0x01, 0xc3, 0x5f, 0xb1, 0xa3, 0xdc,
	0x90, 0x16, 0x6b, 0xb7, 0xd0, 0xc7, 0xf2, 0xf6, 0xbc, 0xe4, 0x47, 0x43, 0x7a, 0x7f, 0x3a, 0x5f,
	0xfe, 0x8b, 0x1d, 0xd5, 0xd1, 0x9a, 0xa8, 0x16, 0xe7, 0x77, 0xe8, 0x6f, 0x44, 0xf5, 0x88, 0xee,
	0x28, 0xfe, 0x86, 0x05, 0x05, 0x5a, 0x2c, 0xa1, 0x11, 0x17, 0x4e, 0x78, 0x32, 0x12, 0x3e, 0xf7,
	0xcd, 0xc1, 0x19, 0x58, 0xfe, 0x81, 0x2d, 0xb2, 0x16, 0x4b, 0x95, 0xb5, 0xf9, 0x2f, 0xb0, 0xe2,
	0xd2, 0xa9, 0x4f, 0x47, 0x6a, 0x74, 0x68, 0x0f, 0xfa, 0xd8, 0x89, 0x8e, 0x59, 0x50, 0xa7, 0xbb,
	0x92, 0x52, 0x75, 0xf6, 0x89, 0x2d, 0xf6, 0xd0, 0x57, 0x6c, 0x2c, 0x7f, 0xcb, 0xe6, 0xc3, 0x2b,
	0x10, 0xde, 0xca, 0x77, 0x97, 0x73, 0xe7, 0xe9, 0xc9, 0xbd, 0x11, 0xdf, 0xb2, 0x11, 0xfb, 0x31,
	0x1f, 0xb0, 0x6c, 0xe6, 0xae, 0xf1, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0xfb, 0x3a,
	0x6f, 0xf6, 0x02, 0x00, 0x00,
}
