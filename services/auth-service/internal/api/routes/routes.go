package routes

import (
	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/Mir00r/auth-service/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes all API routes grouped by category and applies appropriate middlewares.
func SetupRoutes(
	router *gin.Engine,
	publicAuthController *controllers.PublicAuthController,
	protectedAuthController *controllers.ProtectedAuthController,
	internalAuthController *controllers.InternalAuthController,
) {
	// Initialize Public API routes
	initializePublicRoutes(router, publicAuthController)

	// Initialize Protected API routes
	initializeProtectedRoutes(router, protectedAuthController)

	// Initialize Internal API routes
	initializeInternalRoutes(router, internalAuthController)
}

// initializePublicRoutes sets up routes for Public APIs
func initializePublicRoutes(router *gin.Engine, controller *controllers.PublicAuthController) {
	publicGroup := router.Group("/v1/public/auth")
	{
		publicGroup.POST("/login", controller.PublicLogin)
		publicGroup.POST("/register", controller.PublicRegister)
		publicGroup.POST("/password-reset", controller.PasswordReset)
		publicGroup.POST("/confirm-password-reset", controller.ConfirmPasswordReset)
	}
}

// initializeProtectedRoutes sets up routes for Protected APIs
func initializeProtectedRoutes(router *gin.Engine, controller *controllers.ProtectedAuthController) {
	protectedGroup := router.Group("/v1/protected/auth")
	protectedGroup.Use(middlewares.AuthMiddleware()) // Apply JWT validation middleware
	{
		protectedGroup.POST("/logout", controller.ProtectedLogout)
		protectedGroup.GET("/user-profile", controller.ProtectedUserProfile)
		protectedGroup.POST("/mfa/enable", controller.EnableMFA)
		protectedGroup.POST("/mfa/verify", controller.VerifyMFA)
	}
}

// initializeInternalRoutes sets up routes for Internal APIs
func initializeInternalRoutes(router *gin.Engine, controller *controllers.InternalAuthController) {
	internalGroup := router.Group("/v1/internal/auth")
	internalGroup.Use(middlewares.BasicAuthMiddleware) // Apply Basic Auth middleware
	{
		internalGroup.POST("/validate-token", controller.ValidateToken)
		internalGroup.GET("/service-health", controller.ServiceHealth)
	}
}
