package handler

import (
	"go-fiber/initializers"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []models.User
	model := initializers.DB.Debug().Preload("Post").Model(&users)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": initializers.PG.With(model).Request(c.Request()).Response(&[]models.User{}),
	})
}
