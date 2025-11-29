// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"errors"
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_ToAny(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		fn := func(value any) (int, error) { return value.(int), nil }

		// --- When ---
		have, err := ToAny(fn)(42)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, 42, have)
	})

	t.Run("returns error", func(t *testing.T) {
		// --- Given ---
		ErrTst := errors.New("err")
		fn := func(value any) (int, error) { return 0, ErrTst }

		// --- When ---
		have, err := ToAny(fn)(42)

		// --- Then ---
		assert.Same(t, ErrTst, err)
		assert.Equal(t, 0, have)
	})
}

func Test_keyValue(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		// --- Given ---
		m := map[string]any{"A": 1, "B": 2}

		// --- When ---
		have, ok := keyValue("B", m)

		// --- Then ---
		assert.Equal(t, 2, have)
		assert.True(t, ok)
	})

	t.Run("nil map", func(t *testing.T) {
		// --- When ---
		have, ok := keyValue("B", nil)

		// --- Then ---
		assert.Nil(t, have)
		assert.False(t, ok)
	})
}
