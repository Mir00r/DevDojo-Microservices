package api

import (
	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/gorilla/mux"
)

// SetupRoutes defines all API routes categorized by Public, Protected, and Internal APIs
func SetupRoutes(publicAuthController *controllers.PublicAuthController) *mux.Router {

	router := mux.NewRouter()

	// Public APIs
	public := router.PathPrefix("/public").Subrouter()
	public.HandleFunc("/v1/login", publicAuthController.PublicLogin).Methods("POST")
	public.HandleFunc("/v1/register", publicAuthController.PublicRegister).Methods("POST")
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

	return router
}
