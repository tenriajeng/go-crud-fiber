package routes

import (
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")

	var PostHandler handler.PostHandler

	route.Get("/posts", PostHandler.Index)
	route.Get("/posts/:id", PostHandler.Show)
	route.Post("/posts", middleware.Protected, PostHandler.Store)
	route.Put("/posts/:id", middleware.Protected, PostHandler.Update)
	route.Delete("/posts/:id", middleware.Protected, PostHandler.Delete)
}
