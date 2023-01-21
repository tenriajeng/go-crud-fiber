package routes

import (
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func PostRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")
	route.Get("/posts", handler.GetAllPost)
	route.Get("/posts/:id", handler.GetSinglePost)
	route.Post("/posts", middleware.Protected, handler.CreatePost)
	route.Put("/posts/:id", middleware.Protected, handler.UpdatePost)
	route.Delete("/posts/:id", middleware.Protected, handler.DeletePost)
}
