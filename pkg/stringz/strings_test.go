package stringz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToAlphaNumeric(t *testing.T) {

	assert.Equal(t, "abcDEF12345", ToAlphaNumeric("a~b`c!D@EF$%^123^&&*()45"))
}
