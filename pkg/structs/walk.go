package structs

import (
	"reflect"

	"github.com/tartale/go/pkg/reflectx"
)

// WalkFn is a callback that can be used in the Walk method, below.
// While iterating the struct's fields, the callback will be invoked
// for every field/value pair.  If the callback function returns
// an error, the iterating is aborted and the error returned from
// the Walk function itself.
type WalkFn func(field reflect.StructField, value reflect.Value) error

// Walk is a convenience function for the Struct.Walk() method.
func Walk(s any, fn WalkFn) error {
	return New(s).Walk(fn)
}

// Walk iterates the struct's fields using depth-first search, calling the given
// callback function.  It panics if s's kind is not struct.  Supports
// the 'omitnested' and 'flatten' tags on nested struct fields.
// If 'omitnested' is present, the Walk callback will be invoked for the
// nested struct field itself, but not for the nested struct's fields.
// If 'flatten' is present, the Walk callback will be invoked for the
// nested struct's fields, but not for the top-level nested struct itself.
func (s *Struct) Walk(fn WalkFn) error {
	fields := s.structFields()

	for _, field := range fields {
		val := s.reflectValueOfElement.FieldByName(field.Name)

		if err := s.WalkValue(field, val, fn); err != nil {
			return err
		}
	}

	return nil
}

func (s *Struct) WalkValue(field reflect.StructField, val reflect.Value, fn WalkFn) error {
	_, tagOpts := parseTag(field.Tag.Get(s.TagName))

	if !tagOpts.Has("omitnested") && reflectx.IsStruct(val.Interface()) {
		return s.WalkSubStruct(field, val, fn, tagOpts.Has("flatten"))
	}

	if reflectx.IsSlice(val.Interface()) {
		return s.WalkSlice(val, fn)
	}

	return fn(field, val)
}

// WalkSubStruct deals with sub-structures. If the 'flatten' tag is not set, it calls the walk function
// on the current field. In any case, it walks over the nested structure.
func (s *Struct) WalkSubStruct(field reflect.StructField, val reflect.Value, fn WalkFn, flatten bool) error {
	if !flatten {
		if err := fn(field, val); err != nil {
			return err
		}
	}

	return New(val.Interface()).Walk(fn)
}

// WalkSlice walks over each element of the slice, if the element is a struct.
func (s *Struct) WalkSlice(val reflect.Value, fn WalkFn) error {
	for i := 0; i < val.Len(); i++ {
		v := val.Index(i)

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if v.CanAddr() {
			v = v.Addr()
		}

		if reflectx.IsStruct(v.Interface()) {
			if err := New(v.Interface()).Walk(fn); err != nil {
				return err
			}
		}
	}

	return nil
}
