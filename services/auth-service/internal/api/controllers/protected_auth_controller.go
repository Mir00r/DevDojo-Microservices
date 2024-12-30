package controllers

import (
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// ProtectedAuthController handles protected API actions that require authentication and authorization.
type ProtectedAuthController struct {
	AuthService  services.AuthService           // Handles user-related operations
	TokenService services.TokenServiceInterface // Handles token-related operations
	MFAService   services.MFAService            // Handles multi-factor authentication operations
}

// NewProtectedAuthController creates a new instance of ProtectedAuthController.
func NewProtectedAuthController(
	authService services.AuthService,
	tokenService services.TokenServiceInterface,
	mfaService services.MFAService,
) *ProtectedAuthController {
	return &ProtectedAuthController{
		AuthService:  authService,
		TokenService: tokenService,
		MFAService:   mfaService,
	}
}

// ProtectedLogout handles the logout API request.
// This endpoint invalidates the user's current access token, effectively logging them out.
func (ctrl *ProtectedAuthController) ProtectedLogout(c *gin.Context) {
	// Extract the user ID using the helper function
	userID, err := utils.ExtractUserIDFromContext(c)
	if err != nil {
		utils.ErrorResponseCtx(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Extract the token from the Authorization header
	authHeader := c.GetHeader(constants.Authorization)
	if authHeader == "" {
		utils.ErrorResponseCtx(c, http.StatusUnauthorized, constants.ErrMissingAuthHeader)
		return
	}
	tokenString := strings.TrimPrefix(authHeader, constants.Bearer)

	// Call the logout service to invalidate the token
	if err := ctrl.TokenService.Logout(tokenString, userID); err != nil {
		utils.ErrorResponseCtx(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with a success message
	utils.JSONResponseCtx(c, http.StatusOK, constants.MsgLogoutSuccessful)
}

// ProtectedUserProfile fetches the profile of the currently authenticated user.
// This endpoint retrieves the user's profile data based on their user ID.
func (ctrl *ProtectedAuthController) ProtectedUserProfile(c *gin.Context) {
	// Retrieve the user ID from the context
	userID, exists := c.Get("userID")
	if !exists {
		utils.ErrorResponseCtx(c, http.StatusUnauthorized, constants.ErrUserNotFound)
		return
	}

	// Fetch the user's profile using the AuthService
	userProfile, err := ctrl.AuthService.GetUserProfile(userID.(string))
	if err != nil {
		utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrFailedToFetchProfile)
		return
	}

	// Respond with the user's profile data
	utils.JSONResponseCtx(c, http.StatusOK, userProfile)
}

// EnableMFA enables multi-factor authentication (MFA) for the currently authenticated user.
// This endpoint generates an OTP (One-Time Password) to initiate the MFA setup process.
func (ctrl *ProtectedAuthController) EnableMFA(c *gin.Context) {
	// Extract user ID from JWT claims
	claims, _ := utils.ExtractClaimsFromContext(c.Request.Context())
	userID := claims.UserID

	// Call the MFAService to enable MFA and generate an OTP
	otp, err := ctrl.MFAService.EnableMFA(userID)
	if err != nil {
		utils.ErrorResponseCtx(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond with the generated OTP and success message
	utils.JSONResponseCtx(c, http.StatusOK, gin.H{
		"message": "MFA enabled",
		"otp":     otp,
	})
}

// VerifyMFA verifies the provided OTP for the currently authenticated user.
// This endpoint validates the OTP submitted by the user during the MFA process.
func (ctrl *ProtectedAuthController) VerifyMFA(c *gin.Context) {
	// Bind the request payload to the VerifyMFARequest struct
	var req request.VerifyMFARequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseCtx(c, http.StatusBadRequest, constants.ErrInvalidRqPayload)
		return
	}

	// Extract user ID from JWT claims
	claims, _ := utils.ExtractClaimsFromContext(c.Request.Context())
	userID := claims.UserID

	// Call the MFAService to verify the OTP
	if err := ctrl.MFAService.VerifyMFA(userID, req.OTP); err != nil {
		utils.ErrorResponseCtx(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with a success message
	utils.JSONResponseCtx(c, http.StatusOK, constants.MFAVerifySuccessful)
}

func (ctrl *ProtectedAuthController) RefreshToken(c *gin.Context) {
	var req dtos.RefreshTokenRequest

	// Parse JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseCtx(c, http.StatusBadRequest, constants.ErrInvalidRqPayload)
		return
	}

	// Validate the request
	if err := services.ValidateRequest(req); err != nil {
		utils.ErrorResponseCtx(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := ctrl.TokenService.RefreshToken(req)
	if err != nil {
		utils.ErrorResponseCtx(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Send success response
	utils.JSONResponseCtx(c, http.StatusOK, response)
}
