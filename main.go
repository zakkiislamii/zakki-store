package main

import (
	"os"
	"zakki-store/models"
	"zakki-store/routers"
)

func main() {
	models.ConnectDB()
	routers.StartServer().Run(":" + os.Getenv("PGPORT"))
}
