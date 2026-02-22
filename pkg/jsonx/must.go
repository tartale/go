package jsonx

import "encoding/json"

// MustMarshal wraps json.Marshal but panics if there's an error.
func MustMarshal[T any](t T) []byte {
	bytes, err := json.Marshal(&t)
	if err != nil {
		panic(err)
	}

	return bytes
}

// MustMarshalToString is a convenience function
// for MustMarshal that casts the result to a string.
func MustMarshalToString[T any](t T) string {
	return string(MustMarshal(t))
}

// MustUarshal wraps json.Unmarshal but panics if there's an error.
func MustUnmarshal[T any](data []byte, t *T) {
	err := json.Unmarshal(data, t)
	if err != nil {
		panic(err)
	}
}

// MustUnmarshalFromString is a convenience function
// for MustUnmarshal that casts the result to a string.
func MustUnmarshalFromString[T any](data string, t *T) {
	MustUnmarshal([]byte(data), t)
}
