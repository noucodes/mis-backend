package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noucodes/mis-backend/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/users", handlers.GetUsers)
}