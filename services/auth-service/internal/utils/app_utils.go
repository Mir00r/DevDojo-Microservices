package utils

import (
	"errors"
	config "github.com/Mir00r/auth-service/configs"
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httptest"
	"time"
)

// ExtractUserIDFromContext retrieves the user ID from the Gin context.
func ExtractUserIDFromContext(c *gin.Context) (string, error) {
	// Retrieve the user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		return "", errors.New("user ID not found in context")
	}

	// Type assertion to ensure userID is a string
	userIDStr, ok := userID.(string)
	if !ok {
		return "", errors.New("invalid user ID type")
	}

	return userIDStr, nil
}

func TokenExpiry() time.Duration {
	// Parse the "24h" string into a time.Duration
	duration, err := time.ParseDuration(config.AppConfig.JWT.Expiry)
	if err != nil {
		log.Fatalf("Failed to parse JWT expiry duration: %v", err)
	}
	return duration
}

func ConvertTokenExpiry(expiry string) time.Duration {
	// Parse the "24h" string into a time.Duration
	duration, err := time.ParseDuration(expiry)
	if err != nil {
		log.Fatalf("Failed to parse JWT expiry duration: %v", err)
	}
	return duration
}

// CreateTestContext initializes a mock Gin context for testing
func CreateTestContext(w *httptest.ResponseRecorder) *gin.Context {
	// Create a mock Gin context with the given ResponseRecorder
	c, _ := gin.CreateTestContext(w)

	return c
}
