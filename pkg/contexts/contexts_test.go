package contexts

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/primitives"
)

type TestKey string

func TestNil(t *testing.T) {

	ctx := context.Background()
	assert.Nil(t, Value[int](ctx, TestKey("key")))
	assert.Nil(t, Value[uint](ctx, TestKey("key")))
	assert.Nil(t, Value[float32](ctx, TestKey("key")))
	assert.Nil(t, Value[float64](ctx, TestKey("key")))
	assert.Nil(t, Value[string](ctx, TestKey("key")))
}

func TestInts(t *testing.T) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, TestKey("key"), 10)
	assert.Equal(t, primitives.Ref(10), Value[int](ctx, TestKey("key")))

	ctx = context.WithValue(ctx, TestKey("key"), uint(10))
	assert.Equal(t, primitives.Ref(uint(10)), Value[uint](ctx, TestKey("key")))
}

func TestFloats(t *testing.T) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, TestKey("key"), float32(10.0))
	assert.Equal(t, primitives.Ref[float32](10.0), Value[float32](ctx, TestKey("key")))

	ctx = context.WithValue(ctx, TestKey("key"), 10.0)
	assert.Equal(t, primitives.Ref(10.0), Value[float64](ctx, TestKey("key")))

	ctx = context.WithValue(ctx, TestKey("key"), float64(10.0))
	assert.Equal(t, primitives.Ref(float64(10.0)), Value[float64](ctx, TestKey("key")))
}

func TestString(t *testing.T) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, TestKey("key"), "hello")
	assert.Equal(t, primitives.Ref("hello"), Value[string](ctx, TestKey("key")))
}

func TestStruct(t *testing.T) {

	ctx := context.Background()

	type TestStruct struct{}
	var testStruct TestStruct

	ctx = context.WithValue(ctx, TestKey("key"), testStruct)
	assert.Equal(t, &testStruct, Value[TestStruct](ctx, TestKey("key")))
}
