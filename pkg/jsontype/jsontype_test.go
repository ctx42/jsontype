// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"
	"time"

	"github.com/ctx42/convert/pkg/convert"
	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/dump"
	"github.com/ctx42/testing/pkg/must"

	"github.com/ctx42/jsontype/internal/test"
)

func Test_init(t *testing.T) {
	assert.NotNil(t, registry)
	assert.Len(t, 19, registry.reg)
}

func Test_Register(t *testing.T) {
	t.Run("new converter", func(t *testing.T) {
		// --- Given ---
		cnv := func(any) (any, error) { return nil, nil }
		name := t.Name()

		// --- When ---
		have := Register(name, cnv)

		// --- Then ---
		assert.Nil(t, have)
		assert.Same(t, cnv, registry.reg[name])
	})

	t.Run("overwrite existing converter", func(t *testing.T) {
		// --- Given ---
		cnv0 := func(any) (any, error) { return nil, nil }
		cnv1 := func(any) (any, error) { return nil, nil }
		name := t.Name()
		Register(name, cnv0)

		// --- When ---
		have := Register(name, cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)
	})

	t.Run("nil converter is nop", func(t *testing.T) {
		// --- Given ---
		cnv := func(any) (any, error) { return nil, nil }
		name := t.Name()
		Register(name, cnv)

		// --- When ---
		have := Register(name, nil)

		// --- Then ---
		assert.Nil(t, have)
		assert.Same(t, cnv, registry.reg[name])
	})
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
		assert.Equal(t, Uint, have.typ)
		assert.Equal(t, uint(42), have.val)
	})

	t.Run("uint8", func(t *testing.T) {
		// --- When ---
		have := New(uint8(42))

		// --- Then ---
		assert.Equal(t, Uint8, have.typ)
		assert.Equal(t, uint8(42), have.val)
	})

	t.Run("uint16", func(t *testing.T) {
		// --- When ---
		have := New(uint16(42))

		// --- Then ---
		assert.Equal(t, Uint16, have.typ)
		assert.Equal(t, uint16(42), have.val)
	})

	t.Run("uint32", func(t *testing.T) {
		// --- When ---
		have := New(uint32(42))

		// --- Then ---
		assert.Equal(t, Uint32, have.typ)
		assert.Equal(t, uint32(42), have.val)
	})

	t.Run("uint64", func(t *testing.T) {
		// --- When ---
		have := New(uint64(42))

		// --- Then ---
		assert.Equal(t, Uint64, have.typ)
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

	t.Run("byte", func(t *testing.T) {
		// --- When ---
		have := New(byte(42))

		// --- Then ---
		assert.Equal(t, Uint8, have.typ)
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

	t.Run("time.Time", func(t *testing.T) {
		// --- Given ---
		v := time.Date(2000, 1, 2, 3, 4, 5, 0, time.UTC)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Time, have.typ)
		assert.Equal(t, v, have.val)
	})

	t.Run("time.Duration", func(t *testing.T) {
		// --- Given ---
		v := time.Minute

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Duration, have.typ)
		assert.Equal(t, v, have.val)
	})

	t.Run("type from a standard library", func(t *testing.T) {
		// --- Given ---
		v := test.Type{}

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, "test.Type", have.typ)
		assert.Equal(t, v, have.val)
	})

	t.Run("type from external module library", func(t *testing.T) {
		// --- Given ---
		v := dump.Dump{}

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, "dump.Dump", have.typ)
		assert.Equal(t, v, have.val)
	})
}

func Test_NewValue(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		// --- When ---
		have, err := NewValue(nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Nil, have.typ)
		assert.Nil(t, have.val)
	})

	t.Run("registered type", func(t *testing.T) {
		// --- When ---
		have, err := NewValue(42)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Int, have.typ)
		assert.Equal(t, 42, have.val)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- When ---
		have, err := NewValue(Value{})

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "unsupported type: jsontype.Value", err)
		assert.Nil(t, have)
	})
}

func Test_Value_GoTypeName(t *testing.T) {
	// --- Given ---
	val := &Value{typ: String, val: "abc"}

	// --- When ---
	have := val.GoTypeName()

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

func Test_Value_Map(t *testing.T) {
	// --- Given ---
	val := &Value{typ: Uint, val: uint(42)}

	// --- When ---
	have := val.Map()

	// --- Then ---
	want := map[string]any{"type": "uint", "value": uint(42)}
	assert.Equal(t, want, have)
}

func Test_Value_MarshalJSON(t *testing.T) {
	t.Run("success int", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Int, val: 42}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.JSON(t, `{"type":"int","value":42}`, string(have))
	})

	t.Run("success nil", func(t *testing.T) {
		// --- Given ---
		val := must.Value(NewValue(nil))

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.JSON(t, `{"type":"nil","value":null}`, string(have))
	})

	t.Run("error - empty type", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "", val: nil}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil Value", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "func()", val: func() {}}

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
		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("nil", func(t *testing.T) {
		// --- Given ---
		data := `{"type": "nil", "value": null}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "nil", val.typ)
		assert.Nil(t, val.val)
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

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		data := `{"type": "unknown", "value": 42}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "unsupported type: unknown", err)
	})

	t.Run("error - invalid format", func(t *testing.T) {
		// --- Given ---
		data := `{"type": "time.Time", "value": "abc"}`
		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON([]byte(data))

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		wMsg := "jsontype: invalid value: from string to time.Time"
		assert.ErrorEqual(t, wMsg, err)
	})
}

func Test_FromMap(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint", "value": uint(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint, have.typ)
		assert.Equal(t, uint(42), have.val)
	})

	t.Run("error - missing value key", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint"}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		wMsg := "FromMap: missing value field: invalid format"
		assert.ErrorEqual(t, wMsg, err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported value key type", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint", "value": Value{}}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "FromMap: unsupported type: jsontype.Value", err)
		assert.Nil(t, have)
	})

	t.Run("error - missing type key", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"value": 42}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		assert.ErrorEqual(t, "FromMap: missing type field: invalid format", err)
		assert.Nil(t, have)
	})

	t.Run("error - type key not a string", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": 44, "value": uint(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		assert.ErrorEqual(t, "FromMap: type field: invalid format", err)
		assert.Nil(t, have)
	})

	t.Run("error - types do not match", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "int", "value": uint(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "FromMap: types do not match: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil map", func(t *testing.T) {
		// --- When ---
		have, err := FromMap(nil)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		assert.ErrorEqual(t, "FromMap: missing value field: invalid format", err)
		assert.Nil(t, have)
	})
}

func Test_AsValue(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint", "value": uint(42)}

		// --- When ---
		have, err := AsValue(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint, have.typ)
		assert.Equal(t, uint(42), have.val)
	})

	t.Run("error - not a map", func(t *testing.T) {
		// --- When ---
		have, err := AsValue(nil)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvType, err)
		assert.ErrorEqual(t, "AsValue: invalid type", err)
		assert.Nil(t, have)
	})
}

func Test_Value_UnmarshallJSON_success_tabular(t *testing.T) {
	tt := []struct {
		testN string

		typ  string
		json string
		want any
	}{
		{
			"int",
			"int",
			`{"type": "int", "value": 42}`,
			42,
		},
		{
			"int8",
			"int8",
			`{"type": "int8", "value": 42}`,
			int8(42),
		},
		{
			"int16",
			"int16",
			`{"type": "int16", "value": 42}`,
			int16(42),
		},
		{
			"int32",
			"int32",
			`{"type": "int32", "value": 42}`,
			int32(42),
		},
		{
			"int64",
			"int64",
			`{"type": "int64", "value": 42}`,
			int64(42),
		},
		{
			"uint",
			"uint",
			`{"type": "uint", "value": 42}`,
			uint(42),
		},
		{
			"uint8",
			"uint8",
			`{"type": "uint8", "value": 42}`,
			uint8(42),
		},
		{
			"uint16",
			"uint16",
			`{"type": "uint16", "value": 42}`,
			uint16(42),
		},
		{
			"uint32",
			"uint32",
			`{"type": "uint32", "value": 42}`,
			uint32(42),
		},
		{
			"uint64",
			"uint64",
			`{"type": "uint64", "value": 42}`,
			uint64(42),
		},
		{
			"float32",
			"float32",
			`{"type": "float32", "value": 42}`,
			float32(42),
		},
		{
			"float64",
			"float64",
			`{"type": "float64", "value": 4.2}`,
			4.2,
		},
		{
			"byte",
			"byte",
			`{"type": "byte", "value": 42}`,
			byte(42),
		},
		{
			"rune",
			"rune",
			`{"type": "rune", "value": 42}`,
			rune(42),
		},
		{
			"string",
			"string",
			`{"type": "string", "value": "abc"}`,
			"abc",
		},
		{
			"bool",
			"bool",
			`{"type": "bool", "value": true}`,
			true,
		},
		{
			"time.Duration",
			"time.Duration",
			`{"type": "time.Duration", "value": "42s"}`,
			42 * time.Second,
		},
		{
			"time.Time",
			"time.Time",
			`{"type": "time.Time", "value": "2000-01-02T03:04:05.6Z"}`,
			time.Date(2000, time.January, 2, 3, 4, 5, 600000000, time.UTC),
		},
		{
			"nil",
			"nil",
			`{"type": "nil", "value": null}`,
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			val := &Value{}

			// --- When ---
			err := val.UnmarshalJSON([]byte(tc.json))

			// --- Then ---
			assert.NoError(t, err)
			assert.Equal(t, tc.typ, val.typ)
			assert.Equal(t, tc.want, val.val)
		})
	}
}
