// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

// ToAny adapts functions decoding JSON value to concrete type to [Decoder] type.
func ToAny[T any](fn func(value any) (T, error)) Decoder {
	return func(value any) (any, error) { return fn(value) }
}
