package utils

import "errors"

// Error messages
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrForbidden          = errors.New("forbidden")
	ErrNotFound           = errors.New("resource not found")
	ErrInternalServer     = errors.New("internal server error")
)

// WrapError allows adding context to an error message.
func WrapError(err error, context string) error {
	return errors.New(context + ": " + err.Error())
}

// IsError checks if an error is a specific error type.
func IsError(err error, target error) bool {
	return errors.Is(err, target)
}
