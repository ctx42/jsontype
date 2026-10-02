[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/jsontype)](https://goreportcard.com/report/github.com/ctx42/jsontype)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/jsontype/pkg/jsontype)
![Tests](https://github.com/ctx42/jsontype/actions/workflows/go.yml/badge.svg?branch=master)

# jsontype

Keep Go types intact when values stored or sent as JSON are read back into
`any`.

<!-- TOC -->
  * [Why Use `jsontype`?](#why-use-jsontype)
  * [Features](#features)
  * [Installation](#installation)
  * [Usage](#usage)
    * [Marshal and Unmarshal](#marshal-and-unmarshal)
    * [Safe Conversion](#safe-conversion)
    * [Construct Values](#construct-values)
    * [Map Form](#map-form)
  * [Type Registry](#type-registry)
  * [Custom Converters](#custom-converters)
    * [Custom Registry](#custom-registry)
<!-- TOC -->

`jsontype` is a small Go module that preserves Go types when marshaling values
to JSON. It embeds type information directly into the JSON alongside the value.

## Why Use `jsontype`?

Standard Go JSON marshaling loses specific type information for `any`
fields. For example, a `uint64` becomes a `float64` after a
round-trip through JSON if unmarshaled into a `map[string]any`. `jsontype`
solves this by explicitly storing the type name.

This is especially useful for:
- Storing data in databases as JSON.
- Passing typed messages over a bus where the receiver uses `map[string]any`.

## Features

- Keeps the Go type of a value: `{"type":"uint64","value":42}` unmarshals back
  to a `uint64`.
- Converts exactly: a value that would overflow, underflow, or lose precision
  is an error instead of a silently changed number, and 64-bit integers beyond
  2^53 keep every digit.
- Writes `string`, `bool`, `nil`, and `float64` as plain JSON values, which
  JSON already reads back as the same Go types.
- Supports custom type names through converters, registered in the
  package-level registry or in a registry of your own.
- Converts to and from a map form for code that works with `map[string]any`.

## Installation

Requires Go 1.26 or newer. Install using `go get`:

```bash
go get github.com/ctx42/jsontype
```

The package lives in the `pkg/jsontype` directory of the module:

```go
import "github.com/ctx42/jsontype/pkg/jsontype"
```

## Usage

### Marshal and Unmarshal

Create a `Value` instance encapsulating the value and its type.

<!-- gmmce:pkg/jsontype/ExampleValue_MarshalJSON -->
```go
jType := jsontype.New(uint(42))
data, _ := json.Marshal(jType)

fmt.Println(string(data))
// Output:
// {"type":"uint","value":42}
```

When unmarshalling, the library reads the `type` field, finds the matching
converter in the package-level registry, and converts the value with it.

<!-- gmmce:pkg/jsontype/ExampleValue_UnmarshalJSON -->
```go
data := []byte(`{"type": "uint", "value": 42}`)
gType := &jsontype.Value{}
_ = json.Unmarshal(data, gType)

fmt.Printf("%[1]v (%[1]T)\n", gType.GoValue())
// Output:
// 42 (uint)
```

### Safe Conversion

A value that does not fit its type is an error, not a silently changed number.

<!-- gmmce:pkg/jsontype/ExampleValue_UnmarshalJSON_safe -->
```go
data := []byte(`{"type": "uint8", "value":99999}`)

gType := &jsontype.Value{}
err := json.Unmarshal(data, gType)

fmt.Println(err)
// Output:
// jsontype: value out of range: from float64 to uint8
```

### Construct Values

`New` takes the type name from the Go type. `NewValue` also accepts an untyped
`nil` and returns an error for a type without a registered converter.

<!-- gmmce:pkg/jsontype/ExampleNewValue -->
```go
val, err := jsontype.NewValue(int8(42))
fmt.Println(val.GoTypeName(), err)

_, err = jsontype.NewValue(struct{}{})
fmt.Println(err)
// Output:
// int8 <nil>
// jsontype: unsupported type: struct {}
```

### Map Form

`Value.Map` returns the `{"type": ..., "value": ...}` map, and `FromMap`
builds a `Value` back from it. `AsValue` accepts either a `*Value` or such a
map.

<!-- gmmce:pkg/jsontype/ExampleAsValue -->
```go
m := jsontype.New(uint16(42)).Map()

val, err := jsontype.AsValue(m)
if err != nil {
	log.Fatal(err)
}

fmt.Printf("%v (%s)\n", val.GoValue(), val.GoTypeName())
// Output: 42 (uint16)
```

## Type Registry

The package-level registry provides converters for the following types:

- `int`
- `int8`
- `int16`
- `int32`
- `int64`
- `uint`
- `uint8`
- `uint16`
- `uint32`
- `uint64`
- `float32`
- `float64`
- `byte`
- `rune`
- `string`
- `bool`
- `time.Duration`
- `time.Time`
- `nil`

Values of type `string`, `bool`, `nil`, and `float64` are written as plain JSON
values, for example `New("abc")` marshals to `"abc"`, and a plain JSON string,
boolean, `null`, or number unmarshals to one of these types. All other types use
the `{"type": ..., "value": ...}` form, which is accepted for every type.

## Custom Converters

You may register a custom converter for your custom type.

<!-- gmmce:pkg/jsontype/ExampleRegister_custom -->
```go
// Custom converter for a type named "seconds" representing
// duration in seconds.
cnv := func(value float64) (time.Duration, error) {
	return time.Duration(value) * time.Second, nil
}

// Register the converter in the package-level registry. It stays
// registered for the rest of the program and applies to every
// json.Unmarshal of a Value.
jsontype.Register("seconds", convert.ToAnyAny(cnv))

// Custom type named "seconds" representing duration in seconds.
data := []byte(`{"type": "seconds", "value": 42}`)

gType := &jsontype.Value{}
if err := json.Unmarshal(data, gType); err != nil {
	log.Fatal(err)
}

fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())
// Output:
// unmarshalled: 42s (time.Duration)
```

A registered converter must return an error when converting a JSON value to the
Go type would lose precision, overflow, underflow, or is not possible at all,
the same way the [github.com/ctx42/convert](https://github.com/ctx42/convert)
conversion functions do. It receives JSON numbers as `float64`.

### Custom Registry

To keep converters out of the package-level registry, register them in your
own `Registry` and pass it to `Unmarshal`.

<!-- gmmce:pkg/jsontype/ExampleUnmarshal -->
```go
// Register converters in a custom registry instead of the package-level
// one used by json.Unmarshal.
reg := jsontype.NewRegistry()
cnv := func(value float64) (time.Duration, error) {
	return time.Duration(value) * time.Second, nil
}
reg.Register("seconds", convert.ToAnyAny(cnv))

data := []byte(`{"type": "seconds", "value": 42}`)
val := &jsontype.Value{}
if err := jsontype.Unmarshal(reg, data, val); err != nil {
	log.Fatal(err)
}

fmt.Printf("%[1]v (%[1]T)\n", val.GoValue())
// Output: 42s (time.Duration)
```
