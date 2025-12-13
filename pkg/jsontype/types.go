// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac <rzajac@gmail.com>
// SPDX-License-Identifier: MIT

package jsontype

import (
	"fmt"

	"github.com/ctx42/convert/pkg/xcast"
)

// DecodeNil expects value to be nil, otherwise returns an error.
// On success, it returns nil and nil error.
func DecodeNil(value any) (any, error) {
	if value != nil {
		format := "DecodeNil: requires a nil value: %w"
		return nil, fmt.Errorf(format, xcast.ErrInvType)
	}
	return nil, nil
}
