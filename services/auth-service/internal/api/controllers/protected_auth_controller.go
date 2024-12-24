package controllers

import (
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ProtectedAuthController struct {
	AuthService  *services.AuthService
	TokenService *services.TokenService
}

func NewProtectedAuthController(authService *services.AuthService, tokenService *services.TokenService) *ProtectedAuthController {
	return &ProtectedAuthController{AuthService: authService, TokenService: tokenService}
}

// ProtectedLogout handles the logout API request
func (ctrl *ProtectedAuthController) ProtectedLogout(c *gin.Context) {
	// Retrieve the user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		utils.GinErrorResponse(c, http.StatusUnauthorized, "User ID not found")
		return
	}

	// Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Call the logout service
	if err := ctrl.TokenService.Logout(tokenString, userID.(string)); err != nil {
		utils.GinErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.GinJSONResponse(c, http.StatusOK, gin.H{"message": "Logout successful"})
}

func (ctrl *ProtectedAuthController) ProtectedUserProfile(c *gin.Context) {
	// Retrieve the user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		utils.GinErrorResponse(c, http.StatusUnauthorized, "User ID not found")
		return
	}

	// Call the service to fetch the user's profile
	userProfile, err := ctrl.AuthService.GetUserProfile(userID.(string))
	if err != nil {
		utils.GinErrorResponse(c, http.StatusInternalServerError, "Failed to fetch user profile")
		return
	}

	utils.GinJSONResponse(c, http.StatusOK, userProfile)
}
