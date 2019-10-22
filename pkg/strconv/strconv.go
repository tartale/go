// Copyright 2019 Tom Artale. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
//
// Package strconv implements conversions to and from string representations
// of basic data types.
package strconv

import (
	"fmt"
	"strconv"
)

// ParsePrimitive interprets a string s and saves the result
// in the value pointed to by primitivePtr.
//
// This is a wrapper for all the Parse* functions provided
// in strconv, but is useful when you're working with an
// interface{} type and don't want to do a giant switch
// statement. You trade-off the type-safety provided by
// the existing functions.
//
// The base can be provided for the integer-like Parse functions,
// if it is required. If the type of primitivePtr is not an
// integer, the base is ignored.  If the base is not provided,
// it defaults to 0.
//
// The bitSize argument required by some of the Parse functions
// is provided automatically, depending on the type.
//
// Any errors from the underlying Parse functions are reflected
// back to the caller, and the variable is left unchanged.
func ParsePrimitive(s string, primitivePtr interface{}, optionalBase ...int) error {
	base := 0
	if len(optionalBase) == 1 {
		base = optionalBase[0]
	} else if len(optionalBase) > 1 {
		return fmt.Errorf("Unexpected optional parameters: %v", optionalBase)
	}

	switch v := primitivePtr.(type) {
	case *bool:
		result, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		*v = result
	case *float32:
		result, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return err
		}
		*v = float32(result)
	case *float64:
		result, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		*v = result
	case *int:
		result, err := strconv.ParseInt(s, base, 64)
		if err != nil {
			return err
		}
		*v = int(result)
	case *int8:
		result, err := strconv.ParseInt(s, base, 8)
		if err != nil {
			return err
		}
		*v = int8(result)
	case *int16:
		result, err := strconv.ParseInt(s, base, 16)
		if err != nil {
			return err
		}
		*v = int16(result)
	case *int32:
		result, err := strconv.ParseInt(s, base, 32)
		if err != nil {
			return err
		}
		*v = int32(result)
	case *int64:
		result, err := strconv.ParseInt(s, base, 64)
		if err != nil {
			return err
		}
		*v = result
	case *uint:
		result, err := strconv.ParseUint(s, base, 64)
		if err != nil {
			return err
		}
		*v = uint(result)
	case *uint8:
		result, err := strconv.ParseUint(s, base, 8)
		if err != nil {
			return err
		}
		*v = uint8(result)
	case *uint16:
		result, err := strconv.ParseUint(s, base, 16)
		if err != nil {
			return err
		}
		*v = uint16(result)
	case *uint32:
		result, err := strconv.ParseUint(s, base, 32)
		if err != nil {
			return err
		}
		*v = uint32(result)
	case *uint64:
		result, err := strconv.ParseUint(s, base, 64)
		if err != nil {
			return err
		}
		*v = result
	case *string:
		*v = s
	default:
		return fmt.Errorf("Unexpected type: %T; ensure passed argument is a pointer to a primitive", primitivePtr)
	}

	return nil
}
