package controllers

import (
	"errors"
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	request "github.com/Mir00r/auth-service/internal/models/request"
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
		utils.ErrorResponseCtx(c, http.StatusBadRequest, constants.ErrInvalidRqPayload)
		return
	}

	// Authenticate the user and generate tokens
	token, err := ctrl.AuthService.Authenticate(req)
	if err != nil {
		// Handle different error types
		switch {
		case errors.Is(err, constants.ErrInvalidCredentials):
			utils.ErrorResponseCtx(c, http.StatusUnauthorized, constants.ErrInvalidCredentials.Error())
		case errors.Is(err, constants.ErrGenerateTokenVar):
			utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrGenerateTokenVar.Error())
		case errors.Is(err, constants.ErrSaveTokenVar):
			utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrSaveTokenVar.Error())
		default:
			utils.ErrorResponseCtx(c, http.StatusInternalServerError, "An unknown error occurred")
		}
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
	var req request.RegisterRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseCtx(c, http.StatusBadRequest, constants.ErrInvalidRqPayload)
		return
	}

	// Register the user
	if err := ctrl.AuthService.RegisterUser(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrFailedToRegisterUser)
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
	var req request.PasswordResetRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseCtx(c, http.StatusBadRequest, constants.ErrInvalidRqPayload)
		return
	}

	// Initiate the password reset process
	if err := ctrl.TokenService.InitiatePasswordReset(req); err != nil {
		if errors.Is(err, constants.ErrUserNotFoundVar) {
			utils.ErrorResponseCtx(c, http.StatusNotFound, constants.ResourceNotFound)
		} else {
			utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrFailedToInitiatePasswordReset)
		}
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
	var req request.ConfirmPasswordResetRequest

	// Parse and validate the request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrInvalidRqPayload})
		utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrFailedToInitiatePasswordReset)
		return
	}

	// Confirm the password reset
	if err := ctrl.TokenService.ResetPassword(req); err != nil {
		if errors.Is(err, constants.ErrUserNotFoundVar) {
			utils.ErrorResponseCtx(c, http.StatusNotFound, constants.ErrUserNotFound)
		} else {
			utils.ErrorResponseCtx(c, http.StatusInternalServerError, constants.ErrFailedToConfirmPasswordReset)
		}
		return
	}

	// Return a success message
	utils.JSONResponseCtx(c, http.StatusOK, constants.PasswordResetLinkSentSuccessful)
}
