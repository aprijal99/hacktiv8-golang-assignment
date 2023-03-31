package main

import (
	"book-service/database"
	"book-service/router"
)

func main() {
	database.StartDB()

	PORT := ":4000"
	router.StartServer().Run(PORT)
}
