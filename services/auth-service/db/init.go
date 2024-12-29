package database

import (
	"fmt"
	config "github.com/Mir00r/auth-service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// DB is the global database instance
var DB *gorm.DB

// Once ensures the database connection is initialized only once
var once sync.Once

// Connect initializes the database connection using the configuration and ensures the connection is established only once.
// Returns an error if the connection fails.
func Connect() error {
	var dbErr error

	// Use sync.Once to ensure that the connection is created only once
	once.Do(func() {
		// Construct the DSN (Data Source Name) using configuration
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.AppConfig.Database.Host,
			config.AppConfig.Database.Port,
			config.AppConfig.Database.User,
			config.AppConfig.Database.Password,
			config.AppConfig.Database.DBName,
		)

		// Open the database connection
		DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr != nil {
			log.Printf("Failed to connect to the database: %v", dbErr)
			return
		}

		// Log successful connection
		log.Println("Database connection established successfully")
	})

	return dbErr
}

// MigrationPath constructs and returns the absolute path to the database migrations directory.
// This ensures portability and compatibility across different environments.
func MigrationPath() (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get the current working directory: %v", err)
		return "", err
	}

	// Construct the absolute migration path
	absoluteMigrationPath := filepath.Join(cwd, "db", "migrations")
	return absoluteMigrationPath, err
}

// Migrate applies the database migrations using GORM's AutoMigrate function.
// This method should be extended to include all models that need to be migrated.
func Migrate(models ...interface{}) error {
	if DB == nil {
		return fmt.Errorf("database connection is not established")
	}

	// Apply migrations
	if err := DB.AutoMigrate(models...); err != nil {
		log.Printf("Failed to apply database migrations: %v", err)
		return err
	}

	log.Println("Database migrations applied successfully")
	return nil
}
