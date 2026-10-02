// SPDX-FileCopyrightText: (c) 2025 Rafal Zajac
// SPDX-License-Identifier: MIT

package jsontype

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

// durationConverter returns a converter for [time.Duration]. It accepts a
// duration string (e.g. "1m30s") and a number of nanoseconds, which is how
// [json.Marshal] writes a [time.Duration].
func durationConverter() convert.AnyToAny {
	fromNumber := numberConverter(
		nanosecondsToDuration,
		convert.Float64ToDuration,
	)
	return func(value any) (any, error) {
		if str, ok := value.(string); ok {
			return convert.StringToDuration(str)
		}
		return fromNumber(value)
	}
}

// nanosecondsToDuration converts a base-10 number of nanoseconds to
// [time.Duration] without precision loss.
func nanosecondsToDuration(src string) (time.Duration, error) {
	ns, err := convert.StringToInt64(src)
	if err != nil {
		return 0, convert.ChangeErrDstName(err, "time.Duration")
	}
	return time.Duration(ns), nil
}
