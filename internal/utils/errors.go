package utils

import "errors"

var (
	ErrNoRows             = errors.New("no rows in result set")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrMissingSecret      = errors.New("missing JWT secret key")
)
