package helpers

import (
	"errors"
)

var (
	ErrClientError        = errors.New("client error")
	ErrNotFound           = errors.New("not found error")
	ErrInternalServer     = errors.New("internal server error")
	ErrValidation         = errors.New("validation error")
	ErrUnauthorized       = errors.New("unauthorized error")
	ErrUnprocessable      = errors.New("unprocessable entity error")
	ErrEmailAlreadyExists = errors.New("email already exists error")
)
