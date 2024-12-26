package services

import (
	"github.com/Mir00r/auth-service/configs"
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/entities"
	services "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
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
		return nil, constants.ErrInvalidCredentialVar
	}

	if !utils.VerifyPassword(user.Password, req.Password) {
		return nil, constants.ErrInvalidCredentialVar
	}

	//config.LoadConfig()
	token, err := utils.GenerateJWT(user.ID, user.Email, config.AppConfig.JWT.Secret, config.TokenExpiry())
	if err != nil {
		return nil, constants.ErrGenerateTokenVar
	}

	return map[string]string{"access_token": token}, nil
}

// RegisterUser creates a new user account
func (svc *AuthService) RegisterUser(req services.RegisterRequest) error {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return constants.ErrHashPasswordVar
	}

	// Create a new user instance
	newUser := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return svc.UserRepo.CreateUser(newUser)
}

func (svc *AuthService) GetUserProfile(userID string) (*entities.User, error) {
	user, err := svc.UserRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, constants.ErrUserNotFoundVar
	}
	return user, nil
}
