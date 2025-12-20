// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"sync"

	"github.com/ctx42/convert/pkg/convert"
)

// Registry maps type names to their converters.
type Registry struct {
	reg map[string]convert.AnyToAny
	mx  sync.RWMutex
}

// NewRegistry returns a new instance of [Registry].
func NewRegistry() *Registry {
	return &Registry{reg: make(map[string]convert.AnyToAny, 20)}
}

// Register registers a converter for the given type name. When the converter
// for it already exists, it will return it, nil otherwise.
func (reg *Registry) Register(name string, cnv convert.AnyToAny) convert.AnyToAny {
	if cnv == nil {
		return nil
	}
	reg.mx.Lock()
	defer reg.mx.Unlock()

	old := reg.reg[name]
	reg.reg[name] = cnv
	return old
}

// Converter returns a converter for the given type name. When the converter
// for it is not registered, it returns nil.
func (reg *Registry) Converter(typ string) convert.AnyToAny {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	return reg.reg[typ]
}
