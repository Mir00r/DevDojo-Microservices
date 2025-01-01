package errors

import (
	"fmt"
	"net/http"
)

// Specific error types
var (
	ErrInvalidCredentials            = NewAppError(http.StatusUnauthorized, "Invalid credentials", nil)
	ErrGenerateToken                 = NewAppError(http.StatusInternalServerError, "Failed to generate token", nil)
	ErrSaveToken                     = NewAppError(http.StatusInternalServerError, "Failed to save token", nil)
	ErrInvalidPayload                = NewAppError(http.StatusBadRequest, "Invalid request payload", nil)
	ErrFailedToRegisterUser          = NewAppError(http.StatusInternalServerError, "Failed to register user", nil)
	ErrUserNotFound                  = NewAppError(http.StatusNotFound, "User not found", nil)
	ErrHashPassword                  = NewAppError(http.StatusInternalServerError, "Failed to hash password", nil)
	ErrFailedToInitiatePasswordReset = NewAppError(http.StatusInternalServerError, "Failed to initiate password reset", nil)
	ErrInvalidOrExpiredResetToken    = NewAppError(http.StatusBadRequest, "Invalid or expired reset token", nil)
	ErrResetTokenAlreadyUsed         = NewAppError(http.StatusBadRequest, "Reset token already used", nil)
	ErrFailedToUpdatePassword        = NewAppError(http.StatusInternalServerError, "Failed to update password", nil)
	ErrInvalidOrExpiredRefreshToken  = NewAppError(http.StatusUnauthorized, "Invalid or expired refresh token", nil)
)

// AppError represents a generic application error
type AppError struct {
	Code       int    // HTTP status code
	CodeStatus string // HTTP status code message
	Message    string // Error message to be returned to the client
	Err        error  // Underlying error (optional)
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:       code,
		CodeStatus: http.StatusText(code),
		Message:    message,
		Err:        err,
	}
}
