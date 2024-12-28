package repositories

import (
	"github.com/Mir00r/auth-service/internal/models/entities"
	"gorm.io/gorm"
	"time"
)

type MFARepository struct {
	DB *gorm.DB
}

func NewMFARepository(db *gorm.DB) *MFARepository {
	return &MFARepository{DB: db}
}

func (repo *MFARepository) SaveMFA(userID, otp string, expiry time.Time) error {
	mfa := entities.Mfa{
		OTP:       otp,
		UserID:    userID,
		ExpiresAt: expiry,
	}
	//return repo.DB.Exec("INSERT INTO mfa (user_id, otp, expires_at) VALUES (?, ?, ?) ON CONFLICT (user_id) DO UPDATE SET otp = ?, expires_at = ?",
	//	userID, otp, expiry, otp, expiry).Error
	return repo.DB.Create(mfa).Error
}

func (repo *MFARepository) CreateMFA(mfa *entities.Mfa) error {
	if err := repo.DB.Create(mfa).Error; err != nil {
		return err
	}
	return nil
}

func (repo *MFARepository) GetMFA(userID string) (string, time.Time, error) {
	var otp string
	var expiry time.Time
	var mfa entities.Mfa

	//err := repo.DB.Raw("SELECT otp, expires_at FROM mfa WHERE user_id = ?", userID).Scan(&otp, &expiry).Error
	err := repo.DB.Where("user_id = ?", userID).First(&mfa).Error
	return otp, expiry, err
}

func (repo *MFARepository) GetUnusedMFAByUserId(userID string) (*entities.Mfa, error) {
	var mfa entities.Mfa

	// Query the database to find the MFA record for the given user ID
	err := repo.DB.Where("user_id = ? AND used = ?", userID, false).Order("expires_at DESC").First(&mfa).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil if no record is found
		}
		return nil, err // Return error if something else went wrong
	}

	return &mfa, nil
}

func (repo *MFARepository) GetUnusedMFAByUserIdAndOtp(userID string, otp string) (*entities.Mfa, error) {
	var mfa entities.Mfa

	// Query the database to find the MFA record for the given user ID
	err := repo.DB.Where("user_id = ? AND otp = ? AND used = ?", userID, otp, false).Order("expires_at DESC").First(&mfa).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Return nil if no record is found
		}
		return nil, err // Return error if something else went wrong
	}

	return &mfa, nil
}

func (repo *MFARepository) DeleteMFA(userID string) error {
	return repo.DB.Exec("DELETE FROM mfa WHERE user_id = ?", userID).Error
}

func (repo *MFARepository) UpdateUsed(id string, used bool) error {
	return repo.DB.Model(&entities.Mfa{}).
		Where("id = ?", id).
		Update("used", used).
		Error
}
