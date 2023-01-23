package helper

import "github.com/gofiber/fiber/v2"

func JsonResponse(c *fiber.Ctx, status int, data interface{}) error {

	return c.Status(status).JSON(fiber.Map{
		"data": data,
	})
}
