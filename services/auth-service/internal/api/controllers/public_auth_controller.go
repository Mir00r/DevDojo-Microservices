package controllers

import (
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PublicAuthController struct {
	AuthService *services.AuthService
}

func NewPublicAuthController(authService *services.AuthService) *PublicAuthController {
	return &PublicAuthController{AuthService: authService}
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

//
//// PublicPasswordReset handles password reset requests
//func PublicPasswordReset(w http.ResponseWriter, r *http.Request) {
//	var resetRequest services.PasswordResetRequest
//	if err := json.NewDecoder(r.Body).Decode(&resetRequest); err != nil {
//		responses.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
//		return
//	}
//
//	if err := services.InitiatePasswordReset(resetRequest); err != nil {
//		responses.ErrorResponse(w, http.StatusInternalServerError, err.Error())
//		return
//	}
//
//	responses.JSONResponse(w, http.StatusOK, map[string]string{"message": "Password reset initiated"})
//}
