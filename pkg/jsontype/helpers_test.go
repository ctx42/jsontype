// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

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
