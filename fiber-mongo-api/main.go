package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	app.Use(logger.New())

	routes.UserRote(app)

	app.Listen(":8080")
}
