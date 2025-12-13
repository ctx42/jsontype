// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_NewRegistry(t *testing.T) {
	// --- When ---
	have := NewRegistry()

	// --- Then ---
	assert.Len(t, 0, have.reg)
	assert.NotNil(t, have.reg)
}

func Test_Registry_Register(t *testing.T) {
	t.Run("register not registered", func(t *testing.T) {
		// --- Given ---
		dec := func(value any) (any, error) { return value, nil }
		reg := NewRegistry()

		// --- When ---
		have := reg.Register(Int, dec)

		// --- Then ---
		assert.Nil(t, have)
		val, _ := assert.HasKey(t, Int, reg.reg)
		assert.Same(t, dec, val)
	})

	t.Run("register not registered", func(t *testing.T) {
		// --- Given ---
		dec0 := func(value any) (any, error) { return value, nil }
		dec1 := func(value any) (any, error) { return value, nil }
		reg := NewRegistry()
		reg.Register(Int, dec0)

		// --- When ---
		have := reg.Register(Int, dec1)

		// --- Then ---
		assert.Same(t, dec0, have)
		val, _ := assert.HasKey(t, Int, reg.reg)
		assert.Same(t, dec1, val)
	})

	t.Run("register nil decoder", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := reg.Register(Int, nil)

		// --- Then ---
		assert.Nil(t, have)
		assert.Len(t, 0, reg.reg)
	})
}

func Test_Registry_Decoder(t *testing.T) {
	t.Run("registered", func(t *testing.T) {
		// --- Given ---
		dec := func(value any) (any, error) { return value, nil }
		reg := NewRegistry()
		reg.Register(Int, dec)

		// --- When ---
		have := reg.Decoder(Int)

		// --- Then ---
		assert.Same(t, dec, have)
	})

	t.Run("not registered", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := reg.Decoder(Int)

		// --- Then ---
		assert.Nil(t, have)
	})
}
