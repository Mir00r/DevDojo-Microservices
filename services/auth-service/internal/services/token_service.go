package services

import (
	"fmt"
	"github.com/Mir00r/auth-service/configs"
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/errors"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/Mir00r/auth-service/internal/models/entities"
	"net/http"
	"time"

	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
)

// TokenServiceInterface defines the methods for the TokenService
type TokenServiceInterface interface {
	InitiatePasswordReset(req dtos.PasswordResetRequest) error
	ResetPassword(req dtos.ConfirmPasswordResetRequest) error
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

func (svc *TokenService) InitiatePasswordReset(req dtos.PasswordResetRequest) error {
	// Check if the user exists
	user, err := svc.UserRepo.FindUserByEmail(req.Email)
	if err != nil {
		return errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToRetrieveUser, err)
	}
	if user == nil {
		return errors.ErrUserNotFound // Defined as a domain-specific error
	}

	// Generate a reset token
	resetToken, err := utils.GeneratePasswordResetToken(user.ID)
	if err != nil {
		return errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToGenerateResetToken, err)
	}

	// Save the token in the database
	err = svc.TokenRepo.SaveResetToken(resetToken, user.ID)
	if err != nil {
		return errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToSaveResetToken, err)
	}

	// Send the reset email
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.AppConfig.Password.PasswordResetURL, resetToken)
	err = NewEmailService().SendPasswordResetEmail(user.Email, resetLink)
	if err != nil {
		return errors.NewAppError(http.StatusInternalServerError, "Failed to send password reset email", err)
	}

	return nil
}

func (svc *TokenService) ResetPassword(req dtos.ConfirmPasswordResetRequest) error {
	// Verify the reset token
	userID, err := utils.VerifyPasswordResetToken(req.Token)
	if err != nil {
		return errors.ErrInvalidOrExpiredResetToken
	}

	// Check if the token has already been used
	resetToken, err := svc.TokenRepo.FindToken(req.Token)
	if err != nil || resetToken.Used {
		return errors.ErrResetTokenAlreadyUsed
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.ErrHashPassword
	}

	// Update the user's password
	err = svc.UserRepo.UpdatePassword(userID, hashedPassword)
	if err != nil {
		return errors.ErrFailedToUpdatePassword
	}

	// Mark the reset token as used
	err = svc.TokenRepo.MarkTokenAsUsed(req.Token)
	if err != nil {
		return errors.ErrFailedToUpdatePassword
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
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToFindToken, err)
	}
	if token == nil || time.Now().After(token.RefreshTokenExpiresAt) {
		return nil, errors.ErrInvalidOrExpiredRefreshToken // Domain-specific error
	}

	// Find the user associated with the refresh token
	user, err := svc.UserRepo.FindUserByID(token.UserID)
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToRetrieveUser, err)
	}
	if user == nil {
		return nil, errors.ErrUserNotFound // Domain-specific error
	}

	// Generate a new access token
	accessToken, err := utils.GenerateJWT(user.ID, user.Email, config.AppConfig.JWT.Secret, utils.TokenExpiry())
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToGenerateAccessToken, err)
	}

	// Optionally: Rotate the refresh token (generate a new one)
	newRefreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToGenerateNewAccessToken, err)
	}

	// Update the existing token entry in the database
	token.Token = accessToken
	token.RefreshToken = newRefreshToken
	token.Type = constants.RefreshToken
	token.ExpiresAt = time.Now().Add(utils.TokenExpiry())
	token.UpdatedAt = time.Now()
	token.RefreshTokenExpiresAt = time.Now().Add(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry))
	if err := svc.TokenRepo.UpdateToken(token); err != nil {
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToUpdateToken, err)
	}

	// Return the new tokens in the response
	return &dtos.RefreshTokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          newRefreshToken,
		ExpiresIn:             int64(utils.TokenExpiry().Seconds()),                                               // 2 hours
		RefreshTokenExpiresIn: int64(utils.ConvertTokenExpiry(config.AppConfig.JWT.RefreshTokenExpiry).Seconds()), // 24 hours
	}, nil
}
