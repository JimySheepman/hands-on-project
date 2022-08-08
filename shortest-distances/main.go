package main

import (
	"shortest-distances/configs"
	"shortest-distances/dto"
	"shortest-distances/middleware"
	"shortest-distances/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	err := configs.ConfigLoad("./.env")
	utils.IsFatalError(err)

	config := configs.AppConfig()

	app := fiber.New(config)

	middleware.Middlewares(app)

	dns := dto.PostgreConnectionDTO{}
	dns = dns.New()
}
