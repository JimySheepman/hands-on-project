package main

import (
	"notes-api-fibe/database"
	"notes-api-fibe/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// start a new fiber app
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Setup the Router
	router.SetupRoutes(app)

	// listen on port 3000
	app.Listen(":3000")

}
