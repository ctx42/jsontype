// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_DecodeInt_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int
		err  error
		msg  string
	}{
		{
			"int from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeInt: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeInt: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt: requires float64 value in range of int: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt: requires float64 value in range of int: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt8_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int8
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt8(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt8_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int8
		err  error
		msg  string
	}{
		{
			"int8 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeInt8: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeInt8: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt8: requires float64 value in range of int8: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt8: requires float64 value in range of int8: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt8(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt16_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int16
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt16(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt16_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int16
		err  error
		msg  string
	}{
		{
			"int16 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeInt16: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeInt16: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt16: requires float64 value in range of int16: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt16: requires float64 value in range of int16: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt16(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt32(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := DecodeInt32(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int32(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := DecodeInt32(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "DecodeInt32: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, int32(0), have)
	})
}

func Test_DecodeRune(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := DecodeRune(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, rune(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := DecodeRune(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "DecodeRune: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, rune(0), have)
	})
}

func Test_DecodeNamedInt32_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		name string
		have any
		want int32
	}{
		{"float64", "name", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeNamedInt32(tc.name, tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeNamedInt32_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		name string
		have any
		want int32
		err  error
		msg  string
	}{
		{
			"int32 from not expected JSON type",
			"name",
			42,
			0,
			ErrInvType,
			"name: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			"name",
			42.1,
			0,
			ErrInvValue,
			"name: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			"name",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"name: requires float64 value in range of int32: invalid range",
		},
		{
			"int out of range positive",
			"name",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"name: requires float64 value in range of int32: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeNamedInt32(tc.name, tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int64
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeInt64_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want int64
		err  error
		msg  string
	}{
		{
			"int64 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeInt64: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeInt64: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt64: requires float64 value in range of int64: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeInt64: requires float64 value in range of int64: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeInt64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint
		err  error
		msg  string
	}{
		{
			"uint from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeUInt: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeUInt: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt: requires float64 value in range of uint: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt: requires float64 value in range of uint: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt8(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := DecodeUInt8(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint8(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := DecodeUInt8(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "DecodeUInt8: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, uint8(0), have)
	})
}

func Test_DecodeByte(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := DecodeByte(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, byte(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := DecodeByte(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "DecodeByte: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, byte(0), have)
	})
}

func Test_DecodeNamedUInt8_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		name string
		have any
		want uint8
	}{
		{"float64", "name", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeNamedUInt8(tc.name, tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeNamedUInt8_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		name string
		have any
		want uint8
		err  error
		msg  string
	}{
		{
			"uint8 from not expected JSON type",
			"name",
			42,
			0,
			ErrInvType,
			"name: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			"name",
			42.1,
			0,
			ErrInvValue,
			"name: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			"name",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"name: requires float64 value in range of uint8: invalid range",
		},
		{
			"int out of range positive",
			"name",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"name: requires float64 value in range of uint8: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeNamedUInt8(tc.name, tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt16_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint16
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt16(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt16_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint16
		err  error
		msg  string
	}{
		{
			"uint16 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeUInt16: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeUInt16: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt16: requires float64 value in range of uint16: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt16: requires float64 value in range of uint16: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt16(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt32_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint32
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt32(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt32_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint32
		err  error
		msg  string
	}{
		{
			"uint32 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeUInt32: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeUInt32: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt32: requires float64 value in range of uint32: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt32: requires float64 value in range of uint32: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt32(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint64
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeUInt64_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want uint64
		err  error
		msg  string
	}{
		{
			"uint64 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeUInt64: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"DecodeUInt64: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt64: requires float64 value in range of uint64: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeUInt64: requires float64 value in range of uint64: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeUInt64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeFloat32_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float32
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeFloat32(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeFloat32_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float32
		err  error
		msg  string
	}{
		{
			"float32 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeFloat32: requires float64 value got int: invalid type",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeFloat32: requires float64 value in range of float32: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"DecodeFloat32: requires float64 value in range of float32: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeFloat32(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeFloat64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float64
	}{
		{"float64", 42.0, 42},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeFloat64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeFloat64_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want float64
		err  error
		msg  string
	}{
		{
			"float64 from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeFloat64: requires float64 value got int: invalid type",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeFloat64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeString_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want string
	}{
		{"string", "abc", "abc"},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeString(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeString_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want string
		err  error
		msg  string
	}{
		{
			"string from not expected JSON type",
			42,
			"",
			ErrInvType,
			"DecodeString: requires string value got int: invalid type",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeString(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeBool_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want bool
	}{
		{"true", true, true},
		{"false", false, false},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeBool(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeBool_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want bool
		err  error
		msg  string
	}{
		{
			"bool from not expected JSON type",
			42,
			false,
			ErrInvType,
			"DecodeBool: requires boolean value got int: invalid type",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeBool(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeTime_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want time.Time
	}{
		{
			"time",
			"2000-01-02T03:04:05.6Z",
			time.Date(2000, 1, 2, 3, 4, 5, 600000000, time.UTC),
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeTime(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeTime_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want time.Time
		err  error
		msg  string
	}{
		{
			"time from not expected JSON type",
			42,
			time.Time{},
			ErrInvType,
			"DecodeTime: requires boolean value got int: invalid type",
		},
		{
			"invalid time format",
			"abc",
			time.Time{},
			ErrInvFormat,
			"DecodeTime: parsing RFC3339 string value to time.Time: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeTime(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeDuration_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want time.Duration
	}{
		{"one second", "1s", time.Second},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeDuration(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeDuration_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want time.Duration
		err  error
		msg  string
	}{
		{
			"duration from not expected JSON type",
			42,
			time.Duration(0),
			ErrInvType,
			"DecodeDuration: requires boolean value got int: invalid type",
		},
		{
			"invalid duration format",
			"abc",
			time.Duration(0),
			ErrInvFormat,
			"DecodeDuration: parsing string value to time.Duration: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeDuration(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeComplex64_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want complex64
	}{
		{"one second", "4+2i", 4 + 2i},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeComplex64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeComplex64_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want complex64
		err  error
		msg  string
	}{
		{
			"duration from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeComplex64: requires boolean value got int: invalid type",
		},
		{
			"invalid duration format",
			"abc",
			0,
			ErrInvFormat,
			"DecodeComplex64: parsing string value to complex64: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeComplex64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeComplex128_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want complex128
	}{
		{"one second", "4+2i", 4 + 2i},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeComplex128(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeComplex128_error_tabular(t *testing.T) {
	tt := []struct {
		testN string

		have any
		want complex128
		err  error
		msg  string
	}{
		{
			"duration from not expected JSON type",
			42,
			0,
			ErrInvType,
			"DecodeComplex128: requires boolean value got int: invalid type",
		},
		{
			"invalid duration format",
			"abc",
			0,
			ErrInvFormat,
			"DecodeComplex128: parsing string value to complex128: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := DecodeComplex128(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_DecodeNil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := DecodeNil(nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Nil(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := DecodeNil(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		assert.ErrorEqual(t, "DecodeNil: requires nil value: invalid type", err)
		assert.Nil(t, have)
	})
}
