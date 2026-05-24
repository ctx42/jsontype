// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"testing"
)

func Benchmark_MarshalJSON_int(b *testing.B) {
	val := &Value{typ: Int, val: 42}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = val.MarshalJSON()
	}
}

func Benchmark_MarshalJSON_string(b *testing.B) {
	val := &Value{typ: String, val: "hello world"}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = val.MarshalJSON()
	}
}

func Benchmark_MarshalJSON_bool(b *testing.B) {
	val := &Value{typ: Bool, val: true}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = val.MarshalJSON()
	}
}

func Benchmark_MarshalJSON_nil(b *testing.B) {
	val := &Value{typ: Nil, val: nil}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = val.MarshalJSON()
	}
}

func Benchmark_MarshalJSON_float64(b *testing.B) {
	val := &Value{typ: Float64, val: 3.14}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = val.MarshalJSON()
	}
}

func Benchmark_Unmarshal_envelope_int(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`{"type":"int","value":42}`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = Unmarshal(reg, data, val)
	}
}

func Benchmark_Unmarshal_bare_string(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`"hello world"`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = Unmarshal(reg, data, val)
	}
}

func Benchmark_Unmarshal_bare_bool(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`true`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = Unmarshal(reg, data, val)
	}
}

func Benchmark_Unmarshal_bare_null(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`null`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = Unmarshal(reg, data, val)
	}
}

func Benchmark_Unmarshal_bare_float64(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`3.14`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = Unmarshal(reg, data, val)
	}
}

// Benchmark_legacy_* benchmarks measure the allocation floor that the
// transparent-type fast paths avoid, providing a stable baseline for
// regression comparisons.

func Benchmark_legacy_MarshalJSON_string(b *testing.B) {
	val := &Value{typ: String, val: "hello world"}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = json.Marshal(val.Map())
	}
}

func Benchmark_legacy_MarshalJSON_bool(b *testing.B) {
	val := &Value{typ: Bool, val: true}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = json.Marshal(val.Map())
	}
}

func Benchmark_legacy_MarshalJSON_nil(b *testing.B) {
	val := &Value{typ: Nil, val: nil}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = json.Marshal(val.Map())
	}
}

func Benchmark_legacy_MarshalJSON_float64(b *testing.B) {
	val := &Value{typ: Float64, val: 3.14}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = json.Marshal(val.Map())
	}
}

func Benchmark_legacy_Unmarshal_string(b *testing.B) {
	// Old code always decoded into the envelope struct.
	reg := DefaultRegistry()
	data := []byte(`{"type":"string","value":"hello world"}`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = unmarshalEnvelope(reg, data, val)
	}
}

func Benchmark_legacy_Unmarshal_bool(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`{"type":"bool","value":true}`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = unmarshalEnvelope(reg, data, val)
	}
}

func Benchmark_legacy_Unmarshal_null(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`{"type":"nil","value":null}`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = unmarshalEnvelope(reg, data, val)
	}
}

func Benchmark_legacy_Unmarshal_float64(b *testing.B) {
	reg := DefaultRegistry()
	data := []byte(`{"type":"float64","value":3.14}`)
	b.ReportAllocs()
	for b.Loop() {
		val := &Value{}
		_ = unmarshalEnvelope(reg, data, val)
	}
}

// Benchmark_json_Marshal_* benchmarks measure encoding/json overhead on
// Value to show the benefit of implementing json.Marshaler directly.

func Benchmark_json_Marshal_string(b *testing.B) {
	val := &Value{typ: String, val: "hello world"}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = json.Marshal(val)
	}
}

func Benchmark_json_Marshal_int(b *testing.B) {
	val := &Value{typ: Int, val: 42}
	b.ReportAllocs()
	for b.Loop() {
		_, _ = json.Marshal(val)
	}
}
