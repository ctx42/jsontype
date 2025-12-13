[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/jsontype)](https://goreportcard.com/report/github.com/ctx42/jsontype)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/jsontype)
![Tests](https://github.com/ctx42/jsontype/actions/workflows/go.yml/badge.svg?branch=master)

<!-- TOC -->
  * [Installation](#installation)
  * [Example](#example)
  * [Custom Decoders](#custom-decoders)
<!-- TOC -->

`jsontype` is a small Go module that preserves Go types when marshaling to JSON.
It embeds type information directly into the JSON alongside the value.

This is useful in scenarios where you marshal and unmarshal JSON to composite
types, such as `map[string]any`. It ensures that Go types are preserved during
a round-trip.

## Installation

To use `jsontype` in your Go project, install it with:

```bash
go get github.com/ctx42/jsontype
```

## Example

Create a `Value` instance encapsulating the value and its type when marshaled
to JSON.

```go
jType := jsontype.New(uint(42))
data, _ := json.Marshal(jType)

fmt.Println(string(data))
// Output:
// {"type":"uint","value":42}
```

Later when unmarshalling the library is looking at the `type` field, finds the
matching `Decoder` which casts the `value` to specific Go type.

```go
data := []byte(`{"type": "uint", "value": 42}`)

gType := &jsontype.Value{}
_ = json.Unmarshal(data, gType)

fmt.Printf("%[1]v (%[1]T)\n", gType.GoValue())
// Output:
// 42 (uint)
```

The `Decoders` for most built-in types are provided by the module.

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

## Custom Decoders

You may register custom `Decoder` for a type.

```go
// Custom decoder for the "seconds" type.
dec := func (value float64) (time.Duration, error) {
return time.Duration(value) * time.Second, nil
}

// Register decoder.
jsontype.Register("seconds", jsontype.FromConv(dec))

// Custom type named "seconds" representing duration in seconds.
data := []byte(`{"type": "seconds", "value": 42}`)

gType := &jsontype.Value{}
err := json.Unmarshal(data, gType)

_ = err // Check error.

fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())
// Output:
// unmarshalled: 42s (time.Duration)
```

The registered decoder must return an error when conversion a value from a
JSON type to GO type would result in loss of precision, overflow, underflow or
conversion is simply impossible. Similarly to how the
[http://github.com/ctx42/convert](http://github.com/ctx42/convert) converter
functions work.
