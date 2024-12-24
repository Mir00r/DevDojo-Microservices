package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a `.env` file, if it exists
func LoadEnv() {
	if _, err := os.Stat(".env"); err == nil {
		log.Println("Loading environment variables from .env file")
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}
}

func LoadYamlEnv() {
	// Load the configuration file
	configPath := "./internal/configs/config.yaml"
	if envPath := os.Getenv("CONFIG_PATH"); envPath != "" {
		configPath = envPath
	}
	log.Println("Loading environment variables from .yaml file")
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
}

// GetEnv retrieves a string value from environment variables or returns a default value
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetEnvAsInt retrieves an integer value from environment variables or returns a default value
func GetEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Invalid integer for key %s: %v. Using default: %d", key, err, defaultValue)
		return defaultValue
	}

	return value
}
