// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"sync"

	"github.com/ctx42/convert/pkg/convert"
)

// Registry maps type names to their converters.
type Registry struct {
	reg map[string]convert.AnyToAny
	num map[string]bool // Converters receiving numbers as json.Number.
	mx  sync.RWMutex
}

// NewRegistry returns a new instance of [Registry].
func NewRegistry() *Registry {
	return &Registry{
		reg: make(map[string]convert.AnyToAny, 20),
		num: make(map[string]bool, 20),
	}
}

// Register registers a converter for the given type name. Returns the
// previous converter if one was already registered, nil otherwise.
func (reg *Registry) Register(
	name string,
	cnv convert.AnyToAny,
) convert.AnyToAny {

	if cnv == nil {
		return nil
	}
	reg.mx.Lock()
	defer reg.mx.Unlock()

	old := reg.reg[name]
	reg.reg[name] = cnv
	delete(reg.num, name)
	return old
}

// registerNumber registers a converter that receives JSON numbers as
// [json.Number] instead of float64, so it can convert them without
// precision loss. Returns the previous converter if one was already
// registered, nil otherwise.
func (reg *Registry) registerNumber(
	name string,
	cnv convert.AnyToAny,
) convert.AnyToAny {

	reg.mx.Lock()
	defer reg.mx.Unlock()

	old := reg.reg[name]
	reg.reg[name] = cnv
	reg.num[name] = true
	return old
}

// Converter returns a converter for the given type name. When the converter
// for it is not registered, it returns nil.
func (reg *Registry) Converter(typ string) convert.AnyToAny {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	return reg.reg[typ]
}

// converter returns a converter for the given type name and reports whether
// it expects JSON numbers as [json.Number].
func (reg *Registry) converter(typ string) (convert.AnyToAny, bool) {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	return reg.reg[typ], reg.num[typ]
}
