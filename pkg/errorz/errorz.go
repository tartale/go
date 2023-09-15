package errorz

import "errors"

var ErrFatal = errors.New("fatal error")
var ErrNotFound = errors.New("not found")
var ErrInvalidType = errors.New("invalid type")
var ErrInvalidArgument = errors.New("invalid argument")
var ErrBadRequest = errors.New("bad request")
var ErrResponse = errors.New("error in response")
