package main

import (
	"os"

	"rest-api-fiber/database"
	"rest-api-fiber/middleware"
	"rest-api-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()
	app.Use(cors.New())
	app.Use(middleware.CustomeMiddleware())
	database.DBconn = database.InitDb()
	routes.RoutesInit(app)
	port, envExists := os.LookupEnv("PORT")
	if !envExists {
		port = "8080"
	}
	app.Listen(":" + port)

}
