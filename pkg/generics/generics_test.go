package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/primitives"
)

type TestStruct struct{}

func TestNormalizeNil(t *testing.T) {

	assert.Nil(t, Normalize[TestStruct](nil))
}

func TestNormalizeEqualType(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, Normalize[TestStruct](testStruct))
}

func TestNormalizeNotNeeded(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, Normalize[TestStruct](testStruct))
}

func TestNormalizeDereferencesPointer(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, Normalize[TestStruct](&testStruct))
}

func TestNormalizeHandlesConvertibleTypes(t *testing.T) {

	var testInt64 int64 = 10
	assert.Equal(t, primitives.Ref[int](10), Normalize[int](testInt64))
	assert.Equal(t, primitives.Ref[float32](10), Normalize[float32](testInt64))
}

func TestNormalizeHandlesIncompatibleTypes(t *testing.T) {

	var testStruct TestStruct
	assert.Nil(t, Normalize[int](testStruct))
}
