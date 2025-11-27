// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"math"
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_decodeInt_success_tabular(t *testing.T) {
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
			have, err := decodeInt(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt_error_tabular(t *testing.T) {
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
			"decodeInt: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeInt: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt: requires float64 value in range of int: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt: requires float64 value in range of int: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeInt(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt8_success_tabular(t *testing.T) {
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
			have, err := decodeInt8(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt8_error_tabular(t *testing.T) {
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
			"decodeInt8: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeInt8: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt8: requires float64 value in range of int8: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt8: requires float64 value in range of int8: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeInt8(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt16_success_tabular(t *testing.T) {
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
			have, err := decodeInt16(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt16_error_tabular(t *testing.T) {
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
			"decodeInt16: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeInt16: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt16: requires float64 value in range of int16: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt16: requires float64 value in range of int16: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeInt16(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt32(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := decodeInt32(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, int32(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := decodeInt32(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "decodeInt32: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, int32(0), have)
	})
}

func Test_decodeRune(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := decodeRune(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, rune(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := decodeRune(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "decodeRune: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, rune(0), have)
	})
}

func Test_decodeNamedInt32_success_tabular(t *testing.T) {
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

func Test_decodeNamedInt32_error_tabular(t *testing.T) {
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

func Test_decodeInt64_success_tabular(t *testing.T) {
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
			have, err := decodeInt64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeInt64_error_tabular(t *testing.T) {
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
			"decodeInt64: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeInt64: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt64: requires float64 value in range of int64: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeInt64: requires float64 value in range of int64: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeInt64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt_success_tabular(t *testing.T) {
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
			have, err := decodeUInt(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt_error_tabular(t *testing.T) {
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
			"decodeUInt: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeUInt: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt: requires float64 value in range of uint: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt: requires float64 value in range of uint: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeUInt(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt8(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := decodeUInt8(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint8(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := decodeUInt8(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "decodeUInt8: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, uint8(0), have)
	})
}

func Test_decodeByte(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := decodeByte(42.0)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, byte(42), have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := decodeByte(42)

		// --- Then ---
		assert.ErrorIs(t, ErrInvType, err)
		wMsg := "decodeByte: requires float64 value got int: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Equal(t, byte(0), have)
	})
}

func Test_decodeNamedUInt8_success_tabular(t *testing.T) {
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

func Test_decodeNamedUInt8_error_tabular(t *testing.T) {
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

func Test_decodeUInt16_success_tabular(t *testing.T) {
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
			have, err := decodeUInt16(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt16_error_tabular(t *testing.T) {
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
			"decodeUInt16: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeUInt16: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt16: requires float64 value in range of uint16: " +
				"invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt16: requires float64 value in range of uint16: " +
				"invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeUInt16(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt32_success_tabular(t *testing.T) {
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
			have, err := decodeUInt32(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt32_error_tabular(t *testing.T) {
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
			"decodeUInt32: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeUInt32: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt32: requires float64 value in range of uint32: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt32: requires float64 value in range of uint32: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeUInt32(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt64_success_tabular(t *testing.T) {
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
			have, err := decodeUInt64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeUInt64_error_tabular(t *testing.T) {
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
			"decodeUInt64: requires float64 value got int: invalid type",
		},
		{
			"integer from fraction",
			42.1,
			0,
			ErrInvValue,
			"decodeUInt64: requires non-fractional value: invalid value",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt64: requires float64 value in range of uint64: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeUInt64: requires float64 value in range of uint64: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeUInt64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeFloat32_success_tabular(t *testing.T) {
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
			have, err := decodeFloat32(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeFloat32_error_tabular(t *testing.T) {
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
			"decodeFloat32: requires float64 value got int: invalid type",
		},
		{
			"int out of range negative",
			-math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeFloat32: requires float64 value in range of float32: invalid range",
		},
		{
			"int out of range positive",
			math.MaxFloat64,
			0,
			ErrInvRange,
			"decodeFloat32: requires float64 value in range of float32: invalid range",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeFloat32(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeFloat64_success_tabular(t *testing.T) {
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
			have, err := decodeFloat64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeFloat64_error_tabular(t *testing.T) {
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
			"decodeFloat64: requires float64 value got int: invalid type",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeFloat64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeString_success_tabular(t *testing.T) {
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
			have, err := decodeString(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeString_error_tabular(t *testing.T) {
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
			"decodeString: requires string value got int: invalid type",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeString(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeBool_success_tabular(t *testing.T) {
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
			have, err := decodeBool(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeBool_error_tabular(t *testing.T) {
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
			"decodeBool: requires boolean value got int: invalid type",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeBool(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeTime_success_tabular(t *testing.T) {
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
			have, err := decodeTime(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeTime_error_tabular(t *testing.T) {
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
			"decodeTime: requires boolean value got int: invalid type",
		},
		{
			"invalid time format",
			"abc",
			time.Time{},
			ErrInvFormat,
			"decodeTime: parsing RFC3339 string value to time.Time: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeTime(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeDuration_success_tabular(t *testing.T) {
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
			have, err := decodeDuration(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeDuration_error_tabular(t *testing.T) {
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
			"decodeDuration: requires boolean value got int: invalid type",
		},
		{
			"invalid duration format",
			"abc",
			time.Duration(0),
			ErrInvFormat,
			"decodeDuration: parsing string value to time.Duration: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeDuration(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeComplex64_success_tabular(t *testing.T) {
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
			have, err := decodeComplex64(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeComplex64_error_tabular(t *testing.T) {
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
			"decodeComplex64: requires boolean value got int: invalid type",
		},
		{
			"invalid duration format",
			"abc",
			0,
			ErrInvFormat,
			"decodeComplex64: parsing string value to complex64: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeComplex64(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeComplex128_success_tabular(t *testing.T) {
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
			have, err := decodeComplex128(tc.have)

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.want, have)
		})
	}
}

func Test_decodeComplex128_error_tabular(t *testing.T) {
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
			"decodeComplex128: requires boolean value got int: invalid type",
		},
		{
			"invalid duration format",
			"abc",
			0,
			ErrInvFormat,
			"decodeComplex128: parsing string value to complex128: " +
				"invalid format",
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, err := decodeComplex128(tc.have)

			// --- Then ---
			assert.ErrorIs(t, tc.err, err)
			assert.ErrorEqual(t, tc.msg, err)
			assert.Equal(t, tc.want, have)
		})
	}
}
