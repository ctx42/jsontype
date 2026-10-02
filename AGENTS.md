This file provides guidance to Claude Code (claude.ai/code) when working with
code in this repository.

## Commands

```bash
# Run all tests with race detection
go test -v -race ./...

# Run a single test
go test -v -run Test_FunctionName ./pkg/jsontype/

# Build all packages
go build ./...
```

## Architecture

`jsontype` is a Go library that preserves Go type information through JSON
serialization. Standard JSON unmarshaling loses type specificity (e.g.,
`uint64` becomes `float64`). This library solves that by embedding the type
name alongside the value in the JSON output:
`{"type":"uint64","value":42}`.

### Package layout

All library code lives in `pkg/jsontype/`:

- **`jsontype.go`** — `Value` struct (core type), `New[T]()` generic
  constructor, `NewValue()` for dynamic construction, `MarshalJSON()` /
  `UnmarshalJSON()`, `Unmarshal()` (parses raw JSON, looks up the converter
  in a registry, and applies it), `Map()` / `FromMap()` / `AsValue()`
  helpers, and the package-level default registry.
- **`registry.go`** — Thread-safe `Registry` (map of type name → converter)
  with `NewRegistry()`, `Register()`, and `Converter()`.
- **`helpers.go`** — type-agnostic helpers (`marshal`, `keyValue`).
- **`converters.go`** — `NilConverter()` and the converter builders for
  numeric types (`numberConverter`, which parses `json.Number` exactly) and
  `time.Duration` (`durationConverter`).
- **`options.go`** — Functional options pattern (`Option`, `Options`,
  `WithRegistry()`).

### Type conversion

All type conversions delegate to `github.com/ctx42/convert`, which handles
overflow, underflow, and precision loss. Adding a new type requires:
1. Defining a converter function and registering it in `DefaultRegistry()`.
2. Adding a type-name constant (e.g., `const MyType = "mytype"`).

### Testing

Tests live alongside source files. The `internal/test/` package provides
shared test helpers. Examples are in `examples_test.go` and serve as
living documentation.
