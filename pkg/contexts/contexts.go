package contexts

import (
	"context"

	"github.com/tartale/go/pkg/generics"
)

func Value[T any](ctx context.Context, key any) *T {

	val, _ := ValueE[T](ctx, key)
	return val
}

func ValueE[T any](ctx context.Context, key any) (*T, error) {

	if val := ctx.Value(key); val != nil {
		return generics.CastTo[T](val)
	}

	return nil, nil
}
