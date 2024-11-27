package main

import (
	"log"

	"developer-allocation-system/pkg/db"
	"developer-allocation-system/pkg/models"
	"developer-allocation-system/pkg/utils"
)

func main() {
	// Load configuration
	config := utils.LoadConfig()

	// Connect to the database
	dbClient, err := db.NewDatabase(config.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	log.Println("Running database migrations...")
	err = dbClient.AutoMigrate(
		&models.User{},
		&models.Developer{},
		&models.Task{},
	)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully!")
}
