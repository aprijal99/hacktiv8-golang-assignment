package main

import (
	"mygram/db"
	"mygram/router"
)

func main() {
	db.StartDB()

	PORT := ":4000"
	router.StartServer().Run(PORT)
}
