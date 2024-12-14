package routes

import (
	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes defines all API routes categorized by Public, Protected, and Internal APIs
func SetupRoutes(router *gin.Engine, publicAuthController *controllers.PublicAuthController) {

	// Public APIs
	public := router.Group("/public")
	{
		public.POST("/v1/login", publicAuthController.PublicLogin)
		public.POST("/v1/register", publicAuthController.PublicRegister)
	}
	//
	//// Protected APIs
	//protected := router.PathPrefix("/protected").Subrouter()
	//protected.Use(middlewares.JWTMiddleware)
	//protected.HandleFunc("/v1/user-profile", controllers.ProtectedUserProfile).Methods("GET")
	//protected.HandleFunc("/v1/logout", controllers.ProtectedLogout).Methods("POST")
	//
	//// Internal APIs
	//internal := router.PathPrefix("/internal").Subrouter()
	//internal.Use(middlewares.BasicAuthMiddleware)
	//internal.HandleFunc("/v1/validate-token", controllers.InternalValidateToken).Methods("POST")
}
