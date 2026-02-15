package gqlgen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tartale/go/pkg/errorz"
	"github.com/tartale/go/pkg/generics"
	"github.com/vektah/gqlparser/v2/ast"
)

// ArgKey is an object that must be
// provided when attempting to get an argument
// from a gqlgen context. The path represents
// the JSON path into the query, while the name
// identifies the argument.
type ArgKey struct {
	Path string
	Name string
}

// ErrFieldNotFound is returned when no field can be located in
// the gqlgen context at the given path.
var ErrFieldNotFound = fmt.Errorf("%w: field", errorz.ErrNotFound)

// ErrArgumentNotFound is returned when no argument with the given
// name can be located in gqlgen field at the given path. This
// error could indicated that an optional argument was simply not
// provided in the GraphQL query itself.
var ErrArgumentNotFound = fmt.Errorf("%w: argument", errorz.ErrNotFound)

func (a ArgKey) String() string {
	return fmt.Sprintf("%s.%s", a.Path, a.Name)
}

// MustGetArgValue is a convenience function that wraps GetArgValue,
// but panics if an error occurs.
func MustGetArgValue[T any](ctx context.Context, key ArgKey) *T {

	val, err := GetArgValue[T](ctx, key)
	if err != nil {
		panic(err)
	}
	return val
}

// GetArgValue searches the gqlgen context for a query argument
// identified by path and name in the ArgKey, and returns
// the value casted to the given type T. GetArgValue returns an
// error if:
//   - The context is not a gqlgen-provided context
//   - The field specified by the path is not found in the query
//   - The argument is not provided in the identified query field
//   - The value of the argument cannot be cast to the desired type.
func GetArgValue[T any](ctx context.Context, key ArgKey) (*T, error) {

	argumentList, err := GetArgList(ctx, key.Path)
	if err != nil {
		return nil, err
	}
	arg := argumentList.ForName(key.Name)
	if arg == nil {
		return nil, fmt.Errorf("%w '%s'", ErrArgumentNotFound, key)
	}
	val, err := arg.Value.Value(nil)
	if err != nil {
		return nil, err
	}

	return generics.CastTo[T](val)
}

// GetArgList searches the gqlgen context for a query field
// identified by path and returns the list of arguments passed to the query.
func GetArgList(ctx context.Context, path string) (ast.ArgumentList, error) {

	fctx := graphql.GetFieldContext(ctx)
	field := FindField(ctx, path, fctx)
	if field == nil {
		return nil, fmt.Errorf("%w '%s'", ErrFieldNotFound, path)
	}

	return field.Arguments, nil
}

// FindField does a depth-first search of the gqlgen context for the
// field identified by path.
func FindField(ctx context.Context, path string, fctx *graphql.FieldContext) *graphql.CollectedField {

	if !graphql.HasOperationContext(ctx) {
		return nil
	}
	octx := graphql.GetOperationContext(ctx)
	if fctx == nil || octx == nil {
		return nil
	}
	fieldPath := fctx.Path()
	fieldPathStr := fieldPath.String()
	if fieldPathStr == path {
		return &fctx.Field
	}

	collectedFields := graphql.CollectFields(octx, fctx.Field.Selections, nil)
	for _, cf := range collectedFields {
		fullPath := fmt.Sprintf("%s.%s", fieldPathStr, cf.Field.Name)
		if path == fullPath {
			return &cf
		}
		child, err := fctx.Child(ctx, cf)
		if err != nil {
			continue
		}
		childField := FindField(ctx, path, child)
		if childField != nil {
			return childField
		}
	}

	return nil
}
