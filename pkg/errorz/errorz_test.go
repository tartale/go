package errorz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorSentinels(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected string
	}{
		{"fatal", ErrFatal, "fatal error"},
		{"not found", ErrNotFound, "not found"},
		{"invalid type", ErrInvalidType, "invalid type"},
		{"invalid argument", ErrInvalidArgument, "invalid argument"},
		{"bad request", ErrBadRequest, "bad request"},
		{"response", ErrResponse, "error in response"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.err.Error())
		})
	}
}

