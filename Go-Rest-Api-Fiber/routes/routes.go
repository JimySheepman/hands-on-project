package main

import (
	"rest-api-fiber/controllers"

	"github.com/gofiber/fiber/v2"
)

func RotesInit(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SentString("Hello, World!")
	})
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/posts", controllers.GetPosts)
	v1.Post("/posts", controllers.CreatePost)
	v1.Get("/posts/:id", controllers.GetPostById)
	v1.Delete("/posts/:id", controllers.DeletePost)
	v1.Get("/replies", controllers.GetReplies)
	v1.Get("/replies/:id", controllers.GetRepliesByPostId)
	v1.Post("/replies", controllers.CreateReply)
}
