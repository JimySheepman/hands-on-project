package routes

import (
	"shortest-distances/controllers"
	"shortest-distances/service"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, repo service.UserService) {
	route := app.Group("/api/v1")
	controller := controllers.NewController(repo)

	route.Get("/healthcheck", controller.Healthcheck)
	route.Post("/login", controller.Login)
	route.Post("/register", controller.Register)
}
