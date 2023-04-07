package main

import (
	"todolist/database"
	"todolist/router"
)

func main() {
	database.StartDB()

	PORT := ":4000"
	router.StartServer().Run(PORT)
}
