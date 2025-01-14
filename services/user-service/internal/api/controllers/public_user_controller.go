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
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationErrorToString(err)})
		_ = ctx.Error(errors.ErrInvalidPayload) // Propagate error to
		return
	}

	// Call the service to create a new user
	user, err := c.UserService.CreateUser(ctx, req)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_ = ctx.Error(errors.ErrFailedToRegisterUser) // Propagate error to middleware
		return
	}

	// Return the created user response
	//ctx.JSON(http.StatusCreated, user)
	utils.JSONResponseCtx(ctx, http.StatusCreated, user)
}

// GetUser handles fetching user details by ID
func (c *PublicUserController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	// Call the service to fetch user details
	user, err := c.UserService.GetUserByID(ctx, userId)
	if err != nil {
		if err.Error() == "user not found" {
			//ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			_ = ctx.Error(errors.ErrUserNotFound) // Propagate error to middleware
		} else {
			//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			_ = ctx.Error(errors.ErrFailedToFetchUser) // Propagate error to middleware
		}
		return
	}

	// Return the user details
	//ctx.JSON(http.StatusOK, user)
	utils.JSONResponseCtx(ctx, http.StatusOK, user)
}
