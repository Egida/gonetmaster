package router

import (
	"github.com/edwardelton/gonetmaster/api/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/users", user.CreateUser)
}
