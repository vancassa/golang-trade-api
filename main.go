package main

import (
	"trade-api/database"
	"trade-api/router"
)


func main() {
	var PORT = ":3000"

	database.StartDB()

	router.StartApp().Run(PORT)
}
