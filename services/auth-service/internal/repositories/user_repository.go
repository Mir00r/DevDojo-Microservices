package repositories

import (
	"errors"
	"github.com/Mir00r/auth-service/internal/models/entities"
	"gorm.io/gorm"
	"log"
)

// UserRepository defines methods for interacting with the users table
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser inserts a new user record into the database
func (repo *UserRepository) CreateUser(user *entities.User) error {
	if err := repo.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByEmail retrieves a user by their email address
func (repo *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	//log.Printf("Querying user by email: %v", email)

	// Use LOWER() for case-insensitive matching
	if err := repo.DB.Where("LOWER(email) = LOWER(?)", email).First(&user).Error; err != nil {
		log.Printf("Error querying user from DB: %v using email: %v", err, email)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("No user found with email: %v", email)
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindUserByID retrieves a user by their ID
func (repo *UserRepository) FindUserByID(id string) (*entities.User, error) {
	var user entities.User
	if err := repo.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil for user if not found
		}
		return nil, err
	}
	return &user, nil
}

// UpdatePassword updates the user's password in the database
func (repo *UserRepository) UpdatePassword(userID string, hashedPassword string) error {
	return repo.DB.Model(&entities.User{}).
		Where("id = ?", userID).
		Update("password", hashedPassword).
		Error
}
