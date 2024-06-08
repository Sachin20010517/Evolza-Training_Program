package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New(logger.Config{
		Format: "${time} - ${method} ${path} - ${status}\n",
	}))

	// Error handling middleware
	app.Use(func(c *fiber.Ctx) error {
		if err := c.Next(); err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
			return nil
		}
		return nil
	})

	// Define routes for incoming API requests
	app.All("/api/v1/users/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:3001"+c.OriginalURL()[len("/api/v1"):])
	})

	app.All("/api/v1/products/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:3002"+c.OriginalURL()[len("/api/v1"):])
	})

	app.All("/api/v1/orders/*", func(c *fiber.Ctx) error {
		return proxy.Do(c, "http://localhost:3003"+c.OriginalURL()[len("/api/v1"):])
	})

	// Chatbot endpoint
	app.Post("/api/v1/chatbot", func(c *fiber.Ctx) error {
		type Request struct {
			Message string `json:"message"`
		}
		type Response struct {
			Reply string `json:"reply"`
		}

		req := new(Request)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		// Simple chatbot logic
		var reply string
		switch req.Message {
		case "hello":
			reply = "Hello! How can I assist you today?"
		case "bye":
			reply = "Goodbye! Have a great day!"
		default:
			reply = "I'm sorry, I didn't understand that."
		}

		return c.JSON(Response{Reply: reply})
	})

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(":" + port))
}
