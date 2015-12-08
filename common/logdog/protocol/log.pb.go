// Code generated by protoc-gen-go.
// source: log.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	log.proto
	butler.proto

It has these top-level messages:
	LogStreamDescriptor
	Text
	Binary
	Datagram
	LogEntry
	LogIndex
	ButlerMetadata
	ButlerLogBundle
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"
import google_protobuf1 "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The log stream's content type (required).
type LogStreamDescriptor_StreamType int32

const (
	LogStreamDescriptor_TEXT     LogStreamDescriptor_StreamType = 0
	LogStreamDescriptor_BINARY   LogStreamDescriptor_StreamType = 1
	LogStreamDescriptor_DATAGRAM LogStreamDescriptor_StreamType = 2
)

var LogStreamDescriptor_StreamType_name = map[int32]string{
	0: "TEXT",
	1: "BINARY",
	2: "DATAGRAM",
}
var LogStreamDescriptor_StreamType_value = map[string]int32{
	"TEXT":     0,
	"BINARY":   1,
	"DATAGRAM": 2,
}

func (x LogStreamDescriptor_StreamType) String() string {
	return proto.EnumName(LogStreamDescriptor_StreamType_name, int32(x))
}
func (LogStreamDescriptor_StreamType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

// *
// Log stream descriptor data. This is the full set of information that
// describes a logging stream.
type LogStreamDescriptor struct {
	//
	// The stream's prefix (required).
	//
	// Logs originating from the same Butler instance will share a Prefix.
	//
	// A valid prefix value is a StreamName described in:
	// https://github.com/luci/luci-go/common/logdog/types
	Prefix string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	//
	// The log stream's name (required).
	//
	// This is used to uniquely identify a log stream within the scope of its
	// prefix.
	//
	// A valid name value is a StreamName described in:
	// https://github.com/luci/luci-go/common/logdog/types
	Name       string                         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	StreamType LogStreamDescriptor_StreamType `protobuf:"varint,3,opt,name=stream_type,enum=protocol.LogStreamDescriptor_StreamType" json:"stream_type,omitempty"`
	//
	// The stream's content type (required).
	//
	// This must be an HTTP Content-Type value. It is made available to LogDog
	// clients when querying stream metadata. It will also be applied to archived
	// binary log data.
	ContentType string `protobuf:"bytes,4,opt,name=content_type" json:"content_type,omitempty"`
	//
	// The log stream's base timestamp (required).
	//
	// This notes the start time of the log stream. All LogEntries express their
	// timestamp as microsecond offsets from this field.
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
	// The set of associated log tags.
	Tags []*LogStreamDescriptor_Tag `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
	//
	// If set, the stream will be joined together during archival to recreate the
	// original stream and made available at <prefix>/+/<name>.ext.
	BinaryFileExt string `protobuf:"bytes,7,opt,name=binary_file_ext" json:"binary_file_ext,omitempty"`
}

func (m *LogStreamDescriptor) Reset()                    { *m = LogStreamDescriptor{} }
func (m *LogStreamDescriptor) String() string            { return proto.CompactTextString(m) }
func (*LogStreamDescriptor) ProtoMessage()               {}
func (*LogStreamDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LogStreamDescriptor) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogStreamDescriptor) GetTags() []*LogStreamDescriptor_Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

//
// Tag is an arbitrary key/value tag associated with this log stream.
//
// LogDog clients can query for log streams based on tag values.
type LogStreamDescriptor_Tag struct {
	// The tag key (required).
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The tag value.
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LogStreamDescriptor_Tag) Reset()                    { *m = LogStreamDescriptor_Tag{} }
func (m *LogStreamDescriptor_Tag) String() string            { return proto.CompactTextString(m) }
func (*LogStreamDescriptor_Tag) ProtoMessage()               {}
func (*LogStreamDescriptor_Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Text stream content.
type Text struct {
	Lines []*Text_Line `protobuf:"bytes,1,rep,name=lines" json:"lines,omitempty"`
}

func (m *Text) Reset()                    { *m = Text{} }
func (m *Text) String() string            { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()               {}
func (*Text) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Text) GetLines() []*Text_Line {
	if m != nil {
		return m.Lines
	}
	return nil
}

// Contiguous text lines and their delimiters.
type Text_Line struct {
	// The line's text content, not including its delimiter.
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	//
	// The line's delimiter string.
	//
	// If this is an empty string, this line is continued in the next sequential
	// line, and the line's sequence number does not advance.
	Delimiter string `protobuf:"bytes,2,opt,name=delimiter" json:"delimiter,omitempty"`
}

func (m *Text_Line) Reset()                    { *m = Text_Line{} }
func (m *Text_Line) String() string            { return proto.CompactTextString(m) }
func (*Text_Line) ProtoMessage()               {}
func (*Text_Line) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

// Binary stream content.
type Binary struct {
	// The byte offset in the stream of the first byte of data.
	Offset uint64 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	// The binary stream's data.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Binary) Reset()                    { *m = Binary{} }
func (m *Binary) String() string            { return proto.CompactTextString(m) }
func (*Binary) ProtoMessage()               {}
func (*Binary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Datagram stream content type.
type Datagram struct {
	// This datagram data.
	Data    []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Partial *Datagram_Partial `protobuf:"bytes,2,opt,name=partial" json:"partial,omitempty"`
}

func (m *Datagram) Reset()                    { *m = Datagram{} }
func (m *Datagram) String() string            { return proto.CompactTextString(m) }
func (*Datagram) ProtoMessage()               {}
func (*Datagram) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Datagram) GetPartial() *Datagram_Partial {
	if m != nil {
		return m.Partial
	}
	return nil
}

//
// If this is not a partial datagram, this field will include reassembly and
// state details for the full datagram.
type Datagram_Partial struct {
	//
	// The index, starting with zero, of this datagram fragment in the full
	// datagram.
	Index uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// The size of the full datagram
	Size uint64 `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	// If true, this is the last partial datagram in the overall datagram.
	Last bool `protobuf:"varint,3,opt,name=last" json:"last,omitempty"`
}

func (m *Datagram_Partial) Reset()                    { *m = Datagram_Partial{} }
func (m *Datagram_Partial) String() string            { return proto.CompactTextString(m) }
func (*Datagram_Partial) ProtoMessage()               {}
func (*Datagram_Partial) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// *
// An individual log entry.
//
// This contains the superset of transmissible log data. Its content fields
// should be interpreted in the context of the log stream's content type.
type LogEntry struct {
	//
	// The stream time offset for this content.
	//
	// This offset is added to the log stream's base "timestamp" to resolve the
	// timestamp for this specific Content.
	TimeOffset *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=time_offset" json:"time_offset,omitempty"`
	//
	// The message index within the Prefix (required).
	//
	// This is value is unique to this LogEntry across the entire set of entries
	// sharing the stream's Prefix. It is used to designate unambiguous log
	// ordering.
	PrefixIndex uint64 `protobuf:"varint,2,opt,name=prefix_index" json:"prefix_index,omitempty"`
	//
	// The message index within its Stream (required).
	//
	// This value is unique across all entries sharing the same Prefix and Stream
	// Name. It is used to designate unambiguous log ordering within the stream.
	StreamIndex uint64 `protobuf:"varint,3,opt,name=stream_index" json:"stream_index,omitempty"`
	//
	// The sequence number of the first content entry in this LogEntry.
	//
	// Text: This is the line index of the first included line. Line indices begin
	//     at zero.
	// Binary: This is the byte offset of the first byte in the included data.
	// Datagram: This is the index of the datagram. The first datagram has index
	//     zero.
	Sequence uint64 `protobuf:"varint,4,opt,name=sequence" json:"sequence,omitempty"`
	//
	// The content of the message. The field that is populated here must
	// match the log's `stream_type`.
	//
	// Types that are valid to be assigned to Content:
	//	*LogEntry_Text
	//	*LogEntry_Binary
	//	*LogEntry_Datagram
	Content isLogEntry_Content `protobuf_oneof:"content"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isLogEntry_Content interface {
	isLogEntry_Content()
}

type LogEntry_Text struct {
	Text *Text `protobuf:"bytes,10,opt,name=text,oneof"`
}
type LogEntry_Binary struct {
	Binary *Binary `protobuf:"bytes,11,opt,name=binary,oneof"`
}
type LogEntry_Datagram struct {
	Datagram *Datagram `protobuf:"bytes,12,opt,name=datagram,oneof"`
}

func (*LogEntry_Text) isLogEntry_Content()     {}
func (*LogEntry_Binary) isLogEntry_Content()   {}
func (*LogEntry_Datagram) isLogEntry_Content() {}

func (m *LogEntry) GetContent() isLogEntry_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *LogEntry) GetTimeOffset() *google_protobuf1.Duration {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func (m *LogEntry) GetText() *Text {
	if x, ok := m.GetContent().(*LogEntry_Text); ok {
		return x.Text
	}
	return nil
}

func (m *LogEntry) GetBinary() *Binary {
	if x, ok := m.GetContent().(*LogEntry_Binary); ok {
		return x.Binary
	}
	return nil
}

func (m *LogEntry) GetDatagram() *Datagram {
	if x, ok := m.GetContent().(*LogEntry_Datagram); ok {
		return x.Datagram
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, []interface{}{
		(*LogEntry_Text)(nil),
		(*LogEntry_Binary)(nil),
		(*LogEntry_Datagram)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// content
	switch x := m.Content.(type) {
	case *LogEntry_Text:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Text); err != nil {
			return err
		}
	case *LogEntry_Binary:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Binary); err != nil {
			return err
		}
	case *LogEntry_Datagram:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Datagram); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.Content has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 10: // content.text
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Text)
		err := b.DecodeMessage(msg)
		m.Content = &LogEntry_Text{msg}
		return true, err
	case 11: // content.binary
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Binary)
		err := b.DecodeMessage(msg)
		m.Content = &LogEntry_Binary{msg}
		return true, err
	case 12: // content.datagram
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Datagram)
		err := b.DecodeMessage(msg)
		m.Content = &LogEntry_Datagram{msg}
		return true, err
	default:
		return false, nil
	}
}

// *
// LogIndex is an index into an at-rest log storage.
//
// The log stream and log index are generated by the Archivist during archival.
//
// An archived log stream is a series of contiguous LogEntry frames. The index
// maps a log's logical logation in its stream, prefix, and timeline to its
// frame's binary offset in the archived log stream blob.
type LogIndex struct {
	//
	// The LogStreamDescriptor for this log stream (required).
	//
	// The index stores the stream's LogStreamDescriptor so that a client can
	// know the full set of log metadata by downloading its index.
	Desc *LogStreamDescriptor `protobuf:"bytes,1,opt,name=desc" json:"desc,omitempty"`
	//
	// A series of ascending-ordered Entry messages representing snapshots of an
	// archived log stream.
	//
	// Within this set of Entry messages, the "offset", "prefix_index",
	// "stream_index", and "time_offset" fields will be ascending.
	//
	// The frequency of Entry messages is not defined; it is up to the Archivist
	// process to choose a frequency.
	Entries []*LogIndex_Entry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *LogIndex) Reset()                    { *m = LogIndex{} }
func (m *LogIndex) String() string            { return proto.CompactTextString(m) }
func (*LogIndex) ProtoMessage()               {}
func (*LogIndex) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *LogIndex) GetDesc() *LogStreamDescriptor {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *LogIndex) GetEntries() []*LogIndex_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

//
// Entry is a single index entry.
//
// The index is composed of a series of entries, each corresponding to a
// sequential snapshot of of the log stream.
type LogIndex_Entry struct {
	//
	// The sequence number of the first content entry.
	//
	// Text: This is the line index of the first included line. Line indices
	//     begin at zero.
	// Binary: This is the byte offset of the first byte in the included data.
	// Datagram: This is the index of the datagram. The first datagram has index
	//     zero.
	Sequence uint64 `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	//
	// The log index that this entry describes (required).
	//
	// This is used by clients to identify a specific LogEntry within a set of
	// streams sharing a Prefix.
	PrefixIndex uint64 `protobuf:"varint,2,opt,name=prefix_index" json:"prefix_index,omitempty"`
	//
	// The time offset of this log entry (required).
	//
	// This is used by clients to identify a specific LogEntry within a log
	// stream.
	StreamIndex uint64 `protobuf:"varint,3,opt,name=stream_index" json:"stream_index,omitempty"`
	//
	// The time offset of this log entry, in microseconds.
	//
	// This is added to the descriptor's "timestamp" field to identify the
	// specific timestamp of this log. It is used by clients to identify a
	// specific LogEntry by time.
	TimeOffset *google_protobuf1.Duration `protobuf:"bytes,4,opt,name=time_offset" json:"time_offset,omitempty"`
}

func (m *LogIndex_Entry) Reset()                    { *m = LogIndex_Entry{} }
func (m *LogIndex_Entry) String() string            { return proto.CompactTextString(m) }
func (*LogIndex_Entry) ProtoMessage()               {}
func (*LogIndex_Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *LogIndex_Entry) GetTimeOffset() *google_protobuf1.Duration {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func init() {
	proto.RegisterType((*LogStreamDescriptor)(nil), "protocol.LogStreamDescriptor")
	proto.RegisterType((*LogStreamDescriptor_Tag)(nil), "protocol.LogStreamDescriptor.Tag")
	proto.RegisterType((*Text)(nil), "protocol.Text")
	proto.RegisterType((*Text_Line)(nil), "protocol.Text.Line")
	proto.RegisterType((*Binary)(nil), "protocol.Binary")
	proto.RegisterType((*Datagram)(nil), "protocol.Datagram")
	proto.RegisterType((*Datagram_Partial)(nil), "protocol.Datagram.Partial")
	proto.RegisterType((*LogEntry)(nil), "protocol.LogEntry")
	proto.RegisterType((*LogIndex)(nil), "protocol.LogIndex")
	proto.RegisterType((*LogIndex_Entry)(nil), "protocol.LogIndex.Entry")
	proto.RegisterEnum("protocol.LogStreamDescriptor_StreamType", LogStreamDescriptor_StreamType_name, LogStreamDescriptor_StreamType_value)
}

var fileDescriptor0 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x6d, 0x12, 0x27, 0x71, 0xc6, 0x69, 0xbe, 0x7c, 0x5b, 0x24, 0x8c, 0x25, 0x4a, 0xeb, 0x8b,
	0x2a, 0x08, 0xe1, 0xa0, 0xc0, 0x2d, 0x17, 0xa9, 0x52, 0xd1, 0x4a, 0x05, 0xa1, 0xe2, 0x0b, 0xb8,
	0x8a, 0x36, 0xf1, 0xc6, 0x5a, 0xe1, 0x3f, 0xd6, 0x1b, 0xd4, 0xc0, 0x7b, 0x22, 0xf1, 0x14, 0xbc,
	0x02, 0xb3, 0xbb, 0xce, 0x4f, 0x5b, 0x54, 0xb8, 0x4a, 0x3c, 0x73, 0x76, 0xe6, 0xcc, 0x39, 0x07,
	0x3a, 0x49, 0x1e, 0x07, 0x85, 0xc8, 0x65, 0x4e, 0x6c, 0xfd, 0x33, 0xcf, 0x13, 0xef, 0x49, 0x9c,
	0xe7, 0x71, 0xc2, 0x86, 0xba, 0x30, 0x5b, 0x2e, 0x86, 0x92, 0xa7, 0xac, 0x94, 0x34, 0x2d, 0x0c,
	0xd4, 0x3b, 0xbc, 0x0d, 0x88, 0x96, 0x82, 0x4a, 0x9e, 0x67, 0xa6, 0xef, 0xff, 0xa8, 0xc3, 0xc1,
	0x65, 0x1e, 0x7f, 0x90, 0x82, 0xd1, 0x74, 0xc2, 0xca, 0xb9, 0xe0, 0x85, 0xcc, 0x05, 0xe9, 0x41,
	0xab, 0x10, 0x6c, 0xc1, 0xaf, 0xdd, 0xda, 0x51, 0x6d, 0xd0, 0x21, 0x5d, 0xb0, 0x32, 0x9a, 0x32,
	0xb7, 0xae, 0xbf, 0x5e, 0x83, 0x53, 0xea, 0x17, 0x53, 0xb9, 0x2a, 0x98, 0xdb, 0xc0, 0x62, 0x6f,
	0x34, 0x08, 0xd6, 0xb4, 0x82, 0x3f, 0x4c, 0x0c, 0x4c, 0x21, 0x44, 0x3c, 0x79, 0x00, 0xdd, 0x79,
	0x9e, 0x49, 0x96, 0x49, 0xf3, 0xde, 0xd2, 0x43, 0x9f, 0x43, 0x67, 0xc3, 0xde, 0x6d, 0x62, 0xc9,
	0x19, 0x79, 0x81, 0xa1, 0x1f, 0xac, 0xe9, 0x07, 0xe1, 0x1a, 0x41, 0x86, 0x60, 0x49, 0x1a, 0x97,
	0x6e, 0xeb, 0xa8, 0x81, 0xc8, 0xe3, 0xfb, 0x97, 0x87, 0x34, 0x26, 0x0f, 0xe1, 0xbf, 0x19, 0xcf,
	0xa8, 0x58, 0x4d, 0x17, 0x3c, 0x61, 0x53, 0x76, 0x2d, 0xdd, 0xb6, 0x5a, 0xec, 0x1d, 0x43, 0x43,
	0xf5, 0x1d, 0x68, 0x7c, 0x66, 0xab, 0xea, 0xde, 0x7d, 0x68, 0x7e, 0xa5, 0xc9, 0xb2, 0x3a, 0xd8,
	0x7f, 0x01, 0xb0, 0xc3, 0xdf, 0x06, 0x2b, 0x3c, 0xfb, 0x18, 0xf6, 0xf7, 0x08, 0x40, 0xeb, 0xf4,
	0xe2, 0xdd, 0xf8, 0xea, 0x53, 0xbf, 0x86, 0x12, 0xd9, 0x93, 0x71, 0x38, 0x7e, 0x73, 0x35, 0x7e,
	0xdb, 0xaf, 0xfb, 0x21, 0x62, 0x70, 0x05, 0xf1, 0xa1, 0x99, 0xf0, 0x8c, 0x95, 0x38, 0x57, 0xf1,
	0x3c, 0xd8, 0xf2, 0x54, 0xed, 0xe0, 0x12, 0x7b, 0xde, 0x00, 0x2c, 0xf5, 0xbb, 0x5d, 0x6a, 0x38,
	0xfc, 0x0f, 0x9d, 0x88, 0x25, 0x3c, 0xe5, 0x92, 0x89, 0x8a, 0xc7, 0x09, 0xee, 0xd3, 0x37, 0x28,
	0x83, 0xf2, 0xc5, 0xa2, 0x64, 0x52, 0x83, 0x2d, 0x65, 0x50, 0x44, 0x25, 0xd5, 0xb8, 0xae, 0xff,
	0x1d, 0xb9, 0xe0, 0x57, 0x2c, 0x68, 0xba, 0xe9, 0x28, 0x5c, 0x97, 0x3c, 0x83, 0x76, 0x41, 0x85,
	0xe4, 0x34, 0xd1, 0x50, 0xa5, 0xf1, 0x86, 0xd1, 0xfa, 0x49, 0xf0, 0xde, 0x20, 0xbc, 0x57, 0xd0,
	0xae, 0xfe, 0x2a, 0x6e, 0x3c, 0x8b, 0x98, 0xc9, 0xc3, 0xbe, 0x1a, 0x5a, 0xf2, 0x6f, 0x46, 0x1e,
	0xbd, 0x3c, 0xa1, 0xa5, 0xd4, 0x41, 0xb0, 0xfd, 0x5f, 0x35, 0xb0, 0xd1, 0x84, 0xb3, 0x4c, 0x22,
	0xcf, 0x00, 0x1c, 0xe5, 0xea, 0x74, 0x87, 0xac, 0x33, 0x7a, 0x74, 0xc7, 0xd7, 0x49, 0x15, 0x4b,
	0x95, 0x0d, 0x13, 0xbc, 0xa9, 0x59, 0x67, 0x16, 0x60, 0xb5, 0x0a, 0x9c, 0xa9, 0x36, 0x74, 0xb5,
	0x0f, 0x76, 0xc9, 0xbe, 0x2c, 0x59, 0x36, 0x37, 0x19, 0xb2, 0xc8, 0x21, 0x86, 0x42, 0x19, 0x0b,
	0x7a, 0x4d, 0xef, 0xa6, 0xd8, 0xe7, 0x7b, 0xe8, 0x46, 0xcb, 0x64, 0xc0, 0x75, 0x34, 0xa2, 0xbf,
	0x45, 0x18, 0x5d, 0x11, 0x73, 0x02, 0x76, 0x54, 0x09, 0xe1, 0x76, 0x35, 0x8a, 0xdc, 0x95, 0xe8,
	0x7c, 0xef, 0xb4, 0x03, 0xed, 0x2a, 0xc5, 0xfe, 0x4f, 0x73, 0xf1, 0x85, 0xe2, 0x86, 0x0a, 0x5b,
	0x11, 0x26, 0xaf, 0x3a, 0xf5, 0xf1, 0xbd, 0xc1, 0x24, 0x4f, 0xa1, 0x8d, 0x03, 0x04, 0xc7, 0x80,
	0xd4, 0x75, 0x40, 0xdc, 0x1b, 0x78, 0x3d, 0x31, 0xd0, 0x4a, 0x7a, 0x4b, 0x68, 0x1a, 0x49, 0x77,
	0xcf, 0xae, 0xad, 0xe5, 0xf9, 0x67, 0xd1, 0x6e, 0x19, 0x62, 0xfd, 0xc5, 0x90, 0x59, 0x4b, 0x97,
	0x5e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x4a, 0x98, 0xa2, 0x80, 0x04, 0x00, 0x00,
}
