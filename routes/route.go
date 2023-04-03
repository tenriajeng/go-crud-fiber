package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
}

func (r *Route) Url(app *fiber.App) {

	PostRoute(app)
	AuthRoute(app)
	UserRoute(app)
}
