package filter

import (
	"fmt"
	"iter"
	"maps"
	"reflect"
	"regexp"
	"slices"
	"strings"

	"github.com/tartale/go/pkg/jsonx"
	"github.com/tartale/go/pkg/reflectx"
	"github.com/tartale/go/pkg/structs"
)

var (
	quotedFields = regexp.MustCompile(`"(\w+)":`)
	typeOfString = reflect.TypeFor[string]()
)

type Filterer interface {
	ShouldInclude(val any) bool
}

type FiltererOf[T any] interface {
	ShouldInclude(val T) bool
}

type Operator struct {
	Eq      any `json:"eq,omitempty"`
	Ne      any `json:"ne,omitempty"`
	Lte     any `json:"lte,omitempty"`
	Gte     any `json:"gte,omitempty"`
	Lt      any `json:"lt,omitempty"`
	Gt      any `json:"gt,omitempty"`
	Matches any `json:"matches,omitempty"`
}

// GetExpression turns the JSON representation of the
// given object into a boolean expression that can
// be used in the "gval.Evaluate" library. For example,
// if the JSON of the given object looks something
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
func GetExpression(filter any) string {
	filterValue := reflect.ValueOf(filter)
	if !reflectx.IsSlice(filterValue.Interface()) {
		filter = []any{filter}
	}
	filterJson := jsonx.MustMarshalToString(filter)
	expression := Format(filterJson)

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
	if !reflectx.IsSlice(filterValue) {
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

// Format does the full conversion of a JSON string into
// a gval-compatible boolean expression.
func Format(expression string) string {
	expression = removeQuotesOnFields(expression)
	expression = replaceComparisonOperators(expression)
	expression = replaceBrackets(expression)
	expression = replaceLogicOperators(expression)

	return expression
}

// Filter takes a sequence iterator to the filtered type T, and returns
// an iterator function that can be applied to that input sequence.
func Filter(f Filterer, vals iter.Seq[any]) iter.Seq[any] {
	return func(yield func(any) bool) {
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

// FilterAll is a wrapper around Filter that
// accepts and returns slices instead of iterators.
func FilterAll(f Filterer, vals []any) []any {
	filterVals := Filter(f, slices.Values(vals))
	return slices.Collect(filterVals)
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
