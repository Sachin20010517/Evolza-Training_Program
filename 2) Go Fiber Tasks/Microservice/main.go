package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// User represents a simple User model
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// to store users
var userList = map[string]User{}

func main() {

	// Create a new Fiber instance
	app := fiber.New()

	app.Post("/users", createUser)

	app.Get("/users/:id", getUserById)

	app.Get("/users", getAllUsers)

	app.Delete("/users/:id", deleteUser)

	app.Put("/users/:id", updateUser)

	app.Listen(":3000")

}

func createUser(c *fiber.Ctx) error {
	// Create a new User instance
	newUser := new(User)

	// Parse the request body into the User instance
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Store the user in the in-memory map
	userList[newUser.Id] = *newUser

	// Create a response message
	response := fiber.Map{
		"message": "A new user (" + newUser.Name + ") has been created",
	}

	// Return the created user as JSON
	// return c.Status(fiber.StatusCreated).JSON(response)
	responseString := fmt.Sprintf("%v", response) // Convert response to string
	return c.Status(fiber.StatusCreated).SendString(responseString)

}

func getUserById(c *fiber.Ctx) error {
	// Extract the user ID from the URL
	id := c.Params("id")

	// Retrieve the user from the map, save the user in in new variable
	userById, exists := userList[id]

	//If user can't find,
	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("User not found!. Try Again")
	}

	// Return the user as JSON
	return c.JSON(userById)

}

func getAllUsers(c *fiber.Ctx) error {
	//To retrieve all user  with all details at the same time
	//return c.JSON(userList)

	allUsers_slice := []User{}

	for _, user := range userList {
		allUsers_slice = append(allUsers_slice, user)
	}

	return c.JSON(allUsers_slice)

}

func deleteUser(c *fiber.Ctx) error {
	// Extract the user ID from the URL
	id := c.Params("id")

	_, exists := userList[id]

	//If user can't find,
	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("User not found!. Try Again")
	}

	delete(userList, id)

	// Return a success message
	return c.SendString("User " + id + " deleted successfully")
}

func updateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	//check whether user exist or not
	_, exists := userList[id]

	if !exists {
		return c.Status(fiber.StatusNotFound).SendString("User not found!. Try Again")
	}

	// Create a new User instance
	updatedUser := new(User)

	// Parse the request body into the User instance
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Update the user in the map
	updatedUser.Id = id
	userList[id] = *updatedUser

	// Return the updated user as JSON
	return c.JSON(updatedUser)

}
