package routes

import (
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")

	auth := route.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/singup", handler.SingUp)
	auth.Get("/validate", handler.Validate)
}
