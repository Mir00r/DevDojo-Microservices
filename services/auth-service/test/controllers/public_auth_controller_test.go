package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/Mir00r/auth-service/constants"
	"github.com/Mir00r/auth-service/internal/models/dtos"
	"github.com/Mir00r/auth-service/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/Mir00r/auth-service/mocks"
)

func TestPublicLogin_ValidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock dependencies
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockTokenService := mocks.NewMockTokenServiceInterface(ctrl)
	reqBody := dtos.LoginRequest{
		Email:    "john.doe1@example.com",
		Password: "12345",
	}
	respBody := dtos.LoginResponse{
		AccessToken:           "mockAccessToken",
		RefreshToken:          "mockRefreshToken",
		ExpiresIn:             3600,
		RefreshTokenExpiresIn: 7200,
	}
	mockAuthService.EXPECT().Authenticate(reqBody).Return(&respBody, nil)

	// Create a mock Gin context
	w := httptest.NewRecorder()
	c := utils.CreateTestContext(w)

	// Mock the request
	reqBodyBytes, _ := json.Marshal(reqBody)
	c.Request, _ = http.NewRequest("POST", "/v1/public/auth/login", bytes.NewReader(reqBodyBytes))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the controller with the mock context and service
	controller := controllers.NewPublicAuthController(mockAuthService, mockTokenService)
	controller.PublicLogin(c)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	// Deserialize the full response
	var apiResp dtos.LoginAPIResponse
	err := json.Unmarshal(w.Body.Bytes(), &apiResp)
	assert.NoError(t, err)

	// Validate the data field in the response
	assert.Equal(t, respBody.AccessToken, apiResp.Data.AccessToken)
	assert.Equal(t, respBody.RefreshToken, apiResp.Data.RefreshToken)
	assert.Equal(t, respBody.ExpiresIn, apiResp.Data.ExpiresIn)
	assert.Equal(t, respBody.RefreshTokenExpiresIn, apiResp.Data.RefreshTokenExpiresIn)
}

func TestPublicLogin_InvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock dependencies
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockTokenService := mocks.NewMockTokenServiceInterface(ctrl)
	reqBody := dtos.LoginRequest{
		Email:    "invalid@example.com",
		Password: "wrongpassword",
	}
	mockAuthService.EXPECT().Authenticate(reqBody).Return(nil, constants.ErrInvalidCredentials)

	// Create a mock Gin context
	w := httptest.NewRecorder()
	c := utils.CreateTestContext(w)

	// Mock the request
	reqBodyBytes, _ := json.Marshal(reqBody)
	c.Request, _ = http.NewRequest("POST", "/v1/public/auth/login", bytes.NewReader(reqBodyBytes))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the controller with the mock context and service
	controller := controllers.NewPublicAuthController(mockAuthService, mockTokenService)
	controller.PublicLogin(c)

	// Assertions
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), constants.InvalidCredentials)
}

func TestPublicLogin_InvalidRequestBody(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock dependencies
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockTokenService := mocks.NewMockTokenServiceInterface(ctrl)

	// Create a mock Gin context
	w := httptest.NewRecorder()
	c := utils.CreateTestContext(w)

	// Mock the request
	c.Request, _ = http.NewRequest("POST", "/v1/public/auth/login", bytes.NewReader([]byte("invalid-json")))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the controller with the mock context and service
	controller := controllers.NewPublicAuthController(mockAuthService, mockTokenService)
	controller.PublicLogin(c)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), constants.ErrInvalidRqPayload)
}

func TestPublicLogin_ServiceError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock dependencies
	mockAuthService := mocks.NewMockAuthService(ctrl)
	mockTokenService := mocks.NewMockTokenServiceInterface(ctrl)
	reqBody := dtos.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Ensure the second return value is of type error
	mockAuthService.EXPECT().Authenticate(reqBody).Return(nil, constants.ErrGenerateTokenVar)

	// Create a mock Gin context
	w := httptest.NewRecorder()
	c := utils.CreateTestContext(w)

	// Mock the request
	reqBodyBytes, _ := json.Marshal(reqBody)
	c.Request, _ = http.NewRequest("POST", "/v1/public/auth/login", bytes.NewReader(reqBodyBytes))
	c.Request.Header.Set("Content-Type", "application/json")

	// Call the controller with the mock context and service
	controller := controllers.NewPublicAuthController(mockAuthService, mockTokenService)
	controller.PublicLogin(c)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), constants.ErrGenerateToken)
}
