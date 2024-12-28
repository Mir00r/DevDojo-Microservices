package services

import (
	"errors"
	services "github.com/Mir00r/auth-service/internal/models/response"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"time"
)

type InternalAuthService struct {
	UserRepo *repositories.UserRepository
}

func NewInternalAuthService(repo *repositories.UserRepository) *InternalAuthService {
	return &InternalAuthService{UserRepo: repo}
}

// ValidateToken validates a JWT token
func (svc *InternalAuthService) ValidateToken(token string) (*services.ValidateTokenResponse, error) {
	claims, err := utils.VerifyJWT(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return &services.ValidateTokenResponse{
		IsValid: true,
		UserID:  claims.UserID,
		Expires: time.Unix(claims.ExpiresAt.Unix(), 0).Format(time.RFC3339),
		Message: "Token is valid",
	}, nil
}

// CheckHealth returns the health status of the service
func (svc *InternalAuthService) CheckHealth() map[string]string {
	return map[string]string{
		"status":  "healthy",
		"uptime":  time.Now().Format(time.RFC3339),
		"version": "1.0.0",
	}
}
