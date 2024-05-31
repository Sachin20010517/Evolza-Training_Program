package main

import (
	"api-gateway/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Initialize routes
	routes.SetupRoutes(app)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
