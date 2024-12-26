package controllers

import (
	"errors"
	"github.com/Mir00r/auth-service/constants"
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PublicAuthController struct {
	AuthService  *services.AuthService
	TokenService *services.TokenService
}

func NewPublicAuthController(authService *services.AuthService, tokenService *services.TokenService) *PublicAuthController {
	return &PublicAuthController{AuthService: authService, TokenService: tokenService}
}

// PublicLogin handles user login and issues JWT
func (ctrl *PublicAuthController) PublicLogin(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Authenticate the user
	token, err := ctrl.AuthService.Authenticate(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}

// PublicRegister handles user registration
func (ctrl *PublicAuthController) PublicRegister(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Register the user
	err := ctrl.AuthService.RegisterUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (ctrl *PublicAuthController) PasswordReset(c *gin.Context) {
	var req request.PasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the TokenService to initiate the password reset
	err := ctrl.TokenService.InitiatePasswordReset(req)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFoundVar) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate password reset"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent successfully"})
}

func (ctrl *PublicAuthController) ConfirmPasswordReset(c *gin.Context) {
	var req request.ConfirmPasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the TokenService to confirm the password reset
	err := ctrl.TokenService.ResetPassword(req)
	if err != nil {
		if errors.Is(err, constants.ErrUserNotFoundVar) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to confirm password reset"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link confirm successfully"})
}
