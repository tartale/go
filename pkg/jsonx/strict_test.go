package jsonx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrictUnmarshal_RejectsUnknownFields(t *testing.T) {
	type Target struct {
		Name string `json:"name"`
	}

	// "extra" is not a known field and should cause StrictUnmarshal to fail.
	jsonWithExtra := []byte(`{"name":"alice","extra":"field"}`)

	var tgt Target
	err := StrictUnmarshal(jsonWithExtra, &tgt)

	assert.Error(t, err)
}

func TestStrictUnmarshal_AcceptsKnownFieldsOnly(t *testing.T) {
	type Target struct {
		Name string `json:"name"`
	}

	validJSON := []byte(`{"name":"alice"}`)

	var tgt Target
	err := StrictUnmarshal(validJSON, &tgt)

	if assert.NoError(t, err) {
		assert.Equal(t, "alice", tgt.Name)
	}
}

