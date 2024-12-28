package controllers

import (
	request "github.com/Mir00r/auth-service/internal/models/request"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/Mir00r/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InternalAuthController struct {
	InternalAuthService *services.InternalAuthService
}

func NewInternalAuthController(
	internalAuthService *services.InternalAuthService,
) *InternalAuthController {
	return &InternalAuthController{
		InternalAuthService: internalAuthService,
	}
}

// ValidateToken validates a JWT token for inter-service communication
func (ctrl *InternalAuthController) ValidateToken(c *gin.Context) {
	var req request.ValidateTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.GinErrorResponse(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	response, err := ctrl.InternalAuthService.ValidateToken(req.Token)
	if err != nil {
		utils.GinErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.GinJSONResponse(c, http.StatusOK, response)
}

// ServiceHealth checks the health of the authentication service
func (ctrl *InternalAuthController) ServiceHealth(c *gin.Context) {
	health := ctrl.InternalAuthService.CheckHealth()
	utils.GinJSONResponse(c, http.StatusOK, health)
}
