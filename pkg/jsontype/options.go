// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

// Option represents a configuration option.
type Option func(*Options)

// Options represents configuration options.
type Options struct {
	reg *Registry
}

// WithRegistry creates an [Option] that sets the registry.
func WithRegistry(reg *Registry) Option {
	return func(opt *Options) { opt.reg = reg }
}
