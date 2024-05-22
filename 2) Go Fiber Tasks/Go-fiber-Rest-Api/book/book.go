package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx) {
	c.Send("All Books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A Single Books")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Adds a new book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete a book")
}
