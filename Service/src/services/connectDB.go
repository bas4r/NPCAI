package services

import (
	"fmt"
	"log"
	"os"

	"github.com/basarrcan/NPCAI/models"
	"github.com/basarrcan/NPCAI/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	var err error

	// Load config
	config, err := utils.LoadConfig(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to load config: %v\n", err)
	}
	// Connect to the database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s", config.DBHost, config.DBPort, config.DBUserName, config.DBUserPassword, config.DBName, config.DBOptions)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	autoMigrate(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	return db
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Unable automigrate database: %v\n", err)
	}
}
