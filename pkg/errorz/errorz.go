package errorz

import "errors"

// ErrFatal represents a non-recoverable error condition.
//
// Example:
//
//	if err != nil {
//		return fmt.Errorf("%w: %w", errorz.ErrFatal, err)
//	}
var ErrFatal = errors.New("fatal error")

// ErrNotFound represents a resource-not-found condition.
var ErrNotFound = errors.New("not found")

// ErrInvalidType is returned when a value is of an unexpected type.
var ErrInvalidType = errors.New("invalid type")

// ErrInvalidArgument is returned when a function argument is invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// ErrBadRequest represents a malformed or invalid request.
var ErrBadRequest = errors.New("bad request")

// ErrResponse represents an error reported by an external HTTP response.
var ErrResponse = errors.New("error in response")
