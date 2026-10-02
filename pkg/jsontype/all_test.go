// SPDX-FileCopyrightText: (c) 2026 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"testing"

	"github.com/ctx42/testing/pkg/assert"
	"github.com/ctx42/testing/pkg/tester"
)

// isolateRegistry replaces the package-level registry with a new default
// registry and restores the original when the test completes. Use it in
// tests that register converters in the package-level registry.
func isolateRegistry(t tester.T) {
	t.Helper()
	org := registry
	registry = DefaultRegistry()
	t.Cleanup(func() { registry = org })
}

func Test_isolateRegistry(t *testing.T) {
	// --- Given ---
	org := registry

	tspy := tester.New(t, 1).ExpectCleanups(1).Close()

	// --- When ---
	isolateRegistry(tspy)

	// --- Then ---
	assert.NotSame(t, org, registry)
	assert.Len(t, len(org.reg), registry.reg)

	tspy.Finish()
	assert.Same(t, org, registry)
}
