package mathx

import (
	"math"

	"github.com/tartale/go/pkg/constraintz"
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

func Floor[N constraintz.Number](a N) int {
	return int(math.Floor(float64(a)))
}

func Ceil[N constraintz.Number](a N) int {
	return int(math.Ceil(float64(a)))
}

func DivideAndRound[N constraints.Integer](a, b N) N {
	dividend := float64(a)
	divisor := float64(b)

	return N(math.Round(dividend / divisor))
}
