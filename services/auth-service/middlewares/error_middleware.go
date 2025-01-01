package middlewares

import (
	"github.com/Mir00r/auth-service/errors"
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
