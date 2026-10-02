// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_marshal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- Given ---
		v := map[string]any{"A": 1}

		// --- When ---
		have, err := marshal(v)

		// --- Then ---
		assert.NoError(t, err)
		assert.Equal(t, `{"A":1}`, string(have))
	})

	t.Run("error", func(t *testing.T) {
		// --- Given ---
		v := func() {}

		// --- When ---
		have, err := marshal(v)

		// --- Then ---
		assert.ErrorEqual(t, "jsontype: json: unsupported type: func()", err)
		assert.Nil(t, have)
	})
}

func Test_keyValue(t *testing.T) {
	t.Run("key exists", func(t *testing.T) {
		// --- Given ---
		key := "B"

		m := map[string]any{"A": 1, "B": 2}

		// --- When ---
		have, ok := keyValue(key, m)

		// --- Then ---
		assert.Equal(t, 2, have)
		assert.True(t, ok)
	})

	t.Run("nil map", func(t *testing.T) {
		// --- Given ---
		key := "B"

		// --- When ---
		have, ok := keyValue(key, nil)

		// --- Then ---
		assert.Nil(t, have)
		assert.False(t, ok)
	})
}
