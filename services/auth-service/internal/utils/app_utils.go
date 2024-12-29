package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
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
