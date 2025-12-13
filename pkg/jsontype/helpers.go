// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"fmt"

	"github.com/ctx42/convert/pkg/xcast"
	"github.com/ctx42/convert/pkg/xconv"
)

// ToAny adapts functions decoding JSON value to concrete type to [Decoder]
// type.
func ToAny[T any](fn func(value any) (T, error)) Decoder {
	return func(value any) (any, error) { return fn(value) }
}

// FromConv return [Decoder] based on [xconv.Converter].
func FromConv[From, To any](conv xconv.Converter[From, To]) Decoder {
	return func(value any) (any, error) {
		var ok bool
		var from From
		if from, ok = value.(From); !ok {
			format := "%w: expected %T, got %T"
			var to To
			return to, fmt.Errorf(format, xcast.ErrInvType, from, value)
		}
		return conv(from)
	}
}

// keyValue reruns a key value from the given map. Returns the key value and
// true if it exists nil and false if it doesn't or when the map is nil.
func keyValue(key string, m map[string]any) (any, bool) {
	if m == nil || len(m) == 0 {
		return nil, false
	}
	val, ok := m[key]
	return val, ok
}
