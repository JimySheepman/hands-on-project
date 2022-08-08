package configs

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func AppConfig() fiber.Config {
	return fiber.Config{
		AppName:       "dispatcher-api",
		ReadTimeout:   time.Second * time.Duration(30),
		WriteTimeout:  time.Second * time.Duration(30),
		CaseSensitive: true,
		BodyLimit:     64 * 1024 * 1024,
		Concurrency:   256 * 1024,
		IdleTimeout:   10 * time.Second,
	}
}
