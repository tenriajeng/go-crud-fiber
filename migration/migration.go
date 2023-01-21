package main

import (
	"go-fiber/initializers"
	"go-fiber/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabse()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
}
