package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) {
	app.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Makassar",
		// Done: func(c *fiber.Ctx, logString []byte) {
		// 	if c.Response().StatusCode() != fiber.StatusOK {
		// 		reporter.SendToSlack(logString)
		// 	}
		// },
	}))
}
