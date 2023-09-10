package generics

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/tartale/go/pkg/primitives"
	"github.com/tartale/go/pkg/reflectx"
)

var ErrNotCasted = errors.New("unable to cast value")

func MustCastTo[T any](val any) *T {

	v, err := CastTo[T](val)
	if err != nil {
		panic(err)
	}
	return v
}

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

	return nil, fmt.Errorf("%w: input type: %s; output type: %s", ErrNotCasted, vtype, rtype)
}

func MustCast[T any](val T, dest *T) error {

	newVal := MustCastTo[T](val)
	newDest := T(*newVal)
	*dest = newDest

	return nil
}

func Cast[T any](val any, dest *T) error {

	fmt.Printf("%T\n", dest)
	newVal, err := CastTo[T](val)
	if err != nil {
		return err
	}
	newDest := T(*newVal)
	*dest = newDest

	return nil
}
