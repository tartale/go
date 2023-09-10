package constraintz

import "golang.org/x/exp/constraints"

type Primitive interface {
	constraints.Float | constraints.Integer | string | bool
}

type Number interface {
	constraints.Float | constraints.Integer | constraints.Complex
}
