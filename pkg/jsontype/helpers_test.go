// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"
	"time"

	"github.com/ctx42/convert/pkg/convert"
	"github.com/ctx42/testing/pkg/assert"
)

func Test_UnmarshallJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		reg.Register(Uint8, convert.ToAnyAny(convert.Float64ToUint8))
		data := `{"type": "uint8", "value": 42}`
		val := &Value{}

		// --- When ---
		err := UnmarshalJSON(reg, []byte(data), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("nil", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		reg.Register(Nil, NilConverter)
		data := `{"type": "nil", "value": null}`
		val := &Value{}

		// --- When ---
		err := UnmarshalJSON(reg, []byte(data), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "nil", val.typ)
		assert.Nil(t, val.val)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := `{!!!}`
		val := &Value{}

		// --- When ---
		err := UnmarshalJSON(reg, []byte(data), val)

		// --- Then ---
		assert.ErrorContain(t, "invalid character", err)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := `{"type": "unknown", "value": 42}`
		val := &Value{}

		// --- When ---
		err := UnmarshalJSON(reg, []byte(data), val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "unsupported type: unknown", err)
	})

	t.Run("error - invalid format", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		cnv := convert.StringToTime(time.RFC3339Nano)
		reg.Register(Time, convert.ToAnyAny(cnv))
		data := `{"type": "time.Time", "value": "abc"}`
		val := &Value{}

		// --- When ---
		err := UnmarshalJSON(reg, []byte(data), val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		wMsg := "jsontype: invalid value: from string to time.Time"
		assert.ErrorEqual(t, wMsg, err)
	})
}

func Test_keyValue(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"A": 1, "B": 2}

		// --- When ---
		have, ok := keyValue("B", m)

		// --- Then ---
		assert.Equal(t, 2, have)
		assert.True(t, ok)
	})

	t.Run("nil map", func(t *testing.T) {
		// --- When ---
		have, ok := keyValue("B", nil)

		// --- Then ---
		assert.Nil(t, have)
		assert.False(t, ok)
	})
}
