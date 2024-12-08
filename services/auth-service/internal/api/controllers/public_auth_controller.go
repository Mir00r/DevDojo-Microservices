package controllers

import (
	"encoding/json"
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"

	"net/http"
)

type PublicAuthController struct {
	AuthService *services.AuthService
}

// PublicLogin handles user login and issues JWT
func (ctrl *PublicAuthController) PublicLogin(w http.ResponseWriter, r *http.Request) {
	var loginRequest request.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	token, err := ctrl.AuthService.Authenticate(loginRequest)
	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusOK, token)
}

// PublicRegister handles new user registration
func (ctrl *PublicAuthController) PublicRegister(w http.ResponseWriter, r *http.Request) {
	var registerRequest request.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := ctrl.AuthService.RegisterUser(registerRequest); err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
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
