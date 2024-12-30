package services

import (
	"errors"
	"fmt"
	"github.com/Mir00r/auth-service/configs"
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/Mir00r/auth-service/internal/models/entities"
	"time"

	services "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"log"
)

// TokenServiceInterface defines the methods for the TokenService
type TokenServiceInterface interface {
	InitiatePasswordReset(req services.PasswordResetRequest) error
	ResetPassword(req services.ConfirmPasswordResetRequest) error
	Logout(tokenString string, userID string) error
	RefreshToken(req dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse, error)
}

// TokenService is the concrete implementation of TokenServiceInterface
type TokenService struct {
	TokenRepo repositories.TokenRepository
	UserRepo  repositories.UserRepository
}

// NewTokenService initializes a new instance of TokenService
func NewTokenService(repo repositories.TokenRepository, userRepo repositories.UserRepository) TokenServiceInterface {
	return &TokenService{TokenRepo: repo, UserRepo: userRepo}
}

func (svc *TokenService) InitiatePasswordReset(req services.PasswordResetRequest) error {
	// Check if the user exists
	user, err := svc.UserRepo.FindUserByEmail(req.Email)
	//log.Printf("After query the user from DB user: %v\n and error: %v\n", user, err)
	if err != nil {
		return err
	}
	if user == nil {
		return constants.ErrUserNotFoundVar
	}

	// Generate a reset token
	resetToken, err := utils.GeneratePasswordResetToken(user.ID)
	log.Printf("After Generate a reset token: %v\n and the error is: %v\n", resetToken, err)
	if err != nil {
		return err
	}

	// Save the token in the database (if you have a table for reset tokens)
	err = svc.TokenRepo.SaveResetToken(resetToken, user.ID)
	log.Printf("After save the token: %v\n", err)
	if err != nil {
		return err
	}

	// Send the reset email
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.AppConfig.Password.PasswordResetURL, resetToken)
	err = NewEmailService().SendPasswordResetEmail(user.Email, resetLink)
	if err != nil {
		return err
	}

	return nil
}

func (svc *TokenService) ResetPassword(req services.ConfirmPasswordResetRequest) error {
	// Verify the reset token
	userID, err := utils.VerifyPasswordResetToken(req.Token)
	log.Printf("Verify Password Reset token: %v\n", err)
	if err != nil {
		return errors.New("invalid or expired reset token")
	}

	// Check if the token has already been used
	resetToken, err := svc.TokenRepo.FindToken(req.Token)
	if err != nil || resetToken.Used {
		return errors.New("reset token already used or invalid")
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// Update the user's password
	err = svc.UserRepo.UpdatePassword(userID, hashedPassword)
	if err != nil {
		return err
	}

	// Mark the reset token as used
	err = svc.TokenRepo.MarkTokenAsUsed(req.Token)
	if err != nil {
		return err
	}

	return nil
}

// Logout invalidates the current token (via blacklisting or other mechanisms)
func (svc *TokenService) Logout(tokenString string, userID string) error {
	// Optionally check if the token is already blacklisted
	isBlacklisted, err := svc.TokenRepo.IsTokenBlacklisted(tokenString)
	if err != nil {
		return err
	}
	if isBlacklisted {
		return constants.ErrTokenInvalidatedVar
	}

	// Blacklist the token
	token := &entities.Token{
		Token:     tokenString,
		UserID:    userID,
		Type:      "access",                  // Assuming token type is "access"
		ExpiresAt: time.Now().Add(time.Hour), // Set expiration for blacklisted entry
	}
	return svc.TokenRepo.BlacklistToken(token)
}

// RefreshToken generates a new access token using a valid refresh token
func (svc *TokenService) RefreshToken(req dtos.RefreshTokenRequest) (*dtos.RefreshTokenResponse, error) {
	// Validate the refresh token
	token, err := svc.TokenRepo.FindRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}
	if token == nil || time.Now().After(token.ExpiresAt) {
		return nil, constants.ErrInvalidOrExpiredRefreshTokenVar
	}

	// Find the user associated with the refresh token
	user, err := svc.UserRepo.FindUserByID(token.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, constants.ErrUserNotFoundVar
	}

	// Generate a new access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, config.AppConfig.JWT.Secret, utils.TokenExpiry())
	if err != nil {
		return nil, err
	}

	// Optionally: Rotate the refresh token (generate a new one)
	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, constants.ErrGenerateTokenVar
	}

	// Update the existing token entry in the database
	token.Token = accessToken
	token.RefreshToken = newRefreshToken
	token.Type = constants.RefreshToken
	token.ExpiresAt = time.Now().Add(utils.TokenExpiry())
	token.UpdatedAt = time.Now()
	token.RefreshTokenExpiresAt = time.Now().Add(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry))
	if err := svc.TokenRepo.UpdateToken(token); err != nil {
		return nil, err
	}

	// Save the new refresh token in the database
	//err = svc.TokenRepo.CreateToken(&entities.Token{
	//	UserID:                user.ID,
	//	Token:                 accessToken,
	//	RefreshToken:          newRefreshToken,
	//	Type:                  constants.RefreshToken,
	//	ExpiresAt:             time.Now().Add(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry)),
	//	RefreshTokenExpiresAt: time.Now().Add(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry)),
	//})
	//if err != nil {
	//	return nil, constants.ErrSaveTokenVar
	//}

	return &dtos.RefreshTokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          newRefreshToken,
		ExpiresIn:             utils.TokenExpiry().Microseconds(),                                                 // 2 hour
		RefreshTokenExpiresIn: int64(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry).Seconds()), // 24 hour
	}, nil
}
