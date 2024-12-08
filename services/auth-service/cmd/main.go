package main

import (
	"github.com/Mir00r/auth-service/db"
	"github.com/Mir00r/auth-service/internal/api/routes"
	"github.com/Mir00r/auth-service/internal/configs"
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
	//database.Connect()

	// Create a new Gin router
	router := gin.Default()

	// Register routes
	routes.SetupAuthRoutes(router)

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
