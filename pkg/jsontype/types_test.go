// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"

	"github.com/ctx42/convert/pkg/xcast"
	"github.com/ctx42/testing/pkg/assert"
)

func Test_DecodeNil(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// --- When ---
		have, err := DecodeNil(nil)

		// --- Then ---
		assert.NoError(t, err)
		assert.Nil(t, have)
	})

	t.Run("error", func(t *testing.T) {
		// --- When ---
		have, err := DecodeNil(42)

		// --- Then ---
		assert.ErrorIs(t, xcast.ErrInvType, err)
		wMsg := "DecodeNil: requires a nil value: invalid type"
		assert.ErrorEqual(t, wMsg, err)
		assert.Nil(t, have)
	})
}
