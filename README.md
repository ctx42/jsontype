[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/jsontype)](https://goreportcard.com/report/github.com/ctx42/jsontype)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/jsontype)
![Tests](https://github.com/ctx42/jsontype/actions/workflows/go.yml/badge.svg?branch=master)

<!-- TOC -->
  * [Why Use `jsontype`?](#why-use-jsontype)
  * [Installation](#installation)
  * [Example](#example)
  * [Type Registry](#type-registry)
  * [Custom Converters](#custom-converters)
<!-- TOC -->

`jsontype` is a small Go module that preserves Go types when marshaling values 
to JSON. It embeds type information directly into the JSON alongside the value.

## Why Use `jsontype`?

Standard Go JSON marshaling loses specific type information for `interface{}` 
(or `any`) fields. For example, a `uint64` becomes a `float64` after a
round-trip through JSON if unmarshaled into a `map[string]any`. `jsontype`
solves this by explicitly storing the type name.

This is especially useful for:
- Storing data in databases as JSON.
- Passing typed messages over a bus where the receiver uses `map[string]any`.

## Installation

Install using `go get`:

```bash
go get github.com/ctx42/jsontype
```

## Example

Create a `Value` instance encapsulating the value and its type.

```go
jType := jsontype.New(uint(42))
data, _ := json.Marshal(jType)

fmt.Println(string(data))
// Output:
// {"type":"uint","value":42}
```

Later when unmarshalling the library is looking at the `type` field, finds the
matching converter in the package-level registry and converts the value. 


```go
data := []byte(`{"type": "uint", "value": 42}`)

gType := &jsontype.Value{}
_ = json.Unmarshal(data, gType)

fmt.Printf("%[1]v (%[1]T)\n", gType.GoValue())
// Output:
// 42 (uint)
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

## Custom Converters

You may register a custom converter for your custom type.

```go
// Custom converter for a type named "seconds" representing duration in seconds.
cnv := func(value float64) (time.Duration, error) {
    return time.Duration(value) * time.Second, nil
}

// Register converter.
jsontype.Register("seconds", convert.ToAnyAny(cnv))

// Custom type named "seconds" representing duration in seconds.
data := []byte(`{"type": "seconds", "value": 42}`)

gType := &jsontype.Value{}
err := json.Unmarshal(data, gType)

_ = err // Check error.

fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())
// Output:
// unmarshalled: 42s (time.Duration)
```

The registered converter must return an error when conversion of a value from a
JSON type to Go type would result in loss of precision, overflow, underflow or
conversion is simply impossible. The same way how the
[github.com/ctx42/convert](http://github.com/ctx42/convert) converter
functions work.
