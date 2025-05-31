package helpers

import "errors"

var (
	ErrNotFound           = errors.New("not found error")
	ErrInternalServer     = errors.New("internal server error")
	ErrValidation         = errors.New("validation error")
	ErrUnauthorized       = errors.New("unauthorized error")
	ErrEmailAlreadyExists = errors.New("email already exists")
)
