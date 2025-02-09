package controllers

import (
	"github.com/Mir00r/user-service/errors"
	"github.com/Mir00r/user-service/internal/models/dtos"
	"github.com/Mir00r/user-service/internal/services"
	"github.com/Mir00r/user-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PublicUserController handles public user-related APIs
type PublicUserController struct {
	UserService services.UserService
}

// NewPublicUserController initializes a new PublicUserController
func NewPublicUserController(userService services.UserService) *PublicUserController {
	return &PublicUserController{
		UserService: userService,
	}
}

// CreateUser handles the creation of a new user
func (c *PublicUserController) CreateUser(ctx *gin.Context) {
	var req dtos.CreateUserRequest

	// Bind and validate the request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(errors.ErrInvalidPayload) // Propagate error to
		return
	}

	// Call the service to create a new user
	user, err := c.UserService.CreateUser(ctx, req)
	if err != nil {
		_ = ctx.Error(err) // Propagate error to middlewares
		return
	}

	// Return the created user response
	utils.JSONResponseCtx(ctx, http.StatusCreated, user)
}

// GetUser handles fetching user details by ID
func (c *PublicUserController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	// Call the service to fetch user details
	user, err := c.UserService.GetUserByID(ctx, userId)
	if err != nil {
		_ = ctx.Error(err) // Propagate error to middlewares
		return
	}

	// Return the user details
	utils.JSONResponseCtx(ctx, http.StatusOK, user)
}
