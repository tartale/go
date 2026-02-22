package primitives

import (
	"fmt"
	"reflect"

	"github.com/tartale/go/pkg/constraintz"
	"github.com/tartale/go/pkg/errorz"
	"github.com/tartale/go/pkg/reflectx"
)

func CastTo[U constraintz.Primitive](val any) (U, error) {
	var result U
	s := fmt.Sprintf("%v", val)
	Parse(s, &result)

	return result, nil
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
func CastAway(val any) (any, error) {
	valType := reflectx.TypeOfElement(val)
	valKind := valType.Kind()
	switch valKind {
	case reflect.Bool:
		return CastTo[bool](val)
	case reflect.Float32:
		return CastTo[float32](val)
	case reflect.Float64:
		return CastTo[float64](val)
	case reflect.Int:
		return CastTo[int](val)
	case reflect.Int8:
		return CastTo[int8](val)
	case reflect.Int16:
		return CastTo[int16](val)
	case reflect.Int32:
		return CastTo[int32](val)
	case reflect.Int64:
		return CastTo[int64](val)
	case reflect.Uint:
		return CastTo[uint](val)
	case reflect.Uint8:
		return CastTo[uint8](val)
	case reflect.Uint16:
		return CastTo[uint16](val)
	case reflect.Uint32:
		return CastTo[uint32](val)
	case reflect.Uint64:
		return CastTo[uint64](val)
	case reflect.String:
		return CastTo[string](val)
	default:
		err := fmt.Errorf("%w: %T", errorz.ErrInvalidType, val)
		return nil, err
	}
}

func MustCastAway(val any) any {
	result, err := CastAway(val)
	if err != nil {
		panic(err)
	}
	return result
}
