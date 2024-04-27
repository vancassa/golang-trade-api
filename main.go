package main

import (
	"log"
	"os"
	"trade-api/database"
	"trade-api/router"
)


func main() {
	database.StartDB()

	// Use `PORT` provided in environment or default to 3000
  var port = envPortOr("3000")

  log.Fatal(router.StartApp().Run(port))
}

// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func envPortOr(port string) string {
  // If `PORT` variable in environment exists, return it
  if envPort := os.Getenv("PORT"); envPort != "" {
    return ":" + envPort
  }
  // Otherwise, return the value of `port` variable from function argument
  return ":" + port
}