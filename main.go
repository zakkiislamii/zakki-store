package main

import (
	"fmt"
	"os"
	"zakki-store/models"
	"zakki-store/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		os.Exit(1)
	}
	models.ConnectDB()
	err = models.DbMigrate(models.DB)
	if err != nil {
		fmt.Println("Error migrating database:", err)
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "55341" // Default to port 8080 if PORT environment variable is not set
	}
	routers.StartServer().Run(":" + port)
}
