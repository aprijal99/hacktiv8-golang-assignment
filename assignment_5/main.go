package main

import "book/router"

func main() {
	var PORT string = ":4000"

	router.StartServer().Run(PORT)
}
