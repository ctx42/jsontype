// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"math"
	"testing"
	"time"

	"github.com/ctx42/convert/pkg/convert"
	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/dump"
	"github.com/ctx42/testing/pkg/must"

	"github.com/ctx42/jsontype/internal/test"
)

func Test_registry(t *testing.T) {
	assert.NotNil(t, registry)
	assert.Len(t, 19, registry.reg)
}

func Test_Register(t *testing.T) {
	t.Run("new converter", func(t *testing.T) {
		// --- Given ---
		isolateRegistry(t)

		name := t.Name()

		cnv := func(any) (any, error) { return nil, nil }

		// --- When ---
		have := Register(name, cnv)

		// --- Then ---
		assert.Nil(t, have)

		assert.Same(t, cnv, registry.reg[name])
	})

	t.Run("overwrite existing converter", func(t *testing.T) {
		// --- Given ---
		isolateRegistry(t)

		name := t.Name()

		cnv0 := func(any) (any, error) { return nil, nil }
		Register(name, cnv0)

		cnv1 := func(any) (any, error) { return nil, nil }

		// --- When ---
		have := Register(name, cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)
	})

	t.Run("nil converter is nop", func(t *testing.T) {
		// --- Given ---
		isolateRegistry(t)

		name := t.Name()

		cnv := func(any) (any, error) { return nil, nil }
		Register(name, cnv)

		// --- When ---
		have := Register(name, nil)

		// --- Then ---
		assert.Nil(t, have)

		assert.Same(t, cnv, registry.reg[name])
	})
}

func Test_DefaultRegistry(t *testing.T) {
	// --- When ---
	have := DefaultRegistry()

	// --- Then ---
	assert.Len(t, 19, have.reg)

	assert.NotNil(t, have.Converter(Int))
	assert.NotNil(t, have.Converter(Int16))
	assert.NotNil(t, have.Converter(Int32))
	assert.NotNil(t, have.Converter(Int64))
	assert.NotNil(t, have.Converter(Int8))

	assert.NotNil(t, have.Converter(Uint))
	assert.NotNil(t, have.Converter(Uint16))
	assert.NotNil(t, have.Converter(Uint32))
	assert.NotNil(t, have.Converter(Uint64))
	assert.NotNil(t, have.Converter(Uint8))

	assert.NotNil(t, have.Converter(Float32))
	assert.NotNil(t, have.Converter(Float64))

	assert.NotNil(t, have.Converter(Byte))
	assert.NotNil(t, have.Converter(Rune))
	assert.NotNil(t, have.Converter(String))
	assert.NotNil(t, have.Converter(Bool))
	assert.NotNil(t, have.Converter(Time))
	assert.NotNil(t, have.Converter(Duration))
	assert.NotNil(t, have.Converter(Nil))
}

func Test_New(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		// --- Given ---
		v := 42

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int, have.typ)
		assert.Equal(t, 42, have.val)
	})

	t.Run("int8", func(t *testing.T) {
		// --- Given ---
		v := int8(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int8, have.typ)
		assert.Equal(t, int8(42), have.val)
	})

	t.Run("int16", func(t *testing.T) {
		// --- Given ---
		v := int16(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int16, have.typ)
		assert.Equal(t, int16(42), have.val)
	})

	t.Run("int32", func(t *testing.T) {
		// --- Given ---
		v := int32(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int32, have.typ)
		assert.Equal(t, int32(42), have.val)
	})

	t.Run("int64", func(t *testing.T) {
		// --- Given ---
		v := int64(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int64, have.typ)
		assert.Equal(t, int64(42), have.val)
	})

	t.Run("uint", func(t *testing.T) {
		// --- Given ---
		v := uint(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Uint, have.typ)
		assert.Equal(t, uint(42), have.val)
	})

	t.Run("uint8", func(t *testing.T) {
		// --- Given ---
		v := uint8(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Uint8, have.typ)
		assert.Equal(t, uint8(42), have.val)
	})

	t.Run("uint16", func(t *testing.T) {
		// --- Given ---
		v := uint16(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Uint16, have.typ)
		assert.Equal(t, uint16(42), have.val)
	})

	t.Run("uint32", func(t *testing.T) {
		// --- Given ---
		v := uint32(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Uint32, have.typ)
		assert.Equal(t, uint32(42), have.val)
	})

	t.Run("uint64", func(t *testing.T) {
		// --- Given ---
		v := uint64(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Uint64, have.typ)
		assert.Equal(t, uint64(42), have.val)
	})

	t.Run("float32", func(t *testing.T) {
		// --- Given ---
		v := float32(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Float32, have.typ)
		assert.Equal(t, float32(42), have.val)
	})

	t.Run("float64", func(t *testing.T) {
		// --- Given ---
		v := float64(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Float64, have.typ)
		assert.Equal(t, float64(42), have.val)
	})

	t.Run("byte", func(t *testing.T) {
		// --- Given ---
		v := byte(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Uint8, have.typ)
		assert.Equal(t, byte(42), have.val)
	})

	t.Run("rune", func(t *testing.T) {
		// --- Given ---
		v := rune(42)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int32, have.typ)
		assert.Equal(t, int32(42), have.val)
	})

	t.Run("string", func(t *testing.T) {
		// --- Given ---
		v := "abc"

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, String, have.typ)
		assert.Equal(t, "abc", have.val)
	})

	t.Run("bool", func(t *testing.T) {
		// --- Given ---
		v := true

		// --- When ---
		have := New(v)

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

	t.Run("interface holding a value", func(t *testing.T) {
		// --- Given ---
		var v any = 42

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Int, have.typ)
		assert.Equal(t, 42, have.val)
	})

	t.Run("nil interface", func(t *testing.T) {
		// --- Given ---
		var v error

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, Nil, have.typ)
		assert.Nil(t, have.val)
	})

	t.Run("interface holding a nil pointer", func(t *testing.T) {
		// --- Given ---
		var v any = (*int)(nil)

		// --- When ---
		have := New(v)

		// --- Then ---
		assert.Equal(t, "*int", have.typ)
		assert.Equal(t, (*int)(nil), have.val)
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
	type MyType int

	t.Run("nil", func(t *testing.T) {
		// --- When ---
		have, err := NewValue(nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Nil, have.typ)
		assert.Nil(t, have.val)
	})

	t.Run("registered type", func(t *testing.T) {
		// --- Given ---
		v := 42

		// --- When ---
		have, err := NewValue(v)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Int, have.typ)
		assert.Equal(t, 42, have.val)
	})

	t.Run("use custom registry", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.Register("jsontype.MyType", cnv)

		v := MyType(42)

		opt := WithRegistry(reg)

		// --- When ---
		have, err := NewValue(v, opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "jsontype.MyType", have.typ)
		assert.Equal(t, MyType(42), have.val)
	})

	t.Run("error - nil registry", func(t *testing.T) {
		// --- Given ---
		v := 42

		opt := WithRegistry(nil)

		// --- When ---
		have, err := NewValue(v, opt)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrNilRegistry, err)
		assert.ErrorEqual(t, "jsontype: nil registry", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value with nil registry", func(t *testing.T) {
		// --- Given ---
		opt := WithRegistry(nil)

		// --- When ---
		have, err := NewValue(nil, opt)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrNilRegistry, err)
		assert.ErrorEqual(t, "jsontype: nil registry", err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		v := MyType(42)

		// --- When ---
		have, err := NewValue(v)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "jsontype: unsupported type: jsontype.MyType", err)
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
	assert.Equal(t, map[string]any{"type": "uint", "value": uint(42)}, have)
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

	t.Run("success string", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: String, val: "abc"}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, `"abc"`, string(have))
	})

	t.Run("success bool true", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Bool, val: true}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "true", string(have))
	})

	t.Run("success bool false", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Bool, val: false}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "false", string(have))
	})

	t.Run("success nil", func(t *testing.T) {
		// --- Given ---
		val := must.Value(NewValue(nil))

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "null", string(have))
	})

	t.Run("success float64", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Float64, val: 4.2}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "4.2", string(have))
	})

	t.Run("success not addressable value", func(t *testing.T) {
		// --- Given ---
		m := map[string]Value{"k": {typ: Uint8, val: uint8(5)}}

		// --- When ---
		have, err := json.Marshal(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.JSON(t, `{"k": {"type": "uint8", "value": 5}}`, string(have))
	})

	t.Run("error - empty type", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "", val: nil}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: empty type: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("nil pointer", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have, err := json.Marshal(val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "null", string(have))
	})

	t.Run("error - string type mismatch", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: String, val: 42}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: string: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("error - bool type mismatch", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Bool, val: "x"}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: bool: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil type mismatch", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Nil, val: 42}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: nil: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("error - float64 type mismatch", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Float64, val: "x"}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: float64: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: "func()", val: func() {}}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: json: unsupported type: func()", err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported float64 value", func(t *testing.T) {
		// --- Given ---
		val := &Value{typ: Float64, val: math.NaN()}

		// --- When ---
		have, err := val.MarshalJSON()

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: json: unsupported value: NaN", err)
		assert.Nil(t, have)
	})
}

func Test_Value_MarshalJSON_roundtrip_tabular(t *testing.T) {
	tim := time.Date(2000, 1, 2, 3, 4, 5, 600000000, time.UTC)

	tt := []struct {
		testN string

		val *Value
	}{
		{"int max", New(math.MaxInt)},
		{"int min", New(math.MinInt)},
		{"int8 min", New(int8(math.MinInt8))},
		{"int16 max", New(int16(math.MaxInt16))},
		{"int32 min", New(int32(math.MinInt32))},
		{"int64 max", New(int64(math.MaxInt64))},
		{"int64 min", New(int64(math.MinInt64))},
		{"int64 above 2^53", New(int64(9007199254740993))},
		{"uint max", New(uint(math.MaxUint))},
		{"uint8 max", New(uint8(math.MaxUint8))},
		{"uint16 max", New(uint16(math.MaxUint16))},
		{"uint32 max", New(uint32(math.MaxUint32))},
		{"uint64 max", New(uint64(math.MaxUint64))},
		{"float32 fraction", New(float32(0.1))},
		{"float32 max", New(float32(math.MaxFloat32))},
		{"float32 smallest", New(float32(math.SmallestNonzeroFloat32))},
		{"float64 fraction", New(0.1)},
		{"float64 max", New(math.MaxFloat64)},
		{"byte via New", New(byte(42))},
		{"byte type name", &Value{typ: Byte, val: byte(42)}},
		{"rune via New", New('A')},
		{"rune type name", &Value{typ: Rune, val: 'A'}},
		{"string", New("abc")},
		{"bool", New(true)},
		{"nil", &Value{typ: Nil}},
		{"time.Time", New(tim)},
		{"time.Duration", New(time.Minute)},
		{"time.Duration max", New(time.Duration(math.MaxInt64))},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			data := must.Value(json.Marshal(tc.val))

			have := &Value{}

			// --- When ---
			err := json.Unmarshal(data, have)

			// --- Then ---
			assert.NoError(t, err)

			assert.Equal(t, tc.val.typ, have.typ)
			assert.Equal(t, tc.val.val, have.val)
		})
	}
}

func Test_Value_UnmarshalJSON(t *testing.T) {
	t.Run("success envelope", func(t *testing.T) {
		// --- Given ---
		data := []byte(`{"type": "uint8", "value": 42}`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Uint8, val.typ)
		assert.Equal(t, uint8(42), val.val)
	})

	t.Run("success bare string", func(t *testing.T) {
		// --- Given ---
		data := []byte(`"abc"`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, String, val.typ)
		assert.Equal(t, "abc", val.val)
	})

	t.Run("success bare bool true", func(t *testing.T) {
		// --- Given ---
		data := []byte(`true`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, true, val.val)
	})

	t.Run("success bare bool false", func(t *testing.T) {
		// --- Given ---
		data := []byte(`false`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, false, val.val)
	})

	t.Run("success bare null", func(t *testing.T) {
		// --- Given ---
		data := []byte(`null`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Nil, val.typ)
		assert.Nil(t, val.val)
	})

	t.Run("success bare float64", func(t *testing.T) {
		// --- Given ---
		data := []byte(`4.2`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Float64, val.typ)
		assert.Equal(t, 4.2, val.val)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		data := []byte(`{"type": "unknown", "value": 42}`)

		val := &Value{}

		// --- When ---
		err := val.UnmarshalJSON(data)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "jsontype: unsupported type: unknown", err)
	})
}

func Test_Value_UnmarshalJSON_success_tabular(t *testing.T) {
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
		{
			"bare string",
			"string",
			`"abc"`,
			"abc",
		},
		{
			"bare bool",
			"bool",
			`true`,
			true,
		},
		{
			"bare null",
			"nil",
			`null`,
			nil,
		},
		{
			"bare float64",
			"float64",
			`4.2`,
			4.2,
		},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- Given ---
			data := []byte(tc.json)

			val := &Value{}

			// --- When ---
			err := val.UnmarshalJSON(data)

			// --- Then ---
			assert.NoError(t, err)

			assert.Equal(t, tc.typ, val.typ)
			assert.Equal(t, tc.want, val.val)
		})
	}
}

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

	t.Run("success envelope max uint64", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()

		data := []byte(`{"type": "uint64", "value": 18446744073709551615}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Uint64, val.typ)
		assert.Equal(t, uint64(math.MaxUint64), val.val)
	})

	t.Run("success envelope large int64", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()

		data := []byte(`{"type": "int64", "value": -9007199254740993}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Int64, val.typ)
		assert.Equal(t, int64(-9007199254740993), val.val)
	})

	t.Run("success envelope duration in nanoseconds", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()

		data := []byte(`{"type": "time.Duration", "value": 60000000000}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Duration, val.typ)
		assert.Equal(t, time.Minute, val.val)
	})

	t.Run("success envelope float32", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()

		data := []byte(`{"type": "float32", "value": 0.1}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Float32, val.typ)
		assert.Equal(t, float32(0.1), val.val)
	})

	t.Run("success envelope custom converter gets float64", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()
		reg.Register(Uint64, func(value any) (any, error) { return value, nil })

		data := []byte(`{"type": "uint64", "value": 42}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Uint64, val.typ)
		assert.Equal(t, 42.0, val.val)
	})

	t.Run("success envelope without value", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()

		data := []byte(`{"type": "nil"}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Nil, val.typ)
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

	t.Run("success surrounding whitespace", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(" \t\r\ntrue\n")

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Bool, val.typ)
		assert.Equal(t, true, val.val)
	})

	t.Run("success bare float64 trailing newline", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte("1\n")

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, Float64, val.typ)
		assert.Equal(t, 1.0, val.val)
	})

	t.Run("error - infinity", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`-Inf`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		want := "jsontype: invalid character 'I' in numeric literal"
		assert.ErrorEqual(t, want, err)
	})

	t.Run("error - hexadecimal float", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`0x1p4`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		want := "jsontype: invalid character 'x' after top-level value"
		assert.ErrorEqual(t, want, err)
	})

	t.Run("success bare value with nil registry", func(t *testing.T) {
		// --- Given ---
		data := []byte(`"abc"`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(nil, data, val)

		// --- Then ---
		assert.NoError(t, err)

		assert.Equal(t, String, val.typ)
		assert.Equal(t, "abc", val.val)
	})

	t.Run("error - nil registry", func(t *testing.T) {
		// --- Given ---
		data := []byte(`{"type": "uint8", "value": 42}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(nil, data, val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrNilRegistry, err)
		assert.ErrorEqual(t, "jsontype: nil registry", err)
	})

	t.Run("error - nil value", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`"abc"`)

		// --- When ---
		err := Unmarshal(reg, data, nil)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: nil Value: invalid value", err)
	})

	t.Run("error - invalid bare string", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`"abc`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: unexpected end of JSON input", err)
	})

	t.Run("error - malformed bool true token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`tXYZ`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: invalid JSON token: invalid format"
		assert.ErrorEqual(t, want, err)
	})

	t.Run("error - malformed bool false token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`fals`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: invalid JSON token: invalid format"
		assert.ErrorEqual(t, want, err)
	})

	t.Run("error - malformed null token", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`nope`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: invalid JSON token: invalid format"
		assert.ErrorEqual(t, want, err)
	})

	t.Run("error - invalid number", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`1abc`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		want := "jsontype: invalid character 'a' after top-level value"
		assert.ErrorEqual(t, want, err)
	})

	t.Run("error - invalid JSON", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		data := []byte(`{!!!}`)

		val := &Value{}

		// --- When ---
		err := Unmarshal(reg, data, val)

		// --- Then ---
		want := "" +
			"jsontype: invalid character '!' " +
			"looking for beginning of object key string"
		assert.ErrorEqual(t, want, err)
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
		assert.ErrorEqual(t, "jsontype: unsupported type: unknown", err)
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
		want := "jsontype: invalid value: from string to time.Time"
		assert.ErrorEqual(t, want, err)
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
		assert.ErrorContain(t, "jsontype: invalid character '!'", err)
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
		assert.ErrorEqual(t, "jsontype: unsupported type: unknown", err)
	})
	t.Run("error - value unchanged on conversion error", func(t *testing.T) {
		// --- Given ---
		reg := DefaultRegistry()

		data := []byte(`{"type": "uint8", "value": 99999}`)

		val := &Value{typ: String, val: "abc"}

		// --- When ---
		err := unmarshalEnvelope(reg, data, val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvRange, err)

		assert.Equal(t, String, val.typ)
		assert.Equal(t, "abc", val.val)
	})
}

func Test_FromMap(t *testing.T) {
	type MyType int

	t.Run("use custom registry", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.Register("jsontype.MyType", cnv)

		m := map[string]any{"type": "jsontype.MyType", "value": MyType(42)}

		opt := WithRegistry(reg)

		// --- When ---
		have, err := FromMap(m, opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "jsontype.MyType", have.typ)
		assert.Equal(t, MyType(42), have.val)
	})

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

	t.Run("byte alias", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "byte", "value": uint8(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Byte, have.typ)
		assert.Equal(t, uint8(42), have.val)
	})

	t.Run("rune alias", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "rune", "value": int32(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Rune, have.typ)
		assert.Equal(t, int32(42), have.val)
	})

	t.Run("custom type name", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.Register("seconds", cnv)

		m := map[string]any{"type": "seconds", "value": time.Minute}

		opt := WithRegistry(reg)

		// --- When ---
		have, err := FromMap(m, opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "seconds", have.typ)
		assert.Equal(t, time.Minute, have.val)
	})

	t.Run("nil value without nil converter", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "nil", "value": nil}

		opt := WithRegistry(NewRegistry())

		// --- When ---
		have, err := FromMap(m, opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Nil, have.typ)
		assert.Nil(t, have.val)
	})

	t.Run("round trip of unmarshalled byte", func(t *testing.T) {
		// --- Given ---
		val := &Value{}
		must.Nil(json.Unmarshal([]byte(`{"type":"byte","value":65}`), val))

		m := val.Map()

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, val, have)
	})

	t.Run("error - missing value key", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint"}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: missing value field: invalid format"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})

	t.Run("error - unsupported type", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "unknown", "value": uint(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrUnsType, err)
		assert.ErrorEqual(t, "jsontype: unsupported type: unknown", err)
		assert.Nil(t, have)
	})

	t.Run("error - nil registry", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "nil", "value": nil}

		opt := WithRegistry(nil)

		// --- When ---
		have, err := FromMap(m, opt)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrNilRegistry, err)
		assert.ErrorEqual(t, "jsontype: nil registry", err)
		assert.Nil(t, have)
	})

	t.Run("error - value not of built-in type", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint", "value": Value{}}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		want := "" +
			"jsontype: types do not match: uint != jsontype.Value: " +
			"invalid value"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil value of built-in type", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "int", "value": nil}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		want := "jsontype: types do not match: int != nil: invalid value"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})

	t.Run("error - missing type key checked before value", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"value": struct{}{}}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: missing type field: invalid format"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})

	t.Run("error - missing type key", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"value": 42}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: missing type field: invalid format"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})

	t.Run("error - type key not a string", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": 44, "value": uint(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		assert.ErrorEqual(t, "jsontype: type field: invalid format", err)
		assert.Nil(t, have)
	})

	t.Run("error - types do not match", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "int", "value": uint(42)}

		// --- When ---
		have, err := FromMap(m)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		want := "jsontype: types do not match: int != uint: invalid value"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})

	t.Run("error - nil map", func(t *testing.T) {
		// --- When ---
		have, err := FromMap(nil)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvFormat, err)
		want := "jsontype: missing value field: invalid format"
		assert.ErrorEqual(t, want, err)
		assert.Nil(t, have)
	})
}

func Test_builtinGoName_tabular(t *testing.T) {
	tt := []struct {
		testN string

		typ  string
		want string
		ok   bool
	}{
		{"int", Int, Int, true},
		{"int8", Int8, Int8, true},
		{"int16", Int16, Int16, true},
		{"int32", Int32, Int32, true},
		{"int64", Int64, Int64, true},
		{"uint", Uint, Uint, true},
		{"uint8", Uint8, Uint8, true},
		{"uint16", Uint16, Uint16, true},
		{"uint32", Uint32, Uint32, true},
		{"uint64", Uint64, Uint64, true},
		{"float32", Float32, Float32, true},
		{"float64", Float64, Float64, true},
		{"byte", Byte, Uint8, true},
		{"rune", Rune, Int32, true},
		{"string", String, String, true},
		{"bool", Bool, Bool, true},
		{"time.Time", Time, Time, true},
		{"time.Duration", Duration, Duration, true},
		{"nil", Nil, Nil, true},
		{"custom", "seconds", "", false},
	}

	for _, tc := range tt {
		t.Run(tc.testN, func(t *testing.T) {
			// --- When ---
			have, ok := builtinGoName(tc.typ)

			// --- Then ---
			assert.Equal(t, tc.want, have)
			assert.Equal(t, tc.ok, ok)
		})
	}
}

func Test_AsValue(t *testing.T) {
	t.Run("instance of Value", func(t *testing.T) {
		// --- Given ---
		val := New(uint32(42))

		// --- When ---
		have, err := AsValue(val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint32, have.typ)
		assert.Equal(t, uint32(42), have.val)
	})

	t.Run("map", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"type": "uint", "value": uint(42)}

		// --- When ---
		have, err := AsValue(m)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint, have.typ)
		assert.Equal(t, uint(42), have.val)
	})

	t.Run("map with custom registry", func(t *testing.T) {
		// --- Given ---
		type MyType int

		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.Register("jsontype.MyType", cnv)

		m := map[string]any{"type": "jsontype.MyType", "value": MyType(42)}

		opt := WithRegistry(reg)

		// --- When ---
		have, err := AsValue(m, opt)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, "jsontype.MyType", have.typ)
		assert.Equal(t, MyType(42), have.val)
	})

	t.Run("Value passed by value", func(t *testing.T) {
		// --- Given ---
		val := *New(uint32(42))

		// --- When ---
		have, err := AsValue(val)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, Uint32, have.typ)
		assert.Equal(t, uint32(42), have.val)
	})

	t.Run("error - nil Value", func(t *testing.T) {
		// --- Given ---
		var val *Value

		// --- When ---
		have, err := AsValue(val)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvValue, err)
		assert.ErrorEqual(t, "jsontype: nil Value: invalid value", err)
		assert.Nil(t, have)
	})

	t.Run("error - not a map", func(t *testing.T) {
		// --- When ---
		have, err := AsValue(nil)

		// --- Then ---
		assert.ErrorIs(t, convert.ErrInvType, err)
		assert.ErrorEqual(t, "jsontype: invalid type", err)
		assert.Nil(t, have)
	})
}
