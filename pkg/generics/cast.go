package generics

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/tartale/go/pkg/primitives"
	"github.com/tartale/go/pkg/reflectx"
)

// ErrNotCasted is returned when a given value is
// unable to be cast to the desired type.
var ErrNotCasted = errors.New("unable to cast value")

// CastTo takes a value of any type and attempts several
// methods to cast it to type T in an expected way, using reflection
// if necessary. CastTo returns the cast value, and the
// type T cannot be automatically inferred. If inference is
// desired, use Cast.
func CastTo[T any](val any) (*T, error) {
	if val == nil {
		return nil, nil
	}

	// Try some basic casts first
	if result, ok := val.(T); ok {
		return &result, nil
	}

	if result, ok := val.(*T); ok {
		return result, nil
	}

	vval := reflect.ValueOf(val)

	// If val is a pointer, dereference it
	if vval.Kind() == reflect.Ptr {
		vval = vval.Elem()
	}
	val = vval.Interface()

	var (
		result T
		err    error
	)
	rtype := reflect.TypeOf(result)
	vtype := vval.Type()

	// Special case; number to string; the "ConvertibleTo"
	// from a number to string is "true", but does not convert as-expected.
	if reflectx.IsString(result) && reflectx.IsNumber(val) {
		s := fmt.Sprintf("%v", val)
		sval := reflect.ValueOf(s)
		res := sval.Interface().(T)

		return &res, nil
	}

	// Special case; string to number
	if reflectx.IsNumber(result) && reflectx.IsString(val) {
		err := primitives.Parse(val.(string), &result)
		if err == nil {
			return &result, nil
		}
	}

	if vtype.ConvertibleTo(rtype) {
		result = vval.Convert(rtype).Interface().(T)
		return &result, nil
	}

	// If the value can be decoded from a map; then return that
	err = mapstructure.Decode(vval.Interface(), &result)
	if err == nil {
		return &result, nil
	}

	// Try to encode as JSON, and then decode
	jsonBytes, _ := json.Marshal(vval.Interface())
	err = json.Unmarshal(jsonBytes, &result)
	if err == nil {
		return &result, nil
	}

	return nil, fmt.Errorf("%w: input type: %s; output type: %s", ErrNotCasted, vtype, rtype)
}

// MustCastTo is a convenience function that wraps
// CastTo, but panics if an error occurs.
func MustCastTo[T any](val any) *T {
	v, err := CastTo[T](val)
	if err != nil {
		panic(err)
	}
	return v
}

// Cast works similarly to the CastTo function, except
// that the caller provides a pointer to destination variable,
// instead of getting it through a return value. This allows
// type inference.
func Cast[T any](val any, dest *T) error {
	newVal, err := CastTo[T](val)
	if err != nil {
		return err
	}
	newDest := T(*newVal)
	*dest = newDest

	return nil
}

// MustCast is a convenience function that wraps
// Cast, but panics if an error occurs.
func MustCast[T any](val T, dest *T) error {
	newVal := MustCastTo[T](val)
	newDest := T(*newVal)
	*dest = newDest

	return nil
}
