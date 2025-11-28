[![Go Report Card](https://goreportcard.com/badge/github.com/ctx42/jsontype)](https://goreportcard.com/report/github.com/ctx42/jsontype)
[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/ctx42/jsontype)
![Tests](https://github.com/ctx42/jsontype/actions/workflows/go.yml/badge.svg?branch=master)

<!-- TOC -->
  * [Marshaling and Unmarshaling](#marshaling-and-unmarshaling)
  * [Installation](#installation)
  * [Validation](#validation)
  * [Supported types](#supported-types)
<!-- TOC -->

`jsontype` is a small module that helps preserve Go types when marshaling them
to JSON by adding the Go type to the marshalled JSON.

## Marshaling and Unmarshaling

The workhorse of the module is the `jsontype.Value` type which implements
`json.Marshaler` and `json.Unmarshaler` interfaces.

```go
// Create a `jsontype.Value` and marshall it.
jType := jsontype.New(uint(42))
data, _ := json.Marshal(jType)
fmt.Printf("  marshalled: %s\n", string(data))

// Unmarshall the value and show its Go type.
gType := &jsontype.Value{}
_ = json.Unmarshal(data, gType)
fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())

// Output:
//   marshalled: {"type":"uint","value":42}
// unmarshalled: 42 (uint)
```

The module supports also some of the builtin types.

```go
// Create a `jsontype.Value` and marshall it.
jType := jsontype.New(time.Date(2000, 1, 2, 3, 4, 5, 600000000, time.UTC))
data, _ := json.Marshal(jType)
fmt.Printf("  marshalled: %s\n", string(data))

// Unmarshall the value and show its Go type.
gType := &jsontype.Value{}
_ = json.Unmarshal(data, gType)
fmt.Printf("unmarshalled: %[1]v (%[1]T)\n", gType.GoValue())

// Output:
// marshalled: {"type":"time.Time","value":"2000-01-02T03:04:05.6Z"}
// unmarshalled: 2000-01-02 03:04:05.6 +0000 UTC (time.Time)
```

## Installation

To use `jsontype` in your Go project, install it with:

```bash
go get github.com/ctx42/jsontype
```

## Validation

When unmarshaling `jsontype` checks if the value set in JSON is valid in
context the Go type. It will return an error when the value is of invalid type,
is in an invalid format, is in an invalid range for the Go type. Blow example
shows what happens when you try unmarshal too big value to `byte`.

```go
data := []byte(`{"type":"byte","value":99999}`)

gType := &jsontype.Value{}
err := json.Unmarshal(data, gType)

fmt.Println(err)
// Output:
// decodeByte: requires float64 value in range of uint8: invalid range
```

## Supported types

```go
int
int8
int16
int32
rune
int64
uint
uint8
byte
uint16
uint32
uint64
float32
float64
string
bool
time.Time
time.Duration
complex64
complex128
```
