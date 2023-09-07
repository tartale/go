package generics

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var ErrInvalidType = errors.New("invalid type")

func Cast[T any](val any) *T {

	v, _ := CastE[T](val)
	return v
}

func CastE[T any](val any) (*T, error) {

	if val == nil {
		return nil, nil
	}

	// Try some basic casts first
	if tval, ok := val.(T); ok {
		return &tval, nil
	}

	if tval, ok := val.(*T); ok {
		return tval, nil
	}

	reflectVal := reflect.ValueOf(val)

	// If val is a pointer, dereference it
	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	// Get the type of the generic parameter
	// https://stackoverflow.com/a/73932292/1258206
	var (
		tval T
		err  error
	)
	ttype := reflect.TypeOf(tval)
	vtype := reflectVal.Type()

	// If the value is convertablie to T, then convert it and return it
	if vtype.ConvertibleTo(ttype) {
		tval = reflectVal.Convert(ttype).Interface().(T)
		return &tval, nil
	}

	// If the value can be decoded from a map; then return that
	err = mapstructure.Decode(reflectVal.Interface(), &tval)
	if err == nil {
		return &tval, nil
	}

	return nil, fmt.Errorf("%w: input type: %s; output type: %s", ErrInvalidType, vtype, ttype)
}
