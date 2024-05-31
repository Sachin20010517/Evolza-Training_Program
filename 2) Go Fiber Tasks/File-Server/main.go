package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {

	app := fiber.New()

	// Serve static files from the public directory
	app.Static("/", "./public")

	// Middleware for basic authentication for the private directory
	app.Use("/private", basicauth.New(basicauth.Config{
		Users: map[string]string{
			"user": "password", // Replace with your username and password
		},
	}))

	// Route to serve files from the private directory
	app.Static("/private", "./private")

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

/*
What Is Middleware?

The Middleware as a definition can be considered like a broker(or a Bridge) which exist between two things to
facilitate communication and data processing.

Most Common Uses of Middleware :-
                  Authentication and Authorization
                  Error Handling
                  Logging and Monitoring
                  Request Parsing and Validation
                  Compression and Caching
                  Routing and URL Rewriting
*/
