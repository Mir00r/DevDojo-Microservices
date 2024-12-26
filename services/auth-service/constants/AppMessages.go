package constants

import "errors"

// Error message strings
const (
	ErrMissingAuthHeader   = "Authorization header is missing"
	ErrInvalidAuthHeader   = "Invalid Authorization header"
	ErrInvalidToken        = "Invalid token"
	ErrUserNotFound        = "User not found"
	InvalidCredentials     = "Invalid email or password"
	Unauthorized           = "Unauthorized"
	Forbidden              = "Forbidden"
	ResourceNotFound       = "Resource not found"
	InternalServerError    = "Internal server error"
	ErrHashPassword        = "Failed to hash password"
	ErrGenerateToken       = "Failed to generate token"
	ErrInvalidCredential   = "Invalid credentials"
	ErrTokenInvalidOrReset = "Reset token already used or invalid"
	ErrTokenInvalidated    = "Token is already invalidated"
)

// Error variables for use throughout the project
var (
	ErrInvalidCredentials     = errors.New(InvalidCredentials)
	ErrUnauthorized           = errors.New(Unauthorized)
	ErrForbidden              = errors.New(Forbidden)
	ErrNotFound               = errors.New(ResourceNotFound)
	ErrInternalServer         = errors.New(InternalServerError)
	ErrMissingAuthHeaderVar   = errors.New(ErrMissingAuthHeader)
	ErrInvalidAuthHeaderVar   = errors.New(ErrInvalidAuthHeader)
	ErrInvalidTokenVar        = errors.New(ErrInvalidToken)
	ErrUserNotFoundVar        = errors.New(ErrUserNotFound)
	ErrHashPasswordVar        = errors.New(ErrHashPassword)
	ErrGenerateTokenVar       = errors.New(ErrGenerateToken)
	ErrInvalidCredentialVar   = errors.New(ErrInvalidCredential)
	ErrTokenInvalidOrResetVar = errors.New(ErrTokenInvalidOrReset)
	ErrTokenInvalidatedVar    = errors.New(ErrTokenInvalidated)
)

// Success messages
const (
	MsgLogoutSuccessful = "Logout successful"
)

// Api Header
const (
	Authorization = "Authorization"
	Bearer        = "Bearer "
)
