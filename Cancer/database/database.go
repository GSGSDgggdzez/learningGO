package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database struct holds the database connection
type Database struct {
	Db *gorm.DB
}

// DB is a global variable to store database instance
var DB Database

// ConnectDb initializes the database connection
func ConnectDb() {
	// Open a new SQLite database connection with the file name "api.db"
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	// Check if there was an error while opening the database
	if err != nil {
		log.Fatal("failed to connect database! \n", err.Error())
		os.Exit(2)
	}

	// Log successful database connection
	log.Println("Connection Opened to Database")
	// Set logger to show detailed SQL operations
	db.Logger = logger.Default.LogMode(logger.Info)

	// Assign the database connection to our global DB variable
	DB = Database{Db: db}
}
