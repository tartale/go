package slicez

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFirstAndLast_OnNonEmptySlice(t *testing.T) {
	s := []int{1, 2, 3}

	first := GetFirst(s)
	last := GetLast(s)

	if assert.NotNil(t, first) {
		assert.Equal(t, 1, *first)
	}
	if assert.NotNil(t, last) {
		assert.Equal(t, 3, *last)
	}
}

func TestGetFirstAndLast_OnEmptySlice(t *testing.T) {
	var s []int

	assert.Nil(t, GetFirst(s))
	assert.Nil(t, GetLast(s))
}

func TestMustGetFirstAndLast_OnNonEmptySlice(t *testing.T) {
	s := []string{"a", "b"}

	assert.Equal(t, "a", MustGetFirst(s))
	assert.Equal(t, "b", MustGetLast(s))
}

func TestMustGetFirstAndLast_PanicsOnEmptySlice(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic")
		}
	}()

	var s []int
	_ = MustGetFirst(s)
}

func TestMap(t *testing.T) {
	input := []int{1, 2, 3}
	output := Map(input, func(i int, v int) string {
		return fmt.Sprintf("%d:%d", i, v)
	})

	assert.Equal(t, []string{"0:1", "1:2", "2:3"}, output)
}

func TestFill(t *testing.T) {
	result := Fill("x", 3)

	assert.Equal(t, []string{"x", "x", "x"}, result)
}

