package main

import (
	"log" // Import log package for logging

	"github.com/gofiber/fiber/v2"     // Import Fiber package
	"github.com/gofiber/websocket/v2" // Import WebSocket package for Fiber
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// WebSocket route
	app.Get("/ws", websocket.New(websocketHandler))

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

// websocketHandler handles WebSocket connections
func websocketHandler(c *websocket.Conn) {
	defer func() {
		// Recover from panic, if any
		if r := recover(); r != nil {
			log.Println("Recovered in websocketHandler:", r)
		}
	}()

	for {
		// Read message from the client
		mt, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break // Break the loop on read error
		}

		log.Printf("Received: %s", msg) // Log the received message

		// Echo the message back to the client
		err = c.WriteMessage(mt, msg)
		if err != nil {
			log.Println("Write error:", err)
			break // Break the loop on write error
		}
	}
}
