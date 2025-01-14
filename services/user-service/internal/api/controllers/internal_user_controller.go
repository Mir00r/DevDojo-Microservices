package controllers

import (
	"github.com/Mir00r/user-service/errors"
	"github.com/Mir00r/user-service/internal/models/dtos"
	"github.com/Mir00r/user-service/internal/services"
	"github.com/Mir00r/user-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InternalUserController handles internal user-related APIs
type InternalUserController struct {
	UserService services.UserService
}

// NewInternalUserController initializes a new InternalUserController
func NewInternalUserController(userService services.UserService) *InternalUserController {
	return &InternalUserController{
		UserService: userService,
	}
}

// CreateUser creates a new user
func (c *InternalUserController) CreateUser(ctx *gin.Context) {
	var req dtos.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationErrorToString(err)})
		_ = ctx.Error(errors.ErrInvalidPayload) // Propagate error to middleware
		return
	}

	user, err := c.UserService.CreateUser(ctx, req)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_ = ctx.Error(errors.ErrFailedToRegisterUser) // Propagate error to middleware
		return
	}

	//ctx.JSON(http.StatusCreated, user)
	utils.JSONResponseCtx(ctx, http.StatusCreated, user)
}

// GetUserDetails retrieves user details, including internal fields
func (c *InternalUserController) GetUserDetails(ctx *gin.Context) {
	userId := ctx.Param("userId")

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

	ctx.JSON(http.StatusOK, user)
}

// ActivateUser activates a user account
//func (c *InternalUserController) ActivateUser(ctx *gin.Context) {
//	userId := ctx.Param("userId")
//
//	err := c.UserService.ActivateUser(userId)
//	if err != nil {
//		if err.Error() == "user not found" {
//			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		} else {
//			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		}
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "User activated successfully"})
//}
//
//// DeactivateUser deactivates a user account
//func (c *InternalUserController) DeactivateUser(ctx *gin.Context) {
//	userId := ctx.Param("userId")
//
//	err := c.UserService.DeactivateUser(userId)
//	if err != nil {
//		if err.Error() == "user not found" {
//			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//		} else {
//			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		}
//		return
//	}
//
//	ctx.JSON(http.StatusOK, gin.H{"message": "User deactivated successfully"})
//}
//
//// SearchUsers searches for users based on filters
//func (c *InternalUserController) SearchUsers(ctx *gin.Context) {
//	var req dtos.UserSearchRequest
//	if err := ctx.ShouldBindQuery(&req); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationErrorToString(err)})
//		return
//	}
//
//	users, err := c.UserService.SearchUsers(&req)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, users)
//}
