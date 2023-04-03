package handler

import (
	"errors"
	"go-fiber/helper"
	"go-fiber/initializers"
	"go-fiber/middleware"
	"go-fiber/models"
	"go-fiber/validation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PostHandler struct {
}

func (h *PostHandler) Index(c *fiber.Ctx) error {
	var posts []models.Post
	model := initializers.DB.Debug().Preload("User").Model(&posts)

	result := helper.Paginate.With(model).Request(c.Request()).Response(&[]models.Post{})

	return helper.JsonResponse(c, fiber.StatusOK, result)
}

func (h *PostHandler) Store(c *fiber.Ctx) error {
	newPost := new(models.Post)

	err := c.BodyParser(newPost)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	errors := validation.ValidateStruct(*newPost)
	if errors != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, errors)
	}

	authenticatedUser := middleware.AuthenticatedUser

	post := models.Post{Title: newPost.Title, Body: newPost.Body, UserID: authenticatedUser.ID}
	result := initializers.DB.Debug().Create(&post)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, result.Error)
	}

	return helper.JsonResponse(c, fiber.StatusOK, "post success fully created")
}

func (h *PostHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")

	var post models.Post

	result := initializers.DB.Debug().Preload("User").First(&post, id)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, errors.Is(result.Error, gorm.ErrRecordNotFound))
	}

	return helper.JsonResponse(c, fiber.StatusOK, post)
}

func (h *PostHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	newPost := new(models.Post)

	err := c.BodyParser(newPost)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	errors := validation.ValidateStruct(*newPost)
	if errors != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, errors)
	}

	var post models.Post
	result := initializers.DB.Debug().First(&post, id)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	result = initializers.DB.Debug().Model(&post).Updates(models.Post{
		Title: newPost.Title,
		Body:  newPost.Body,
	})

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	return helper.JsonResponse(c, fiber.StatusOK, "post success fully updated")
}

func (h *PostHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	result := initializers.DB.Debug().Delete(&models.Post{}, id)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, result.Error.Error())
	}

	return helper.JsonResponse(c, fiber.StatusOK, "post success fully deleted")
}
