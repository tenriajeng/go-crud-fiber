package routes

import (
	"go-fiber/handler"
	"go-fiber/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App) {
	// routes
	route := app.Group("/api")
	var AuthHandler handler.AuthHandler

	auth := route.Group("/auth")
	auth.Post("/login", AuthHandler.Login)
	auth.Post("/singup", AuthHandler.SingUp)
	auth.Get("/validate", middleware.Protected, AuthHandler.Validate)
}
