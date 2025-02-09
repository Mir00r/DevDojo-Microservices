package controllers

import (
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/errors"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PublicAuthController manages public-facing authentication APIs
type PublicAuthController struct {
	AuthService  services.AuthService           // Handles authentication-related logic
	TokenService services.TokenServiceInterface // Handles token-related logic
}

// NewPublicAuthController initializes a new PublicAuthController instance
func NewPublicAuthController(authService services.AuthService, tokenService services.TokenServiceInterface) *PublicAuthController {
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
	var req dtos.LoginRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.ErrInvalidPayload) // Propagate error to middlewares
		return
	}

	// Authenticate the user
	token, err := ctrl.AuthService.Authenticate(req)
	if err != nil || token == nil {
		_ = c.Error(err) // Propagate error to middlewares
		return
	}

	// Return the generated token
	utils.JSONResponseCtx(c, http.StatusOK, token)
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
	var req dtos.RegisterRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.ErrInvalidPayload) // Propagate error to middlewares
		return
	}

	// Register the user
	if err := ctrl.AuthService.RegisterUser(req); err != nil {
		_ = c.Error(err) // Propagate error to middlewares
		return
	}

	// Return a success message
	utils.JSONResponseCtx(c, http.StatusCreated, constants.MsgUserRegSuccessful)
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
	var req dtos.PasswordResetRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.ErrInvalidPayload) // Propagate error to middlewares
		return
	}

	// Initiate the password reset process
	if err := ctrl.TokenService.InitiatePasswordReset(req); err != nil {
		_ = c.Error(err) // Propagate error to middlewares
		return
	}

	// Return a success message
	utils.JSONResponseCtx(c, http.StatusOK, constants.PasswordResetLinkSentSuccessful)
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
	var req dtos.ConfirmPasswordResetRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(errors.ErrInvalidPayload) // Propagate error to middlewares
		return
	}

	// Confirm the password reset
	if err := ctrl.TokenService.ResetPassword(req); err != nil {
		_ = c.Error(err) // Propagate error to middlewares
		return
	}

	// Return a success message
	utils.JSONResponseCtx(c, http.StatusOK, constants.PasswordResetLinkSentSuccessful)
}
