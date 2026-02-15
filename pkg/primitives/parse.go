package primitives

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/tartale/go/pkg/errorz"
	"github.com/tartale/go/pkg/reflectx"
)

func MustParseTo[T any](s string, optionalBase ...int) T {
	result, err := ParseTo[T](s, optionalBase...)
	if err != nil {
		panic(err)
	}

	return result
}

// ParseTo attempts to parse the string s into a primitive
// of type T. This is a wrapper for all the Parse* functions provided
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
// back to the caller, and the returned variable will be the
// zero-value of type T.
func ParseTo[T any](s string, optionalBase ...int) (T, error) {

	base := 0
	var result T
	if !reflectx.IsPrimitive(result) {
		return result, fmt.Errorf("%w: %T is not a primitive type", errorz.ErrInvalidType, result)
	}

	rval := reflect.ValueOf(result)
	if len(optionalBase) == 1 {
		base = optionalBase[0]
	} else if len(optionalBase) > 1 {
		return result, fmt.Errorf("%w: unexpected optional parameters: %v", errorz.ErrInvalidArgument, optionalBase)
	}

	switch rval.Interface().(type) {
	case bool:
		r, err := strconv.ParseBool(s)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(r).Interface().(T)

	case float32:
		r, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(float32(r)).Interface().(T)

	case float64:
		r, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(float64(r)).Interface().(T)

	case int:
		r, err := strconv.ParseInt(s, base, 32)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(int(r)).Interface().(T)

	case int8:
		r, err := strconv.ParseInt(s, base, 8)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(int8(r)).Interface().(T)

	case int16:
		r, err := strconv.ParseInt(s, base, 16)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(int16(r)).Interface().(T)

	case int32:
		r, err := strconv.ParseInt(s, base, 32)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(int32(r)).Interface().(T)

	case int64:
		r, err := strconv.ParseInt(s, base, 64)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(int64(r)).Interface().(T)

	case uint:
		r, err := strconv.ParseUint(s, base, 32)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(uint(r)).Interface().(T)

	case uint8:
		r, err := strconv.ParseUint(s, base, 8)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(uint8(r)).Interface().(T)

	case uint16:
		r, err := strconv.ParseUint(s, base, 16)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(uint16(r)).Interface().(T)

	case uint32:
		r, err := strconv.ParseUint(s, base, 32)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(uint32(r)).Interface().(T)

	case uint64:
		r, err := strconv.ParseUint(s, base, 64)
		if err != nil {
			return result, err
		}
		result = reflect.ValueOf(uint64(r)).Interface().(T)

	case string:
		result = reflect.ValueOf(s).Interface().(T)

	default:
		return result, fmt.Errorf("%w: %T", errorz.ErrInvalidType, result)
	}

	return result, nil
}

func MustParse[T any](s string, dest *T, optionalBase ...int) {

	newVal := MustParseTo[T](s, optionalBase...)
	*dest = newVal
}

func Parse[T any](s string, dest *T, optionalBase ...int) error {

	newVal, err := ParseTo[T](s, optionalBase...)
	if err != nil {
		return err
	}
	*dest = newVal

	return nil
}
