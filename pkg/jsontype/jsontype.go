// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

// Package jsontype preserves Go type information during JSON marshaling.
package jsontype

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/ctx42/convert/pkg/convert"
)

// registry is package level [Registry].
var registry *Registry

// Register registers a converter for the given type name. Returns the
// previous converter if one was already registered, nil otherwise.
func Register(typ string, cnv convert.AnyToAny) convert.AnyToAny {
	if cnv == nil {
		return nil
	}
	return registry.Register(typ, cnv)
}

func init() { registry = DefaultRegistry() }

// List of type names supported by the package out of the box.
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

// DefaultRegistry returns default registry configuration.
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

	reg.Register(Float32, convert.ToAnyAny(convert.Float64ToFloat32))
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

// New returns new instance of [Value] for the given value. The type name is
// set to the name returned from `reflect.TypeFor[T]().String()`.
func New[T any](value T) *Value {
	return &Value{typ: reflect.TypeFor[T]().String(), val: value}
}

// NewValue works like [New], but it supports untyped nil as the value and
// checks if the type has a registered converter. Returns error when the type
// has no registered converter or the registry set with [WithRegistry] is nil.
func NewValue(val any, opts ...Option) (*Value, error) {
	if val == nil {
		return &Value{typ: Nil, val: nil}, nil
	}
	def := &Options{reg: registry}
	for _, opt := range opts {
		opt(def)
	}
	if def.reg == nil {
		return nil, fmt.Errorf("jsontype: %w", convert.ErrNilRegistry)
	}
	typ := reflect.TypeOf(val).String()
	if cnv := def.reg.Converter(typ); cnv == nil {
		return nil, fmt.Errorf("jsontype: %w: %s", convert.ErrUnsType, typ)
	}
	return &Value{typ: typ, val: val}, nil
}

// GoTypeName returns the Go type name of the value.
func (val *Value) GoTypeName() string { return val.typ }

// GoValue returns the underlying Go value.
func (val *Value) GoValue() any { return val.val }

// Map returns map representation of the [Value].
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
func (val *Value) UnmarshalJSON(bytes []byte) error {
	return Unmarshal(registry, bytes, val)
}

// FromMap constructs an instance of [Value] from its map representation. It
// expects the map to have the same structure as the one returned from the
// [Value.Map] method. The options are passed to [NewValue].
func FromMap(m map[string]any, opts ...Option) (val *Value, err error) {
	var v any
	var ok bool

	if v, ok = keyValue("value", m); !ok {
		format := "jsontype: missing value field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}
	if val, err = NewValue(v, opts...); err != nil {
		return nil, err
	}

	if v, ok = keyValue("type", m); !ok {
		format := "jsontype: missing type field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}

	var typ string
	if typ, ok = v.(string); !ok {
		format := "jsontype: type field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}

	if typ != val.typ {
		format := "jsontype: types do not match: %s != %s: %w"
		return nil, fmt.Errorf(format, typ, val.typ, convert.ErrInvValue)
	}
	return val, nil
}

// AsValue converts a map in the format returned by [Value.Map] into a [Value].
// If v is already a *Value, it returns that value directly. Returns error if
// conversion is not possible. The options are passed to [FromMap].
func AsValue(v any, opts ...Option) (*Value, error) {
	if val, ok := v.(*Value); ok {
		return val, nil
	}
	if val, ok := v.(map[string]any); ok {
		return FromMap(val, opts...)
	}
	return nil, fmt.Errorf("jsontype: %w", convert.ErrInvType)
}
