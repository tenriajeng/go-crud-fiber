package handler

import (
	"errors"
	"go-fiber/initializers"
	"go-fiber/middleware"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllPost(c *fiber.Ctx) error {
	var posts []models.Post
	model := initializers.DB.Preload("User").Model(&posts)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": initializers.PG.With(model).Request(c.Request()).Response(&[]models.Post{}),
	})
}

func CreatePost(c *fiber.Ctx) error {
	var body struct {
		Title string
		Body  string
	}

	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	authenticatedUser := middleware.AuthenticatedUser

	post := models.Post{Title: body.Title, Body: body.Body, UserID: authenticatedUser.ID}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.JSON(fiber.Map{
		"data": post,
	})
}

func GetSinglePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post

	result := initializers.DB.Preload("User").First(&post, id)

	if result.Error != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": errors.Is(result.Error, gorm.ErrRecordNotFound),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": post,
	})
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var body struct {
		Title string
		Body  string
	}

	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": errors.Is(result.Error, gorm.ErrRecordNotFound),
		})
	}

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": "failed update data",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": post,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": errors.Is(result.Error, gorm.ErrRecordNotFound),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "post completely deleted",
	})
}
