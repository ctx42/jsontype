// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

// Option configures [NewValue], [FromMap], and [AsValue].
type Option func(*Options)

// Options holds the configuration set by [Option] functions.
type Options struct {
	reg *Registry
}

// WithRegistry creates an [Option] that sets the registry.
func WithRegistry(reg *Registry) Option {
	return func(opt *Options) { opt.reg = reg }
}

// newOptions returns options defaulting to the package-level registry with
// the given options applied.
func newOptions(opts ...Option) *Options {
	ops := &Options{reg: registry}
	for _, opt := range opts {
		opt(ops)
	}
	return ops
}
