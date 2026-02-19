package jsonx

import "encoding/json"

// MustMarshalJSON takes an object of any type,
// marshals it to JSON format, and returns the raw
// bytes, panicking if an error occurs.
func MustMarshalJSON[T any](t T) []byte {
	bytes, err := json.Marshal(&t)
	if err != nil {
		panic(err)
	}

	return bytes
}

// MustMarshalJSONToString is a convenience function
// for MustMarshalJSON that casts the result to a string.
func MustMarshalJSONToString[T any](t T) string {
	return string(MustMarshalJSON(t))
}
