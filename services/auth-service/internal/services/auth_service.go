package services

import (
	config "github.com/Mir00r/auth-service/configs"
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/Mir00r/auth-service/internal/models/entities"
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"time"
)

// AuthService encapsulates the business logic for authentication
type AuthService struct {
	UserRepo  *repositories.UserRepository  // Repository for interacting with user data
	TokenRepo *repositories.TokenRepository // Repository for interacting with token data
}

// NewAuthService initializes a new AuthService instance
func NewAuthService(
	repo *repositories.UserRepository,
	tokenRepo *repositories.TokenRepository,
) *AuthService {
	return &AuthService{
		UserRepo:  repo,
		TokenRepo: tokenRepo,
	}
}

// Authenticate validates user credentials and generates a JWT token
//
// This function performs the following steps:
// 1. Validates the provided email and password against the database.
// 2. Generates a JWT token for the authenticated user.
// 3. Returns the generated token or an error if authentication fails.
//
// Parameters:
// - req: LoginRequest containing email and password.
//
// Returns:
// - A map containing the access token.
// - An error if authentication fails.
func (svc *AuthService) Authenticate(req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	// Retrieve the user from the database by email
	user, err := svc.UserRepo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, constants.ErrInvalidCredentials
	}

	// Verify the provided password against the hashed password
	if !utils.VerifyPassword(user.Password, req.Password) {
		return nil, constants.ErrInvalidCredentials
	}

	// Generate a JWT token for the authenticated user
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, config.AppConfig.JWT.Secret, utils.TokenExpiry())
	if err != nil {
		return nil, constants.ErrGenerateTokenVar
	}

	// Generate a refresh token for the user
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, constants.ErrGenerateTokenVar
	}

	// Save the refresh token in the database
	err = svc.TokenRepo.CreateToken(&entities.Token{
		UserID:                user.ID,
		Token:                 refreshToken,
		RefreshToken:          refreshToken,
		Type:                  constants.AccessToken,
		ExpiresAt:             time.Now().Add(utils.TokenExpiry()),
		RefreshTokenExpiresAt: time.Now().Add(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry)),
	})
	if err != nil {
		return nil, constants.ErrSaveTokenVar
	}

	// Return the LoginResponse
	return &dtos.LoginResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		ExpiresIn:             int64(utils.TokenExpiry().Seconds()),
		RefreshTokenExpiresIn: int64(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry).Seconds()),
	}, nil
}

// RegisterUser creates a new user account in the system
//
// This function performs the following steps:
// 1. Hashes the provided password.
// 2. Creates a new user instance and saves it in the database.
//
// Parameters:
// - req: RegisterRequest containing user registration details (name, email, password).
//
// Returns:
// - An error if registration fails.
func (svc *AuthService) RegisterUser(req request.RegisterRequest) error {
	// Hash the user's password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return constants.ErrHashPasswordVar
	}

	// Create a new user entity
	newUser := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Save the new user in the database
	if err := svc.UserRepo.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

// GetUserProfile retrieves the profile of a user by their ID
//
// This function performs the following steps:
// 1. Looks up the user in the database by their ID.
// 2. Returns the user's profile or an error if the user does not exist.
//
// Parameters:
// - userID: The unique identifier of the user.
//
// Returns:
// - A pointer to the User entity.
// - An error if the user does not exist or retrieval fails.
func (svc *AuthService) GetUserProfile(userID string) (*entities.User, error) {
	// Retrieve the user by their ID
	user, err := svc.UserRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}

	// Check if the user exists
	if user == nil {
		return nil, constants.ErrUserNotFoundVar
	}

	return user, nil
}
