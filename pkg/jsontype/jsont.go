// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

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

	registry.Register(Int, ToAny(decodeInt))
	registry.Register(Int8, ToAny(decodeInt8))
	registry.Register(Int16, ToAny(decodeInt16))
	registry.Register(Int32, ToAny(decodeInt32))
	registry.Register(Rune, ToAny(decodeRune))
	registry.Register(Int64, ToAny(decodeInt64))
	registry.Register(UInt, ToAny(decodeUInt))
	registry.Register(UInt8, ToAny(decodeUInt8))
	registry.Register(Byte, ToAny(decodeByte))
	registry.Register(UInt16, ToAny(decodeUInt16))
	registry.Register(UInt32, ToAny(decodeUInt32))
	registry.Register(UInt64, ToAny(decodeUInt64))
	registry.Register(Float32, ToAny(decodeFloat32))
	registry.Register(Float64, ToAny(decodeFloat64))
	registry.Register(String, ToAny(decodeString))
	registry.Register(Bool, ToAny(decodeBool))
	registry.Register(Time, ToAny(decodeTime))
	registry.Register(Duration, ToAny(decodeDuration))
	registry.Register(Complex64, ToAny(decodeComplex64))
	registry.Register(Complex128, ToAny(decodeComplex128))
}

// TypeName represents a type name.
type TypeName string

// List of types which can be encoded to JSON by [Value] and later on decoded
// without losing the Go type in the process. The decoders for all the listed
// types are by default added to global registry in init function.
const (
	Int        TypeName = "int"
	Byte       TypeName = "byte"
	Int8       TypeName = "int8"
	Int16      TypeName = "int16"
	Int32      TypeName = "int32"
	Rune       TypeName = "rune"
	Int64      TypeName = "int64"
	UInt       TypeName = "uint"
	UInt8      TypeName = "uint8"
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

func (v *Value) MarshalJSON() ([]byte, error) {
	if v == nil || v.typ == "" {
		return nil, ErrInvValue
	}
	data := map[string]any{
		"type":  v.typ,
		"value": v.val,
	}
	return json.Marshal(data)
}

func (v *Value) UnmarshalJSON(bytes []byte) error {
	val := struct {
		Type  TypeName `json:"type"`
		Value any      `json:"value"`
	}{}

	if err := json.Unmarshal(bytes, &val); err != nil {
		return err
	}

	var err error
	v.typ = val.Type
	dec := registry.Decoder(val.Type)
	if dec == nil {
		return fmt.Errorf("%w: %s", ErrUnkType, val.Type)
	}
	if v.val, err = dec(val.Value); err != nil {
		return err
	}
	return nil
}
