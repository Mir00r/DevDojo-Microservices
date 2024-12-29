package controllers

import (
	"errors"
	"github.com/Mir00r/auth-service/constants"
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PublicAuthController manages public-facing authentication APIs
type PublicAuthController struct {
	AuthService  *services.AuthService  // Handles authentication-related logic
	TokenService *services.TokenService // Handles token-related logic
}

// NewPublicAuthController initializes a new PublicAuthController instance
func NewPublicAuthController(authService *services.AuthService, tokenService *services.TokenService) *PublicAuthController {
	return &PublicAuthController{
		AuthService:  authService,
		TokenService: tokenService,
	}
}

// PublicLogin handles user login requests and issues JWT tokens
// @Summary Login a user and return an access token
// @Tags Public Authentication
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login request payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /v1/public/auth/login [post]
func (ctrl *PublicAuthController) PublicLogin(c *gin.Context) {
	var req request.LoginRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRqPayload})
		return
	}

	// Authenticate the user and generate tokens
	token, err := ctrl.AuthService.Authenticate(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the generated token
	c.JSON(http.StatusOK, token)
}

// PublicRegister handles user registration requests
// @Summary Register a new user
// @Tags Public Authentication
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Registration request payload"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/public/auth/register [post]
func (ctrl *PublicAuthController) PublicRegister(c *gin.Context) {
	var req request.RegisterRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRqPayload})
		return
	}

	// Register the user
	if err := ctrl.AuthService.RegisterUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message
	c.JSON(http.StatusCreated, gin.H{"message": constants.MsgUserRegSuccessful})
}

// PasswordReset initiates a password reset process for a user
// @Summary Initiate password reset
// @Tags Public Authentication
// @Accept json
// @Produce json
// @Param request body request.PasswordResetRequest true "Password reset request payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/public/auth/password-reset [post]
func (ctrl *PublicAuthController) PasswordReset(c *gin.Context) {
	var req request.PasswordResetRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRqPayload})
		return
	}

	// Initiate the password reset process
	if err := ctrl.TokenService.InitiatePasswordReset(req); err != nil {
		if errors.Is(err, constants.ErrUserNotFoundVar) {
			c.JSON(http.StatusNotFound, gin.H{"error": constants.ResourceNotFound})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToInitiatePasswordReset})
		}
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": constants.PasswordResetLinkSentSuccessful})
}

// ConfirmPasswordReset finalizes the password reset process
// @Summary Confirm password reset
// @Tags Public Authentication
// @Accept json
// @Produce json
// @Param request body request.ConfirmPasswordResetRequest true "Confirm password reset request payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/public/auth/confirm-password-reset [post]
func (ctrl *PublicAuthController) ConfirmPasswordReset(c *gin.Context) {
	var req request.ConfirmPasswordResetRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRqPayload})
		return
	}

	// Confirm the password reset
	if err := ctrl.TokenService.ResetPassword(req); err != nil {
		if errors.Is(err, constants.ErrUserNotFoundVar) {
			c.JSON(http.StatusNotFound, gin.H{"error": constants.ErrUserNotFound})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrFailedToConfirmPasswordReset})
		}
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": constants.PasswordResetLinkSentSuccessful})
}
