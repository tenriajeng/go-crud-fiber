package main

import (
	"go-fiber/config"
	"go-fiber/initializers"
	"go-fiber/middleware"
	"go-fiber/routes"
	"time"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabse()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: config.Config("APP_NAME"),
		// Prefork:           true,
		// EnablePrintRoutes: true,
		IdleTimeout: 5 * time.Second,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	middleware.Logger(app)

	middleware.Cache(app)

	routes.PostRoute(app)
	routes.AuthRoute(app)
	routes.UserRoute(app)

	app.Listen(":" + config.Config("PORT"))
}
