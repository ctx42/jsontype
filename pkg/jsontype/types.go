// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// decodeInt expects a value of type float64 from JSON and returns it as int.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for int.
func decodeInt(value any) (int, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeInt: requires float64 value got %T: %w"
		return 0, fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeInt: requires non-fractional value: %w"
		return 0, fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt || f64 > math.MaxInt {
		format := "decodeInt: requires float64 value in range of int: %w"
		return 0, fmt.Errorf(format, ErrInvRange)
	}
	return int(f64), nil
}

// decodeInt8 expects a value of type float64 from JSON and returns it as int8.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for int8.
func decodeInt8(value any) (int8, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeInt8: requires float64 value got %T: %w"
		return int8(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeInt8: requires non-fractional value: %w"
		return int8(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt8 || f64 > math.MaxInt8 {
		format := "decodeInt8: requires float64 value in range of int8: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return int8(0), err
	}
	return int8(f64), nil
}

// decodeInt16 expects a value of type float64 from JSON and returns it as
// int16. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int16.
func decodeInt16(value any) (int16, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeInt16: requires float64 value got %T: %w"
		return int16(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeInt16: requires non-fractional value: %w"
		return int16(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt16 || f64 > math.MaxInt16 {
		format := "decodeInt16: requires float64 value in range of int16: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return int16(0), err
	}
	return int16(f64), nil
}

// decodeInt32 expects a value of type float64 from JSON and returns it as
// int32. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int32.
func decodeInt32(value any) (int32, error) {
	return decodeNamedInt32("decodeInt32", value)
}

// decodeRune expects a value of type float64 from JSON and returns it as rune.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for rune.
func decodeRune(value any) (rune, error) {
	return decodeNamedInt32("decodeRune", value)
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

// decodeInt64 expects a value of type float64 from JSON and returns it as
// int64. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for int64.
func decodeInt64(value any) (int64, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeInt64: requires float64 value got %T: %w"
		return int64(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeInt64: requires non-fractional value: %w"
		return int64(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < math.MinInt64 || f64 > math.MaxInt64 {
		format := "decodeInt64: requires float64 value in range of int64: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return int64(0), err
	}
	return int64(f64), nil
}

// decodeUInt expects a value of type float64 from JSON and returns it as uint.
// Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint.
func decodeUInt(value any) (uint, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeUInt: requires float64 value got %T: %w"
		return uint(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeUInt: requires non-fractional value: %w"
		return uint(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint {
		format := "decodeUInt: requires float64 value in range of uint: %w"
		return uint(0), fmt.Errorf(format, ErrInvRange)
	}
	return uint(f64), nil
}

// decodeUInt8 expects a value of type float64 from JSON and returns it as
// uint8. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint8.
func decodeUInt8(value any) (uint8, error) {
	return decodeNamedUInt8("decodeUInt8", value)
}

// decodeByte expects a value of type float64 from JSON and returns it as byte.
// Returns error when value is not float64, value cannot be represented as an
// integer without loss of information or is out of range for byte.
func decodeByte(value any) (byte, error) {
	return decodeNamedUInt8("decodeByte", value)
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

// decodeUInt16 expects a value of type float64 from JSON and returns it as
// uint16. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint16.
func decodeUInt16(value any) (uint16, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeUInt16: requires float64 value got %T: %w"
		return uint16(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeUInt16: requires non-fractional value: %w"
		return uint16(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint16 {
		format := "decodeUInt16: requires float64 value in range of uint16: %w"
		return uint16(0), fmt.Errorf(format, ErrInvRange)
	}
	return uint16(f64), nil
}

// decodeUInt32 expects a value of type float64 from JSON and returns it as
// uint32. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint32.
func decodeUInt32(value any) (uint32, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeUInt32: requires float64 value got %T: %w"
		return uint32(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeUInt32: requires non-fractional value: %w"
		return uint32(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint32 {
		format := "decodeUInt32: requires float64 value in range of uint32: %w"
		return uint32(0), fmt.Errorf(format, ErrInvRange)
	}
	return uint32(f64), nil
}

// decodeUInt64 expects a value of type float64 from JSON and returns it as
// uint64. Returns error when value is not float64, value cannot be represented
// as an integer without loss of information or is out of range for uint64.
func decodeUInt64(value any) (uint64, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeUInt64: requires float64 value got %T: %w"
		return uint64(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 != math.Trunc(f64) {
		format := "decodeUInt64: requires non-fractional value: %w"
		return uint64(0), fmt.Errorf(format, ErrInvValue)
	}
	if f64 < 0 || f64 > math.MaxUint64 {
		format := "decodeUInt64: requires float64 value in range of uint64: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return uint64(0), err
	}
	return uint64(f64), nil
}

// decodeFloat32 expects a value of type float64 from JSON and returns it as
// float32. Returns error when value is not float64 or value is out of range
// for float32.
func decodeFloat32(value any) (float32, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeFloat32: requires float64 value got %T: %w"
		return float32(0), fmt.Errorf(format, value, ErrInvType)
	}
	if f64 < 0 || f64 > math.MaxFloat32 {
		format := "decodeFloat32: requires float64 value " +
			"in range of float32: %w"
		err := fmt.Errorf(format, ErrInvRange)
		return float32(0), err
	}
	return float32(f64), nil
}

// decodeFloat64 expects a value of type float64 from JSON and returns it as
// float64. Returns error when value is not float64.
func decodeFloat64(value any) (float64, error) {
	var ok bool
	var f64 float64
	if f64, ok = value.(float64); !ok {
		format := "decodeFloat64: requires float64 value got %T: %w"
		return float64(0), fmt.Errorf(format, value, ErrInvType)
	}
	return f64, nil

}

// decodeString expects a value of type string from JSON and returns it.
// Returns error when value is not string.
func decodeString(value any) (string, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "decodeString: requires string value got %T: %w"
		return "", fmt.Errorf(format, value, ErrInvType)
	}
	return str, nil
}

// decodeBool expects a value of type boolean from JSON and returns it. Returns
// error when value is not boolean.
func decodeBool(value any) (bool, error) {
	var ok bool
	var str bool
	if str, ok = value.(bool); !ok {
		format := "decodeBool: requires boolean value got %T: %w"
		return false, fmt.Errorf(format, value, ErrInvType)
	}
	return str, nil
}

// decodeTime expects JSON string value, with valid RFC3339 time and returns
// its [time.Time] representation. Returns error when value is not string or
// value cannot be parsed as RFC3339 [time.Time].
func decodeTime(value any) (time.Time, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "decodeTime: requires boolean value got %T: %w"
		return time.Time{}, fmt.Errorf(format, value, ErrInvType)
	}
	tim, err := time.Parse(time.RFC3339, str)
	if err != nil {
		format := "decodeTime: parsing RFC3339 string value to time.Time: %w"
		return time.Time{}, fmt.Errorf(format, ErrInvFormat)
	}
	return tim, nil
}

// decodeDuration expects JSON string value, with valid duration and parses it.
// Returns error when value is not string or value cannot be parsed as
// [time.Duration].
func decodeDuration(value any) (time.Duration, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "decodeDuration: requires boolean value got %T: %w"
		return time.Duration(0), fmt.Errorf(format, value, ErrInvType)
	}
	dur, err := time.ParseDuration(str)
	if err != nil {
		format := "decodeDuration: parsing string value to time.Duration: %w"
		return time.Duration(0), fmt.Errorf(format, ErrInvFormat)
	}
	return dur, nil
}

// decodeComplex64 expects JSON string value, with valid complex number and
// parses it. Returns error when value is not string or value cannot be parsed
// as complex64.
func decodeComplex64(value any) (complex64, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "decodeComplex64: requires boolean value got %T: %w"
		return 0i + 0, fmt.Errorf(format, value, ErrInvType)
	}
	cpx, err := strconv.ParseComplex(str, 64)
	if err != nil {
		format := "decodeComplex64: parsing string value to complex64: %w"
		return 0i + 0, fmt.Errorf(format, ErrInvFormat)
	}
	return complex64(cpx), nil
}

// decodeComplex128 expects JSON string value, with valid complex number and
// parses it. Returns error when value is not string or value cannot be parsed
// as complex128.
func decodeComplex128(value any) (complex128, error) {
	var ok bool
	var str string
	if str, ok = value.(string); !ok {
		format := "decodeComplex128: requires boolean value got %T: %w"
		return 0i + 0, fmt.Errorf(format, value, ErrInvType)
	}
	cpx, err := strconv.ParseComplex(str, 128)
	if err != nil {
		format := "decodeComplex128: parsing string value to complex128: %w"
		return 0i + 0, fmt.Errorf(format, ErrInvFormat)
	}
	return cpx, nil
}
