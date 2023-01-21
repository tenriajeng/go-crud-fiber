package routes

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")

	auth := route.Group("/auth")
	auth.Post("/login", handler.Login)
}
