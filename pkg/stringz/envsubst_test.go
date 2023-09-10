package stringz

import (
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tartale/go/pkg/errorz"
)

func TestStringPtr(t *testing.T) {

	os.Setenv("FOO", "foo")

	input := "/${FOO}/bar"
	expected := "/foo/bar"
	err := Envsubst(&input)
	assert.Nil(t, err)
	assert.Equal(t, expected, input)
}

func TestStringPtrReflectVal(t *testing.T) {

	os.Setenv("FOO", "foo")

	input := "/${FOO}/bar"
	expected := "/foo/bar"
	rvinput := reflect.ValueOf(input).Interface()
	err := Envsubst(&rvinput)
	assert.Nil(t, err)
	assert.Equal(t, expected, rvinput.(string))
}

func TestTypedStringPtr(t *testing.T) {

	os.Setenv("FOO", "foo")
	type TestType string

	input := TestType("/${FOO}/bar")
	expected := TestType("/foo/bar")
	err := Envsubst(&input)
	assert.Nil(t, err)
	assert.Equal(t, expected, input)
}

func TestTypedStringPtrReflectVal(t *testing.T) {

	os.Setenv("FOO", "foo")
	type TestType string

	input := TestType("/${FOO}/bar")
	expected := TestType("/foo/bar")
	rvinput := reflect.ValueOf(input).Interface()
	err := Envsubst(&rvinput)
	assert.Nil(t, err)
	assert.Equal(t, expected, TestType(rvinput.(string)))
}

func TestNil(t *testing.T) {

	err := Envsubst(nil)
	assert.ErrorIs(t, err, errorz.ErrInvalidArgument)
}

func TestNonStringPtr(t *testing.T) {

	input := 1
	err := Envsubst(&input)
	assert.ErrorIs(t, err, errorz.ErrInvalidType)
}

func TestTypedNonStringPtr(t *testing.T) {

	type TestType int
	input := TestType(1)
	err := Envsubst(&input)
	assert.ErrorIs(t, err, errorz.ErrInvalidType)
}
