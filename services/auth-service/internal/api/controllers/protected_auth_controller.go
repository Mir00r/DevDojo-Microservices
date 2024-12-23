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
	// Extract the token from the Authorization header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.GinErrorResponse(c, http.StatusUnauthorized, "Authorization header is missing")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		utils.GinErrorResponse(c, http.StatusUnauthorized, "Invalid Authorization header")
		return
	}

	// Verify the JWT and extract claims
	claims, err := utils.VerifyJWT(tokenString)
	if err != nil {
		utils.GinErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		return
	}

	// Extract user_id from the claims
	userID := claims.UserID
	if userID == "" {
		utils.GinErrorResponse(c, http.StatusUnauthorized, "User ID not found in token")
		return
	}

	// Call the logout service
	if err := ctrl.TokenService.Logout(tokenString, userID); err != nil {
		utils.GinErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.GinJSONResponse(c, http.StatusOK, gin.H{"message": "Logout successful"})
}
