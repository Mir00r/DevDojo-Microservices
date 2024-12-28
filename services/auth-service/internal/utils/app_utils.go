package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserId(c *gin.Context) any {
	userID, exists := c.Get("userID")
	if !exists {
		GinErrorResponse(c, http.StatusUnauthorized, "User ID not found")
		return ""
	}
	return userID
}
