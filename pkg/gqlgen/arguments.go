package gqlgen

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tartale/go/pkg/generics"
	"github.com/vektah/gqlparser/v2/ast"
)

type ArgKey struct {
	Path string
	Name string
}

func (a ArgKey) String() string {
	return fmt.Sprintf("%s.%s", a.Path, a.Name)
}

var ErrNotFound = errors.New("not found")

func GetArgValue[T any](ctx context.Context, key ArgKey) *T {

	val, _ := GetArgValueE[T](ctx, key)
	return val
}

func GetArgValueE[T any](ctx context.Context, key ArgKey) (*T, error) {

	argumentList := GetArgList(ctx, key.Path)
	arg := argumentList.ForName(key.Name)
	if arg == nil {
		return nil, fmt.Errorf("%w: %s", ErrNotFound, key)

	}
	val, err := arg.Value.Value(nil)
	if err != nil {
		return nil, err
	}

	return generics.CastE[T](val)
}

func GetArgList(ctx context.Context, path string) ast.ArgumentList {

	fctx := graphql.GetFieldContext(ctx)
	field := FindField(ctx, path, fctx)
	if field == nil {
		return nil
	}

	return field.Arguments
}

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
