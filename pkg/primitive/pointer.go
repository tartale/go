package primitive

// Ref returns a pointer to the given primitive literal.
// Basically, it is for a situation like this:
//
//	var foo *int
//	foo = &1               // illegal
//	foo = primitive.Ref(1) // legal
func Ref[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](val T) *T {
	return &val
}
