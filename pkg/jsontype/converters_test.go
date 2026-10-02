// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"errors"
	"math"
	"testing"
	"time"

	"github.com/ctx42/convert/pkg/convert"
	"github.com/ctx42/testing/pkg/assert"
)

func Test_NilConverter(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := NilConverter(nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Nil(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := NilConverter(42)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvType, err)
		wMsg := "jsontype: requires a nil value: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Nil(t, have)
	})
}

func Test_numberConverter(t *testing.T) {
	t.Run("json.Number", func(t *testing.T) {
		// --- Given ---
		cnv := numberConverter(convert.StringToUint64, convert.Float64ToUint64)
		num := json.Number("18446744073709551615")

		// --- When ---
		have, err := cnv(num)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint64(math.MaxUint64), have)
	})

	t.Run("float64", func(t *testing.T) {
		// --- Given ---
		cnv := numberConverter(convert.StringToUint64, convert.Float64ToUint64)
		f64 := 42.0

		// --- When ---
		have, err := cnv(f64)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, uint64(42), have)
	})

	t.Run("error - json.Number out of range", func(t *testing.T) {
		// --- Given ---
		cnv := numberConverter(convert.StringToUint8, convert.Float64ToUint8)
		num := json.Number("256")

		// --- When ---
		have, err := cnv(num)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvRange, err)
		want := "value out of range: from float64 to uint8"
		assert.ErrorEqual(t, want, err)
		assert.Equal(t, uint8(0), have)
	})

	t.Run("error - json.Number other error", func(t *testing.T) {
		// --- Given ---
		str := func(string) (int, error) { return 0, errors.New("my error") }
		cnv := numberConverter(str, convert.Float64ToInt)
		num := json.Number("42")

		// --- When ---
		have, err := cnv(num)

		// --- Then ---
		assert.ErrorEqual(t, "my error", err)
		assert.Equal(t, 0, have)
	})

	t.Run("error - invalid type", func(t *testing.T) {
		// --- Given ---
		cnv := numberConverter(convert.StringToUint8, convert.Float64ToUint8)
		str := "42"

		// --- When ---
		have, err := cnv(str)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvType, err)
		assert.Equal(t, uint8(0), have)
	})
}

func Test_durationConverter(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		// --- Given ---
		cnv := durationConverter()
		str := "1m30s"

		// --- When ---
		have, err := cnv(str)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 90*time.Second, have)
	})

	t.Run("json.Number", func(t *testing.T) {
		// --- Given ---
		cnv := durationConverter()
		num := json.Number("9223372036854775807")

		// --- When ---
		have, err := cnv(num)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, time.Duration(math.MaxInt64), have)
	})

	t.Run("float64", func(t *testing.T) {
		// --- Given ---
		cnv := durationConverter()
		f64 := 60e9

		// --- When ---
		have, err := cnv(f64)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, time.Minute, have)
	})

	t.Run("error - invalid string", func(t *testing.T) {
		// --- Given ---
		cnv := durationConverter()
		str := "abc"

		// --- When ---
		have, err := cnv(str)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.Equal(t, time.Duration(0), have)
	})

	t.Run("error - json.Number fraction", func(t *testing.T) {
		// --- Given ---
		cnv := durationConverter()
		num := json.Number("1.5")

		// --- When ---
		have, err := cnv(num)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrFraction, err)
		want := "must be a whole number: from float64 to time.Duration"
		assert.ErrorEqual(t, want, err)
		assert.Equal(t, time.Duration(0), have)
	})
}

func Test_nanosecondsToDuration(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		src := "60000000000"

		// --- When ---
		have, err := nanosecondsToDuration(src)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, time.Minute, have)
	})

	t.Run("error - out of range", func(t *testing.T) {
		// --- Given ---
		src := "9223372036854775808"

		// --- When ---
		have, err := nanosecondsToDuration(src)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvRange, err)
		want := "value out of range: from string to time.Duration"
		assert.ErrorEqual(t, want, err)
		assert.Equal(t, time.Duration(0), have)
	})
}
