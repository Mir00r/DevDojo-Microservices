package services

import (
	"context"
	"github.com/Mir00r/user-service/errors"
	"github.com/Mir00r/user-service/internal/models/dtos"
	"github.com/Mir00r/user-service/internal/models/entities"
	"github.com/Mir00r/user-service/internal/repositories"
	utils2 "github.com/Mir00r/user-service/utils"
	"log"
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
	// Hash the password
	hashedPassword, err := utils2.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	defaultRole := "user"
	parsedDateTime, err := utils2.ConvertStringToTime(req.DateOfBirth, "2006-01-02")
	if err != nil {
		log.Fatalf("Error parsing date-time: %v", err)
	}

	user := &entities.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    hashedPassword,
		Phone:       req.Phone,
		Role:        utils2.GetOrDefault(req.Role, &defaultRole),
		DateOfBirth: &parsedDateTime,
		Address:     req.Address,
	}

	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return dtos.ToUserResponse(createdUser), nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx context.Context, userID string) (*dtos.UserResponse, error) {
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
		user.Phone = req.Phone
	}
	if req.ProfilePicture != "" {
		user.ProfilePicture = req.ProfilePicture
	}

	updatedUser, err := s.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return dtos.ToUserResponse(updatedUser), nil
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(ctx context.Context, userID string) error {
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
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.ErrUserNotFound
	}

	return s.repo.RemoveRoleFromUser(ctx, userID, "user")
}
