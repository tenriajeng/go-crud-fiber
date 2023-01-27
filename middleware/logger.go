package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger(app *fiber.App) {

	config := logger.Config{
		Format:     "TIME:${time} PID:${pid} PORT:${port} STATUS:${status} - METHOD:${method} PATH:${path} LATENCY:${latency} URL:${host}${url}\n",
		TimeZone:   "Asia/Makassar",
		TimeFormat: "02-Jan-2006 15:04:05",
		// Done: func(c *fiber.Ctx, logString []byte) {
		// 	if c.Response().StatusCode() != fiber.StatusOK {
		// 		utils.SendToDiscord(logString)
		// 	}
		// },
	}

	app.Use(logger.New(config))
}
