package dtos

import "time"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken           string `json:"accessToken"`
	RefreshToken          string `json:"refreshToken"`
	ExpiresIn             int64  `json:"expiresIn"`             // Time in seconds until the token expires
	RefreshTokenExpiresIn int64  `json:"refreshTokenExpiresIn"` // Time in seconds until the token expires
}

// LoginAPIResponse represents the entire response structure
type LoginAPIResponse struct {
	Error   bool          `json:"error"`
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    LoginResponse `json:"data"`
}

type UserResponse struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone,omitempty"`
	IsActive       bool       `json:"isActive"`
	IsVerified     bool       `json:"isVerified"`
	ProfilePicture string     `json:"profilePicture,omitempty"`
	Role           *string    `json:"role"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	LastLogin      *time.Time `json:"lastLogin,omitempty"`
	DateOfBirth    *time.Time `json:"dateOfBirth,omitempty"`
	Address        *string    `json:"address,omitempty"`
	Locale         string     `json:"locale"`
	Timezone       string     `json:"timezone"`
	TenantID       string     `json:"tenantId,omitempty"`
	MFAEnabled     bool       `json:"isMfaEnabled"`
}

type UserAPIResponse struct {
	Error   bool         `json:"error"`
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}
