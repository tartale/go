package maps

import (
	"github.com/tartale/go/pkg/primitives"
	"github.com/tartale/go/pkg/reflectx"
)

// CastPrimitives takes an object and returns a map
// of all the fields of the object as KV pairs.
// Any field that is a primitive is cast down
// to its underlying type, to ensure that type aliases
// are considered equivalent to their underlying types.
func CastPrimitives[K comparable, V any](input map[K]V) map[K]V {
	for k, v := range input {
		if reflectx.IsPrimitive(v) {
			input[k] = primitives.MustCastAway(v).(V)
		}
	}
	return input
}
