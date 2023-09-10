package errorz

import "errors"

var ErrNotFound = errors.New("not found")
var ErrInvalidType = errors.New("invalid type")
var ErrInvalidArgument = errors.New("invalid argument")
