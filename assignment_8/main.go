package main

import (
	"product/database"
	"product/router"
)

func main() {
	database.StartDB()

	PORT := ":4000"
	router.StartServer().Run(PORT)
}
