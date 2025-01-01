#!/bin/bash

# Script to initialize a Go microservice project

# Exit if any command fails
set -e

# Function to display usage instructions
usage() {
    echo "Usage: $0 <service-name> <location> <github-username>"
    exit 1
}

# Check if arguments are provided
if [ $# -ne 3 ]; then
    usage
fi

SERVICE_NAME=$1
LOCATION=$2
GITHUB_USERNAME=$3
SERVICE_PATH="${LOCATION}/${SERVICE_NAME}-service"
MODULE_NAME="github.com/${GITHUB_USERNAME}/${SERVICE_NAME}-service"

# Function to normalize the path based on the OS
normalize_path() {
    local path=$1

    # Check the OS
    case "$(uname -s)" in
        Darwin | Linux) # macOS or Linux
            echo "$path"
            ;;
        CYGWIN* | MINGW32* | MSYS* | MINGW*) # Windows
            echo "$(cygpath -w "$path")" # Convert to Windows-compatible path
            ;;
        *)
            echo "Unsupported OS. Exiting..."
            exit 1
            ;;
    esac
}

# Normalize the location path
NORMALIZED_LOCATION=$(normalize_path "$LOCATION")
SERVICE_PATH="${NORMALIZED_LOCATION}/${SERVICE_NAME}-service"
MODULE_NAME="github.com/${GITHUB_USERNAME}/${SERVICE_NAME}-service"

# Check if location exists, if not create it
if [ ! -d "$NORMALIZED_LOCATION" ]; then
    echo "Provided location '$NORMALIZED_LOCATION' does not exist. Creating it..."
    mkdir -p "$NORMALIZED_LOCATION"
    echo "Location '$NORMALIZED_LOCATION' created successfully."
fi

# Create folder structure
create_structure() {
    echo "Creating folder structure for service: $SERVICE_NAME at $LOCATION"

    # Main directories
    mkdir -p ${SERVICE_PATH}/{cmd,config/env,constants,containers,errors,middlewares,routes,internal/{api/{controllers},services,models/{dtos,entities},repositories},utils,test,scripts,db/{migrations,seeds},build/{helm},docs,logs}

    # Add placeholders
    echo "// Main entry point for the service" > ${SERVICE_PATH}/cmd/main.go
    echo "// Application configuration file" > ${SERVICE_PATH}/config/app_config.go
    echo "// Constants used across the application" > ${SERVICE_PATH}/constants/app_constant.go
    echo "// Dependency injection container" > ${SERVICE_PATH}/containers/container.go
    echo "// Centralized error definitions" > ${SERVICE_PATH}/errors/errors.go
    echo "// Middleware for handling errors and requests" > ${SERVICE_PATH}/middlewares/error_middleware.go
    echo "// Middleware for handling jwt requests" > ${SERVICE_PATH}/middlewares/jwt_middleware.go
    echo "// Define application routes" > ${SERVICE_PATH}/routes/routes.go
    echo "// Placeholder for DTO models" > ${SERVICE_PATH}/internal/models/dtos/sample_dto.go
    echo "// Placeholder for database entities" > ${SERVICE_PATH}/internal/models/entities/sample_entities.go
    echo "// Placeholder for repositories" > ${SERVICE_PATH}/internal/repositories/sample_repository.go
#    echo "// Utility functions for the application" > ${SERVICE_PATH}/internal/utils/bcrypt.go
#    echo "// JWT utility functions" > ${SERVICE_PATH}/internal/utils/jwt.go
    echo "// Database migration for users table" > ${SERVICE_PATH}/db/migrations/001_sample_table.up.sql
    echo "// Database migration logic" > ${SERVICE_PATH}/db/migration.go
    echo "// Database initialization logic" > ${SERVICE_PATH}/db/init.go
    echo "// Seed data for the database" > ${SERVICE_PATH}/db/seeds/sample_data.sql
    echo "// Dockerfile for the service" > ${SERVICE_PATH}/build/Dockerfile
    echo "// Makefile for building the service" > ${SERVICE_PATH}/build/Makefile
    echo "// OpenAPI/Swagger specification for the service" > ${SERVICE_PATH}/docs/openapi.yaml
    echo "// OpenAPI/Swagger specification for the service" > ${SERVICE_PATH}/scripts/migration_gen.sh

    # Create .gitignore for logs
    echo "*" > ${SERVICE_PATH}/logs/.gitignore
    echo "!README.md" >> ${SERVICE_PATH}/logs/.gitignore

    # Create root README.md
    echo "# ${SERVICE_NAME} Service" > ${SERVICE_PATH}/README.md
    echo "This is the ${SERVICE_NAME} microservice built using Go." >> ${SERVICE_PATH}/README.md
}


# Initialize Go project
initialize_go_project() {
    echo "Initializing Go module for service: $SERVICE_NAME"

    cd ${SERVICE_PATH}

    # Initialize Go module
    go mod init ${MODULE_NAME}

    # Install common dependencies
    echo "Installing common Go dependencies..."
    go get -u github.com/gin-gonic/gin
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/postgres
    go get -u github.com/spf13/viper
    go get -u github.com/sirupsen/logrus
    go get -u github.com/dgrijalva/jwt-go
    go get -u gorm.io/driver/postgres
    go get -u github.com/golang-migrate/migrate/v4

    # Create a basic main.go
    cat <<EOF > cmd/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "${MODULE_NAME}/config"
    "${MODULE_NAME}/routes"
)

func main() {
	// Step 1: Load Configuration
	configPath := getConfigPath()
	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Step 2: Initialize Database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Step 3: Run Database Migrations
	migrationPath, _ := database.MigrationPath()
	log.Printf("Resolved migration path: %s", migrationPath)
	if err := database.RunMigrations(migrationPath, config.AppConfig.Database.DSN); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Step 4: Initialize Dependencies
	appContainer := containers.NewContainer()

	// Step 5: Setup Router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Step 6: Start Server
	startServer(router)
}

// getConfigPath determines the configuration file path
func getConfigPath() string {
	configPath := "./configs/config.yaml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}
	return configPath
}

// startServer starts the Gin server on the configured port
func startServer(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port
	}

	log.Printf("Starting server on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
EOF

# Create a basic container.go
    cat <<EOF > containers/container.go
package containers

import (
)

// Container struct holds all application dependencies
type Container struct {
}

// NewContainer initializes all dependencies and returns a Container instance
func NewContainer() *Container {
	// Initialize repositories

	// Initialize services

	// Initialize controllers

	return &Container{
	}
}
EOF


    # Create a basic configuration loader
    cat <<EOF > config/app_config.go
package config

import (
    "gopkg.in/yaml.v3"
    "log"
    "os"
)

var AppConfig Config

type Config struct {
	Server           ServerConfig           `yaml:"server"`
	JWT              JWTConfig              `yaml:"jwt"`
	Database         DatabaseConfig         `yaml:"database"`
	Redis            RedisConfig            `yaml:"redis"`
	Password         PasswordConfig         `yaml:"password"`
	InternalSecurity InternalSecurityConfig `yaml:"internal-security"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type JWTConfig struct {
	Secret             string `yaml:"secret"`
	Expiry             string `yaml:"expiry"`
	RefreshTokenExpiry string `yaml:"refresh-token-expiry"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	DSN      string `yaml:"dsn"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type PasswordConfig struct {
	PasswordResetURL string `yaml:"PasswordResetURL"`
}

type InternalSecurityConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

func LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		return err
	}
	log.Println("Configuration loaded successfully")
	return nil
}
EOF

    # Create a basic inits file
    cat <<EOF > db/init.go

    package database

    import (
    	"fmt"
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
EOF

# Create a basic migration file
    cat <<EOF > db/migration.go
package database
   import (
   	"github.com/golang-migrate/migrate/v4"
   	"log"
   	"os"
   	"strings"

   	_ "github.com/golang-migrate/migrate/v4/database/postgres"
   	_ "github.com/golang-migrate/migrate/v4/source/file"
   )

   func RunMigrations(migrationPath string, dsn string) error {
   	// Convert the Windows path to a valid URL format
   	if os.PathSeparator == '\\' {
   		migrationPath = strings.ReplaceAll(migrationPath, "\\", "/")
   		migrationPath = "file://" + migrationPath
   	} else {
   		migrationPath = "file://" + migrationPath
   	}
   	log.Printf("Using migration path: %s", migrationPath)
   	log.Printf("DSN path: %s", dsn)

   	m, err := migrate.New(migrationPath, dsn)
   	if err != nil {
   		log.Fatalf("Failed to initialize migrations: %v", err)
   	}

   	// Apply all up migrations
   	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
   		log.Fatalf("Failed to run up migrations: %v", err)
   	}

   	log.Println("Migrations applied successfully")
   	return err
   }
EOF


# Create a basic errors.go
    cat <<EOF > errors/errors.go
package errors

import (
	"fmt"
	"net/http"
)

// Specific error types
var (
	ErrInvalidCredentials            = NewAppError(http.StatusUnauthorized, "Invalid credentials", nil)
	ErrGenerateToken                 = NewAppError(http.StatusInternalServerError, "Failed to generate token", nil)
	ErrSaveToken                     = NewAppError(http.StatusInternalServerError, "Failed to save token", nil)
	ErrInvalidPayload                = NewAppError(http.StatusBadRequest, "Invalid request payload", nil)
	ErrFailedToRegisterUser          = NewAppError(http.StatusInternalServerError, "Failed to register user", nil)
	ErrUserNotFound                  = NewAppError(http.StatusNotFound, "User not found", nil)
	ErrHashPassword                  = NewAppError(http.StatusInternalServerError, "Failed to hash password", nil)
	ErrFailedToInitiatePasswordReset = NewAppError(http.StatusInternalServerError, "Failed to initiate password reset", nil)
	ErrInvalidOrExpiredResetToken    = NewAppError(http.StatusBadRequest, "Invalid or expired reset token", nil)
	ErrResetTokenAlreadyUsed         = NewAppError(http.StatusBadRequest, "Reset token already used", nil)
	ErrFailedToUpdatePassword        = NewAppError(http.StatusInternalServerError, "Failed to update password", nil)
	ErrInvalidOrExpiredRefreshToken  = NewAppError(http.StatusUnauthorized, "Invalid or expired refresh token", nil)
)

// AppError represents a generic application error
type AppError struct {
	Code       int    // HTTP status code
	CodeStatus string // HTTP status code message
	Message    string // Error message to be returned to the client
	Err        error  // Underlying error (optional)
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError creates a new AppError
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:       code,
		CodeStatus: http.StatusText(code),
		Message:    message,
		Err:        err,
	}
}
EOF

# Create a basic error_middleware.go
    cat <<EOF > middlewares/error_middleware.go
package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		// Check if there were any errors
		err := c.Errors.Last()
		if err != nil {
			appErr, ok := err.Err.(*errors.AppError)
			if ok {
				// Handle known application errors
				c.JSON(appErr.Code, gin.H{
					"error":      true,
					"code":       appErr.Code,
					"codeStatus": http.StatusText(appErr.Code),
					"message":    appErr.Message,
				})
			} else {
				// Log and handle unknown errors
				log.Printf("Unexpected error: %v", err.Err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error":      true,
					"code":       appErr.Code,
					"codeStatus": http.StatusText(appErr.Code),
					"message":    "An unexpected error occurred",
				})
			}
		}
	}
}
EOF

# Create a basic jwt_middleware.go
    cat <<EOF > middlewares/jwt_middleware.go
package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// AuthMiddleware validates the JWT and injects user claims into the Gin context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader(constants.Authorization)
		if authHeader == "" {
			utils.GinErrorResponse(c, http.StatusUnauthorized, constants.ErrMissingAuthHeader)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, constants.Bearer)
		if tokenString == "" {
			utils.GinErrorResponse(c, http.StatusUnauthorized, constants.ErrInvalidAuthHeader)
			c.Abort()
			return
		}

		// Verify the JWT and extract claims
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			log.Printf("JWT verification error: %v", err)
			utils.GinErrorResponse(c, http.StatusUnauthorized, constants.ErrInvalidToken)
			c.Abort()
			return
		}
		log.Printf("Extracted user ID: %s", claims.UserID)

		// Add claims to the Go context
		ctx := utils.AddClaimsToContext(c.Request.Context(), claims)

		// Inject claims into the Gin context
		c.Set("userID", claims.UserID)
		c.Request = c.Request.WithContext(ctx) // Attach the modified context to the request

		log.Printf("Extracted user ID: %s", claims.UserID)
		c.Next()
	}
}

// BasicAuthMiddleware validates requests using Basic Auth or API keys
func BasicAuthMiddleware(c *gin.Context) {
	// Fetch expected credentials from environment variables
	expectedUsername := config.AppConfig.InternalSecurity.UserName
	expectedPassword := config.AppConfig.InternalSecurity.Password
	//expectedAPIKey := os.Getenv("API_KEY") // API key support

	// Check Basic Auth credentials
	username, password, hasAuth := c.Request.BasicAuth()
	if hasAuth && username == expectedUsername && password == expectedPassword {
		c.Next()
		return
	}

	// Unauthorized if no valid credentials
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
}
EOF

    # Create a basic routes file
    cat <<EOF > routes/routes.go
package routes

import (
    "github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
    public := router.Group("/v1/public")
    {
        public.GET("/health", func(c *gin.Context) {
            c.JSON(200, gin.H{"status": "healthy"})
        })
    }
}
EOF

    # Run go mod tidy to clean up dependencies
    go mod tidy
}

# Call the functions
create_structure
initialize_go_project

# Print success message
echo "Go microservice project '${SERVICE_NAME}' initialized successfully at '${SERVICE_PATH}' with module '${MODULE_NAME}'"
