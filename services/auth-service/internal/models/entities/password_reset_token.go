package entities

import (
	"gorm.io/gorm"
	"time"
)

// PasswordResetToken represents a password reset token entity
type PasswordResetToken struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Token     string         `gorm:"type:text;not null" json:"token"`
	UserID    string         `gorm:"type:uuid;not null" json:"user_id"`
	Used      bool           `gorm:"default:false" json:"used"` // Indicates whether the token has been used
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName overrides the default table name
func (PasswordResetToken) TableName() string {
	return "auth.password_reset_token"
}
