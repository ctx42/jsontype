// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
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

// Register registers converter for a given type name.
func Register(typ string, cnv convert.AnyToAny) convert.AnyToAny {
	if cnv == nil {
		return nil
	}
	return registry.Register(typ, cnv)
}

func init() {
	registry = NewRegistry()

	registry.Register(Byte, convert.ToAnyAny(convert.Float64ToByte))
	registry.Register(Uint8, convert.ToAnyAny(convert.Float64ToUint8))
	registry.Register(Uint16, convert.ToAnyAny(convert.Float64ToUint16))
	registry.Register(Uint32, convert.ToAnyAny(convert.Float64ToUint32))
	registry.Register(Uint64, convert.ToAnyAny(convert.Float64ToUint64))
	registry.Register(Uint, convert.ToAnyAny(convert.Float64ToUint))

	registry.Register(Int8, convert.ToAnyAny(convert.Float64ToInt8))
	registry.Register(Int16, convert.ToAnyAny(convert.Float64ToInt16))
	registry.Register(Rune, convert.ToAnyAny(convert.Float64ToRune))
	registry.Register(Int32, convert.ToAnyAny(convert.Float64ToInt32))
	registry.Register(Int64, convert.ToAnyAny(convert.Float64ToInt64))
	registry.Register(Int, convert.ToAnyAny(convert.Float64ToInt))

	registry.Register(Float32, convert.ToAnyAny(convert.Float64ToFloat32))
	registry.Register(Float64, convert.ToAnyAny(convert.Float64ToFloat64))

	cnv := convert.StringToTime(time.RFC3339Nano)
	registry.Register(Time, convert.ToAnyAny(cnv))
	registry.Register(Duration, convert.ToAnyAny(convert.StringToDuration))

	registry.Register(String, convert.ToAnyAny(convert.StringToString))
	registry.Register(Bool, convert.ToAnyAny(convert.BoolToBool))
	registry.Register(Duration, convert.ToAnyAny(convert.StringToDuration))

	registry.Register(Nil, NilConverter)
}

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
// has no registered converter.
func NewValue(val any) (*Value, error) { return newValue(registry, val) }

func newValue(reg *Registry, value any) (*Value, error) {
	if value == nil {
		return &Value{typ: Nil, val: nil}, nil
	}
	typ := reflect.TypeOf(value).String()
	if cnv := reg.Converter(typ); cnv == nil {
		return nil, fmt.Errorf("%w: %s", convert.ErrUnkType, typ)
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
		return nil, convert.ErrInvValue
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
	cnv := reg.Converter(tmp.Type)
	if cnv == nil {
		return fmt.Errorf("%w: %s", convert.ErrUnkType, tmp.Type)
	}
	val.typ = tmp.Type
	if val.val, err = cnv(tmp.Value); err != nil {
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
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}
	if val, err = NewValue(v); err != nil {
		return nil, fmt.Errorf("FromMap: %w", err)
	}

	if v, ok = keyValue("type", m); !ok {
		format := "FromMap: missing type field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}

	var typ string
	if typ, ok = v.(string); !ok {
		format := "FromMap: type field: %w"
		return nil, fmt.Errorf(format, convert.ErrInvFormat)
	}

	if typ != val.typ {
		format := "FromMap: types do not match: %w"
		return nil, fmt.Errorf(format, convert.ErrInvValue)
	}
	return val, nil
}

// FromMapAny expects the argument to be `map[string]any` in the format
// returned by [Value.Map] and constructs an instance of [Value] from it.
func FromMapAny(v any) (*Value, error) {
	if val, ok := v.(map[string]any); ok {
		return FromMap(val)
	}
	return nil, fmt.Errorf("FromMapAny: %w", convert.ErrInvType)
}
