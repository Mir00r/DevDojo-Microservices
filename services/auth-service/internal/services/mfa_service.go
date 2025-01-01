package services

import (
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/errors"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/Mir00r/auth-service/internal/models/entities"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"log"
	"net/http"
	"time"
)

type MFAService interface {
	EnableMFA(userID string) (*dtos.EnableMFAResponse, error)
	VerifyMFA(userID, otp string) error
}

// MFAService handles multi-factor authentication logic
type mfaService struct {
	MFARepo  repositories.MFARepository  // Repository for MFA-related operations
	UserRepo repositories.UserRepository // Repository for user-related operations
}

// NewMFAService creates a new instance of MFAService
func NewMFAService(mfaRepo repositories.MFARepository, userRepo repositories.UserRepository) MFAService {
	return &mfaService{
		MFARepo:  mfaRepo,
		UserRepo: userRepo,
	}
}

// EnableMFA generates and stores an OTP for enabling MFA and sends it to the user's email
func (svc *mfaService) EnableMFA(userID string) (*dtos.EnableMFAResponse, error) {
	// Validate user existence
	user, err := svc.UserRepo.FindUserByID(userID)
	if err != nil {
		log.Printf("Failed to fetch user: %v", err)
		return nil, errors.NewAppError(http.StatusNotFound, constants.ErrUserNotFound, err)
	}

	// Generate OTP
	otp := utils.GenerateOTP()

	// Prepare MFA entity
	mfa := &entities.Mfa{
		OTP:       otp,
		UserID:    userID,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	// Save OTP in the database
	if err := svc.MFARepo.CreateMFA(mfa); err != nil {
		log.Printf("Failed to save MFA record: %v", err)
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToEnableMFA, err)
	}

	// Send OTP to the user's email
	if err := utils.SendOTPEmail(user.Email, otp); err != nil {
		log.Printf("Failed to send OTP email: %v", err)
		return nil, errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToSendOTPEmail, err)
	}

	// Return the EnableMFAResponse object
	return &dtos.EnableMFAResponse{
		Message: "MFA enabled successfully",
		OTP:     otp,
	}, nil
}

// VerifyMFA checks if the provided OTP matches the stored OTP for the user and marks it as used
func (svc *mfaService) VerifyMFA(userID, otp string) error {
	// Retrieve the latest unused MFA record for the user
	mfa, err := svc.MFARepo.GetUnusedMFAByUserId(userID)
	if err != nil {
		return errors.NewAppError(http.StatusNotFound, constants.ErrFailedToVerifyMFA, err)
	}
	if mfa == nil {
		return errors.NewAppError(http.StatusNotFound, "No unused MFA record found for user: "+userID, err)
	}

	// Check if the OTP has expired
	if time.Now().After(mfa.ExpiresAt) {
		return errors.NewAppError(http.StatusUnauthorized, constants.ErrOTPExpired, err)
	}

	// Validate the OTP
	if otp != mfa.OTP {
		return errors.NewAppError(http.StatusUnauthorized, constants.ErrInvalidOTP, err)
	}

	// Mark the OTP as used
	if err := svc.MFARepo.UpdateUsed(mfa.ID, true); err != nil {
		return errors.NewAppError(http.StatusInternalServerError, constants.ErrFailedToMarkMFA, err)
	}

	return nil
}
