package core_errors

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrConflict = errors.New("conflict")
)