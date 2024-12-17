package services

import (
	"errors"
	config "github.com/Mir00r/auth-service/internal/configs"
	"github.com/Mir00r/auth-service/internal/models/entities"
	services "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"time"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: repo}
}

// Authenticate validates user credentials and returns JWT token
func (svc *AuthService) Authenticate(req services.LoginRequest) (map[string]string, error) {
	user, err := svc.UserRepo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.VerifyPassword(user.Password, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	config.LoadConfig()
	token, err := utils.GenerateJWT(user.ID, user.Email, config.GetConfig().JWTSecret, time.Duration(config.GetConfig().TokenExpiry))
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return map[string]string{"access_token": token}, nil
}

// RegisterUser creates a new user account
func (svc *AuthService) RegisterUser(req services.RegisterRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Create a new user instance
	newUser := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return svc.UserRepo.CreateUser(newUser)
}

var ErrUserNotFound = errors.New("user not found")
