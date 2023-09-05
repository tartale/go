package gqlgen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

func GetArgValue[T any](ctx context.Context, path string, argName string) *T {

	argumentList := GetArgList(ctx, path)
	arg := argumentList.ForName(argName)
	if arg == nil {
		return nil
	}
	val, err := arg.Value.Value(nil)
	if err != nil {
		return nil
	}

	return val.(*T)
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
