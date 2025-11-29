// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

// ToAny adapts functions decoding JSON value to concrete type to [Decoder] type.
func ToAny[T any](fn func(value any) (T, error)) Decoder {
	return func(value any) (any, error) { return fn(value) }
}

// keyValue reruns a key value from given map. Returns the key value and true
// if it exists nil and false if it doesn't or when map is nil.
func keyValue(key string, m map[string]any) (any, bool) {
	if m == nil || len(m) == 0 {
		return nil, false
	}
	val, ok := m[key]
	return val, ok
}
