package mathx

import (
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
