package routes

import (
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")

	user := route.Group("/users")
	user.Get("/", middleware.Protected, handler.GetAllUser)
}
