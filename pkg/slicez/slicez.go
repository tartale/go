package slicez

import (
	"fmt"

	"github.com/tartale/go/pkg/errorz"
)

func GetFirst[T any](slc []T) *T {

	if len(slc) == 0 {
		return nil
	}
	return &slc[0]
}

func MustGetFirst[T any](slc []T) T {

	if len(slc) == 0 {
		panic(fmt.Errorf("%w: attempt to get first element of empty slice", errorz.ErrFatal))
	}
	return slc[0]
}

func GetLast[T any](slc []T) *T {

	if len(slc) == 0 {
		return nil
	}
	return &slc[len(slc)-1]
}

func MustGetLast[T any](slc []T) T {

	if len(slc) == 0 {
		panic(fmt.Errorf("%w: attempt to get first element of empty slice", errorz.ErrFatal))
	}
	return slc[len(slc)-1]
}

func Map[In, Out any](inputSlice []In, fn func(index int, inputValue In) Out) (outputSlice []Out) {

	for index, inputValue := range inputSlice {
		outputSlice = append(outputSlice, fn(index, inputValue))
	}

	return outputSlice
}

func Fill[T any](t T, n int) []T {

	result := make([]T, n)
	for i := 0; i < n; i++ {
		result[i] = t
	}

	return result
}
