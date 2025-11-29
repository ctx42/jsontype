// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// DecodeInt expects a value of type float64 from JSON and returns it as int.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for int.
func DecodeInt(value any) (int, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeInt: requires float64 value got %T: %w"
		return 0, fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeInt: requires non-fractional value: %w"
		return 0, fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt || f64 > math.MaxInt {
		format := "DecodeInt: requires float64 value in range of int: %w"
		return 0, fmt.Errorf(format, ErrInvRange)
	}
	return int(f64), nil
}

// DecodeInt8 expects a value of type float64 from JSON and returns it as int8.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for int8.
func DecodeInt8(value any) (int8, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeInt8: requires float64 value got %T: %w"
		return int8(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeInt8: requires non-fractional value: %w"
		return int8(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt8 || f64 > math.MaxInt8 {
		format := "DecodeInt8: requires float64 value in range of int8: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return int8(0), err
	}
	return int8(f64), nil
}

// DecodeInt16 expects a value of type float64 from JSON and returns it as
// int16. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int16.
func DecodeInt16(value any) (int16, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeInt16: requires float64 value got %T: %w"
		return int16(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeInt16: requires non-fractional value: %w"
		return int16(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt16 || f64 > math.MaxInt16 {
		format := "DecodeInt16: requires float64 value in range of int16: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return int16(0), err
	}
	return int16(f64), nil
}

// DecodeInt32 expects a value of type float64 from JSON and returns it as
// int32. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int32.
func DecodeInt32(value any) (int32, error) {
	return decodeNamedInt32("DecodeInt32", value)
}

// DecodeRune expects a value of type float64 from JSON and returns it as rune.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for rune.
func DecodeRune(value any) (rune, error) {
	return decodeNamedInt32("DecodeRune", value)
}

// decodeNamedInt32 expects a value of type float64 from JSON and returns it as
// int32. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int32.
func decodeNamedInt32(name string, value any) (int32, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "%s: requires float64 value got %T: %w"
		return int32(0), fmt.Errorf(format, name, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "%s: requires non-fractional value: %w"
		return int32(0), fmt.Errorf(format, name, ErrInvValue)
	}
	if f64 < math.MinInt32 || f64 > math.MaxInt32 {
		format := "%s: requires float64 value in range of int32: %w"
		err := fmt.Errorf(format, name, ErrInvRange)
		return int32(0), err
	}
	return int32(f64), nil
}

// DecodeInt64 expects a value of type float64 from JSON and returns it as
// int64. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int64.
func DecodeInt64(value any) (int64, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeInt64: requires float64 value got %T: %w"
		return int64(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeInt64: requires non-fractional value: %w"
		return int64(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt64 || f64 > math.MaxInt64 {
		format := "DecodeInt64: requires float64 value in range of int64: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return int64(0), err
	}
	return int64(f64), nil
}

// DecodeUInt expects a value of type float64 from JSON and returns it as uint.
// Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint.
func DecodeUInt(value any) (uint, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeUInt: requires float64 value got %T: %w"
		return uint(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeUInt: requires non-fractional value: %w"
		return uint(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint {
		format := "DecodeUInt: requires float64 value in range of uint: %w"
		return uint(0), fmt.Errorf(format, ErrInvRange)
	}
	return uint(f64), nil
}

// DecodeUInt8 expects a value of type float64 from JSON and returns it as
// uint8. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint8.
func DecodeUInt8(value any) (uint8, error) {
	return decodeNamedUInt8("DecodeUInt8", value)
}

// DecodeByte expects a value of type float64 from JSON and returns it as byte.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for byte.
func DecodeByte(value any) (byte, error) {
	return decodeNamedUInt8("DecodeByte", value)
}

// decodeNamedUInt8 expects a value of type float64 from JSON and returns it as
// uint8. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint8.
func decodeNamedUInt8(name string, value any) (uint8, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "%s: requires float64 value got %T: %w"
		return uint8(0), fmt.Errorf(format, name, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "%s: requires non-fractional value: %w"
		return uint8(0), fmt.Errorf(format, name, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint8 {
		format := "%s: requires float64 value in range of uint8: %w"
		return uint8(0), fmt.Errorf(format, name, ErrInvRange)
	}
	return uint8(f64), nil
}

// DecodeUInt16 expects a value of type float64 from JSON and returns it as
// uint16. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint16.
func DecodeUInt16(value any) (uint16, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeUInt16: requires float64 value got %T: %w"
		return uint16(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeUInt16: requires non-fractional value: %w"
		return uint16(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint16 {
		format := "DecodeUInt16: requires float64 value in range of uint16: %w"
		return uint16(0), fmt.Errorf(format, ErrInvRange)
	}
	return uint16(f64), nil
}

// DecodeUInt32 expects a value of type float64 from JSON and returns it as
// uint32. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint32.
func DecodeUInt32(value any) (uint32, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeUInt32: requires float64 value got %T: %w"
		return uint32(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeUInt32: requires non-fractional value: %w"
		return uint32(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint32 {
		format := "DecodeUInt32: requires float64 value in range of uint32: %w"
		return uint32(0), fmt.Errorf(format, ErrInvRange)
	}
	return uint32(f64), nil
}

// DecodeUInt64 expects a value of type float64 from JSON and returns it as
// uint64. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint64.
func DecodeUInt64(value any) (uint64, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeUInt64: requires float64 value got %T: %w"
		return uint64(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "DecodeUInt64: requires non-fractional value: %w"
		return uint64(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint64 {
		format := "DecodeUInt64: requires float64 value in range of uint64: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return uint64(0), err
	}
	return uint64(f64), nil
}

// DecodeFloat32 expects a value of type float64 from JSON and returns it as
// float32. Returns error when value is not float64 or value is out of range
// for float32.
func DecodeFloat32(value any) (float32, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeFloat32: requires float64 value got %T: %w"
		return float32(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 < 0 || f64 > math.MaxFloat32 {
		format := "DecodeFloat32: requires float64 value " +
			"in range of float32: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return float32(0), err
	}
	return float32(f64), nil
}

// DecodeFloat64 expects a value of type float64 from JSON and returns it as
// float64. Returns error when value is not float64.
func DecodeFloat64(value any) (float64, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "DecodeFloat64: requires float64 value got %T: %w"
		return float64(0), fmt.Errorf(format, value, ErrInvType)
	}
	return f64, nil

}

// DecodeString expects a value of type string from JSON and returns it.
// Returns error when value is not string.
func DecodeString(value any) (string, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "DecodeString: requires string value got %T: %w"
		return "", fmt.Errorf(format, value, ErrInvType)
	}
	return str, nil
}

// DecodeBool expects a value of type boolean from JSON and returns it. Returns
// error when value is not boolean.
func DecodeBool(value any) (bool, error) {
	var ok bool
	var str bool
	if str, ok = value.(bool); !ok {
		format := "DecodeBool: requires boolean value got %T: %w"
		return false, fmt.Errorf(format, value, ErrInvType)
	}
	return str, nil
}

// DecodeTime expects JSON string value, with valid RFC3339 time and returns
// its [time.Time] representation. Returns error when value is not string or
// value cannot be parsed as RFC3339 [time.Time].
func DecodeTime(value any) (time.Time, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "DecodeTime: requires boolean value got %T: %w"
		return time.Time{}, fmt.Errorf(format, value, ErrInvType)
	}
	tim, err := time.Parse(time.RFC3339, str)
	if err != nil {
		format := "DecodeTime: parsing RFC3339 string value to time.Time: %w"
		return time.Time{}, fmt.Errorf(format, ErrInvFormat)
	}
	return tim, nil
}

// DecodeDuration expects JSON string value, with valid duration and parses it.
// Returns error when value is not string or value cannot be parsed as
// [time.Duration].
func DecodeDuration(value any) (time.Duration, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "DecodeDuration: requires boolean value got %T: %w"
		return time.Duration(0), fmt.Errorf(format, value, ErrInvType)
	}
	dur, err := time.ParseDuration(str)
	if err != nil {
		format := "DecodeDuration: parsing string value to time.Duration: %w"
		return time.Duration(0), fmt.Errorf(format, ErrInvFormat)
	}
	return dur, nil
}

// DecodeComplex64 expects JSON string value, with valid complex number and
// parses it. Returns error when value is not string or value cannot be parsed
// as complex64.
func DecodeComplex64(value any) (complex64, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "DecodeComplex64: requires boolean value got %T: %w"
		return 0i + 0, fmt.Errorf(format, value, ErrInvType)
	}
	cpx, err := strconv.ParseComplex(str, 64)
	if err != nil {
		format := "DecodeComplex64: parsing string value to complex64: %w"
		return 0i + 0, fmt.Errorf(format, ErrInvFormat)
	}
	return complex64(cpx), nil
}

// DecodeComplex128 expects JSON string value, with valid complex number and
// parses it. Returns error when value is not string or value cannot be parsed
// as complex128.
func DecodeComplex128(value any) (complex128, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "DecodeComplex128: requires boolean value got %T: %w"
		return 0i + 0, fmt.Errorf(format, value, ErrInvType)
	}
	cpx, err := strconv.ParseComplex(str, 128)
	if err != nil {
		format := "DecodeComplex128: parsing string value to complex128: %w"
		return 0i + 0, fmt.Errorf(format, ErrInvFormat)
	}
	return cpx, nil
}

// DecodeNil expects value to be nil, otherwise returns an error.
// On success, it returns nil and nil error.
func DecodeNil(value any) (any, error) {
	if value != nil {
		format := "DecodeNil: requires nil value: %w"
		return nil, fmt.Errorf(format, ErrInvType)
	}
	return nil, nil
}
