package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Define API routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the home page!")
	})

	app.Get("/api", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the API endpoint!")
	})

	// Serve a default empty favicon
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
