package routes

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")
	route.Get("/posts", handler.GetAllPost)
	route.Get("/posts/:id", handler.GetSinglePost)
	route.Post("/posts", handler.CreatePost)
	route.Put("/posts/:id", handler.UpdatePost)
	route.Delete("/posts/:id", handler.DeletePost)
}
