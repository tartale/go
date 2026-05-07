package structs

import (
	"encoding/json"
	"reflect"

	"github.com/tartale/go/pkg/filter"
	"github.com/tartale/go/pkg/jsonx"
	"github.com/tartale/go/pkg/maps"
)

var typeOfOperator = reflect.TypeFor[*filter.Operator]()

// DynamicFilter is an object that allows a caller to
// filter a list of objects by their fields, using a
// JSON-compatible expression language.
type DynamicFilter[T any] struct {
	Any any
}

// NewDynamicFilter creates a dynamic struct that mirrors
// the input type T, and can be used for filtering
// lists of objects of type T by the fields of T.
func NewDynamicFilter[T any]() DynamicFilter[T] {
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
	structWrapper := New(t)
	structWrapper.TagName = "json"
	structWrapper.Walk(filterWalkFn)
	newStructType := reflect.StructOf(newFields)
	sliceOfNewStructType := reflect.SliceOf(newStructType)
	// Each DynamicFilter has two additional fields "And" and "Or"
	// of type []DynamicFilter; these provide the ability for
	// the DynamicFilter to represent compound boolean expressions.
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
	dynamicFilter := DynamicFilter[T]{reflect.New(sliceOfNewStructType).Interface()}

	return dynamicFilter
}

// NewDynamicFilterFromJson is a convenience method to create a dynamic
// DynamicFilter struct, and instantiate an object of that type
// with the provided JSON string.
func NewDynamicFilterFromJson[T any](inputJson string) DynamicFilter[T] {
	dynamicFilter := NewDynamicFilter[T]()
	jsonx.MustUnmarshalFromString(inputJson, &dynamicFilter)

	return dynamicFilter
}

// MarshalJSON overrides the default JSON marshal function
// so that the inner type is marshalled instead of the
// DynamicFilter outer wrapper.
func (df DynamicFilter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(df.Any)
}

// UnmarshalJSON overrides the default JSON unmarshal function
// so that the inner type is unmarshalled instead of the
// DynamicFilter outer wrapper.
func (df DynamicFilter[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, df.Any)
}

// ShouldInclude accepts an object of type T and determines
// whether it passes the DynamicFilter for type T.
func (df DynamicFilter[T]) ShouldInclude(val any) bool {
	expression := df.GetExpression()
	structWrapper := New(val)
	structWrapper.TagName = "json"
	mapOfValues := structWrapper.Map()
	mapOfValues = maps.CastPrimitives(mapOfValues)
	eval := filter.MustEvaluate(expression, mapOfValues)

	return eval.(bool)
}

func (df DynamicFilter[T]) GetExpression() string {
	return filter.GetExpression(df.Any)
}
