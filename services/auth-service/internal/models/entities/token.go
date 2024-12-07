package entities

import (
	"gorm.io/gorm"
	"time"
)

// Token represents the token entity in the system.
type Token struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"` // UUID as the primary key
	UserID    string         `gorm:"type:uuid;not null;index" json:"user_id"`                  // Foreign key to User
	Token     string         `gorm:"type:text;not null" json:"token"`                          // The token string
	Type      string         `gorm:"type:varchar(50);not null" json:"type"`                    // Token type (e.g., "access", "refresh")
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`                               // Token expiration timestamp
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`                         // Automatically set at creation
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                           // Soft delete support
}

// TableName overrides the default table name
func (Token) TableName() string {
	return "auth.tokens"
}
