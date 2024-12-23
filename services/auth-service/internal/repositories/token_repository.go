package repositories

import (
	"errors"
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

// BlacklistToken saves a token to the blacklist
func (repo *TokenRepository) BlacklistToken(token *entities.Token) error {
	return repo.DB.Create(token).Error
}

// IsTokenBlacklisted checks if a token is blacklisted
func (repo *TokenRepository) IsTokenBlacklisted(tokenString string) (bool, error) {
	var token entities.Token
	err := repo.DB.Where("token = ?", tokenString).First(&token).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
