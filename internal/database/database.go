package database

import (
	"fmt"
	"log"

	"github.com/lskeey/go-filehub/config"
	"github.com/lskeey/go-filehub/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection and runs auto-migration.
func Connect(cfg config.Config) {
	var err error

	// Data Source Name (DSN) for connecting to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	// Open a connection to the database
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successfully established.")

	// Auto-migrate the schema
	// This will create the 'users' and 'files' tables if they don't exist
	err = DB.AutoMigrate(
		&models.User{},
		&models.File{},
	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	log.Println("Database migrated successfully.")
}
