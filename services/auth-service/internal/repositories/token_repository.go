package repositories

import (
	"github.com/Mir00r/auth-service/internal/models/entities"
	"gorm.io/gorm"
)

type TokenRepository struct {
	DB *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{DB: db}
}

// SaveResetToken saves a password reset token in the database
func (repo *TokenRepository) SaveResetToken(token string, userID string) error {
	resetToken := entities.PasswordResetToken{
		Token:  token,
		UserID: userID,
	}
	return repo.DB.Create(&resetToken).Error
}

func (repo *TokenRepository) MarkTokenAsUsed(token string) error {
	return repo.DB.Model(&entities.PasswordResetToken{}).
		Where("token = ?", token).
		Update("used", true).
		Error
}

// FindToken retrieves a password reset token by its value
func (repo *TokenRepository) FindToken(token string) (*entities.PasswordResetToken, error) {
	var resetToken entities.PasswordResetToken
	if err := repo.DB.Where("token = ?", token).First(&resetToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // Token not found
		}
		return nil, err
	}
	return &resetToken, nil
}
