// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"sync"
)

// Registry maps type names to their decoders.
type Registry struct {
	reg map[string]Decoder
	mx  sync.RWMutex
}

// NewRegistry returns a new instance of [Registry].
func NewRegistry() *Registry {
	return &Registry{reg: make(map[string]Decoder, 20)}
}

// Register registers a [Decoder] for the [TypeName]. When the decoder for it
// already exists, it will return it, nil otherwise.
func (reg *Registry) Register(name string, dec Decoder) Decoder {
	if dec == nil {
		return nil
	}
	reg.mx.Lock()
	defer reg.mx.Unlock()

	old := reg.reg[name]
	reg.reg[name] = dec
	return old
}

// Decoder returns a [Decoder] for the given [TypeName]. When the decoder for
// it is not registered, it returns nil.
func (reg *Registry) Decoder(typ string) Decoder {
	reg.mx.RLock()
	defer reg.mx.RUnlock()
	return reg.reg[typ]
}
