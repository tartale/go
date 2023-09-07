package generics

import (
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/primitives"
)

type TestStruct struct {
	Foo string `json:"foo,omitempty"`
	Bar string `json:"bar,omitempty"`
}

func TestCastNil(t *testing.T) {

	assert.Nil(t, Cast[TestStruct](nil))
}

func TestCastEqualType(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, Cast[TestStruct](testStruct))
}

func TestCastNotNeeded(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, Cast[TestStruct](testStruct))
}

func TestCastDereferencesPointer(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, Cast[TestStruct](&testStruct))
}

func TestCastHandlesConvertibleTypes(t *testing.T) {

	var testInt64 int64 = 10
	assert.Equal(t, primitives.Ref[int](10), Cast[int](testInt64))
	assert.Equal(t, primitives.Ref[float32](10), Cast[float32](testInt64))
}

func TestCastHandlesIncompatibleTypes(t *testing.T) {

	var testStruct TestStruct
	assert.Nil(t, Cast[int](testStruct))
}

func TestCastHandlesMapstructure(t *testing.T) {

	var (
		testStruct    = TestStruct{Foo: "foo", Bar: "bar"}
		testStructMap map[string]any
	)

	mapstructure.Decode(&testStruct, &testStructMap)

	assert.Equal(t, &testStruct, Cast[TestStruct](testStructMap))
}

func TestCastHandlesSlices(t *testing.T) {

	var testStructs []*TestStruct
	testStructs = append(testStructs, &TestStruct{})
	assert.Equal(t, &testStructs, Cast[[]*TestStruct](testStructs))
}
