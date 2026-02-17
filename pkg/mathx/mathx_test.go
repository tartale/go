package mathx

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	assert.Equal(t, 10, Abs(-10))
	assert.Equal(t, 10.0, Abs(-10.0))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 10, Max(-10, 10))
	assert.Equal(t, 10.0, Max(-10.0, 10.0))
}

func TestMin(t *testing.T) {
	assert.Equal(t, -10, Min(-10, 10))
	assert.Equal(t, -10.0, Min(-10.0, 10.0))
}

func TestDivide(t *testing.T) {
	var intResult int
	Divide(10, 3, &intResult)
	assert.Equal(t, 3, intResult)

	var float32Result float32
	Divide(10, 3, &float32Result)
	assert.InDelta(t, float32(3.33), float32Result, 0.01)
}

func TestDivideWithOps(t *testing.T) {
	var intResult int
	Divide(10, 3, &intResult, math.Ceil)
	assert.Equal(t, 4, intResult)
}
