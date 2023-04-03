package handler

import (
	"go-fiber/config"
	"go-fiber/helper"
	"go-fiber/initializers"
	"go-fiber/models"
	"go-fiber/validation"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
}

func (h *AuthHandler) SingUp(c *fiber.Ctx) error {

	newUser := new(models.User)

	err := c.BodyParser(newUser)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	errors := validation.ValidateStruct(*newUser)
	if errors != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, errors)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, "Failed hash password")
	}

	user := models.User{Email: newUser.Email, Password: string(hash), Username: helper.RandomString(10)}

	result := initializers.DB.Debug().Create(&user)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusOK, "User already exists")
	}

	return helper.JsonResponse(c, fiber.StatusOK, user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	newUser := new(models.User)

	err := c.BodyParser(newUser)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	errors := validation.ValidateStruct(*newUser)
	if errors != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, errors)
	}

	var user models.User
	initializers.DB.Debug().First(&user, "email = ?", newUser.Email)

	if user.ID == 0 {
		return helper.JsonResponse(c, fiber.StatusBadRequest, "invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(newUser.Password))

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, "invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Config("SECRET")))

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, "Failed to create token")
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour * 30)

	c.Cookie(cookie)

	return helper.JsonResponse(c, fiber.StatusOK, user)
}

func (h *AuthHandler) Validate(c *fiber.Ctx) error {
	user := c.Cookies("Authorization")

	return helper.JsonResponse(c, fiber.StatusOK, user)
}
