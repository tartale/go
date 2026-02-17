package jsonx

import (
	"bytes"
	"encoding/json"
)

// StrictUnmarshal unmarshals JSON into v and fails if unknown fields are present.
func StrictUnmarshal(data []byte, v any) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	return dec.Decode(v)
}
