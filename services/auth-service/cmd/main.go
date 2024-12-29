package main

import (
	config "github.com/Mir00r/auth-service/configs"
	"github.com/Mir00r/auth-service/containers"
	database "github.com/Mir00r/auth-service/db"
	"github.com/Mir00r/auth-service/internal/api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Step 1: Load Configuration
	configPath := getConfigPath()
	if err := config.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Step 2: Initialize Database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Step 3: Run Database Migrations
	migrationPath, _ := database.MigrationPath()
	log.Printf("Resolved migration path: %s", migrationPath)
	if err := database.RunMigrations(migrationPath, config.AppConfig.Database.DSN); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Step 4: Initialize Dependencies
	appContainer := containers.NewContainer()

	// Step 5: Setup Router
	router := gin.Default()
	routes.SetupRoutes(router,
		appContainer.PublicAuthController, appContainer.ProtectedAuthController,
		appContainer.InternalAuthController,
	)

	// Step 6: Start Server
	startServer(router)
}

// getConfigPath determines the configuration file path
func getConfigPath() string {
	configPath := "./configs/config.yaml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}
	return configPath
}

// startServer starts the Gin server on the configured port
func startServer(router *gin.Engine) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Default port
	}

	log.Printf("Starting server on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
