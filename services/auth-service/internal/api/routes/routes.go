package routes

import (
	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/Mir00r/auth-service/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

// SetupRoutes defines all API routes categorized by Public, Protected, and Internal APIs
func SetupRoutes(router *gin.Engine,
	publicAuthController *controllers.PublicAuthController,
	protectedAuthController *controllers.ProtectedAuthController,
	internalAuthController *controllers.InternalAuthController,
) {

	// Public APIs
	public := router.Group("/v1/public/auth")
	{
		public.POST("/login", publicAuthController.PublicLogin)
		public.POST("/register", publicAuthController.PublicRegister)
		public.POST("/password-reset", publicAuthController.PasswordReset)
		public.POST("/confirm-password-reset", publicAuthController.ConfirmPasswordReset)
	}

	// Protected APIs
	protected := router.Group("/v1/protected/auth")
	protected.Use(middlewares.AuthMiddleware()) // Apply JWT validation middleware
	{
		protected.POST("/logout", protectedAuthController.ProtectedLogout)
		protected.GET("/user-profile", protectedAuthController.ProtectedUserProfile)
		protected.POST("/mfa/enable", protectedAuthController.EnableMFA)
		protected.POST("/mfa/verify", protectedAuthController.VerifyMFA)
	}

	// Internal APIs
	internal := router.Group("/v1/internal/auth")
	internal.Use(middlewares.BasicAuthMiddleware) // Apply Basic Auth middleware
	{
		internal.POST("/validate-token", internalAuthController.ValidateToken)
		internal.GET("/service-health", internalAuthController.ServiceHealth)
	}
}
