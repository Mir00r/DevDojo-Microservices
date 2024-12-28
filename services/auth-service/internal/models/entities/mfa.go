package entities

import (
	"gorm.io/gorm"
	"time"
)

// Mfa represents the token entity in the system.
type Mfa struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"` // UUID as the primary key
	UserID    string         `gorm:"type:uuid;not null;index" json:"user_id"`                  // Foreign key to User
	OTP       string         `gorm:"type:text;not null" json:"otp"`                            // The token string
	Used      bool           `gorm:"default:false" json:"used"`                                // Indicates whether the token has been used
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`                               // Token expiration timestamp
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`                         // Automatically set at creation
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                           // Soft delete support
}

// TableName overrides the default table name
func (Mfa) TableName() string {
	return "auth.mfa"
}
