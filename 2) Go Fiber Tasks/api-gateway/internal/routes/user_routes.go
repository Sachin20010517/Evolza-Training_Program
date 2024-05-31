package routes

import (
	"api-gateway/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes defines all the routes for the API gateway
func SetupRoutes(app *fiber.App) {
	// Route for user-related API requests
	app.All("/api/v1/users/*", handlers.HandleUserRequest)
}
