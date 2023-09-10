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

type IntType int
type Int64Type int64
type Float32Type float32

func TestCastNil(t *testing.T) {

	assert.Nil(t, MustCastTo[TestStruct](nil))
}

func TestCastEqualType(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, MustCastTo[TestStruct](testStruct))
}

func TestCastNotNeeded(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, MustCastTo[TestStruct](testStruct))
}

func TestCastDereferencesPointer(t *testing.T) {

	var testStruct TestStruct
	assert.Equal(t, &testStruct, MustCastTo[TestStruct](&testStruct))
}

func TestCastHandlesConvertibleTypes(t *testing.T) {

	var testInt64 int64 = 10
	assert.Equal(t, primitives.Ref[int](10), MustCastTo[int](testInt64))
	assert.Equal(t, primitives.Ref[float32](10), MustCastTo[float32](testInt64))
}

func TestCastHandlesNumberToString(t *testing.T) {

	var (
		expected string
		actual   *string
	)
	expected = "10"

	var testInt int = 10
	actual = MustCastTo[string](testInt)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)

	var testInt64 int64 = 10
	actual = MustCastTo[string](testInt64)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)

	expected = "10.1"
	var testFloat32 float32 = 10.1
	actual = MustCastTo[string](testFloat32)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)

}

func TestCastHandlesTypedNumberToString(t *testing.T) {

	var (
		expected string
		actual   *string
	)
	expected = "10"

	var testInt IntType = 10
	actual = MustCastTo[string](testInt)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)

	var testInt64 Int64Type = 10
	actual = MustCastTo[string](testInt64)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)

	expected = "10.1"
	var testFloat32 Float32Type = 10.1
	actual = MustCastTo[string](testFloat32)
	assert.NotNil(t, actual)
	assert.Equal(t, expected, *actual)

}

func TestCastHandlesStringToNumber(t *testing.T) {

	testString := "10"

	var expectedInt int = 10
	actualInt := MustCastTo[int](testString)
	assert.NotNil(t, actualInt)
	assert.Equal(t, expectedInt, *actualInt)

	var expectedInt64 int64 = 10
	actualInt64 := MustCastTo[int64](testString)
	assert.NotNil(t, actualInt64)
	assert.Equal(t, expectedInt64, *actualInt64)

	testString = "10.1"
	var expectedFloat32 float32 = 10.1
	actualFloat32 := MustCastTo[float32](testString)
	assert.NotNil(t, actualFloat32)
	assert.Equal(t, expectedFloat32, *actualFloat32)

}

func TestCastHandlesIncompatibleTypes(t *testing.T) {

	var testStruct TestStruct
	assert.Panics(t, func() { MustCastTo[int](testStruct) })
}

func TestCastHandlesMapstructure(t *testing.T) {

	var (
		testStruct    = TestStruct{Foo: "foo", Bar: "bar"}
		testStructMap map[string]any
	)

	mapstructure.Decode(&testStruct, &testStructMap)

	assert.Equal(t, &testStruct, MustCastTo[TestStruct](testStructMap))
}

func TestCastHandlesSlices(t *testing.T) {

	var testStructs []*TestStruct
	testStructs = append(testStructs, &TestStruct{})
	assert.Equal(t, &testStructs, MustCastTo[[]*TestStruct](testStructs))
}
