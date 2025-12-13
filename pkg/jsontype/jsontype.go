// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package jsontype preserves Go type information during JSON marshaling.
package jsontype

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/ctx42/convert/pkg/xcast"
)

// registry is package level [Registry].
var registry *Registry

// Register registers [Decoder] for a given type name.
func Register(typ string, dec Decoder) Decoder {
	if dec == nil {
		return nil
	}
	return registry.Register(typ, dec)
}

func init() {
	registry = NewRegistry()

	registry.Register(Byte, FromConv(xcast.Float64ToByte))
	registry.Register(Uint8, FromConv(xcast.Float64ToUint8))
	registry.Register(Uint16, FromConv(xcast.Float64ToUint16))
	registry.Register(Uint32, FromConv(xcast.Float64ToUint32))
	registry.Register(Uint64, FromConv(xcast.Float64ToUint64))
	registry.Register(Uint, FromConv(xcast.Float64ToUint))

	registry.Register(Int8, FromConv(xcast.Float64ToInt8))
	registry.Register(Int16, FromConv(xcast.Float64ToInt16))
	registry.Register(Rune, FromConv(xcast.Float64ToRune))
	registry.Register(Int32, FromConv(xcast.Float64ToInt32))
	registry.Register(Int64, FromConv(xcast.Float64ToInt64))
	registry.Register(Int, FromConv(xcast.Float64ToInt))

	registry.Register(Float32, FromConv(xcast.Float64ToFloat32))
	registry.Register(Float64, FromConv(xcast.Float64ToFloat64))

	registry.Register(Time, FromConv(xcast.StringToTime(time.RFC3339Nano)))
	registry.Register(Duration, FromConv(xcast.StringToDuration))

	registry.Register(String, FromConv(xcast.StringToString))
	registry.Register(Bool, FromConv(xcast.BoolToBool))
	registry.Register(Duration, FromConv(xcast.StringToDuration))

	registry.Register(Nil, DecodeNil)
}

// List of types which can be encoded to JSON by [Value] and later on decoded
// without losing the Go type in the process. The decoders for all the listed
// types are by default added to the global registry in init function.
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

// Decoder is a function that knows how to decode a JSON type to a Go type.
type Decoder func(value any) (any, error)

// Converter represents a function that attempts lossless conversion from a
// source value of type From to a target value of type To. On success, it
// returns the converted value and a nil error. On failure (e.g., truncation,
// overflow, or semantic loss), it returns the zero value of To along with a
// non-nil error describing the issue.
type Converter[From, To any] func(from From) (to To, err error)

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
// checks if the type has a registered decoder. Returns error when the type has
// no registered decoder.
func NewValue(val any) (*Value, error) { return newValue(registry, val) }

func newValue(reg *Registry, value any) (*Value, error) {
	if value == nil {
		return &Value{typ: Nil, val: nil}, nil
	}
	typ := reflect.TypeOf(value).String()
	if dec := reg.Decoder(typ); dec == nil {
		return nil, fmt.Errorf("%w: %s", xcast.ErrUnkType, typ)
	}
	return &Value{typ: typ, val: value}, nil
}

func (val *Value) GoTypeName() string { return val.typ }
func (val *Value) GoValue() any       { return val.val }

// Map returns map representation of the [Value].
func (val *Value) Map() map[string]any {
	return map[string]any{"type": val.typ, "value": val.val}
}

func (val *Value) MarshalJSON() ([]byte, error) {
	if val == nil || val.typ == "" {
		return nil, xcast.ErrInvValue
	}
	return json.Marshal(val.Map())
}

func (val *Value) UnmarshalJSON(bytes []byte) error {
	return val.unmarshalJSON(registry, bytes)
}

func (val *Value) unmarshalJSON(reg *Registry, bytes []byte) error {
	tmp := struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	}{}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	var err error
	dec := reg.Decoder(tmp.Type)
	if dec == nil {
		return fmt.Errorf("%w: %s", xcast.ErrUnkType, tmp.Type)
	}
	val.typ = tmp.Type
	if val.val, err = dec(tmp.Value); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	return nil
}

// FromMap constructs an instance of [Value] from its map representation. It
// expects the map to have the same structure as the one returned from the
// [Value.Map] method.
func FromMap(m map[string]any) (val *Value, err error) {
	var v any
	var ok bool

	if v, ok = keyValue("value", m); !ok {
		format := "FromMap: missing value field: %w"
		return nil, fmt.Errorf(format, xcast.ErrInvFormat)
	}
	if val, err = NewValue(v); err != nil {
		return nil, fmt.Errorf("FromMap: %w", err)
	}

	if v, ok = keyValue("type", m); !ok {
		format := "FromMap: missing type field: %w"
		return nil, fmt.Errorf(format, xcast.ErrInvFormat)
	}

	var typ string
	if typ, ok = v.(string); !ok {
		format := "FromMap: type field: %w"
		return nil, fmt.Errorf(format, xcast.ErrInvFormat)
	}

	if typ != val.typ {
		format := "FromMap: types do not match: %w"
		return nil, fmt.Errorf(format, xcast.ErrInvValue)
	}
	return val, nil
}

// FromMapAny expects the argument to be `map[string]any` in the format
// returned by [Value.Map] and constructs an instance of [Value] from it.
func FromMapAny(v any) (*Value, error) {
	if val, ok := v.(map[string]any); ok {
		return FromMap(val)
	}
	return nil, fmt.Errorf("FromMapAny: %w", xcast.ErrInvType)
}
