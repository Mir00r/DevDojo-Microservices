package utils

import (
	"errors"
	"github.com/Mir00r/user-service/configs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http/httptest"
	"regexp"
	"strings"
	"time"
	"unicode"
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
	duration, err := time.ParseDuration(configs.AppConfig.JWT.Expiry)
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

// IsValidEmail checks if the provided email is valid
func IsValidEmail(email string) bool {
	// Simple email regex validation
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// IsStrongPassword checks if the provided password meets strength criteria
func IsStrongPassword(password string) bool {
	var hasMinLen, hasUpper, hasLower, hasDigit, hasSpecial bool

	// Check minimum length
	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:'\",.<>?/`~", char):
			hasSpecial = true
		}
	}

	// Password is strong if it satisfies all conditions
	return hasMinLen && hasUpper && hasLower && hasDigit && hasSpecial
}

// IsValidUUID checks if the provided string is a valid UUID
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// IsValidPhone checks if the provided phone number is valid
func IsValidPhone(phone string) bool {
	// Regex for validating phone numbers (basic example)
	re := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	return re.MatchString(phone)
}

// IsValidPhoneWithRgx checks if the provided phone number is valid with provided regex
func IsValidPhoneWithRgx(phone string, rgx string) bool {
	// Regex for validating phone numbers (basic example)
	re := regexp.MustCompile(rgx)
	return re.MatchString(phone)
}
