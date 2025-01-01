package utils

import (
	"encoding/json"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSONResponse sends a standardized JSON response
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

// ErrorResponse sends a standardized error response
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	JSONResponse(w, status, map[string]string{"error": message})
}

// GinErrorResponse sends an error response using Gin
func GinErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

// GinJSONResponse sends a JSON response using Gin
func GinJSONResponse(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}

// ErrorResponseCtx sends a JSON error response with a status code
func ErrorResponseCtx(c *gin.Context, status int, message string) {
	c.JSON(status, dtos.APIResponse{
		Error:      true,
		Code:       status,
		CodeStatus: http.StatusText(status),
		Message:    message,
	})
}

// JSONResponseCtx sends a JSON success response with a status code and payload
func JSONResponseCtx(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, dtos.APIResponse{
		Code:    status,
		Error:   false,
		Message: http.StatusText(status),
		Data:    payload,
	})
}
