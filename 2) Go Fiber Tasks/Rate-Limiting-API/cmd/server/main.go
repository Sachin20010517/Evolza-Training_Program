package main

import (
	"Rate-Limiting-API/Rate-Limiting-API/internal/config"
	"Rate-Limiting-API/Rate-Limiting-API/internal/middleware"

	"Rate-Limiting-API/Rate-Limiting-API/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create a new Fiber instance
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(middleware.RateLimiter())

	// Setup API routes
	routes.Setup(app)

	// Start the server
	app.Listen(cfg.Server.Port)
}
