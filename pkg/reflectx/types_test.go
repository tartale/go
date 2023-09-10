package reflectx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntType int
type Int64Type int64
type Float32Type float32

func TestIsNumber(t *testing.T) {

	assert.True(t, IsNumber(1))
	assert.True(t, IsNumber(1.0))
	assert.True(t, IsNumber(IntType(1)))

}
