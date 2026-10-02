// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
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
	assert.Len(t, 0, have.num)
	assert.NotNil(t, have.num)
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

	t.Run("overwrite registered", func(t *testing.T) {
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

	t.Run("clears number flag", func(t *testing.T) {
		// --- Given ---
		cnv0 := func(value any) (any, error) { return value, nil }

		cnv1 := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.registerNumber(Int, cnv0)

		// --- When ---
		have := reg.Register(Int, cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)

		assert.Len(t, 0, reg.num)
	})

	t.Run("zero value registry", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		var reg Registry

		// --- When ---
		have := reg.Register(Int, cnv)

		// --- Then ---
		assert.Nil(t, have)

		assert.Same(t, cnv, reg.Converter(Int))
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

func Test_Registry_registerNumber(t *testing.T) {
	t.Run("register not registered", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()

		// --- When ---
		have := reg.registerNumber(Int, cnv)

		// --- Then ---
		assert.Nil(t, have)

		val, _ := assert.HasKey(t, Int, reg.reg)
		assert.Same(t, cnv, val)
		assert.True(t, reg.num[Int])
	})

	t.Run("overwrite registered", func(t *testing.T) {
		// --- Given ---
		cnv0 := func(value any) (any, error) { return value, nil }

		cnv1 := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.Register(Int, cnv0)

		// --- When ---
		have := reg.registerNumber(Int, cnv1)

		// --- Then ---
		assert.Same(t, cnv0, have)

		val, _ := assert.HasKey(t, Int, reg.reg)
		assert.Same(t, cnv1, val)
		assert.True(t, reg.num[Int])
	})

	t.Run("zero value registry", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		var reg Registry

		// --- When ---
		have := reg.registerNumber(Int, cnv)

		// --- Then ---
		assert.Nil(t, have)

		assert.Same(t, cnv, reg.Converter(Int))
		assert.True(t, reg.num[Int])
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

func Test_Registry_converter(t *testing.T) {
	t.Run("number converter", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.registerNumber(Int, cnv)

		// --- When ---
		have, num := reg.converter(Int)

		// --- Then ---
		assert.Same(t, cnv, have)
		assert.True(t, num)
	})

	t.Run("regular converter", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.Register(Int, cnv)

		// --- When ---
		have, num := reg.converter(Int)

		// --- Then ---
		assert.Same(t, cnv, have)
		assert.False(t, num)
	})

	t.Run("not registered", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		// --- When ---
		have, num := reg.converter(Int)

		// --- Then ---
		assert.Nil(t, have)
		assert.False(t, num)
	})
}

func Test_Registry_alloc(t *testing.T) {
	t.Run("zero value", func(t *testing.T) {
		// --- Given ---
		var reg Registry

		// --- When ---
		reg.alloc()

		// --- Then ---
		assert.NotNil(t, reg.reg)
		assert.NotNil(t, reg.num)
	})

	t.Run("keeps existing maps", func(t *testing.T) {
		// --- Given ---
		cnv := func(value any) (any, error) { return value, nil }

		reg := NewRegistry()
		reg.registerNumber(Int, cnv)

		// --- When ---
		reg.alloc()

		// --- Then ---
		assert.Same(t, cnv, reg.reg[Int])
		assert.True(t, reg.num[Int])
	})
}
