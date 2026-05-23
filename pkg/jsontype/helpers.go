// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"fmt"

	"github.com/ctx42/convert/pkg/convert"
)

// Unmarshal unmarshals JSON representation of the value using [Registry].
// Transparent types (string, bool, nil, float64) are detected by their JSON
// shape and require no envelope. The envelope form is still accepted for all
// types (back-compatibility).
func Unmarshal(reg *Registry, data []byte, val *Value) error {
	var first byte
	for _, b := range data {
		if b != ' ' && b != '\t' && b != '\r' && b != '\n' {
			first = b
			break
		}
	}
	switch first {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
		val.typ, val.val = String, s
		return nil
	case 't', 'f':
		var b bool
		if err := json.Unmarshal(data, &b); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
		val.typ, val.val = Bool, b
		return nil
	case 'n':
		val.typ, val.val = Nil, nil
		return nil
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-':
		var f float64
		if err := json.Unmarshal(data, &f); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
		val.typ, val.val = Float64, f
		return nil
	}
	return unmarshalEnvelope(reg, data, val)
}

// unmarshalEnvelope decodes the {"type":"...","value":...} form.
func unmarshalEnvelope(reg *Registry, data []byte, val *Value) error {
	tmp := struct {
		Type  string `json:"type"`
		Value any    `json:"value"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	cnv := reg.Converter(tmp.Type)
	if cnv == nil {
		return fmt.Errorf("%w: %s", convert.ErrUnsType, tmp.Type)
	}
	val.typ = tmp.Type
	var err error
	if val.val, err = cnv(tmp.Value); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	return nil
}

// keyValue returns the value represented by the key from the given map. Returns
// the value and true if it exists. Returns nil and false if it doesn't or when
// the map is empty or nil.
func keyValue(key string, m map[string]any) (any, bool) {
	if len(m) == 0 {
		return nil, false
	}
	val, ok := m[key]
	return val, ok
}
