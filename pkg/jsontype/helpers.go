// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/ctx42/convert/pkg/convert"
)

// Unmarshal unmarshals JSON representation of the value using [Registry].
// Transparent types (string, bool, nil, float64) are detected by their JSON
// shape and require no envelope. The envelope form is still accepted for all
// types (back-compatibility). Whitespace around the JSON value is ignored.
func Unmarshal(reg *Registry, data []byte, val *Value) error {
	data = bytes.Trim(data, " \t\r\n")
	var first byte
	if len(data) > 0 {
		first = data[0]
	}
	switch first {
	case '"':
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
		val.typ, val.val = String, s
		return nil
	case 't':
		if string(data) != "true" {
			return fmt.Errorf("jsontype: invalid JSON token")
		}
		val.typ, val.val = Bool, true
		return nil
	case 'f':
		if string(data) != "false" {
			return fmt.Errorf("jsontype: invalid JSON token")
		}
		val.typ, val.val = Bool, false
		return nil
	case 'n':
		if string(data) != "null" {
			return fmt.Errorf("jsontype: invalid JSON token")
		}
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

// unmarshalEnvelope decodes the {"type": "...", "value": ...} form.
func unmarshalEnvelope(reg *Registry, data []byte, val *Value) error {
	tmp := struct {
		Type  string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	cnv, num := reg.converter(tmp.Type)
	if cnv == nil {
		return fmt.Errorf("jsontype: %w: %s", convert.ErrUnsType, tmp.Type)
	}
	var value any
	if len(tmp.Value) > 0 {
		dec := json.NewDecoder(bytes.NewReader(tmp.Value))
		if num {
			dec.UseNumber()
		}
		if err := dec.Decode(&value); err != nil {
			return fmt.Errorf("jsontype: %w", err)
		}
	}
	ret, err := cnv(value)
	if err != nil {
		return fmt.Errorf("jsontype: %w", err)
	}
	val.typ, val.val = tmp.Type, ret
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
