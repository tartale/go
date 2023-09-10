package reflectx

import (
	"github.com/tartale/go/pkg/constraintz"
)

func Switch[T constraintz.Primitive](val any, fn func(val T)) {

	switch v := val.(type) {
	default:
		fn(v.(T))
	}

	// case *bool:
	// 	result, err := ParseTo[bool](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = result

	// case *float32:
	// 	result, err := ParseTo[float32](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = float32(result)

	// case *float64:
	// 	result, err := ParseTo[float64](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = float64(result)

	// case *int:
	// 	result, err := ParseTo[int](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = int(result)

	// case *int8:
	// 	result, err := ParseTo[int8](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = int8(result)

	// case *int16:
	// 	result, err := ParseTo[int16](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = int16(result)

	// case *int32:
	// 	result, err := ParseTo[int32](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = int32(result)

	// case *int64:
	// 	result, err := ParseTo[int64](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = int64(result)

	// case *uint:
	// 	result, err := ParseTo[uint](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = uint(result)

	// case *uint8:
	// 	result, err := ParseTo[uint8](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = uint8(result)

	// case *uint16:
	// 	result, err := ParseTo[uint16](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = uint16(result)

	// case *uint32:
	// 	result, err := ParseTo[uint32](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = uint32(result)

	// case *uint64:
	// 	result, err := ParseTo[uint64](s, optionalBase...)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	*v = uint64(result)

	// case *string:
	// 	*v = s

	// default:
	// 	return fmt.Errorf("unexpected type: %T; ensure passed argument is a pointer to a primitive", primitivePtr)
	// }

}
