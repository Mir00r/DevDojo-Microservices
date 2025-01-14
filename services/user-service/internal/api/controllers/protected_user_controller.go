package controllers

import (
	"github.com/Mir00r/user-service/errors"
	"github.com/Mir00r/user-service/internal/models/dtos"
	"github.com/Mir00r/user-service/internal/services"
	"github.com/Mir00r/user-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ProtectedUserController handles protected user-related APIs
type ProtectedUserController struct {
	UserService services.UserService
}

// NewProtectedUserController initializes a new ProtectedUserController
func NewProtectedUserController(userService services.UserService) *ProtectedUserController {
	return &ProtectedUserController{
		UserService: userService,
	}
}

// GetAllUsers retrieves all users (Admin only)
func (c *ProtectedUserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserService.GetAllUsers(ctx, 0, 10)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		_ = ctx.Error(errors.ErrFailedToFetchUser) // Propagate error to middleware
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// GetUserByID retrieves a user's details by ID (Admin/User self-access)
func (c *ProtectedUserController) GetUserByID(ctx *gin.Context) {
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

// UpdateUser updates a user's details (Admin/User self-access)
func (c *ProtectedUserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var req dtos.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		//ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ValidationErrorToString(err)})
		_ = ctx.Error(errors.ErrInvalidPayload) // Propagate error to middleware
		return
	}

	user, err := c.UserService.UpdateUser(ctx, userId, req)
	if err != nil {
		if err.Error() == "user not found" {
			//ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			_ = ctx.Error(errors.ErrUserNotFound) // Propagate error to middleware
		} else {
			//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			_ = ctx.Error(errors.ErrFailedToUpdateUser) // Propagate error to middleware
		}
		return
	}

	//ctx.JSON(http.StatusOK, user)
	utils.JSONResponseCtx(ctx, http.StatusCreated, user)
}

// DeleteUser deletes a user (Admin only)
func (c *ProtectedUserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("userId")

	err := c.UserService.DeleteUser(ctx, userId)
	if err != nil {
		if err.Error() == "user not found" {
			//ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			_ = ctx.Error(errors.ErrUserNotFound) // Propagate error to middleware
		} else {
			//ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			_ = ctx.Error(errors.ErrFailedToDeleteUser) // Propagate error to middleware
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// // GetUserRoles retrieves roles assigned to a user (Admin only)
//
//	func (c *ProtectedUserController) GetUserRoles(ctx *gin.Context) {
//		userId := ctx.Param("userId")
//
//		roles, err := c.UserService.GetUserRoles(ctx, userId)
//		if err != nil {
//			if err.Error() == "user not found" {
//				ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
//			} else {
//				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//			}
//			return
//		}
//
//		ctx.JSON(http.StatusOK, roles)
//	}
//
// AssignRoles assigns roles to a user (Admin only)
func (c *ProtectedUserController) AssignRoles(ctx *gin.Context) {
	userId := ctx.Param("userId")

	var req dtos.AssignRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_ = ctx.Error(errors.ErrInvalidPayload) // Propagate error to middleware
		return
	}

	err := c.UserService.AssignRole(ctx, userId, req.Role)
	if err != nil {
		if err.Error() == "user not found" {
			//ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			_ = ctx.Error(errors.ErrUserNotFound) // Propagate error to middleware
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			_ = ctx.Error(errors.ErrFailedToAssignUserRole) // Propagate error to middleware
		}
		return
	}

	//ctx.JSON(http.StatusOK, gin.H{"message": "Roles assigned successfully"})
	utils.JSONResponseCtx(ctx, http.StatusCreated, "Roles assigned successfully")
}

// RemoveRole removes a role from a user (Admin only)
func (c *ProtectedUserController) RemoveRole(ctx *gin.Context) {
	userId := ctx.Param("userId")
	//roleId := ctx.Param("roleId")

	err := c.UserService.RemoveRole(ctx, userId)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else if err.Error() == "role not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	//ctx.JSON(http.StatusOK, gin.H{"message": "Role removed successfully"})
	utils.JSONResponseCtx(ctx, http.StatusCreated, "Roles removed successfully")
}
