package main

import (
	"github.com/Mir00r/auth-service/db"
	"github.com/Mir00r/auth-service/internal/api/controllers"
	"github.com/Mir00r/auth-service/internal/api/routes"
	"github.com/Mir00r/auth-service/internal/configs"
	"github.com/Mir00r/auth-service/internal/repositories"
	"github.com/Mir00r/auth-service/internal/services"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	dsn := config.GetConfig().DatabaseDSN
	database.InitDatabase(dsn)
	// Run migrations
	database.RunMigrations(config.GetConfig().MigrationPath, dsn)

	// Create a new Gin router
	router := gin.Default()

	// Initialize the controller
	userRepo := repositories.NewUserRepository(database.DB)
	tokenRepo := repositories.NewTokenRepository(database.DB)
	authService := services.NewAuthService(userRepo)
	tokenService := services.NewTokenService(tokenRepo, userRepo)
	publicAuthController := controllers.NewPublicAuthController(authService, tokenService)
	protectedAuthController := controllers.NewProtectedAuthController(authService, tokenService)

	// Register routes
	routes.SetupRoutes(router, publicAuthController, protectedAuthController)

	// Get the port from environment variable or default to 8081
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Start the server
	log.Printf("Starting server on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
