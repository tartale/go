package primitives

import (
	"github.com/tartale/go/pkg/constraintz"
)

// Ref returns a pointer to the given primitive literal.
// Basically, it is for a situation like this:
//
//	var foo *int
//	foo = &1               // illegal
//	foo = primitives.Ref(1) // legal
func Ref[T constraintz.Primitive](val T) *T {
	return &val
}
