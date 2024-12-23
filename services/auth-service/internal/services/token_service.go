package services

import (
	"errors"
	"fmt"
	"github.com/Mir00r/auth-service/configs"
	"github.com/Mir00r/auth-service/internal/models/entities"
	"time"

	services "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"log"
)

type TokenService struct {
	TokenRepo *repositories.TokenRepository
	UserRepo  *repositories.UserRepository
}

func NewTokenService(repo *repositories.TokenRepository, userRepo *repositories.UserRepository) *TokenService {
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
		return ErrUserNotFound
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
		return errors.New("token is already invalidated")
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
