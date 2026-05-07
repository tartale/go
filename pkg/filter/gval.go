package filter

import "github.com/PaesslerAG/gval"

// MustEvaluate is a wrapper around the gval.Evaluate function
// that panics if an error is returned.
func MustEvaluate(expression string, parameter any, opts ...gval.Language) any {
	result, err := gval.Evaluate(expression, parameter, opts...)
	if err != nil {
		panic(err)
	}
	return result
}
