package routes

import (
	"github.com/Mir00r/user-service/internal/api/controllers"
	"github.com/Mir00r/user-service/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all API routes grouped by category and applies appropriate middlewares.
func SetupRoutes(
	router *gin.Engine,
	publicUserController *controllers.PublicUserController,
	protectedUserController *controllers.ProtectedUserController,
	internalUserController *controllers.InternalUserController,
) {
	// Attach exception middleware
	router.Use(middlewares.ErrorHandler())

	// Initialize Public API routes
	initializePublicRoutes(router, publicUserController)

	// Initialize Protected API routes
	initializeProtectedRoutes(router, protectedUserController)

	// Initialize Internal API routes
	initializeInternalRoutes(router, internalUserController)
}

// initializePublicRoutes sets up routes for Public APIs
func initializePublicRoutes(router *gin.Engine, controller *controllers.PublicUserController) {
	publicGroup := router.Group("/v1/public/user")
	{
		publicGroup.POST("/register", controller.CreateUser)
		publicGroup.GET("/:userId/details", controller.GetUser)
	}
}

// initializeProtectedRoutes sets up routes for Protected APIs
func initializeProtectedRoutes(router *gin.Engine, controller *controllers.ProtectedUserController) {
	protectedGroup := router.Group("/v1/protected/user")
	protectedGroup.Use(middlewares.AuthMiddleware()) // Apply JWT validation middleware
	{
		protectedGroup.GET("", controller.GetAllUsers)           // Admin only
		protectedGroup.GET("/:userId", controller.GetUserByID)   // Admin/User (self-access)
		protectedGroup.PUT("/:userId", controller.UpdateUser)    // Admin/User (self-access)
		protectedGroup.DELETE("/:userId", controller.DeleteUser) // Admin only
		//protectedGroup.GET("/:userId/roles", controller.GetUserRoles)   // Admin only
		//protectedGroup.POST("/:userId/roles", controller.AssignRoles)   // Admin only
		//protectedGroup.DELETE("/:userId/roles/:roleId", controller.RemoveRole) // Admin only
	}
}

// initializeInternalRoutes sets up routes for Internal APIs
func initializeInternalRoutes(router *gin.Engine, controller *controllers.InternalUserController) {
	internalGroup := router.Group("/v1/internal/user")
	internalGroup.Use(middlewares.BasicAuthMiddleware) // Apply Basic Auth middleware
	{
		internalGroup.POST("", controller.CreateUser)                    // Create a new user
		internalGroup.GET("/:userId/details", controller.GetUserDetails) // Fetch user details (with all internal fields)
		//internalGroup.PUT("/:userId/activate", controllers.ActivateUser)     // Activate user account
		//internalGroup.PUT("/:userId/deactivate", controllers.DeactivateUser) // Deactivate user account
		//internalGroup.GET("/search", controllers.SearchUsers)                // Search users by filters
	}
}
