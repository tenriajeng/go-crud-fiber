package main

import (
	"go-fiber/initializers"
	"go-fiber/middleware"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabse()
}

func main() {
	app := fiber.New()

	middleware.Logger(app)

	middleware.Cache(app)

	routes.PostRoute(app)
	routes.AuthRoute(app)
	routes.UserRoute(app)

	app.Listen(":3000")
}
