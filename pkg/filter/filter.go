package filter

import (
	"encoding/json"
	"fmt"
	"iter"
	"reflect"
	"regexp"
	"strings"

	"github.com/PaesslerAG/gval"
	"github.com/tartale/go/pkg/jsonx"
	"github.com/tartale/go/pkg/primitives"
	"github.com/tartale/go/pkg/reflectx"
	"github.com/tartale/go/pkg/structs"
	"golang.org/x/exp/maps"
)

var (
	quotedFields   = regexp.MustCompile(`"(\w+)":`)
	typeOfString   = reflect.TypeFor[string]()
	typeOfOperator = reflect.TypeFor[*Operator]()
)

type Operator struct {
	Eq      any `json:"eq,omitempty"`
	Ne      any `json:"ne,omitempty"`
	Lte     any `json:"lte,omitempty"`
	Gte     any `json:"gte,omitempty"`
	Lt      any `json:"lt,omitempty"`
	Gt      any `json:"gt,omitempty"`
	Matches any `json:"matches,omitempty"`
}

type TypeFilter[T any] struct {
	any
}

func NewTypeFilter[T any]() TypeFilter[T] {
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
	newFields = append(newFields, andType, orType)
	newStructType = reflect.StructOf(newFields)
	sliceOfNewStructType = reflect.SliceOf(newStructType)
	// Fixup the "and" and "or" types to ensure they the addition of themselves,
	// then remake the struct type and slice of one more time
	newFields[len(newFields)-2].Type = sliceOfNewStructType
	newFields[len(newFields)-1].Type = sliceOfNewStructType
	newStructType = reflect.StructOf(newFields)
	sliceOfNewStructType = reflect.SliceOf(newStructType)
	typeFilter := TypeFilter[T]{reflect.New(sliceOfNewStructType).Interface()}

	return typeFilter
}

func NewTypeFilterFromJson[T any](inputJson string) TypeFilter[T] {
	typeFilter := NewTypeFilter[T]()
	jsonx.MustUnmarshalFromString(inputJson, &typeFilter)

	return typeFilter
}

func (f TypeFilter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.any)
}

func (f TypeFilter[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, f.any)
}

func (f TypeFilter[T]) ShouldInclude(val any) bool {
	expression := f.GetExpression()
	mapOfValues := GetMapOfValues(val)
	eval := mustEvaluate(expression, mapOfValues)

	return eval.(bool)
}

func (f TypeFilter[T]) GetExpression() string {
	filterableJson := jsonx.MustMarshalToString(f)
	expression := format(filterableJson)

	return expression
}

func (f TypeFilter[T]) Filter(vals iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range vals {
			if !f.ShouldInclude(v) {
				continue
			}
			if !yield(v) {
				break
			}
		}
	}
}

func GetMapOfValues(val any) map[string]any {
	structWrapper := structs.New(val)
	structWrapper.TagName = "json"
	mapOfValues := structWrapper.Map()
	for k, v := range mapOfValues {
		if reflectx.IsPrimitive(v) {
			mapOfValues[k] = primitives.MustCastAway(v)
		}
	}
	return mapOfValues
}

func mustEvaluate(expression string, parameter interface{}, opts ...gval.Language) any {
	result, err := gval.Evaluate(expression, parameter, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

////////////// legacy //////////////

// GetExpression takes a filter object (i.e. an instance
// of a struct that has fields of type filter.Operator) and
// converts it to a string that can be fed into the
// gval.Evaluate function.
func GetExpression(filter any) string {
	filterValue := reflect.ValueOf(filter)
	if !structs.IsSlice(filterValue) {
		filter = []any{filter}
	}
	filterBytes, err := json.Marshal(filter)
	if err != nil {
		panic(fmt.Errorf("unexpected error when marshaling filter: %w", err))
	}

	filterJson := string(filterBytes)
	expression := format(filterJson)

	return expression
}

// GetValues turns an input object into a map of field names
// to values that can be fed into the gval.Evaluate function.
//
// The resulting map only has keys that are part of the passed
// filter object.
//
// Example:
//
//		filter:     {kind: {eq: "SERIES"}}
//		input:      {kind: "MOVIE", title: "Back to the Future"}
//	  values:     {kind => "MOVIE"}
func GetValues(filter, input any) map[string]any {
	filterValue := reflect.ValueOf(filter)
	if !structs.IsSlice(filterValue) {
		filter = []any{filter}
	}

	values := map[string]any{}
	for i := 0; i < filterValue.Len(); i++ {
		f := filterValue.Index(i).Interface()
		v := getValues(f, input)
		maps.Copy(values, v)
	}

	return values
}

func getValues(filter, input any) map[string]any {
	values := map[string]any{}
	filterWalkFn := func(filterField reflect.StructField, filterValue reflect.Value) error {
		if filterValue.IsNil() {
			return nil
		}
		switch filterValue.Interface().(type) {
		case *Operator:

			inputField, ok := structs.New(input).FieldOk(filterField.Name)
			if !ok {
				panic(fmt.Errorf("filter contains a field that is not in the input: %s", filterField.Name))
			}
			inputFieldName := inputField.TagRoot("json")
			inputFieldValue := inputField.Value()
			inputFieldReflectValue := reflect.ValueOf(inputFieldValue)
			if inputFieldReflectValue.Kind() == reflect.String {
				inputFieldValue = inputFieldReflectValue.Convert(typeOfString).Interface()
			}
			values[inputFieldName] = inputFieldValue
		}

		return nil
	}

	structs.Walk(filter, filterWalkFn)

	return values
}

func removeQuotesOnFields(s string) string {
	return quotedFields.ReplaceAllString(s, "$1")
}

func replaceComparisonOperators(s string) string {
	s = regexp.MustCompile(`{eq(.*?)}`).ReplaceAllString(s, " == $1 ")
	s = regexp.MustCompile(`{ne(.*?)}`).ReplaceAllString(s, " != $1 ")
	s = regexp.MustCompile(`{lte(.*?)}`).ReplaceAllString(s, " <= $1 ")
	s = regexp.MustCompile(`{gte(.*?)}`).ReplaceAllString(s, " >= $1 ")
	s = regexp.MustCompile(`{lt(.*?)}`).ReplaceAllString(s, " < $1 ")
	s = regexp.MustCompile(`{gt(.*?)}`).ReplaceAllString(s, " > $1 ")
	s = regexp.MustCompile(`{matches(.*?)}`).ReplaceAllString(s, " =~ $1 ")

	return s
}

func replaceBrackets(s string) string {
	return strings.NewReplacer(
		`[`, `(`,
		`]`, `)`,
		`{`, `(`,
		`}`, `)`,
	).Replace(s)
}

func replaceLogicOperators(s string) string {
	return strings.NewReplacer(
		`,(or`, ` || (`,
		`,(and`, ` && (`,
		`,`, ` && `,
	).Replace(s)
}

func format(expression string) string {
	expression = removeQuotesOnFields(expression)
	expression = replaceComparisonOperators(expression)
	expression = replaceBrackets(expression)
	expression = replaceLogicOperators(expression)

	return expression
}
