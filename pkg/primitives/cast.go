package primitives

import (
	"fmt"
	"reflect"

	"github.com/tartale/go/pkg/constraintz"
	"github.com/tartale/go/pkg/errorz"
	"github.com/tartale/go/pkg/reflectx"
)

func CastTo[T, U constraintz.Primitive](val T) (U, error) {
	var result U
	s := fmt.Sprintf("%v", val)
	Parse(s, &result)

	return result, nil
}

func Cast[T, U constraintz.Primitive](val T, dest *U) error {
	var result U
	result, err := CastTo[T, U](val)
	if err != nil {
		return err
	}
	*dest = result
	return nil
}

// CastAway takes a primitive type, figures out its
// underlying type, and casts the value to that underlying
// type. This is for situations where you want to turn
// the value for a type definition into its raw primitive type.
//
// For example:
//
//	type MyType int
//	myType := 1
//
//	// This fails because 'myType' is not an int
//	assert.Equal(1, myType)
//
//	myPrimitive, _ := CastAway(myType)
//
//	// This passes because 'myPrimitive' is now a plain-old int
//	assert.Equal(1, myPrimitive)
func CastAway[T constraintz.Primitive](val T) (any, error) {
	valType := reflectx.TypeOfElement(val)
	valKind := valType.Kind()
	switch valKind {
	case reflect.Bool:
		return CastTo[T, bool](val)
	case reflect.Float32:
		return CastTo[T, float32](val)
	case reflect.Float64:
		return CastTo[T, float64](val)
	case reflect.Int:
		return CastTo[T, int](val)
	case reflect.Int8:
		return CastTo[T, int8](val)
	case reflect.Int16:
		return CastTo[T, int16](val)
	case reflect.Int32:
		return CastTo[T, int32](val)
	case reflect.Int64:
		return CastTo[T, int64](val)
	case reflect.Uint:
		return CastTo[T, uint](val)
	case reflect.Uint8:
		return CastTo[T, uint8](val)
	case reflect.Uint16:
		return CastTo[T, uint16](val)
	case reflect.Uint32:
		return CastTo[T, uint32](val)
	case reflect.Uint64:
		return CastTo[T, uint64](val)
	case reflect.String:
		return CastTo[T, string](val)
	default:
		err := fmt.Errorf("%w: %T", errorz.ErrInvalidType, val)
		return nil, err
	}
}
