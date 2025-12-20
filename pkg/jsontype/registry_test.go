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
		cnv := func(value any) (any, error) { return value, nil }
		reg := NewRegistry()

		// --- When ---
		have := reg.Register(Int, cnv)

		// --- Then ---
		assert.Nil(t, have)
		val, _ := assert.HasKey(t, Int, reg.reg)
		assert.Same(t, cnv, val)
	})

	t.Run("register not registered", func(t *testing.T) {
		// --- Given ---
		cnv0 := func(value any) (any, error) { return value, nil }
		cnv1 := func(value any) (any, error) { return value, nil }
		reg := NewRegistry()
		reg.Register(Int, cnv0)

		// --- When ---
		have := reg.Register(Int, cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)
		val, _ := assert.HasKey(t, Int, reg.reg)
		assert.Same(t, cnv1, val)
	})

	t.Run("register nil converter", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := reg.Register(Int, nil)

		// --- Then ---
		assert.Nil(t, have)
		assert.Len(t, 0, reg.reg)
	})
}

func Test_Registry_Converter(t *testing.T) {
	t.Run("registered", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }
		reg := NewRegistry()
		reg.Register(Int, cnv)

		// --- When ---
		have := reg.Converter(Int)

		// --- Then ---
		assert.Same(t, cnv, have)
	})

	t.Run("not registered", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have := reg.Converter(Int)

		// --- Then ---
		assert.Nil(t, have)
	})
}
