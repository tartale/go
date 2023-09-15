package mathx

import (
	"math"

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

func DivideAndRound[N constraints.Integer](a, b N) N {
	dividend := float64(a)
	divisor := float64(b)

	return N(math.Round(dividend / divisor))
}
