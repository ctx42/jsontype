// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

// Package jsontype preserves Go type information during JSON marshaling.
package jsontype

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

// registry is package level [Registry].
var registry *Registry

func init() {
	registry = NewRegistry()

	registry.Register(Int, ToAny(DecodeInt))
	registry.Register(Int8, ToAny(DecodeInt8))
	registry.Register(Int16, ToAny(DecodeInt16))
	registry.Register(Int32, ToAny(DecodeInt32))
	registry.Register(Rune, ToAny(DecodeRune))
	registry.Register(Int64, ToAny(DecodeInt64))
	registry.Register(UInt, ToAny(DecodeUInt))
	registry.Register(UInt8, ToAny(DecodeUInt8))
	registry.Register(Byte, ToAny(DecodeByte))
	registry.Register(UInt16, ToAny(DecodeUInt16))
	registry.Register(UInt32, ToAny(DecodeUInt32))
	registry.Register(UInt64, ToAny(DecodeUInt64))
	registry.Register(Float32, ToAny(DecodeFloat32))
	registry.Register(Float64, ToAny(DecodeFloat64))
	registry.Register(String, ToAny(DecodeString))
	registry.Register(Bool, ToAny(DecodeBool))
	registry.Register(Time, ToAny(DecodeTime))
	registry.Register(Duration, ToAny(DecodeDuration))
	registry.Register(Complex64, ToAny(DecodeComplex64))
	registry.Register(Complex128, ToAny(DecodeComplex128))
	registry.Register(Nil, ToAny(DecodeNil))
}

// TypeName represents a type name.
type TypeName string

func (name TypeName) String() string { return string(name) }

// List of types which can be encoded to JSON by [Value] and later on decoded
// without losing the Go type in the process. The decoders for all the listed
// types are by default added to global registry in init function.
const (
	Int        TypeName = "int"
	Int8       TypeName = "int8"
	Int16      TypeName = "int16"
	Int32      TypeName = "int32"
	Rune       TypeName = "rune"
	Int64      TypeName = "int64"
	UInt       TypeName = "uint"
	UInt8      TypeName = "uint8"
	Byte       TypeName = "byte"
	UInt16     TypeName = "uint16"
	UInt32     TypeName = "uint32"
	UInt64     TypeName = "uint64"
	Float32    TypeName = "float32"
	Float64    TypeName = "float64"
	String     TypeName = "string"
	Bool       TypeName = "bool"
	Time       TypeName = "time.Time"
	Duration   TypeName = "time.Duration"
	Complex64  TypeName = "complex64"
	Complex128 TypeName = "complex128"
	Nil        TypeName = "nil"
)

// Package level sentinel errors.
var (
	// ErrInvValue indicates invalid value. Is used for example when a JSON
	// value cannot be cast a Go type without loosing information.
	ErrInvValue = errors.New("invalid value")

	// ErrInvType indicates invalid type for a requested action. Is used for
	// example when JSON type does not match required Go type.
	ErrInvType = errors.New("invalid type")

	// ErrInvRange indicates value is out of bounds for given action. Used for
	// example when trying to cast negative value to unsigned type.
	ErrInvRange = errors.New("invalid range")

	// ErrInvFormat indicates invalid JSON value format. Used for example when
	// parsing invalid string to time.
	ErrInvFormat = errors.New("invalid format")

	// ErrUnkType indicates unknown type. Used when marshalling or
	// unmarshalling not registered type.
	ErrUnkType = errors.New("unknown type")
)

// SimpleType represents types that can be used with simple constructor [New].
type SimpleType interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 |
		complex64 | complex128 |
		string |
		bool |
		time.Time | time.Duration
}

// Decoder is a function which knows how to decode a JSON type to a Go type.
type Decoder func(value any) (any, error)

// Value represents a value and its type.
type Value struct {
	typ TypeName // Name of the type.
	val any      // The value to encode.
}

// New returns new instance of [Value] for given [SimpleType].
func New[T SimpleType](val T) *Value {
	return &Value{
		typ: TypeName(reflect.TypeFor[T]().String()),
		val: val,
	}
}

// NewByte returns new instance of [Value] representing a byte.
func NewByte(val byte) *Value { return &Value{typ: Byte, val: val} }

// NewRune returns new instance of [Value] representing a byte.
func NewRune(val rune) *Value { return &Value{typ: Rune, val: val} }

// NewValue returns an instance of [Value] based on the given argument. For
// untyped nil it returns [Nil] type and nil. Returns an error when a [Decoder]
// for the type has not been found.
func NewValue(val any) (*Value, error) {
	if val == nil {
		return &Value{typ: Nil, val: nil}, nil
	}
	typ := TypeName(reflect.TypeOf(val).String())
	if dec := registry.Decoder(typ); dec == nil {
		return nil, fmt.Errorf("%w: %s", ErrUnkType, typ)
	}
	value := &Value{typ: typ, val: val}
	return value, nil
}

// FromMap constructs an instance of [Value] from its map representation. It
// expects the map to have the same structure as the one returned from
// [Value.Map] method.
func FromMap(m map[string]any) (val *Value, err error) {
	var v any
	var ok bool

	if v, ok = keyValue("value", m); !ok {
		return nil, fmt.Errorf("FromMap: missing value field: %w", ErrInvFormat)
	}
	if val, err = NewValue(v); err != nil {
		return nil, fmt.Errorf("FromMap: %w", err)
	}

	if v, ok = keyValue("type", m); !ok {
		return nil, fmt.Errorf("FromMap: missing type field: %w", ErrInvFormat)
	}

	var typ string
	if typ, err = DecodeString(v); err != nil {
		return nil, fmt.Errorf("FromMap: type field: %w", ErrInvFormat)
	}

	if TypeName(typ) != val.typ {
		return nil, fmt.Errorf("FromMap: types do not match: %w", ErrInvValue)
	}
	return val, nil
}

// FromAny expects the argument to be `map[string]any` in the same format
// [Value.Map] returns and constructs an instance of [Value] from it.
func FromAny(v any) (*Value, error) {
	if val, ok := v.(map[string]any); ok {
		return FromMap(val)
	}
	return nil, fmt.Errorf("FromAny: %w", ErrInvType)
}

func (val *Value) GoType() TypeName { return val.typ }
func (val *Value) GoValue() any     { return val.val }

// Map returns map representation of the [Value].
func (val *Value) Map() map[string]any {
	return map[string]any{"type": val.typ.String(), "value": val.val}
}

func (val *Value) MarshalJSON() ([]byte, error) {
	if val == nil || val.typ == "" {
		return nil, ErrInvValue
	}
	return json.Marshal(val.Map())
}

func (val *Value) UnmarshalJSON(bytes []byte) error {
	tmp := struct {
		Type  TypeName `json:"type"`
		Value any      `json:"value"`
	}{}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return err
	}

	var err error
	val.typ = tmp.Type
	dec := registry.Decoder(tmp.Type)
	if dec == nil {
		return fmt.Errorf("%w: %s", ErrUnkType, tmp.Type)
	}
	if val.val, err = dec(tmp.Value); err != nil {
		return err
	}
	return nil
}
