package controllers

import (
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InternalAuthController handles internal APIs related to authentication
type InternalAuthController struct {
	InternalAuthService *services.InternalAuthService
}

// NewInternalAuthController initializes a new InternalAuthController with its dependencies
func NewInternalAuthController(internalAuthService *services.InternalAuthService) *InternalAuthController {
	return &InternalAuthController{
		InternalAuthService: internalAuthService,
	}
}

// ValidateToken validates a JWT token for inter-service communication
// @Summary Validates a JWT token
// @Description Validates a JWT token provided in the request body
// @Tags Internal APIs
// @Accept json
// @Produce json
// @Param body request.ValidateTokenRequest true "Token validation request payload"
// @Success 200 {object} response.ValidateTokenResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /internal/v1/validate-token [post]
func (ctrl *InternalAuthController) ValidateToken(c *gin.Context) {
	var req request.ValidateTokenRequest

	// Validate the incoming request payload
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GinErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Call the service layer to validate the token
	response, err := ctrl.InternalAuthService.ValidateToken(req.Token)
	if err != nil {
		utils.GinErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with a success response
	utils.GinJSONResponse(c, http.StatusOK, response)
}

// ServiceHealth checks the health status of the authentication service
// @Summary Checks the health of the authentication service
// @Description Returns the health status of the service
// @Tags Internal APIs
// @Accept json
// @Produce json
// @Success 200 {object} response.ServiceHealthResponse
// @Router /internal/v1/service-health [get]
func (ctrl *InternalAuthController) ServiceHealth(c *gin.Context) {
	// Fetch the health status from the service layer
	health := ctrl.InternalAuthService.CheckHealth()

	// Respond with the health status
	utils.GinJSONResponse(c, http.StatusOK, health)
}
