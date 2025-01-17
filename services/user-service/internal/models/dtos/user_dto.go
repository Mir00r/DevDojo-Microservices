package dtos

import (
	"github.com/Mir00r/user-service/internal/models/entities"
	"time"
)

// CreateUserRequest is used for creating a new user
type CreateUserRequest struct {
	Name           string  `json:"name" validate:"required,min=3,max=100"`
	Email          string  `json:"email" validate:"required,email"`
	Password       string  `json:"password" validate:"required,min=8"`
	Phone          string  `json:"phone" validate:"omitempty,e164"` // Optional, E.164 format
	Role           *string `json:"role" validate:"omitempty"`
	DateOfBirth    string  `json:"dateOfBirth" validate:"omitempty,datetime=2006-01-02"` // Optional, ISO date format
	Address        *string `json:"address" validate:"omitempty,max=500"`
	ProfilePicture *string `json:"profilePicture" validate:"omitempty,url"` // Optional, valid URL
	Locale         *string `json:"locale" validate:"omitempty,len=5"`       // Optional, e.g., 'en-US'
	Timezone       *string `json:"timezone" validate:"omitempty,max=50"`
}

// UpdateUserRequest is used for updating an existing user
type UpdateUserRequest struct {
	Name           string  `json:"name" validate:"omitempty,min=3,max=100"`
	Phone          string  `json:"phone" validate:"omitempty,e164"`
	DateOfBirth    *string `json:"dateOfBirth" validate:"omitempty,datetime=2006-01-02"`
	Address        *string `json:"address" validate:"omitempty,max=500"`
	ProfilePicture string  `json:"profilePicture" validate:"omitempty,url"`
	Locale         *string `json:"locale" validate:"omitempty,len=5"`
	Timezone       *string `json:"timezone" validate:"omitempty,max=50"`
}

// AssignRoleRequest is used for assigning roles to a user
type AssignRoleRequest struct {
	Role string `json:"role" validate:"required,oneof=user admin moderator"`
}

// UserResponse is used for retrieving user details
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

// ToUserResponse converts a User model to a UserResponse DTO
func ToUserResponse(user *entities.User) *UserResponse {
	if user == nil {
		return nil
	}

	return &UserResponse{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Phone:          user.Phone,
		IsActive:       user.IsActive,
		IsVerified:     user.IsVerified,
		ProfilePicture: user.ProfilePicture,
		Role:           user.Role,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		LastLogin:      user.LastLogin,
		DateOfBirth:    user.DateOfBirth,
		Address:        user.Address,
		TenantID:       user.TenantID,
		Locale:         user.Locale,
		Timezone:       user.Timezone,
		MFAEnabled:     user.MFAEnabled,
	}
}

// PaginatedUserResponse is used for returning a list of users with pagination
type PaginatedUserResponse struct {
	Users      []UserResponse `json:"users"`
	TotalCount int64          `json:"totalCount"`
	Page       int            `json:"page"`
	PerPage    int            `json:"perPage"`
}
