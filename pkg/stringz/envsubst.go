package stringz

import (
	"fmt"
	"reflect"

	"github.com/a8m/envsubst"
	"github.com/tartale/go/pkg/errorz"
	"github.com/tartale/go/pkg/generics"
	"github.com/tartale/go/pkg/reflectx"
)

// Envsubst expands environment variables in the string pointed to by val.
//
// The argument must be a non-nil pointer to a string-like value; otherwise
// an error is returned. On success, the pointed-to value is replaced with
// the expanded string.
//
// Example:
//
//	s := "hello ${WORLD}"
//	os.Setenv("WORLD", "earth")
//	if err := stringz.Envsubst(&s); err != nil {
//		log.Fatal(err)
//	}
//	// s == "hello earth"
func Envsubst(val any) error {
	if val == nil {
		return fmt.Errorf("%w: argument is nil", errorz.ErrInvalidArgument)
	}

	tval := reflect.TypeOf(val)
	kval := tval.Kind()
	if kval != reflect.Pointer {
		return fmt.Errorf("%w: argument is not a pointer", errorz.ErrInvalidArgument)
	}

	vval := reflect.ValueOf(val).Elem()
	tval = vval.Type()
	ival := vval.Interface()
	if !reflectx.IsString(ival) {
		return fmt.Errorf("%w: argument is not a string", errorz.ErrInvalidType)
	}

	s, err := generics.CastTo[string](ival)
	if err != nil {
		return err
	}
	r, err := envsubst.String(*s)
	if err != nil {
		return err
	}
	rval := reflect.ValueOf(r)
	rval = rval.Convert(tval)
	vval.Set(rval)

	return nil
}
