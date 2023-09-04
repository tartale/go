package mathx

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed | constraints.Float](a T) T {
	if a >= 0 {
		return a
	}

	return -(a)
}

func Max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
}
