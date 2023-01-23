package handler

import (
	"go-fiber/helper"
	"go-fiber/initializers"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllUser(c *fiber.Ctx) error {
	var users []models.User
	model := initializers.DB.Debug().Model(&users)

	result := helper.PG.With(model).Request(c.Request()).Response(&[]models.User{})

	return helper.JsonResponse(c, fiber.StatusOK, result)

}
