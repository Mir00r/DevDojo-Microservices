package services

import (
	"errors"
	services "github.com/Mir00r/auth-service/internal/models/response"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"time"
)

type InternalAuthService interface {
	ValidateToken(token string) (*services.ValidateTokenResponse, error)
	CheckHealth() map[string]string
}

// InternalAuthService handles internal authentication-related operations.
type internalAuthService struct {
	UserRepo repositories.UserRepository // Repository for interacting with the User data
}

// NewInternalAuthService creates a new instance of InternalAuthService with the required dependencies.
// This uses Dependency Injection to ensure testability and modularity.
func NewInternalAuthService(userRepo repositories.UserRepository) InternalAuthService {
	return &internalAuthService{
		UserRepo: userRepo,
	}
}

// ValidateToken verifies the provided JWT token and returns its validity and associated user information.
// Parameters:
// - token: The JWT token string to be validated.
// Returns:
// - A pointer to ValidateTokenResponse containing the validation result, user ID, and expiration details.
// - An error if the token is invalid or any other issue occurs during validation.
func (svc *internalAuthService) ValidateToken(token string) (*services.ValidateTokenResponse, error) {
	// Verify the JWT token and extract claims
	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Build and return the validation response
	return &services.ValidateTokenResponse{
		IsValid: true,
		UserID:  claims.UserID,
		Expires: time.Unix(claims.ExpiresAt.Unix(), 0).Format(time.RFC3339), // Convert expiration to RFC3339 format
		Message: "Token is valid",
	}, nil
}

// CheckHealth provides the health status of the authentication service.
// Returns:
// - A map containing the service health status, uptime, and version.
func (svc *internalAuthService) CheckHealth() map[string]string {
	return map[string]string{
		"status":  "healthy",                       // Indicates the service is operational
		"uptime":  time.Now().Format(time.RFC3339), // Current server time in RFC3339 format
		"version": "1.0.0",                         // Service version
	}
}
