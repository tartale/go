package filter

import (
	"encoding/json"
	"iter"
	"reflect"
	"regexp"
	"strings"

	"github.com/PaesslerAG/gval"
	"github.com/tartale/go/pkg/jsonx"
	"github.com/tartale/go/pkg/primitives"
	"github.com/tartale/go/pkg/reflectx"
	"github.com/tartale/go/pkg/structs"
)

var (
	quotedFields   = regexp.MustCompile(`"(\w+)":`)
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
	Any any
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

func (tf TypeFilter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(tf.Any)
}

func (tf TypeFilter[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, tf.Any)
}

func (tf TypeFilter[T]) ShouldInclude(val any) bool {
	expression := tf.GetExpression()
	mapOfValues := GetMapOfValues(val)
	eval := mustEvaluate(expression, mapOfValues)

	return eval.(bool)
}

func (tf TypeFilter[T]) GetExpression() string {
	filterableJson := jsonx.MustMarshalToString(tf.Any)
	expression := format(filterableJson)

	return expression
}

func (tf TypeFilter[T]) Filter(vals iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range vals {
			if !tf.ShouldInclude(v) {
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
