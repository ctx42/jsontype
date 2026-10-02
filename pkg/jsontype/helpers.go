// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"fmt"
)

// marshal returns the JSON encoding of v, adding context to the error.
func marshal(v any) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("jsontype: %w", err)
	}
	return data, nil
}

// keyValue returns the value represented by the key from the given map. Returns
// the value and true if it exists. Returns nil and false if it doesn't or when
// the map is empty or nil.
func keyValue(key string, m map[string]any) (any, bool) {
	val, ok := m[key]
	return val, ok
}
