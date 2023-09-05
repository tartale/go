package generics

import "reflect"

func Normalize[T any](val any) *T {

	if val == nil {
		return nil
	}

	reflectVal := reflect.ValueOf(val)

	// If val is a pointer, dereference it
	if reflectVal.Kind() == reflect.Ptr {
		reflectVal = reflectVal.Elem()
	}

	// Get the type of the generic parameter
	// https://stackoverflow.com/a/73932292/1258206
	var zero [0]T
	ttype := reflect.TypeOf(zero).Elem()

	// If the types are identical, no normalization needed
	if reflectVal.Type() == ttype {
		tval := reflectVal.Interface().(T)
		return &tval
	}

	// Ensure that the requested and actual types are compatable
	if reflectVal.Type().ConvertibleTo(ttype) {
		tval := reflectVal.Convert(ttype).Interface().(T)
		return &tval
	}

	return nil
}
