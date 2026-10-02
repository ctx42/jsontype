// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ctx42/convert/pkg/convert"
)

// NilConverter expects value to be nil, otherwise returns an error. On success,
// it returns nil and nil error.
func NilConverter(value any) (any, error) {
	if value != nil {
		format := "jsontype: requires a nil value: %w"
		return nil, fmt.Errorf(format, convert.ErrInvType)
	}
	return nil, nil
}

// numberConverter returns a converter for a numeric type T. It converts a
// [json.Number] with str, which parses the number's text without precision
// loss, and any other value with f64, which expects a float64. Errors for a
// [json.Number] name float64 as the source type, as for any JSON number.
func numberConverter[T any](
	str convert.SrcToDst[string, T],
	f64 convert.SrcToDst[float64, T],
) convert.AnyToAny {

	fromFloat := convert.ToAnyAny(f64)
	return func(value any) (any, error) {
		num, ok := value.(json.Number)
		if !ok {
			return fromFloat(value)
		}
		ret, err := str(string(num))
		if err != nil {
			var cnvErr convert.Error
			if errors.As(err, &cnvErr) {
				cnvErr.Src = "float64"
				return ret, cnvErr
			}
			return ret, err
		}
		return ret, nil
	}
}
