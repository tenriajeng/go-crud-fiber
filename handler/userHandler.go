package handler

import (
	"go-fiber/helper"
	"go-fiber/initializers"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func (h *UserHandler) Index(c *fiber.Ctx) error {
	var users []models.User
	model := initializers.DB.Debug().Model(&users)

	result := helper.Paginate.With(model).Request(c.Request()).Response(&[]models.User{})

	return helper.JsonResponse(c, fiber.StatusOK, result)
}
