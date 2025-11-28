// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"
	"time"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_init(t *testing.T) {
	assert.NotNil(t, registry)
	assert.Len(t, 20, registry.reg)
}

func Test_New(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		// --- When ---
		have := New(42)

		// --- Then ---
		assert.Equal(t, Int, have.typ)
		assert.Equal(t, 42, have.val)
	})

	t.Run("int8", func(t *testing.T) {
		// --- When ---
		have := New(int8(42))

		// --- Then ---
		assert.Equal(t, Int8, have.typ)
		assert.Equal(t, int8(42), have.val)
	})

	t.Run("int16", func(t *testing.T) {
		// --- When ---
		have := New(int16(42))

		// --- Then ---
		assert.Equal(t, Int16, have.typ)
		assert.Equal(t, int16(42), have.val)
	})

	t.Run("int32", func(t *testing.T) {
		// --- When ---
		have := New(int32(42))

		// --- Then ---
		assert.Equal(t, Int32, have.typ)
		assert.Equal(t, int32(42), have.val)
	})

	t.Run("int64", func(t *testing.T) {
		// --- When ---
		have := New(int64(42))

		// --- Then ---
		assert.Equal(t, Int64, have.typ)
		assert.Equal(t, int64(42), have.val)
	})

	t.Run("uint", func(t *testing.T) {
		// --- When ---
		have := New(uint(42))

		// --- Then ---
		assert.Equal(t, UInt, have.typ)
		assert.Equal(t, uint(42), have.val)
	})

	t.Run("uint8", func(t *testing.T) {
		// --- When ---
		have := New(uint8(42))

		// --- Then ---
		assert.Equal(t, UInt8, have.typ)
		assert.Equal(t, uint8(42), have.val)
	})

	t.Run("uint16", func(t *testing.T) {
		// --- When ---
		have := New(uint16(42))

		// --- Then ---
		assert.Equal(t, UInt16, have.typ)
		assert.Equal(t, uint16(42), have.val)
	})

	t.Run("uint32", func(t *testing.T) {
		// --- When ---
		have := New(uint32(42))

		// --- Then ---
		assert.Equal(t, UInt32, have.typ)
		assert.Equal(t, uint32(42), have.val)
	})

	t.Run("uint64", func(t *testing.T) {
		// --- When ---
		have := New(uint64(42))

		// --- Then ---
		assert.Equal(t, UInt64, have.typ)
		assert.Equal(t, uint64(42), have.val)
	})

	t.Run("float32", func(t *testing.T) {
		// --- When ---
		have := New(float32(42))

		// --- Then ---
		assert.Equal(t, Float32, have.typ)
		assert.Equal(t, float32(42), have.val)
	})

	t.Run("float64", func(t *testing.T) {
		// --- When ---
		have := New(float64(42))

		// --- Then ---
		assert.Equal(t, Float64, have.typ)
		assert.Equal(t, float64(42), have.val)
	})

	t.Run("complex64", func(t *testing.T) {
		// --- When ---
		have := New(complex64(4i + 2))

		// --- Then ---
		assert.Equal(t, Complex64, have.typ)
		assert.Equal(t, complex64(4i+2), have.val)
	})

	t.Run("complex128", func(t *testing.T) {
		// --- When ---
		have := New(4i + 2)

		// --- Then ---
		assert.Equal(t, Complex128, have.typ)
		assert.Equal(t, 4i+2, have.val)
	})

	t.Run("byte", func(t *testing.T) {
		// --- When ---
		have := New(byte(42))

		// --- Then ---
		assert.Equal(t, UInt8, have.typ)
		assert.Equal(t, byte(42), have.val)
	})

	t.Run("rune", func(t *testing.T) {
		// --- When ---
		have := New(rune(42))

		// --- Then ---
		assert.Equal(t, Int32, have.typ)
		assert.Equal(t, int32(42), have.val)
	})

	t.Run("string", func(t *testing.T) {
		// --- When ---
		have := New("abc")

		// --- Then ---
		assert.Equal(t, String, have.typ)
		assert.Equal(t, "abc", have.val)
	})

	t.Run("bool", func(t *testing.T) {
		// --- When ---
		have := New(true)

		// --- Then ---
		assert.Equal(t, Bool, have.typ)
		assert.Equal(t, true, have.val)
	})

	t.Run("duration", func(t *testing.T) {
		// --- Given ---
		v := time.Minute

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Duration, have.typ)
		assert.Equal(t, v, have.val)
	})

	t.Run("time.Time", func(t *testing.T) {
		// --- Given ---
		v := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Time, have.typ)
		assert.Equal(t, v, have.val)
	})
}

func Test_NewByte(t *testing.T) {
	// --- When ---
	have := NewByte(42)

	// --- Then ---
	assert.Equal(t, Byte, have.typ)
	assert.Equal(t, byte(42), have.val)
}

func Test_NewRune(t *testing.T) {
	// --- When ---
	have := NewRune('*')

	// --- Then ---
	assert.Equal(t, Rune, have.typ)
	assert.Equal(t, rune(42), have.val)
}

func Test_Value_GoType(t *testing.T) {
	// --- Given ---
	val := &Value{typ: String, val: "abc"}

	// --- When ---
	have := val.GoType()

	// --- Then ---
	assert.Equal(t, String, have)
}

func Test_Value_GoValue(t *testing.T) {
	// --- Given ---
	val := &Value{typ: String, val: "abc"}

	// --- When ---
	have := val.GoValue()

	// --- Then ---
	assert.Equal(t, "abc", have)
}

func Test_Value_MarshalJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "type", val: 42}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.JSON(t, `{"type":"type","value":42}`, string(have))
	})

	t.Run("error - empty type", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "", val: nil}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, ErrInvValue, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil Value", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, ErrInvValue, err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "func", val: func() {}}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorEqual(t, "json: unsupported type: func()", err)
		assert.Nil(t, have)
	})
}

func Test_Value_UnmarshallJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		data := `{"type": "uint8", "value": 42}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, UInt8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		data := `{!!!}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.ErrorContain(t, "invalid character", err)
	})

	t.Run("error - unknown type", func(t *testing.T) {
		// --- Given ---
		data := `{"type": "unknown", "value": 42}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.ErrorIs(t, ErrUnkType, err)
		assert.ErrorEqual(t, "unknown type: unknown", err)
	})

	t.Run("error - invalid format", func(t *testing.T) {
		// --- Given ---
		data := `{"type": "time.Time", "value": "abc"}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.ErrorIs(t, ErrInvFormat, err)
		wMsg := "decodeTime: parsing RFC3339 string value to time.Time: " +
			"invalid format"
		assert.ErrorEqual(t, wMsg, err)
	})
}
