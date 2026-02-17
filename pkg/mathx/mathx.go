package mathx

import (
	"math"

	"github.com/tartale/go/pkg/constraintz"
	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value of a.
//
// Example:
//
//	value := mathx.Abs(-10) // value == 10
func Abs[T constraints.Signed | constraints.Float](a T) T {
	if a >= 0 {
		return a
	}

	return -(a)
}

// Max returns the larger of a and b.
func Max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

// Min returns the smaller of a and b.
func Min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// Floor converts a to float64, applies math.Floor and returns the result as int.
func Floor[N constraintz.Number](a N) int {
	return int(math.Floor(float64(a)))
}

// Ceil converts a to float64, applies math.Ceil and returns the result as int.
func Ceil[N constraintz.Number](a N) int {
	return int(math.Ceil(float64(a)))
}

// DivideTo divides a by b, applies any ops to the intermediate float64 result,
// and returns the final value as type OUT.
//
// Example:
//
//	rounded := mathx.DivideTo[int, int, int](5, 2, math.Round) // rounded == 2
func DivideTo[IN1, IN2, OUT constraintz.Number](a IN1, b IN2, ops ...func(float64) float64) OUT {
	dividend := float64(a)
	divisor := float64(b)
	result := dividend / divisor
	for _, op := range ops {
		result = op(result)
	}

	return OUT(result)
}

// Divide divides a by b, applies any ops, and stores the result in result. By
// passing the result as a pointer to be populated, this function can be called
// and allows the types to be inferred.
func Divide[IN1, IN2, OUT constraintz.Number](a IN1, b IN2, result *OUT, ops ...func(float64) float64) {
	*result = DivideTo[IN1, IN2, OUT](a, b, ops...)
}
