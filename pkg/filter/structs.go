package filter

import (
	"encoding/json"
	"reflect"

	"github.com/tartale/go/pkg/jsonx"
	"github.com/tartale/go/pkg/maps"
	"github.com/tartale/go/pkg/structs"
)

var typeOfOperator = reflect.TypeFor[*Operator]()

// StructFilter is an object that allows a caller to
// filter a list of objects by their fields, using a
// JSON-compatible expression language.
type StructFilter[T any] struct {
	Any any
}

// NewStructFilter creates a struct filter that mirrors
// the input type T, and can be used for filtering
// lists of objects of type T by the fields of T.
// The inputJson string is the expression that
// will be used to evaluate inclusion of an item of type T.
func NewStructFilter[T any](inputJson string) StructFilter[T] {
	structFilter := newStructFilter[T]()
	jsonx.MustUnmarshalFromString(inputJson, &structFilter)

	return structFilter
}

// MarshalJSON overrides the default JSON marshal function
// so that the inner type is marshalled instead of the
// StructFilter outer wrapper.
func (df StructFilter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(df.Any)
}

// UnmarshalJSON overrides the default JSON unmarshal function
// so that the inner type is unmarshalled instead of the
// StructFilter outer wrapper.
func (df StructFilter[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, df.Any)
}

// ShouldInclude accepts an object of type T and determines
// whether it passes the StructFilter for type T.
func (df StructFilter[T]) ShouldInclude(val any) bool {
	expression := GetExpression(df.Any)
	structWrapper := structs.New(val)
	structWrapper.TagName = "json"
	mapOfValues := structWrapper.Map()
	mapOfValues = maps.CastPrimitives(mapOfValues)
	eval := MustEvaluate(expression, mapOfValues)

	return eval.(bool)
}

// newStructFilter creates a struct filter that mirrors
// the input type T, and can be used for filtering
// lists of objects of type T by the fields of T.
func newStructFilter[T any]() StructFilter[T] {
	// Walk the input struct and make a new struct that
	// has all same field names, but with the type *Operator
	// instead of the original type.
	var newFields []reflect.StructField
	filterWalkFn := func(filterField reflect.StructField, filterValue reflect.Value) error {
		newField := filterField
		newField.Type = typeOfOperator
		newFields = append(newFields, newField)
		return nil
	}
	var t T
	structWrapper := structs.New(t)
	structWrapper.TagName = "json"
	structWrapper.Walk(filterWalkFn)
	newStructType := reflect.StructOf(newFields)
	sliceOfNewStructType := reflect.SliceOf(newStructType)
	// Each StructFilter has two additional fields "And" and "Or"
	// of type []StructFilter; these provide the ability for
	// the StructFilter to represent compound boolean expressions.
	andType := reflect.StructField{
		Name: "And",
		Type: sliceOfNewStructType,
		Tag:  reflect.StructTag(`json:"and,omitempty"`),
	}
	orType := reflect.StructField{
		Name: "Or",
		Type: sliceOfNewStructType,
		Tag:  reflect.StructTag(`json:"or,omitempty"`),
	}
	// Add the "And" and "Or" fields, and then recreate the
	// dynamic struct and slice of dynamic struct.
	newFields = append(newFields, andType, orType)
	newStructType = reflect.StructOf(newFields)
	sliceOfNewStructType = reflect.SliceOf(newStructType)
	// Fixup the "And" and "Or" types to ensure the addition of themselves,
	// then remake the struct type and slice of one more time
	newFields[len(newFields)-2].Type = sliceOfNewStructType
	newFields[len(newFields)-1].Type = sliceOfNewStructType
	newStructType = reflect.StructOf(newFields)
	sliceOfNewStructType = reflect.SliceOf(newStructType)
	structFilter := StructFilter[T]{reflect.New(sliceOfNewStructType).Interface()}

	return structFilter
}
