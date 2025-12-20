// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

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
