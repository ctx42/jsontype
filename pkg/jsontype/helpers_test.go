// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"
	"time"

	"github.com/ctx42/convert/pkg/convert"
	"github.com/ctx42/testing/pkg/assert"
)

func Test_Unmarshal(t *testing.T) {
	t.Run("success envelope", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		reg.Register(Uint8, convert.ToAnyAny(convert.Float64ToUint8))
		data := `{"type": "uint8", "value": 42}`
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(data), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("success envelope nil", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		reg.Register(Nil, NilConverter)
		data := `{"type": "nil", "value": null}`
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(data), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "nil", val.typ)
		assert.Nil(t, val.val)
	})

	t.Run("success bare string", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`"abc"`), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, String, val.typ)
		assert.Equal(t, "abc", val.val)
	})

	t.Run("success bare bool true", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`true`), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, true, val.val)
	})

	t.Run("success bare bool false", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`false`), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, false, val.val)
	})

	t.Run("success bare null", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`null`), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Nil, val.typ)
		assert.Nil(t, val.val)
	})

	t.Run("success bare float64", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`4.2`), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Float64, val.typ)
		assert.Equal(t, 4.2, val.val)
	})

	t.Run("success bare negative float64", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`-1.5`), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Float64, val.typ)
		assert.Equal(t, -1.5, val.val)
	})

	t.Run("error - malformed bool true token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`tXYZ`), val)

		// --- Then ---
		assert.ErrorContain(t, "invalid JSON token", err)
	})

	t.Run("error - malformed bool false token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`fals`), val)

		// --- Then ---
		assert.ErrorContain(t, "invalid JSON token", err)
	})

	t.Run("error - malformed null token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(`nope`), val)

		// --- Then ---
		assert.ErrorContain(t, "invalid JSON token", err)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := `{!!!}`
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(data), val)

		// --- Then ---
		assert.ErrorContain(t, "invalid character", err)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := `{"type": "unknown", "value": 42}`
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, []byte(data), val)

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
		err := Unmarshal(reg, []byte(data), val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		wMsg := "jsontype: invalid value: from string to time.Time"
		assert.ErrorEqual(t, wMsg, err)
	})
}

func Test_unmarshalEnvelope(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		reg.Register(Uint8, convert.ToAnyAny(convert.Float64ToUint8))
		data := `{"type": "uint8", "value": 42}`
		val := &Value{}

		// --- When ---
		err := unmarshalEnvelope(reg, []byte(data), val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := unmarshalEnvelope(reg, []byte(`{!!!}`), val)

		// --- Then ---
		assert.ErrorContain(t, "invalid character", err)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		val := &Value{}

		// --- When ---
		err := unmarshalEnvelope(reg, []byte(`{"type":"unknown","value":1}`), val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "unsupported type: unknown", err)
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
