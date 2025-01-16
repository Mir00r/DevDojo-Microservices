package main

import (
	"github.com/Mir00r/user-service/configs"
	"github.com/Mir00r/user-service/containers"
	database "github.com/Mir00r/user-service/db"
	"github.com/Mir00r/user-service/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Step 1: Load Configuration
	configPath := getConfigPath()
	if err := configs.LoadConfig(configPath); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Step 2: Initialize Database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Step 3: Run Database Migrations
	migrationPath, _ := database.MigrationPath()
	log.Printf("Resolved migration path: %s", migrationPath)
	if err := database.RunMigrations(migrationPath, configs.AppConfig.Database.DSN); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Step 4: Initialize Dependencies
	appContainer := containers.NewContainer()

	// Step 5: Setup Router
	router := gin.Default()
	routes.SetupRoutes(router, appContainer.PublicUserController, appContainer.ProtectedUserController, appContainer.InternalUserController)

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
		port = "8082" // Default port
	}

	log.Printf("Starting server on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
