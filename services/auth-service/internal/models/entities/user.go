package entities

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user entity in the system.
type User struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"` // UUID as the primary key
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`                   // User's full name
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`      // Unique email
	Password  string         `gorm:"type:varchar(255);not null" json:"-"`                      // Hashed password
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`                         // Automatically set at creation
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`                         // Automatically updated
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                           // Soft delete support
}

// TableName overrides the default table name
func (User) TableName() string {
	return "auth.auth"
}
