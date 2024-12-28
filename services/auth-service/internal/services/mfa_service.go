package services

import (
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/entities"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/utils"
	"log"
	"time"
)

type MFAService struct {
	MFARepo  *repositories.MFARepository
	UserRepo *repositories.UserRepository
}

func NewMFAService(mfaRepo *repositories.MFARepository, userRepo *repositories.UserRepository) *MFAService {
	return &MFAService{
		MFARepo:  mfaRepo,
		UserRepo: userRepo,
	}
}

// EnableMFA generates and stores an OTP for enabling MFA
func (svc *MFAService) EnableMFA(userID string) (string, error) {
	// Generate OTP
	otp := utils.GenerateOTP()

	mfa := &entities.Mfa{
		OTP:       otp,
		UserID:    userID,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	// Save OTP in the database
	if err := svc.MFARepo.CreateMFA(mfa); err != nil {
		return "", err
	}

	// Send OTP to the user's email
	user, err := svc.UserRepo.FindUserByID(userID)
	if err != nil {
		return "", err
	}
	if err := utils.SendOTPEmail(user.Email, otp); err != nil {
		return "", err
	}

	return otp, nil
}

// VerifyMFA checks if the provided OTP matches the stored OTP for the user
func (svc *MFAService) VerifyMFA(userID, otp string) error {
	// Retrieve the stored OTP
	mfa, err := svc.MFARepo.GetUnusedMFAByUserId(userID)
	if err != nil {
		log.Printf("Error fetching MFA record: %v", err)
		return err
	}
	if mfa == nil {
		log.Printf("No MFA record found for user: %s", userID)
		return err
	}

	// Check if the OTP is expired
	if time.Now().After(mfa.ExpiresAt) {
		return constants.ErrOTPExpiredVar
	}

	// Verify the OTP
	if otp != mfa.OTP {
		return constants.ErrInvalidOTPVar
	}

	// Delete the OTP after successful verification
	//return svc.MFARepo.DeleteMFA(userID)

	// Update the OTP used status after successful verification
	return svc.MFARepo.UpdateUsed(mfa.ID, true)
}
