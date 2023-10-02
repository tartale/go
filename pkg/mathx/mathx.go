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

func DivideTo[IN1, IN2, OUT constraintz.Number](a IN1, b IN2, ops ...func(float64) float64) OUT {
	dividend := float64(a)
	divisor := float64(b)
	result := dividend / divisor
	for _, op := range ops {
		result = op(result)
	}

	return OUT(result)
}

func Divide[IN1, IN2, OUT constraintz.Number](a IN1, b IN2, result *OUT, ops ...func(float64) float64) {

	*result = DivideTo[IN1, IN2, OUT](a, b, ops...)
}
