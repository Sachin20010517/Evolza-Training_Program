package main

import "github.com/gofiber/fiber/v2"

func main() {

	app := fiber.New()

	// Define route for creating a post
	app.Post("/posts", createUser)

	app.Get("/get", getAllUsers)

	app.Put("/put", updateUser)

	app.Delete("/delete", deleteUser)

	app.Listen(":3000")

}

//Implement handler functions for posting
func createUser(c *fiber.Ctx) error {
	return c.SendString("A new user has been created successfully ")
}

func getAllUsers(c *fiber.Ctx) error {
	return c.SendString("Retrieving all users")
}

func updateUser(c *fiber.Ctx) error {
	return c.SendString("The user has been updated successfully ")
}

func deleteUser(c *fiber.Ctx) error {
	return c.SendString("The user has been deleted successfully ")
}
