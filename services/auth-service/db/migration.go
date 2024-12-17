package database

import (
	"github.com/golang-migrate/migrate/v4"
	"log"
	"os"
	"strings"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(migrationPath string, dsn string) {
	// Convert the Windows path to a valid URL format
	if os.PathSeparator == '\\' {
		migrationPath = strings.ReplaceAll(migrationPath, "\\", "/")
		migrationPath = "file://" + migrationPath
	} else {
		migrationPath = "file://" + migrationPath
	}
	log.Printf("Using migration path: %s", migrationPath)

	m, err := migrate.New(migrationPath, dsn)
	if err != nil {
		log.Fatalf("Failed to initialize migrations: %v", err)
	}

	// Apply all up migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run up migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
