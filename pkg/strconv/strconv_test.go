package strconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePrimitive_ValidStrings(t *testing.T) {
	var testBool bool
	var testFloat32 float32
	var testFloat64 float64
	var testInt int
	var testInt8 int8
	var testInt16 int16
	var testInt32 int32
	var testInt64 int64
	var testUint uint
	var testUint8 uint8
	var testUint16 uint16
	var testUint32 uint32
	var testUint64 uint64
	var testString string
	var err error

	err = ParsePrimitive("true", &testBool)
	assert.Nil(t, err)
	assert.Equal(t, true, testBool)

	err = ParsePrimitive("3.1415", &testFloat32)
	assert.Nil(t, err)
	assert.Equal(t, float32(3.1415), testFloat32)

	err = ParsePrimitive("3.14159265358979323846", &testFloat64)
	assert.Nil(t, err)
	assert.Equal(t, 3.14159265358979323846, testFloat64)

	err = ParsePrimitive("10", &testInt)
	assert.Nil(t, err)
	assert.Equal(t, 10, testInt)

	err = ParsePrimitive("10", &testInt8)
	assert.Nil(t, err)
	assert.Equal(t, int8(10), testInt8)

	err = ParsePrimitive("10", &testInt16)
	assert.Nil(t, err)
	assert.Equal(t, int16(10), testInt16)

	err = ParsePrimitive("10", &testInt32)
	assert.Nil(t, err)
	assert.Equal(t, int32(10), testInt32)

	err = ParsePrimitive("10", &testInt64)
	assert.Nil(t, err)
	assert.Equal(t, int64(10), testInt64)

	err = ParsePrimitive("10", &testUint)
	assert.Nil(t, err)
	assert.Equal(t, uint(10), testUint)

	err = ParsePrimitive("10", &testUint8)
	assert.Nil(t, err)
	assert.Equal(t, uint8(10), testUint8)

	err = ParsePrimitive("10", &testUint16)
	assert.Nil(t, err)
	assert.Equal(t, uint16(10), testUint16)

	err = ParsePrimitive("10", &testUint32)
	assert.Nil(t, err)
	assert.Equal(t, uint32(10), testUint32)

	err = ParsePrimitive("10", &testUint64)
	assert.Nil(t, err)
	assert.Equal(t, uint64(10), testUint64)

	err = ParsePrimitive("10", &testString)
	assert.Nil(t, err)
	assert.Equal(t, "10", testString)
}

func TestParsePrimitive_InvalidStrings(t *testing.T) {
	var testBool bool
	var testFloat32 float32
	var testFloat64 float64
	var testInt int
	var testInt8 int8
	var testInt16 int16
	var testInt32 int32
	var testInt64 int64
	var testUint uint
	var testUint8 uint8
	var testUint16 uint16
	var testUint32 uint32
	var testUint64 uint64
	var err error

	err = ParsePrimitive("foobar", &testBool)
	assert.NotNil(t, err)
	assert.Equal(t, false, testBool)

	err = ParsePrimitive("foobar", &testFloat32)
	assert.NotNil(t, err)
	assert.Equal(t, float32(0.0), testFloat32)

	err = ParsePrimitive("foobar", &testFloat64)
	assert.NotNil(t, err)
	assert.Equal(t, float64(0.0), testFloat64)

	err = ParsePrimitive("foobar", &testInt)
	assert.NotNil(t, err)
	assert.Equal(t, 0, testInt)

	err = ParsePrimitive("foobar", &testInt8)
	assert.NotNil(t, err)
	assert.Equal(t, int8(0), testInt8)

	err = ParsePrimitive("foobar", &testInt16)
	assert.NotNil(t, err)
	assert.Equal(t, int16(0), testInt16)

	err = ParsePrimitive("foobar", &testInt32)
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), testInt32)

	err = ParsePrimitive("foobar", &testInt64)
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), testInt64)

	err = ParsePrimitive("foobar", &testUint)
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), testUint)

	err = ParsePrimitive("foobar", &testUint8)
	assert.NotNil(t, err)
	assert.Equal(t, uint8(0), testUint8)

	err = ParsePrimitive("foobar", &testUint16)
	assert.NotNil(t, err)
	assert.Equal(t, uint16(0), testUint16)

	err = ParsePrimitive("foobar", &testUint32)
	assert.NotNil(t, err)
	assert.Equal(t, uint32(0), testUint32)

	err = ParsePrimitive("foobar", &testUint64)
	assert.NotNil(t, err)
	assert.Equal(t, uint64(0), testUint64)
}

func TestParsePrimitive_NonBase10(t *testing.T) {
	var testBool bool
	var testFloat32 float32
	var testFloat64 float64
	var testInt int
	var testInt8 int8
	var testInt16 int16
	var testInt32 int32
	var testInt64 int64
	var testUint uint
	var testUint8 uint8
	var testUint16 uint16
	var testUint32 uint32
	var testUint64 uint64
	var testString string
	var err error

	err = ParsePrimitive("true", &testBool, 2)
	assert.Nil(t, err)
	assert.Equal(t, true, testBool)

	err = ParsePrimitive("3.1415", &testFloat32, 2)
	assert.Nil(t, err)
	assert.Equal(t, float32(3.1415), testFloat32)

	err = ParsePrimitive("3.14159265358979323846", &testFloat64, 2)
	assert.Nil(t, err)
	assert.Equal(t, 3.14159265358979323846, testFloat64)

	err = ParsePrimitive("10", &testInt, 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, testInt)

	err = ParsePrimitive("10", &testInt8, 2)
	assert.Nil(t, err)
	assert.Equal(t, int8(2), testInt8)

	err = ParsePrimitive("10", &testInt16, 2)
	assert.Nil(t, err)
	assert.Equal(t, int16(2), testInt16)

	err = ParsePrimitive("10", &testInt32, 2)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), testInt32)

	err = ParsePrimitive("10", &testInt64, 2)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), testInt64)

	err = ParsePrimitive("10", &testUint, 2)
	assert.Nil(t, err)
	assert.Equal(t, uint(2), testUint)

	err = ParsePrimitive("10", &testUint8, 2)
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), testUint8)

	err = ParsePrimitive("10", &testUint16, 2)
	assert.Nil(t, err)
	assert.Equal(t, uint16(2), testUint16)

	err = ParsePrimitive("10", &testUint32, 2)
	assert.Nil(t, err)
	assert.Equal(t, uint32(2), testUint32)

	err = ParsePrimitive("10", &testUint64, 2)
	assert.Nil(t, err)
	assert.Equal(t, uint64(2), testUint64)

	err = ParsePrimitive("10", &testString, 2)
	assert.Nil(t, err)
	assert.Equal(t, "10", testString)
}

func TestParsePrimitive_TooManyOptionalArgs(t *testing.T) {

	var testInt int
	var err error

	err = ParsePrimitive("10", &testInt, 10, 20)
	assert.NotNil(t, err)
	assert.Equal(t, 0, testInt)
}

func TestParsePrimitive_NotPointer(t *testing.T) {

	var testInt int
	var err error

	err = ParsePrimitive("10", testInt)
	assert.NotNil(t, err)
	assert.Equal(t, 0, testInt)
}

func TestParsePrimitive_NotPrimitive(t *testing.T) {

	type FooBar struct {
		foo string
		bar int
	}
	var testStruct FooBar
	var err error

	err = ParsePrimitive("10", &testStruct)
	assert.NotNil(t, err)
	assert.Equal(t, FooBar{}, testStruct)
}
