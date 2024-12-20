package services

import (
	"errors"
	"fmt"
	config "github.com/Mir00r/auth-service/internal/configs"

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
	resetLink := fmt.Sprintf("%s/reset-password?token=%s", config.GetConfig().PasswordResetURL, resetToken)
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
