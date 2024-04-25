package main

import (
	"zakki-store/models"
	"zakki-store/routers"
)

func main() {
	models.ConnectDB()
	routers.StartServer().Run(":8080")
}
