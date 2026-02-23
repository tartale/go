package filter

import (
	"encoding/json"
	"iter"
	"reflect"
	"regexp"
	"slices"
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

// TypeFilter is an object that allows a caller to
// filter a list of objects by their fields, using a
// JSON-compatible expression language.
type TypeFilter[T any] struct {
	Any any
}

// NewTypeFilter creates a dynamic struct that mirrors
// the input type T, and can be used for filtering
// lists of objects of type T by the fields of T.
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
	// Each TypeFilter has two additional fields "And" and "Or"
	// of type []TypeFilter; these provide the ability for
	// the TypeFilter to represent compound boolean expressions.
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
	typeFilter := TypeFilter[T]{reflect.New(sliceOfNewStructType).Interface()}

	return typeFilter
}

// NewTypeFilterFromJson is a convenience method to create a dynamic
// TypeFilter struct, and instantiate an object of that type
// with the provided JSON string.
func NewTypeFilterFromJson[T any](inputJson string) TypeFilter[T] {
	typeFilter := NewTypeFilter[T]()
	jsonx.MustUnmarshalFromString(inputJson, &typeFilter)

	return typeFilter
}

// MarshalJSON overrides the default JSON marshal function
// so that the inner type is marshalled instead of the
// TypeFilter outer wrapper.
func (tf TypeFilter[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(tf.Any)
}

// MarshalJSON overrides the default JSON unmarshal function
// so that the inner type is unmarshalled instead of the
// TypeFilter outer wrapper.
func (tf TypeFilter[T]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, tf.Any)
}

// ShouldInclude accepts an object of type T and determines
// whether it passes the TypeFilter for type T.
func (tf TypeFilter[T]) ShouldInclude(val any) bool {
	expression := tf.GetExpression()
	mapOfValues := GetMapOfValues(val)
	eval := mustEvaluate(expression, mapOfValues)

	return eval.(bool)
}

// GetExpression turns the JSON representation of this
// TypeFilter object into a boolean expression that can
// be used in the "gval.Evaluate" library. For example,
// if the JSON of the TypeFilter object looks something
// like this:
//
//	`[{"kind": {"eq": "MOVIE"}}]`
//
// then the equivalent expression will be something like this:
//
//	`(kind == "MOVIE")`
//
// Note that, due to the simplistic nature of the conversion,
// there may be innocuous artififacts in the resulting
// expression, such as extra unnecessary parentheses. However,
// the expression will be usable and correct.
func (tf TypeFilter[T]) GetExpression() string {
	filterableJson := jsonx.MustMarshalToString(tf.Any)
	expression := format(filterableJson)

	return expression
}

// Filter takes a sequence iterator to the filtered type T, and returns
// an iterator function that can be applied to that input sequence.
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

// FilterAll is a wrapper around Filter that
// accepts and returns slices instead of iterators.
func (tf TypeFilter[T]) FilterAll(vals []T) []T {
	filterVals := tf.Filter(slices.Values(vals))
	return slices.Collect(filterVals)
}

// GetMapOfValues takes an object and returns a map
// of all the fields of the object as KV pairs.
// Any field that is a primitive and cast down
// to its underlying type, to ensure that type aliases
// are considered equivalent to their underlying types.
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

// mustEvaluate is a wrapper around the gval.Evaluate function
// that panics if an error is returned.
func mustEvaluate(expression string, parameter any, opts ...gval.Language) any {
	result, err := gval.Evaluate(expression, parameter, opts...)
	if err != nil {
		panic(err)
	}
	return result
}

// removeQuotesOnFields removes all quotes from the input string;
// this is used in the process of converting a JSON string into a
// gval-compatible expression.
func removeQuotesOnFields(s string) string {
	return quotedFields.ReplaceAllString(s, "$1")
}

// replaceComparisonOperators changes all of the names of the
// operators in the given string to their gval-compatible equivalents
// (e.g. 'eq' is converted to '=='); this is used in the process of
// converting a JSON string into a gval-compatible expression.
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

// replaceBrackets converts all of the brackets in a JSON string
// into parentheses; this is used in the process of
// converting a JSON string into a gval-compatible expression.
func replaceBrackets(s string) string {
	return strings.NewReplacer(
		`[`, `(`,
		`]`, `)`,
		`{`, `(`,
		`}`, `)`,
	).Replace(s)
}

// replaceLogicOperators changes all of the names of the logic
// operators in the given string to their gval-compatible equivalents
// (e.g. 'or' is converted to '||'); this is used in the process of
// converting a JSON string into a gval-compatible expression.
// One quirk of this replacement process is that a "," is interpreted
// as an implicit "and" operator (because this is a somewhat natural
// way to interpret a list in JSON). For example, an input JSON string:
//
//	`[{"title": {"matches": "Back to the .*"}}, {"movieYear": {"eq": 1985}}]`
//
// would be converted to:
//
//	`(title =~ "Back to the .*") and (movieYear == 1985)`
func replaceLogicOperators(s string) string {
	return strings.NewReplacer(
		`,(or`, ` || (`,
		`,(and`, ` && (`,
		`,`, ` && `,
	).Replace(s)
}

// format does the full conversion of a JSON string into
// a gval-compatible boolean expression.
func format(expression string) string {
	expression = removeQuotesOnFields(expression)
	expression = replaceComparisonOperators(expression)
	expression = replaceBrackets(expression)
	expression = replaceLogicOperators(expression)

	return expression
}
