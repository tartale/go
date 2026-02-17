// Package slicez provides helpers for working with slices.
package slicez

import (
	"fmt"

	"github.com/tartale/go/pkg/errorz"
)

// GetFirst returns a pointer to the first element of slc, or nil if slc is empty.
//
// Example:
//
//	vals := []int{1, 2, 3}
//	first := slicez.GetFirst(vals) // *int pointing to 1
func GetFirst[T any](slc []T) *T {
	if len(slc) == 0 {
		return nil
	}
	return &slc[0]
}

// MustGetFirst returns the first element of slc, panicking if slc is empty.
func MustGetFirst[T any](slc []T) T {
	if len(slc) == 0 {
		panic(fmt.Errorf("%w: attempt to get first element of empty slice", errorz.ErrFatal))
	}
	return slc[0]
}

// GetLast returns a pointer to the last element of slc, or nil if slc is empty.
func GetLast[T any](slc []T) *T {
	if len(slc) == 0 {
		return nil
	}
	return &slc[len(slc)-1]
}

// MustGetLast returns the last element of slc, panicking if slc is empty.
func MustGetLast[T any](slc []T) T {
	if len(slc) == 0 {
		panic(fmt.Errorf("%w: attempt to get first element of empty slice", errorz.ErrFatal))
	}
	return slc[len(slc)-1]
}

// Map applies fn to each element of inputSlice and returns the resulting slice.
//
// Example:
//
//	ints := []int{1, 2}
//	strs := slicez.Map(ints, func(i, v int) string {
//		return fmt.Sprintf("%d:%d", i, v)
//	})
func Map[In, Out any](inputSlice []In, fn func(index int, inputValue In) Out) (outputSlice []Out) {
	for index, inputValue := range inputSlice {
		outputSlice = append(outputSlice, fn(index, inputValue))
	}

	return outputSlice
}

// Fill returns a slice of length n where each element is t.
func Fill[T any](t T, n int) []T {
	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = t
	}

	return result
}
