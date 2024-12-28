package controllers

import (
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type ProtectedAuthController struct {
	AuthService  *services.AuthService
	TokenService *services.TokenService
	MFAService   *services.MFAService
}

func NewProtectedAuthController(
	authService *services.AuthService,
	tokenService *services.TokenService,
	mfaService *services.MFAService,
) *ProtectedAuthController {
	return &ProtectedAuthController{
		AuthService:  authService,
		TokenService: tokenService,
		MFAService:   mfaService,
	}
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

// EnableMFA enables multi-factor authentication for the authenticated user
func (ctrl *ProtectedAuthController) EnableMFA(c *gin.Context) {
	// Extract user ID from JWT claims
	claims, _ := utils.ExtractClaimsFromContext(c.Request.Context())
	userID := claims.UserID

	// Call the service to enable MFA
	otp, err := ctrl.MFAService.EnableMFA(userID)
	if err != nil {
		utils.GinErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.GinJSONResponse(c, http.StatusOK, gin.H{"message": "MFA enabled", "otp": otp})
}

// VerifyMFA verifies the provided OTP for the authenticated user
func (ctrl *ProtectedAuthController) VerifyMFA(c *gin.Context) {
	var req request.VerifyMFARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GinErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Extract user ID from JWT claims
	claims, _ := utils.ExtractClaimsFromContext(c.Request.Context())
	userID := claims.UserID

	// Call the service to verify MFA
	if err := ctrl.MFAService.VerifyMFA(userID, req.OTP); err != nil {
		utils.GinErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.GinJSONResponse(c, http.StatusOK, gin.H{"message": "MFA verification successful"})
}
