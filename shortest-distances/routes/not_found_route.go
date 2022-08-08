package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(app *fiber.App) {
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})
}
