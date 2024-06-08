package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var mu sync.Mutex

// Message defines the structure of the chat messages
type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func main() {
	app := fiber.New()

	app.Static("/", "./")

	// WebSocket route
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		mu.Lock()
		clients[c] = true
		mu.Unlock()

		defer func() {
			mu.Lock()
			delete(clients, c)
			mu.Unlock()
			c.Close()
		}()

		var msg Message
		for {
			if err := c.ReadJSON(&msg); err != nil {
				log.Println("read:", err)
				break
			}
			broadcast <- msg
		}
	}))

	// Start listening for incoming chat messages
	go handleMessages()

	log.Fatal(app.Listen(":3000"))
}

func handleMessages() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
