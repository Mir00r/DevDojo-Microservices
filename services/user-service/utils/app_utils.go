package utils

import (
	"errors"
	"github.com/Mir00r/user-service/config"
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

// Helper function to get a default value if input is nil
func GetOrDefault(input *string, defaultValue *string) *string {
	if input == nil || *input == "" {
		return defaultValue
	}
	return input
}

// ConvertStringToTime converts a string to a time.Time object based on the given layout.
// Example layout: "2006-01-02 15:04:05" for "YYYY-MM-DD HH:mm:ss" or "2006-01-02" for "YYYY-MM-DD".
func ConvertStringToTime(dateStr, layout string) (time.Time, error) {
	if layout == "" {
		layout = "2006-01-02" // Default to "YYYY-MM-DD" if no layout is provided
	}

	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, errors.New("invalid date format or value")
	}

	return parsedTime, nil
}
