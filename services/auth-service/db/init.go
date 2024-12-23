package database

import (
	"fmt"
	"github.com/Mir00r/auth-service/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

// DB is the global database connection
var DB *gorm.DB

// InitDatabase initializes the database connection and runs migrations
//func InitDatabase(dsn string) {
//	var err error
//
//	// Open the database connection
//	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//
//	// Run auto-migrations
//	//err = DB.AutoMigrate(&entities.User{}, &entities.Token{}, &entities.PasswordResetToken{})
//	//if err != nil {
//	//	log.Fatalf("Failed to run migrations: %v", err)
//	//}
//
//	log.Println("Database connected and migrated successfully")
//}

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.DBName,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Run auto-migrations
	//err = DB.AutoMigrate(&entities.User{}, &entities.Token{})
	//if err != nil {
	//	log.Fatalf("Failed to run migrations: %v", err)
	//}

	log.Println("Database connected and migrated successfully")
}

func MigrationPath() string {
	// Construct the absolute path
	cwd, _ := os.Getwd()
	absoluteMigrationPath := filepath.Join(cwd, "db", "migrations")
	return absoluteMigrationPath
}

func DSN() string {
	// Construct DB DSN
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.DBName,
	)
	return dsn
}
