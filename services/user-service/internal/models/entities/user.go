package entities

import (
	"time"
)

// User represents the user entity in the system.
type User struct {
	ID             string     `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name           string     `gorm:"type:varchar(100);not null" json:"name"`
	Email          string     `gorm:"type:varchar(100);unique;not null" json:"email"`
	Password       string     `gorm:"type:varchar(255);not null" json:"-"`
	Phone          string     `gorm:"type:varchar(15)" json:"phone,omitempty"`
	IsActive       bool       `gorm:"type:boolean;default:true" json:"is_active"`
	IsVerified     bool       `gorm:"type:boolean;default:false" json:"is_verified"`
	ProfilePicture string     `gorm:"type:varchar(255)" json:"profile_picture,omitempty"`
	Role           *string    `gorm:"type:varchar(50);default:'user'" json:"role"`
	CreatedAt      time.Time  `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt      time.Time  `gorm:"type:timestamp;default:now()" json:"updated_at"`
	DeletedAt      *time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`
	LastLogin      *time.Time `gorm:"type:timestamp" json:"last_login,omitempty"`
	DateOfBirth    *time.Time `gorm:"type:date" json:"date_of_birth,omitempty"`
	Address        *string    `gorm:"type:text" json:"address,omitempty"`
	TenantID       string     `gorm:"type:uuid" json:"tenant_id,omitempty"`
	Locale         string     `gorm:"type:varchar(10);default:'en-US'" json:"locale"`
	Timezone       string     `gorm:"type:varchar(50);default:'UTC'" json:"timezone"`
	MFAEnabled     bool       `gorm:"type:boolean;default:false" json:"mfa_enabled"`
	MFASecret      *string    `gorm:"type:varchar(255)" json:"mfa_secret,omitempty"`
}

// TableName overrides the default table name
func (User) TableName() string {
	return "my_user.users"
}
