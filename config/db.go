// config/db.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Get DSN for main application database
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")

	// Connect to the default database to check if task_management_db exists
	defaultDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to default database: %v", err)
	}

	// Check if the target database exists; if not, create it
	createDBIfNotExists(defaultDB)

	// Now connect to the actual task_management_db
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to task_management_db: %v", err)
	}
}

// createDBIfNotExists checks if the database exists and creates it if it doesn't
func createDBIfNotExists(db *gorm.DB) {
	var exists bool
	db.Raw("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = 'task_management_db')").Scan(&exists)

	if !exists {
		// Create the database
		if err := db.Exec("CREATE DATABASE task_management_db").Error; err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}
		fmt.Println("Database task_management_db created successfully.")
	} else {
		fmt.Println("Database task_management_db already exists.")
	}
}
