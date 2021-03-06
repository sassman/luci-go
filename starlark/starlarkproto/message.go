// Copyright 2018 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package starlarkproto

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"

	"go.starlark.net/starlark"
)

// Message is a Starlark value that implements a struct-like type structured
// like a protobuf message.
//
// Implements starlark.Value, starlark.HasAttrs and starlark.HasSetField
// interfaces.
//
// TODO(vadimsh): Currently not safe for a cross-goroutine use without external
// locking, even when frozen.
type Message struct {
	typ    *MessageType        // type information
	fields starlark.StringDict // populated fields, keyed by proto field name
	frozen bool                // true after Freeze()
}

// NewMessage instantiates a new empty message of the given type.
func NewMessage(typ *MessageType) *Message {
	return &Message{
		typ:    typ,
		fields: starlark.StringDict{},
	}
}

// Public API used by the hosting environment.

// MessageType returns detailed type information about the message.
func (m *Message) MessageType() *MessageType { return m.typ }

// ToProto returns a new populated proto message of an appropriate type.
//
// Returns an error if the data inside the Starlark representation of
// the message has a wrong type.
func (m *Message) ToProto() (proto.Message, error) {
	ptr := m.typ.NewProtoMessage() // ~ ptr := &ProtoMessage{}
	msg := ptr.Elem()              // ~ msg := *ptr (a reference)

	for name, val := range m.fields {
		fd, ok := m.typ.fields[name]
		if !ok {
			panic("should not happen, SetField and Attr checks the structure already")
		}
		if err := assign(fd.onProtoReflection(msg, reflectToProto), val); err != nil {
			return nil, fmt.Errorf("bad value for field %q of %q - %s", name, m.Type(), err)
		}
	}

	return ptr.Interface().(proto.Message), nil
}

// FromProto populates fields of this message based on values in proto.Message.
//
// Returns an error on type mismatch.
func (m *Message) FromProto(p proto.Message) error {
	ptr := reflect.ValueOf(p)
	if ptr.Type() != m.typ.Type() {
		return fmt.Errorf("bad message type: got %s, expect %s", ptr.Type(), m.typ.Type())
	}

	msg := ptr.Elem()
	for name, fd := range m.typ.fields {
		// Get the field's value from the proto message as reflect.Value. For unused
		// oneof alternatives this is an invalid zero value, we skip them right
		// away. For other fields it is reflect.Value (of fd.typ type) that MAY be
		// nil inside (for unset fields). toStarlarkValue converts such values to
		// starlark.None.
		val := fd.onProtoReflection(msg, reflectFromProto)
		if !val.IsValid() {
			continue
		}
		// Convert the Go value to the corresponding Starlark value and assign it to
		// the field in 'm'.
		sv, err := toStarlarkValue(val)
		if err != nil {
			return fmt.Errorf("cannot recognize value of field %s: %s", name, err)
		}
		if err := m.SetField(name, sv); err != nil {
			return err
		}
	}

	return nil
}

// FromDict populates fields of this message based on values in starlark.Dict.
//
// Doesn't reset the message. Basically does this:
//
//   for k in d:
//     setattr(msg, k, d[k])
//
// Returns an error on type mismatch.
func (m *Message) FromDict(d *starlark.Dict) error {
	iter := d.Iterate()
	defer iter.Done()

	var k starlark.Value
	for iter.Next(&k) {
		key, ok := k.(starlark.String)
		if !ok {
			return fmt.Errorf("got %s dict key, expecting a string", k.Type())
		}
		v, _, _ := d.Get(k)
		if err := m.SetField(key.GoString(), v); err != nil {
			return err
		}
	}

	return nil
}

// Basic starlark.Value interface.

// String implements starlark.Value.
func (m *Message) String() string {
	msg, err := m.ToProto()
	if err != nil {
		return fmt.Sprintf("<!Bad %s: %s!>", m.Type(), err)
	}
	return msg.String()
}

// Type implements starlark.Value.
func (m *Message) Type() string {
	// The receiver is nil when doing type checks with starlark.UnpackArgs. It
	// asks the nil message for its type for the error message.
	if m == nil {
		return "proto.Message"
	}
	return m.typ.name
}

// Freeze implements starlark.Value.
func (m *Message) Freeze() {
	if !m.frozen {
		m.fields.Freeze()
		m.frozen = true
	}
}

// Truth implements starlark.Value.
func (m *Message) Truth() starlark.Bool { return starlark.True }

// Hash implements starlark.Value.
func (m *Message) Hash() (uint32, error) {
	return 0, fmt.Errorf("proto message %q is not hashable", m.Type())
}

// HasAttrs and HasSetField interfaces that make the message look like a struct.

// Attr implements starlark.HasAttrs.
func (m *Message) Attr(name string) (starlark.Value, error) {
	// The field was already set?
	val, ok := m.fields[name]
	if ok {
		return val, nil
	}

	// The field wasn't set, but it is defined by the proto schema? Need to
	// generate and return the default value then, except for oneof alternatives
	// that do not have defaults. This is needed to make sure callers are
	// explicitly picking a oneof alternative by assigning a value to it, rather
	// than have it picked implicitly be reading an attribute (which is weird).
	if fd, ok := m.typ.fields[name]; ok {
		if !fd.defaultable {
			return starlark.None, nil
		}
		def, err := newDefaultValue(fd.typ)
		if err != nil {
			return nil, err
		}
		// Lazy initialization of fields is an implementation detail. From the
		// caller's point of view, all fields had their default values even before
		// the object was frozen. Lazy-initialize the field, even if we are frozen,
		// but make sure it is frozen itself too.
		//
		// TODO(vadimsh): This is not thread safe and should be improved if a frozen
		// *Message is shared between goroutines. Generally frozen values are
		// assumed to be safe for cross-goroutine use, which is not the case here.
		// If this becomes important, we can force-initialize and freeze all default
		// fields in Freeze().
		if m.frozen {
			def.Freeze()
		}
		m.fields[name] = def
		return def, nil
	}

	return nil, fmt.Errorf("proto message %q has no field %q", m.Type(), name)
}

// AttrNames implements starlark.HasAttrs.
func (m *Message) AttrNames() []string {
	return m.typ.fieldNames
}

// SetField implements starlark.HasSetField.
func (m *Message) SetField(name string, val starlark.Value) error {
	fd, ok := m.typ.fields[name]
	if !ok {
		return fmt.Errorf("proto message %q has no field %q", m.Type(), name)
	}

	// Setting a field to None removes it completely.
	if val == starlark.None {
		if err := m.checkMutable(); err != nil {
			return err
		}
		delete(m.fields, name)
		return nil
	}

	// If assigning to a messaged-valued field (singular or repeated), recognize
	// dicts and Nones and use them to instantiate values (perhaps empty) of the
	// corresponding proto type. This allows to construct deeply nested protobuf
	// messages just by using lists, dicts and primitive values. Python does this
	// too.
	val, err := maybeMakeMessages(fd.typ, val)
	if err != nil {
		return fmt.Errorf("when constructing %q in proto %q - %s", name, m.Type(), err)
	}

	// Do a light type check. It doesn't "recurse" into lists or tuples. So it is
	// still possible to assign e.g. a list of strings to a "repeated int64"
	// field. This will be discovered later in ToProto when trying to construct
	// a proto message from Starlark values.
	if err := checkAssignable(fd.typ, val); err != nil {
		return fmt.Errorf("can't assign value of type %q to field %q in proto %q - %s", val.Type(), name, m.Type(), err)
	}
	if err := m.checkMutable(); err != nil {
		return err
	}
	m.fields[name] = val

	// onChanged hooks is used by oneof's to clear alternatives that weren't
	// picked.
	if fd.onChanged != nil {
		fd.onChanged(m.fields)
	}

	return nil
}

// checkMutable returns an error if the message is frozen.
func (m *Message) checkMutable() error {
	if m.frozen {
		return fmt.Errorf("cannot modify frozen proto message %q", m.Type())
	}
	return nil
}

// maybeMakeMessages recognizes when a dict is assigned to a message field or
// when a list or tuple of dicts or Nones is assigned to a repeated message
// field.
//
// It converts dicts or Nones to *Message of corresponding type using NewMessage
// and FromDict and returns them as Starlark values to use in place of passed
// value.
//
// Returns 'val' as is in other cases. Returns an error if given a dict, but
// it can't be used to initialize a message (e.g. has wrong schema).
func maybeMakeMessages(typ reflect.Type, val starlark.Value) (starlark.Value, error) {
	if dict, ok := val.(*starlark.Dict); ok && isProtoType(typ) {
		t, err := GetMessageType(typ)
		if err != nil {
			return nil, err
		}
		msg := NewMessage(t)
		return msg, msg.FromDict(dict)
	}

	if seq, ok := val.(starlark.Sequence); ok && typ.Kind() == reflect.Slice && isProtoType(typ.Elem()) && shouldMakeMessages(seq) {
		t, err := GetMessageType(typ.Elem())
		if err != nil {
			return nil, err
		}

		iter := seq.Iterate()
		defer iter.Done()

		out := make([]starlark.Value, 0, seq.Len())

		var v starlark.Value
		for iter.Next(&v) {
			switch val := v.(type) {
			case starlark.NoneType:
				out = append(out, NewMessage(t))
			case *starlark.Dict:
				msg := NewMessage(t)
				if err := msg.FromDict(val); err != nil {
					return nil, err
				}
				out = append(out, msg)
			default:
				out = append(out, v)
			}
		}

		// Note: we return a list (and not a tuple) so that the proto object remains
		// mutable. Callers might want to add more items there.
		return starlark.NewList(out), nil
	}

	return val, nil
}

// isProtoType is true if typ is &Struct{}.
func isProtoType(typ reflect.Type) bool {
	return typ.Kind() == reflect.Ptr && typ.Elem().Kind() == reflect.Struct
}

// shouldMakeMessages returns true if seq has at least one dict or None that
// should be converted to a proto message.
func shouldMakeMessages(seq starlark.Sequence) bool {
	iter := seq.Iterate()
	defer iter.Done()
	var v starlark.Value
	for iter.Next(&v) {
		if v == starlark.None {
			return true
		}
		if _, ok := v.(*starlark.Dict); ok {
			return true
		}
	}
	return false
}
