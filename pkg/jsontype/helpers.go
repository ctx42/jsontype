// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"fmt"

	"github.com/ctx42/convert/pkg/convert"
)

// Unmarshal unmarshals JSON representation of the value using [Registry].
func Unmarshal(reg *Registry, bytes []byte, val *Value) error {
	tmp := struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	}{}
	if err := json.Unmarshal(bytes, &tmp); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}

	var err error
	cnv := reg.Converter(tmp.Type)
	if cnv == nil {
		return fmt.Errorf("%w: %s", convert.ErrUnsType, tmp.Type)
	}
	val.typ = tmp.Type
	if val.val, err = cnv(tmp.Value); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	return nil
}

// keyValue reruns the value represented by the key from the given map. Returns
// the value and true if it exists. Returns nil and false if it doesn't or when
// the map is empty or nil.
func keyValue(key string, m map[string]any) (any, bool) {
	if len(m) == 0 {
		return nil, false
	}
	val, ok := m[key]
	return val, ok
}
