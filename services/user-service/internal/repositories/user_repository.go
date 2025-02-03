package repositories

import (
	"context"
	"errors"
	"github.com/Mir00r/user-service/internal/models/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	// CRUD operations
	CreateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	GetUserByID(ctx context.Context, userID string) (*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAllUsers(ctx context.Context, limit, offset int) ([]entities.User, int64, error)
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, userID string) error

	// Role management
	AssignRoleToUser(ctx context.Context, userID string, role string) error
	RemoveRoleFromUser(ctx context.Context, userID string, role string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *userRepository) CreateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID retrieves a user by ID
func (r *userRepository) GetUserByID(ctx context.Context, userID string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail retrieves a user by email
func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all auth with pagination
func (r *userRepository) GetAllUsers(ctx context.Context, limit, offset int) ([]entities.User, int64, error) {
	var users []entities.User
	var totalCount int64

	if err := r.db.WithContext(ctx).Model(&entities.User{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

// UpdateUser updates a user's details
func (r *userRepository) UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error) {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser soft-deletes a user by marking deleted_at
func (r *userRepository) DeleteUser(ctx context.Context, userID string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", userID).Delete(&entities.User{}).Error; err != nil {
		return err
	}
	return nil
}

// AssignRoleToUser assigns a role to a user
func (r *userRepository) AssignRoleToUser(ctx context.Context, userID string, role string) error {
	return r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", userID).Update("role", role).Error
}

// RemoveRoleFromUser removes a role from a user
func (r *userRepository) RemoveRoleFromUser(ctx context.Context, userID string, role string) error {
	return r.db.WithContext(ctx).Model(&entities.User{}).Where("id = ?", userID).Update("role", "user").Error
}
