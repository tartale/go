package contexts

import (
	"context"

	"github.com/tartale/go/pkg/generics"
)

// Value retrieves a typed value from ctx for key, returning nil if it is absent
// or cannot be cast to T.
func Value[T any](ctx context.Context, key any) *T {
	val, _ := ValueE[T](ctx, key)
	return val
}

// ValueE retrieves a typed value from ctx for key and returns it along with
// any casting error.
func ValueE[T any](ctx context.Context, key any) (*T, error) {
	if val := ctx.Value(key); val != nil {
		return generics.CastTo[T](val)
	}

	return nil, nil
}
