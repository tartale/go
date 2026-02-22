package reflectx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	IntType     int
	Int64Type   int64
	Float32Type float32
)

type TestStruct struct {
	String string
}

func TestIsNumber(t *testing.T) {
	assert.True(t, IsNumber(1))
	assert.True(t, IsNumber(1.0))
	assert.True(t, IsNumber(IntType(1)))
}

func TestValueOfElement(t *testing.T) {
	testStruct := TestStruct{String: "hello"}
	valueOfTestStruct := reflect.ValueOf(testStruct)

	testStructPtr := &testStruct
	valueOfTestStructPtrElement := ValueOfElement(testStructPtr)
	assert.Equal(t, valueOfTestStruct.Kind(), valueOfTestStructPtrElement.Kind())
	assert.Equal(t, valueOfTestStruct.String(), valueOfTestStructPtrElement.String())

	testStructInterface := any(testStruct)
	valueOfTestStructInterfaceElement := ValueOfElement(testStructInterface)
	assert.Equal(t, valueOfTestStruct.Kind(), valueOfTestStructInterfaceElement.Kind())
	assert.Equal(t, valueOfTestStruct.String(), valueOfTestStructInterfaceElement.String())
}

func TestTypeOfElement(t *testing.T) {
	testStruct := TestStruct{}
	testStructPtr := &testStruct
	typeOfTestStruct := reflect.TypeOf(testStruct)
	typeOfTestStructPtrElement := TypeOfElement(testStructPtr)
	assert.Equal(t, typeOfTestStruct, typeOfTestStructPtrElement)
}
