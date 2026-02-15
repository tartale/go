package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTo_ValidStrings(t *testing.T) {
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

	testBool, err = ParseTo[bool]("true")
	assert.Nil(t, err)
	assert.Equal(t, true, testBool)

	testFloat32, err = ParseTo[float32]("3.1415")
	assert.Nil(t, err)
	assert.Equal(t, float32(3.1415), testFloat32)

	testFloat64, err = ParseTo[float64]("3.14159265358979323846")
	assert.Nil(t, err)
	assert.Equal(t, 3.14159265358979323846, testFloat64)

	testInt, err = ParseTo[int]("10")
	assert.Nil(t, err)
	assert.Equal(t, 10, testInt)

	testInt8, err = ParseTo[int8]("10")
	assert.Nil(t, err)
	assert.Equal(t, int8(10), testInt8)

	testInt16, err = ParseTo[int16]("10")
	assert.Nil(t, err)
	assert.Equal(t, int16(10), testInt16)

	testInt32, err = ParseTo[int32]("10")
	assert.Nil(t, err)
	assert.Equal(t, int32(10), testInt32)

	testInt64, err = ParseTo[int64]("10")
	assert.Nil(t, err)
	assert.Equal(t, int64(10), testInt64)

	testUint, err = ParseTo[uint]("10")
	assert.Nil(t, err)
	assert.Equal(t, uint(10), testUint)

	testUint8, err = ParseTo[uint8]("10")
	assert.Nil(t, err)
	assert.Equal(t, uint8(10), testUint8)

	testUint16, err = ParseTo[uint16]("10")
	assert.Nil(t, err)
	assert.Equal(t, uint16(10), testUint16)

	testUint32, err = ParseTo[uint32]("10")
	assert.Nil(t, err)
	assert.Equal(t, uint32(10), testUint32)

	testUint64, err = ParseTo[uint64]("10")
	assert.Nil(t, err)
	assert.Equal(t, uint64(10), testUint64)

	testString, err = ParseTo[string]("10")
	assert.Nil(t, err)
	assert.Equal(t, "10", testString)
}

func TestParseTo_InvalidStrings(t *testing.T) {
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

	testBool, err = ParseTo[bool]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, false, testBool)

	testFloat32, err = ParseTo[float32]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, float32(0.0), testFloat32)

	testFloat64, err = ParseTo[float64]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, float64(0.0), testFloat64)

	testInt, err = ParseTo[int]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, 0, testInt)

	testInt8, err = ParseTo[int8]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, int8(0), testInt8)

	testInt16, err = ParseTo[int16]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, int16(0), testInt16)

	testInt32, err = ParseTo[int32]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, int32(0), testInt32)

	testInt64, err = ParseTo[int64]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, int64(0), testInt64)

	testUint, err = ParseTo[uint]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, uint(0), testUint)

	testUint8, err = ParseTo[uint8]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, uint8(0), testUint8)

	testUint16, err = ParseTo[uint16]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, uint16(0), testUint16)

	testUint32, err = ParseTo[uint32]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, uint32(0), testUint32)

	testUint64, err = ParseTo[uint64]("foobar")
	assert.NotNil(t, err)
	assert.Equal(t, uint64(0), testUint64)
}

func TestParseTo_NonBase10(t *testing.T) {
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

	testBool, err = ParseTo[bool]("true", 2)
	assert.Nil(t, err)
	assert.Equal(t, true, testBool)

	testFloat32, err = ParseTo[float32]("3.1415", 2)
	assert.Nil(t, err)
	assert.Equal(t, float32(3.1415), testFloat32)

	testFloat64, err = ParseTo[float64]("3.14159265358979323846", 2)
	assert.Nil(t, err)
	assert.Equal(t, 3.14159265358979323846, testFloat64)

	testInt, err = ParseTo[int]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, 2, testInt)

	testInt8, err = ParseTo[int8]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, int8(2), testInt8)

	testInt16, err = ParseTo[int16]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, int16(2), testInt16)

	testInt32, err = ParseTo[int32]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, int32(2), testInt32)

	testInt64, err = ParseTo[int64]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), testInt64)

	testUint, err = ParseTo[uint]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, uint(2), testUint)

	testUint8, err = ParseTo[uint8]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, uint8(2), testUint8)

	testUint16, err = ParseTo[uint16]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, uint16(2), testUint16)

	testUint32, err = ParseTo[uint32]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, uint32(2), testUint32)

	testUint64, err = ParseTo[uint64]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, uint64(2), testUint64)

	testString, err = ParseTo[string]("10", 2)
	assert.Nil(t, err)
	assert.Equal(t, "10", testString)
}

func TestParseTo_TooManyOptionalArgs(t *testing.T) {

	var testInt int
	var err error

	testInt, err = ParseTo[int]("10", 10, 20)
	assert.NotNil(t, err)
	assert.Equal(t, 0, testInt)
}
