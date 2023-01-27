package handler

import (
	"go-fiber/config"
	"go-fiber/helper"
	"go-fiber/initializers"
	"go-fiber/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SingUp(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}

	if c.BodyParser(&body) != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, "Failed read body")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		return helper.JsonResponse(c, fiber.StatusInternalServerError, "Failed hash password")
	}

	user := models.User{Email: body.Email, Password: string(hash), Username: helper.RandomString(10)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		return helper.JsonResponse(c, fiber.StatusOK, "User already exists")
	}

	return helper.JsonResponse(c, fiber.StatusOK, user)
}

func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string
		Password string
	}

	if c.BodyParser(&body) != nil {
		return helper.JsonResponse(c, fiber.StatusBadRequest, "Failed read body")
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		return helper.JsonResponse(c, fiber.StatusBadRequest, "invalid email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

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

func Validate(c *fiber.Ctx) error {
	user := c.Cookies("Authorization")

	return helper.JsonResponse(c, fiber.StatusOK, user)
}
