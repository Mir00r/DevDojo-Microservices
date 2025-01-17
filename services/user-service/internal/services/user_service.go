package services

import (
	"context"
	"github.com/Mir00r/user-service/errors"
	"github.com/Mir00r/user-service/internal/models/dtos"
	"github.com/Mir00r/user-service/internal/models/entities"
	"github.com/Mir00r/user-service/internal/repositories"
	utils2 "github.com/Mir00r/user-service/utils"
)

type UserService interface {
	CreateUser(ctx context.Context, req dtos.CreateUserRequest) (*dtos.UserResponse, error)
	GetUserByID(ctx context.Context, userID string) (*dtos.UserResponse, error)
	GetAllUsers(ctx context.Context, limit, offset int) (*dtos.PaginatedUserResponse, error)
	UpdateUser(ctx context.Context, userID string, req dtos.UpdateUserRequest) (*dtos.UserResponse, error)
	DeleteUser(ctx context.Context, userID string) error
	AssignRole(ctx context.Context, userID string, role string) error
	RemoveRole(ctx context.Context, userID string) error
}

type userService struct {
	repo repositories.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// CreateUser creates a new user
func (s *userService) CreateUser(ctx context.Context, req dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	// Validate email
	if !utils2.IsValidEmail(req.Email) {
		return nil, errors.ErrInvalidEmail
	}

	// Validate password strength
	if !utils2.IsStrongPassword(req.Password) {
		return nil, errors.ErrWeakPassword
	}

	// Validate date of birth
	parsedDateTime, err := utils2.ConvertStringToTime(req.DateOfBirth, "2006-01-02")
	if err != nil {
		return nil, errors.ErrInvalidDateOfBirth
	}

	// Check if email already exists
	existingUser, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.ErrEmailAlreadyExists
	}

	// Hash the password
	hashedPassword, err := utils2.HashPassword(req.Password)
	if err != nil {
		return nil, errors.ErrPasswordHashing
	}

	defaultRole := "User"
	user := &entities.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    hashedPassword,
		Phone:       req.Phone,
		Role:        utils2.GetOrDefault(req.Role, &defaultRole),
		DateOfBirth: &parsedDateTime,
		Address:     req.Address,
	}

	// Save user
	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return dtos.ToUserResponse(createdUser), nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx context.Context, userID string) (*dtos.UserResponse, error) {
	// Validate user ID
	if !utils2.IsValidUUID(userID) {
		return nil, errors.ErrInvalidUserID
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	return dtos.ToUserResponse(user), nil
}

// GetAllUsers retrieves a paginated list of users
func (s *userService) GetAllUsers(ctx context.Context, limit, offset int) (*dtos.PaginatedUserResponse, error) {
	if limit <= 0 || offset < 0 {
		return nil, errors.ErrInvalidPagination
	}

	users, totalCount, err := s.repo.GetAllUsers(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	userResponses := make([]dtos.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = *dtos.ToUserResponse(&user)
	}

	return &dtos.PaginatedUserResponse{
		Users:      userResponses,
		TotalCount: totalCount,
	}, nil
}

// UpdateUser updates a user's information
func (s *userService) UpdateUser(ctx context.Context, userID string, req dtos.UpdateUserRequest) (*dtos.UserResponse, error) {
	// Validate user ID
	if !utils2.IsValidUUID(userID) {
		return nil, errors.ErrInvalidUserID
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.ErrUserNotFound
	}

	// Update fields
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Phone != "" {
		if !utils2.IsValidPhone(req.Phone) {
			return nil, errors.ErrInvalidPhone
		}
		user.Phone = req.Phone
	}
	if req.ProfilePicture != "" {
		user.ProfilePicture = req.ProfilePicture
	}

	// Save updated user
	updatedUser, err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return dtos.ToUserResponse(updatedUser), nil
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(ctx context.Context, userID string) error {
	// Validate user ID
	if !utils2.IsValidUUID(userID) {
		return errors.ErrInvalidUserID
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.ErrUserNotFound
	}

	return s.repo.DeleteUser(ctx, userID)
}

// AssignRole assigns a role to a user
func (s *userService) AssignRole(ctx context.Context, userID string, role string) error {
	// Validate inputs
	if !utils2.IsValidUUID(userID) {
		return errors.ErrInvalidUserID
	}
	if role == "" {
		return errors.ErrInvalidRole
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.ErrUserNotFound
	}

	return s.repo.AssignRoleToUser(ctx, userID, role)
}

// RemoveRole removes a user's role and sets it to default
func (s *userService) RemoveRole(ctx context.Context, userID string) error {
	// Validate user ID
	if !utils2.IsValidUUID(userID) {
		return errors.ErrInvalidUserID
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.ErrUserNotFound
	}

	return s.repo.RemoveRoleFromUser(ctx, userID, "user")
}
