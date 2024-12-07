package main

import (
	"github.com/Mir00r/auth-service/db"
	"github.com/Mir00r/auth-service/internal/configs"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize database
	dsn := config.GetConfig().DatabaseDSN
	database.InitDatabase(dsn)
	//database.Connect()

	// Other setup (e.g., API routes)
}
