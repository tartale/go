package contexts

import (
	"context"
	"reflect"
)

func Value[T any](ctx context.Context, key any) *T {
	if val := ctx.Value(key); val != nil {
		if reflect.ValueOf(val).Kind() == reflect.Ptr {
			return val.(*T)
		}
		v := val.(T)
		return &v
	}

	return nil
}
