package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRef(t *testing.T) {

	var (
		intVal     *int
		int16Val   *int16
		int32Val   *int32
		int64Val   *int64
		float32Val *float32
		float64Val *float64
		stringVal  *string
	)

	intVal = Ref[int](1)
	assert.Equal(t, 1, *intVal)

	int16Val = Ref[int16](2)
	assert.Equal(t, int16(2), *int16Val)

	int32Val = Ref[int32](3)
	assert.Equal(t, int32(3), *int32Val)

	int64Val = Ref[int64](4)
	assert.Equal(t, int64(4), *int64Val)

	float32Val = Ref[float32](5.0)
	assert.Equal(t, float32(5.0), *float32Val)

	float64Val = Ref[float64](6.0)
	assert.Equal(t, float64(6.0), *float64Val)

	stringVal = Ref[string]("foo")
	assert.Equal(t, "foo", *stringVal)
}
