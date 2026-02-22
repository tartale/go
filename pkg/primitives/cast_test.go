package primitives

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyString string

func TestCastAway(t *testing.T) {
	var myString MyString = "hello"
	myCast, err := CastAway(myString)

	assert.NoError(t, err)
	assert.NotEqual(t, myString, myCast)
	assert.Equal(t, "hello", myCast)
}
