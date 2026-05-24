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
		data := []byte(`{"type": "uint8", "value": 42}`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("success envelope nil", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		reg.Register(Nil, NilConverter)
		data := []byte(`{"type": "nil", "value": null}`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "nil", val.typ)
		assert.Nil(t, val.val)
	})

	t.Run("success bare string", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`"abc"`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, String, val.typ)
		assert.Equal(t, "abc", val.val)
	})

	t.Run("success bare bool true", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`true`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, true, val.val)
	})

	t.Run("success bare bool false", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`false`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, false, val.val)
	})

	t.Run("success bare null", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`null`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Nil, val.typ)
		assert.Nil(t, val.val)
	})

	t.Run("success bare float64", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`4.2`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Float64, val.typ)
		assert.Equal(t, 4.2, val.val)
	})

	t.Run("success bare negative float64", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`-1.5`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Float64, val.typ)
		assert.Equal(t, -1.5, val.val)
	})

	t.Run("error - malformed bool true token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`tXYZ`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: invalid JSON token", err)
	})

	t.Run("error - malformed bool false token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`fals`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: invalid JSON token", err)
	})

	t.Run("error - malformed null token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`nope`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: invalid JSON token", err)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`{!!!}`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorContain(t, "invalid character", err)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`{"type": "unknown", "value": 42}`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "unsupported type: unknown", err)
	})

	t.Run("error - invalid format", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		cnv := convert.StringToTime(time.RFC3339Nano)
		reg.Register(Time, convert.ToAnyAny(cnv))
		data := []byte(`{"type": "time.Time", "value": "abc"}`)
		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

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
		data := []byte(`{"type": "uint8", "value": 42}`)
		val := &Value{}

		// --- When ---
		err := unmarshalEnvelope(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`{!!!}`)
		val := &Value{}

		// --- When ---
		err := unmarshalEnvelope(reg, data, val)

		// --- Then ---
		assert.ErrorContain(t, "invalid character", err)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()
		data := []byte(`{"type":"unknown","value":1}`)
		val := &Value{}

		// --- When ---
		err := unmarshalEnvelope(reg, data, val)

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
