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
	"sort"
	"sync"

	"github.com/golang/protobuf/proto"

	"go.starlark.net/starlark"
)

type reflectionDirection bool

const (
	reflectToProto   reflectionDirection = false
	reflectFromProto reflectionDirection = true
)

var typeRegistry struct {
	m     sync.RWMutex
	types map[reflect.Type]*MessageType
}

// MessageType contains information about the structure of a proto message.
//
// It is extracted via reflection from a proto message struct type.
type MessageType struct {
	name       string               // fully qualified proto message name
	ptr        reflect.Type         // ~ *Struct{}
	fields     map[string]fieldDesc // keyed by proto field name
	fieldNames []string             // sorted list of keys in 'fields'
}

// fieldDesc holds type information about some proto field in a message.
type fieldDesc struct {
	typ reflect.Type // type of the value

	// Defaultable specifies whether it is OK to auto-instantiate a default
	// value for a field when reading it for the first time.
	//
	// If defaultable is true, an unset field will be set on a default value of
	// the corresponding type when reading it for the first time. On the next
	// read, an existing value will be returned.
	//
	// If defaultable is false, reading an unset field results in None. Callers
	// will have to set it themselves. This is used by oneof implementation.
	defaultable bool

	// onChanged is called each time the field is assigned to by the starlark
	// code.
	//
	// It receives a dict with currently defined fields of the parent message. It
	// can do whatever it wants to this dict. This is used by oneof fields to
	// clear unpicked alternatives.
	onChanged func(dict starlark.StringDict)

	// onProtoReflection is called when the starlark message representation is
	// reflected to or from a proto struct (depending on a value of 'dir').
	//
	// It takes a reference to the proto struct (of reflect.Struct kind, NOT a
	// pointer) and returns a reference to the reflect.Value (of type 'typ') to
	// write or read the actual data to/from.
	//
	// When reflecting to proto (dir == reflectToProto), may have side effects:
	// the callback is allowed to modify the proto struct in a way it sees fit to
	// get the assignable value. This is important for oneof fields that require
	// instantiation of a helper wrapping struct to get a reference to an
	// assignable field.
	//
	// When reflecting from proto (dir == reflectFromProto), may return zero value
	// (use IsValid to check) if the given field must not be used to read data.
	// Again, this is important for alternative oneof fields to indicate that some
	// other alternative is realized and this field must be skipped.
	onProtoReflection func(msg reflect.Value, dir reflectionDirection) reflect.Value
}

// GetMessageType extracts type description for protobuf message of given type.
//
// 'typ' is expected to represent a pointer to a protobuf struct, as returned
// by proto.MessageType(...). Returns an error otherwise.
func GetMessageType(typ reflect.Type) (*MessageType, error) {
	typeRegistry.m.RLock()
	cached := typeRegistry.types[typ]
	typeRegistry.m.RUnlock()
	if cached != nil {
		return cached, nil
	}

	zero := reflect.Zero(typ) // (*Struct)(nil)
	name := proto.MessageName(zero.Interface().(proto.Message))
	if name == "" {
		return nil, fmt.Errorf("%q is not a registered proto message type", typ.Name())
	}

	typeRegistry.m.Lock()
	defer typeRegistry.m.Unlock()
	if typeRegistry.types == nil {
		typeRegistry.types = map[reflect.Type]*MessageType{}
	}

	strct := typ.Elem() // Struct{}
	props := proto.GetProperties(strct)

	fields := map[string]fieldDesc{}
	names := []string{}

	// Handle non-oneof fields.
	for _, prop := range props.Prop {
		prop := prop
		// Oneof wrapper types generated by Go protoc have no tags and should not
		// appear in the list of message fields. Oneofs are handled below based on
		// props.OneofTypes map.
		if prop.Tag == 0 {
			continue
		}
		f, ok := strct.FieldByName(prop.Name)
		if !ok {
			panic("property from proto.GetProperties is not defined in the go struct")
		}
		fields[prop.OrigName] = fieldDesc{
			typ:         f.Type,
			defaultable: true,
			onProtoReflection: func(msg reflect.Value, _ reflectionDirection) reflect.Value {
				return msg.FieldByName(prop.Name)
			},
		}
		names = append(names, prop.OrigName)
	}

	// Oneof fields are very special.
	//
	// Imagine we have the following message:
	//
	// message Msg {
	//   ...
	//   oneof wrapper {
	//     A field_a = 1;
	//     B field_b = 2;
	//   }
	// }
	//
	// Protoc will generate the following types for it:
	//
	// type Msg struct {
	//    ...
	//    Wrapper isMsg_Wrapper // assume index of this field in Msg struct is 7
	// }
	//
	// type Msg_FieldA struct {
	//    FieldA *A
	// }
	//
	// type Msg_FieldB struct {
	//    FieldB *B
	// }
	//
	// (where *Msg_FieldA and *Msg_FieldB implement isMsg_Wrapper interface)
	//
	// When examining this via 'props' we'll see 'wrapper' listed among
	// props.Prop. We skip it completely since we want to expose only 'a' and 'b'.
	//
	// Next we'll see props.OneofTypes to be a map:
	//
	// 'field_a': {
	//   Type: *testprotos.Msg_FieldA
	//   Field: 7
	//   Prop: <useless>
	// },
	// 'field_b': {
	//   Type: *testprotos.Msg_FieldB
	//   Field: 7
	//   Prop: <useless>
	// }
	//
	// (where 7, as mentioned above, is an index of Wrapper field in Msg go
	// struct).
	//
	// Important details here:
	//   * Names of the Go fields inside Msg_* structs are not mentioned anywhere,
	//     e.g. 'FieldA' is not mentioned. We workaround this by asserting that
	//     Msg_* structs have only one field and then just referring to this field
	//     by index 0.
	//   * Wrapper field index (7), can be used to associate a bunch of oneof
	//     alternatives with each other (e.g. when setting one, clear another).
	//     So we'll build a reverse map {7 -> ['field_a', 'field_b']} for this.
	oneofGroups := map[int]*oneofGroup{} // keyed by the field index
	for name, desc := range props.OneofTypes {
		ofg := oneofGroups[desc.Field]
		if ofg == nil {
			ofg = &oneofGroup{
				fieldIdx: desc.Field,
			}
			oneofGroups[desc.Field] = ofg
		}
		ofg.registerAlternative(name, desc)
	}
	for _, ofg := range oneofGroups {
		for name, fd := range ofg.fieldDescs() {
			fields[name] = fd
			names = append(names, name)
		}
	}

	sort.Strings(names)

	newTyp := &MessageType{
		name:       name,
		ptr:        typ,
		fields:     fields,
		fieldNames: names,
	}
	typeRegistry.types[typ] = newTyp
	return newTyp, nil
}

// Name returns fully qualified proto message name.
func (m *MessageType) Name() string {
	return m.name
}

// Type returns proto message type (pointer to a proto struct).
func (m *MessageType) Type() reflect.Type {
	return m.ptr
}

// NewProtoMessage constructs &ProtoMessage{} and returns it as reflect.Value.
func (m *MessageType) NewProtoMessage() reflect.Value {
	return reflect.New(m.ptr.Elem())
}

////////////////////////////////////////////////////////////////////////////////
// Oneof implementation helpers.

// oneofGroup is a set of fields belonging to one oneof definition.
//
// No more than one of them can be set at any time.
type oneofGroup struct {
	fieldIdx     int                 // index of the wrapping field in the message struct
	alternatives []*oneofAlternative // all possible alternatives the field can be set to
}

// oneofAlternative represents one possible field of a oneof group.
type oneofAlternative struct {
	name     string       // e.g. "field_a" for the example above
	outerTyp reflect.Type // e.g. *testprotos.Msg_FieldA
	innerTyp reflect.Type // e.g. *A
}

func (g *oneofGroup) registerAlternative(name string, p *proto.OneofProperties) {
	// Given *Msg_FieldA{FieldA: *A} grab *A as innerTyp.
	strct := p.Type.Elem()
	if strct.NumField() != 1 {
		panic(fmt.Sprintf("expecting %s to have one and only one field", p.Type))
	}
	g.alternatives = append(g.alternatives, &oneofAlternative{
		name:     name,
		outerTyp: p.Type,
		innerTyp: strct.Field(0).Type,
	})
}

func (g *oneofGroup) fieldDescs() map[string]fieldDesc {
	out := make(map[string]fieldDesc, len(g.alternatives))
	for _, alt := range g.alternatives {
		alt := alt
		out[alt.name] = fieldDesc{
			typ:         alt.innerTyp,
			defaultable: false, // we don't want getters to implicitly pick alternatives
			onChanged: func(dict starlark.StringDict) {
				g.onAlternativePicked(alt, dict)
			},
			onProtoReflection: func(msg reflect.Value, dir reflectionDirection) reflect.Value {
				if dir == reflectToProto {
					return g.onWritingAlternative(alt, msg)
				}
				return g.onReadingAlternative(alt, msg)
			},
		}
	}
	return out
}

// onAlternativePicked is called when one oneof alternative was set to.
//
// It clears all others.
func (g *oneofGroup) onAlternativePicked(alt *oneofAlternative, dict starlark.StringDict) {
	for _, a := range g.alternatives {
		if a != alt {
			delete(dict, a.name)
		}
	}
}

// onWritingAlternative is called when writing a value to the proto message.
//
// It instantiates the oneof wrapping struct of a correct type (based on the
// picked alternative) and returns a reference to the value to be set there.
func (g *oneofGroup) onWritingAlternative(alt *oneofAlternative, msg reflect.Value) reflect.Value {
	wrap := reflect.New(alt.outerTyp.Elem()) // ~ wrap = &testprotos.Msg_FieldA{}
	msg.Field(g.fieldIdx).Set(wrap)          // ~ msg.Wrapper = wrap
	return wrap.Elem().Field(0)              // ~ ref to wrap.A, to be set to
}

// onReadingAlternative is called when reading a value from the proto message.
//
// If returns a valid reflect.Value if the given alternative is realized in
// the 'msg' or an invalid zero reflect.Value otherwise.
func (g *oneofGroup) onReadingAlternative(alt *oneofAlternative, msg reflect.Value) reflect.Value {
	switch wrap := msg.Field(g.fieldIdx); {
	case wrap.IsNil():
		return reflect.Value{} // no alternatives are set
	case wrap.Elem().Type() != alt.outerTyp:
		return reflect.Value{} // some other alternative is set
	default:
		return wrap.Elem().Elem().Field(0) // ~ (*msg.Wrapper).A
	}
}
