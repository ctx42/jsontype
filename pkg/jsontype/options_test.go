// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
)

func Test_WithRegistry(t *testing.T) {
	// --- Given ---
	reg := NewRegistry()
	ops := &Options{}

	// --- When ---
	WithRegistry(reg)(ops)

	// --- Then ---
	assert.Same(t, reg, ops.reg)
}

func Test_newOptions(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		// --- When ---
		have := newOptions()

		// --- Then ---
		assert.Same(t, registry, have.reg)
	})

	t.Run("with options", func(t *testing.T) {
		// --- Given ---
		reg := NewRegistry()

		opt := WithRegistry(reg)

		// --- When ---
		have := newOptions(opt)

		// --- Then ---
		assert.Same(t, reg, have.reg)
	})
}
