// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"errors"
	"math"
	"testing"

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
