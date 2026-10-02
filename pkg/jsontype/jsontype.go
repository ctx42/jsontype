// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package jsontype preserves Go type information during JSON marshaling.
package jsontype

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/ctx42/convert/pkg/convert"
)

// registry is the package-level [Registry].
var registry = DefaultRegistry()

// Register registers a converter for the given type name in the
// package-level registry. It is the registry [json.Unmarshal] uses through
// [Value.UnmarshalJSON], and the default for [NewValue] and [FromMap]; it is
// safe for concurrent use. See [Registry.Register].
func Register(typ string, cnv convert.AnyToAny) convert.AnyToAny {
	return registry.Register(typ, cnv)
}

// List of type names supported by the package out of the box. [Byte] and
// [Rune] are accepted when unmarshalling, but [New] and [NewValue] never
// produce them, as Go reports byte and rune as uint8 and int32.
const (
	Int   = "int"
	Int16 = "int16"
	Int32 = "int32"
	Int64 = "int64"
	Int8  = "int8"

	Uint   = "uint"
	Uint16 = "uint16"
	Uint32 = "uint32"
	Uint64 = "uint64"
	Uint8  = "uint8"

	Float32 = "float32"
	Float64 = "float64"

	Byte     = "byte"
	Rune     = "rune"
	String   = "string"
	Bool     = "bool"
	Time     = "time.Time"
	Duration = "time.Duration"
	Nil      = "nil"
)

// DefaultRegistry returns a new [Registry] with converters for all the
// built-in type names.
func DefaultRegistry() *Registry {
	reg := NewRegistry()

	reg.registerNumber(
		Byte,
		numberConverter(convert.StringToByte, convert.Float64ToByte),
	)
	reg.registerNumber(
		Uint8,
		numberConverter(convert.StringToUint8, convert.Float64ToUint8),
	)
	reg.registerNumber(
		Uint16,
		numberConverter(convert.StringToUint16, convert.Float64ToUint16),
	)
	reg.registerNumber(
		Uint32,
		numberConverter(convert.StringToUint32, convert.Float64ToUint32),
	)
	reg.registerNumber(
		Uint64,
		numberConverter(convert.StringToUint64, convert.Float64ToUint64),
	)
	reg.registerNumber(
		Uint,
		numberConverter(convert.StringToUint, convert.Float64ToUint),
	)

	reg.registerNumber(
		Int8,
		numberConverter(convert.StringToInt8, convert.Float64ToInt8),
	)
	reg.registerNumber(
		Int16,
		numberConverter(convert.StringToInt16, convert.Float64ToInt16),
	)
	reg.registerNumber(
		Rune,
		numberConverter(convert.StringToRune, convert.Float64ToRune),
	)
	reg.registerNumber(
		Int32,
		numberConverter(convert.StringToInt32, convert.Float64ToInt32),
	)
	reg.registerNumber(
		Int64,
		numberConverter(convert.StringToInt64, convert.Float64ToInt64),
	)
	reg.registerNumber(
		Int,
		numberConverter(convert.StringToInt, convert.Float64ToInt),
	)

	reg.registerNumber(
		Float32,
		numberConverter(convert.StringToFloat32, convert.Float64ToFloat32),
	)
	reg.Register(Float64, convert.ToAnyAny(convert.Float64ToFloat64))

	cnv := convert.StringToTime(time.RFC3339Nano)
	reg.Register(Time, convert.ToAnyAny(cnv))
	reg.registerNumber(Duration, durationConverter())

	reg.Register(String, convert.ToAnyAny(convert.StringToString))
	reg.Register(Bool, convert.ToAnyAny(convert.BoolToBool))

	reg.Register(Nil, NilConverter)
	return reg
}

// Value represents a value and its type.
type Value struct {
	typ string // Name of the type.
	val any    // The value to encode.
}

// New returns a new instance of [Value] for the given value. The type name is
// set to the name returned from `reflect.TypeFor[T]().String()`. When T is an
// interface type, the name of the value's dynamic type is used instead, or
// [Nil] when the interface is nil.
func New[T any](value T) *Value {
	typ := reflect.TypeFor[T]()
	if typ.Kind() != reflect.Interface {
		return &Value{typ: typ.String(), val: value}
	}
	if any(value) == nil {
		return &Value{typ: Nil, val: nil}
	}
	return &Value{typ: reflect.TypeOf(value).String(), val: value}
}

// NewValue works like [New], but it supports untyped nil as the value and
// checks if the type of any other value has a registered converter. Returns
// an error when the type has no registered converter or the registry set
// with [WithRegistry] is nil.
func NewValue(val any, opts ...Option) (*Value, error) {
	ops := newOptions(opts...)
	if ops.reg == nil {
		return nil, fmt.Errorf("jsontype: %w", convert.ErrNilRegistry)
	}
	if val == nil {
		return &Value{typ: Nil, val: nil}, nil
	}
	typ := reflect.TypeOf(val).String()
	if ops.reg.Converter(typ) == nil {
		return nil, fmt.Errorf("jsontype: %w: %s", convert.ErrUnsType, typ)
	}
	return &Value{typ: typ, val: val}, nil
}

// GoTypeName returns the Go type name of the value.
func (val *Value) GoTypeName() string { return val.typ }

// GoValue returns the underlying Go value.
func (val *Value) GoValue() any { return val.val }

// Map returns the map representation of the [Value].
func (val *Value) Map() map[string]any {
	return map[string]any{"type": val.typ, "value": val.val}
}

// MarshalJSON implements [json.Marshaler]. Transparent types (string,
// bool, nil, float64) are emitted as bare JSON primitives; all other
// types use the {"type": "...", "value": ...} envelope. It has a value
// receiver, so a [Value] is encoded the same way whether or not
// [json.Marshal] can take its address (e.g. a map value).
func (val Value) MarshalJSON() ([]byte, error) {
	if val.typ == "" {
		return nil, fmt.Errorf("jsontype: empty type: %w", convert.ErrInvValue)
	}
	switch val.typ {
	case String:
		s, ok := val.val.(string)
		if !ok {
			format := "jsontype: %s: %w"
			return nil, fmt.Errorf(format, val.typ, convert.ErrInvValue)
		}
		return json.Marshal(s)

	case Bool:
		b, ok := val.val.(bool)
		if !ok {
			format := "jsontype: %s: %w"
			return nil, fmt.Errorf(format, val.typ, convert.ErrInvValue)
		}
		if b {
			return []byte("true"), nil
		}
		return []byte("false"), nil

	case Nil:
		if val.val != nil {
			format := "jsontype: %s: %w"
			return nil, fmt.Errorf(format, val.typ, convert.ErrInvValue)
		}
		return []byte("null"), nil

	case Float64:
		f, ok := val.val.(float64)
		if !ok {
			format := "jsontype: %s: %w"
			return nil, fmt.Errorf(format, val.typ, convert.ErrInvValue)
		}
		return marshal(f)
	}
	return marshal(val.Map())
}

// UnmarshalJSON uses the package-level registry. To unmarshal with a custom
// registry, call [Unmarshal] directly.
func (val *Value) UnmarshalJSON(data []byte) error {
	return Unmarshal(registry, data, val)
}

// Unmarshal unmarshals the JSON representation of the value using the
// [Registry].
// Transparent types (string, bool, nil, float64) are detected by their JSON
// shape and require no envelope. The envelope form is still accepted for all
// types (back-compatibility). Whitespace around the JSON value is ignored.
// The registry is used only for the envelope form; for it, a nil registry
// returns [convert.ErrNilRegistry].
func Unmarshal(reg *Registry, data []byte, val *Value) error {
	if val == nil {
		return fmt.Errorf("jsontype: nil Value: %w", convert.ErrInvValue)
	}
	data = bytes.Trim(data, " \t\r\n")
	var first byte
	if len(data) > 0 {
		first = data[0]
	}
	switch first {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
		val.typ, val.val = String, s
		return nil

	case 't':
		if string(data) != "true" {
			format := "jsontype: invalid JSON token: %w"
			return fmt.Errorf(format, convert.ErrInvFormat)
		}
		val.typ, val.val = Bool, true
		return nil

	case 'f':
		if string(data) != "false" {
			format := "jsontype: invalid JSON token: %w"
			return fmt.Errorf(format, convert.ErrInvFormat)
		}
		val.typ, val.val = Bool, false
		return nil

	case 'n':
		if string(data) != "null" {
			format := "jsontype: invalid JSON token: %w"
			return fmt.Errorf(format, convert.ErrInvFormat)
		}
		val.typ, val.val = Nil, nil
		return nil

	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		var f float64
		if err := json.Unmarshal(data, &f); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
		val.typ, val.val = Float64, f
		return nil
	}
	return unmarshalEnvelope(reg, data, val)
}

// unmarshalEnvelope decodes the {"type": "...", "value": ...} form.
func unmarshalEnvelope(reg *Registry, data []byte, val *Value) error {
	tmp := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}
	if reg == nil {
		return fmt.Errorf("jsontype: %w", convert.ErrNilRegistry)
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	cnv, num := reg.converter(tmp.Type)
	if cnv == nil {
		return fmt.Errorf("jsontype: %w: %s", convert.ErrUnsType, tmp.Type)
	}
	var value any
	if len(tmp.Value) > 0 {
		dec := json.NewDecoder(bytes.NewReader(tmp.Value))
		if num {
			dec.UseNumber()
		}
		if err := dec.Decode(&value); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
	}
	ret, err := cnv(value)
	if err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	val.typ, val.val = tmp.Type, ret
	return nil
}

// FromMap constructs an instance of [Value] from its map representation. It
// expects the map to have the same structure as the one returned from the
// [Value.Map] method. The type name must have a converter in the registry
// set with [WithRegistry], except for [Nil]. A built-in type name must also
// match the Go type of the value, where [Byte] and [Rune] match uint8 and
// int32. A custom type name is trusted, as its converter decodes JSON, not
// Go values.
func FromMap(m map[string]any, opts ...Option) (*Value, error) {
	v, ok := keyValue("value", m)
	if !ok {
		format := "jsontype: missing value field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}
	field, ok := keyValue("type", m)
	if !ok {
		format := "jsontype: missing type field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}
	typ, ok := field.(string)
	if !ok {
		format := "jsontype: type field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}

	ops := newOptions(opts...)
	if ops.reg == nil {
		return nil, fmt.Errorf("jsontype: %w", convert.ErrNilRegistry)
	}
	if typ != Nil && ops.reg.Converter(typ) == nil {
		return nil, fmt.Errorf("jsontype: %w: %s", convert.ErrUnsType, typ)
	}

	if goName, ok := builtinGoName(typ); ok {
		valName := Nil
		if v != nil {
			valName = reflect.TypeOf(v).String()
		}
		if valName != goName {
			format := "jsontype: types do not match: %s != %s: %w"
			return nil, fmt.Errorf(format, typ, valName, convert.ErrInvValue)
		}
	}
	return &Value{typ: typ, val: v}, nil
}

// builtinGoName returns the Go type name of the values held under the
// built-in type name and reports whether the name is a built-in one.
func builtinGoName(typ string) (string, bool) {
	switch typ {
	case Byte:
		return Uint8, true

	case Rune:
		return Int32, true

	case Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64,
		Float32, Float64, String, Bool, Time, Duration, Nil:
		return typ, true
	}
	return "", false
}

// AsValue converts a map in the format returned by [Value.Map] into a [Value].
// If v is already a non-nil *Value, it returns that value directly; for a
// [Value] it returns a pointer to its copy. Returns an error if conversion is
// not possible. The options are passed to [FromMap].
func AsValue(v any, opts ...Option) (*Value, error) {
	switch val := v.(type) {
	case *Value:
		if val == nil {
			format := "jsontype: nil Value: %w"
			return nil, fmt.Errorf(format, convert.ErrInvValue)
		}
		return val, nil

	case Value:
		return &val, nil

	case map[string]any:
		return FromMap(val, opts...)
	}
	return nil, fmt.Errorf("jsontype: %w", convert.ErrInvType)
}
