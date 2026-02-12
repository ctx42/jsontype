// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
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
