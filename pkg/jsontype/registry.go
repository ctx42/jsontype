// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"sync"
)

// Registry maps type names to their decoders.
type Registry struct {
	reg map[TypeName]Decoder
	mx  sync.RWMutex
}

// NewRegistry returns new instance of [Registry].
func NewRegistry() *Registry {
	return &Registry{reg: make(map[TypeName]Decoder, 20)}
}

// Register registers a [Decoder] for the [TypeName]. When the decoder already
// exists for the type name it will return it, nil otherwise.
func (reg *Registry) Register(typ TypeName, dec Decoder) Decoder {
	reg.mx.Lock()
	defer reg.mx.Unlock()

	have := reg.reg[typ]
	reg.reg[typ] = dec
	return have
}

// Decoder returns a [Decoder] for given [TypeName]. When the decoder for given
// type name is not registered it returns nil.
func (reg *Registry) Decoder(typ TypeName) Decoder {
	reg.mx.RLock()
	defer reg.mx.RUnlock()

	return reg.reg[typ]
}
