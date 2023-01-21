package handler

import (
	"errors"
	"go-fiber/initializers"
	"go-fiber/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetAllPost(c *fiber.Ctx) error {

	var posts []models.Post
	initializers.DB.Find(&posts)

	return c.JSON(posts)

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

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		return c.JSON(result.Error)
	}

	return c.JSON(fiber.Map{
		"data": post,
	})
}

func GetSinglePost(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		return c.JSON(fiber.Map{
			"data": errors.Is(result.Error, gorm.ErrRecordNotFound),
		})
	}

	return c.JSON(fiber.Map{
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
		return c.JSON(result.Error)
	}

	result = initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if result.Error != nil {
		return c.JSON(result.Error)
	}

	return c.JSON(fiber.Map{
		"data": post,
	})
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		return c.JSON(result.Error)
	}

	return c.JSON(fiber.Map{
		"data": "post completely deleted",
	})
}
