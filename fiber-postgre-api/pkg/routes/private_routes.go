package routes

import (
	"fiber-postgre-api/app/controllers"
	"fiber-postgre-api/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)   // create a new book
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)    // update one book by ID
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
