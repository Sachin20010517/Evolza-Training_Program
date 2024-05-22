package main

import "github.com/gofiber/fiber/v2"

func main() {

	// //create instance of fiber
	// app := fiber.New()

	// //create httpHandler
	// app.Get("/testApi", func(ctx *fiber.Ctx) error {
	// 	return ctx.Status(200).JSON(fiber.Map{
	// 		"success": true,
	// 		"message": "This is my first first go-fiber api project",
	// 	})
	// })

	// //listen on port
	// app.Listen(":3000")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World. Im Sachin")
	})

	app.Listen(":3000")
}
