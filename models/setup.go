package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func ConnectDB() {
	// Load environment variables from file
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("Failed to read environment file:", err)
		return
	}
	log.Println("Successfully read environment file")

	// Get connection values from environment
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Construct connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open database connection
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Failed to open connection:", err)
		return
	}

	// Check database connection
	err = DB.Ping()
	if err != nil {
		log.Println("Failed to ping database:", err)
		return
	}
	log.Println("Database connection successful")
}
