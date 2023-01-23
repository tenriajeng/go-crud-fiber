package handler

import (
	"errors"
	"go-fiber/helper"
	"go-fiber/initializers"
	"go-fiber/middleware"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllPost(c *fiber.Ctx) error {
	var posts []models.Post
	model := initializers.DB.Debug().Preload("User").Model(&posts)

	result := helper.PG.With(model).Request(c.Request()).Response(&[]models.Post{})

	return helper.JsonResponse(c, fiber.StatusOK, result)
}

func CreatePost(c *fiber.Ctx) error {
	var body struct {
		Title string
		Body  string
	}

	err := c.BodyParser(&body)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	authenticatedUser := middleware.AuthenticatedUser
	post := models.Post{Title: body.Title, Body: body.Body, UserID: authenticatedUser.ID}
	result := initializers.DB.Debug().Create(&post)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, result.Error)
	}

	return helper.JsonResponse(c, fiber.StatusOK, "post success fully created")
}

func GetSinglePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post

	result := initializers.DB.Debug().Preload("User").First(&post, id)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusOK, errors.Is(result.Error, gorm.ErrRecordNotFound))
	}

	return helper.JsonResponse(c, fiber.StatusOK, post)
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		Title string
		Body  string
	}

	err := c.BodyParser(&body)
	if err != nil {
		return helper.JsonResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusOK, errors.Is(result.Error, gorm.ErrRecordNotFound))
	}

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusOK, "failed update data")
	}

	return helper.JsonResponse(c, fiber.StatusOK, "post success fully updated")
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusOK, errors.Is(result.Error, gorm.ErrRecordNotFound))
	}

	return helper.JsonResponse(c, fiber.StatusOK, "post success fully updated")
}
