package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User represents a simple User model
type User struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Age  int                `json:"age" bson:"age"`
}

// Global variable to hold the user collection
var userCollection *mongo.Collection

func main() {
	// Replace these with your actual credentials
	username := "<sachinayeshmantha>"
	password := "Ab2KA8z3OfR3vCQP"
	database := "microservice-1"

	// Create a new MongoDB client and connect to the server
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@cluster0.3vcczsb.mongodb.net/%s?retryWrites=true&w=majority&authMechanism=SCRAM-SHA-256",
			username, password, database))

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Set a context with a timeout to avoid blocking
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Ping the primary to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Select the database and collection
	userCollection = client.Database(database).Collection("users")

	// Create a new Fiber instance
	app := fiber.New()

	// Define routes
	app.Post("/users", createUser)
	app.Get("/users/:id", getUserById)
	app.Get("/users", getAllUsers)
	app.Delete("/users/:id", deleteUser)
	app.Put("/users/:id", updateUser)

	// Start the server on port 3000
	app.Listen(":3000")
}

// createUser handles the creation of a new user
func createUser(c *fiber.Ctx) error {
	// Create a new User instance
	newUser := new(User)

	// Parse the request body into the User instance
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Insert the new user into the database
	result, err := userCollection.InsertOne(context.Background(), newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Set the generated ID to the new user
	newUser.Id = result.InsertedID.(primitive.ObjectID)

	// Create a response message
	response := fiber.Map{
		"message": fmt.Sprintf("A new user (%s) has been created", newUser.Name),
		"user":    newUser,
	}

	// Return the created user as JSON
	return c.Status(fiber.StatusCreated).JSON(response)
}

// getUserById handles retrieving a user by their ID
func getUserById(c *fiber.Ctx) error {
	// Extract the user ID from the URL
	idParam := c.Params("id")

	// Convert the ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	// Retrieve the user from the database
	var user User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return the user as JSON
	return c.JSON(user)
}

// getAllUsers handles retrieving all users
func getAllUsers(c *fiber.Ctx) error {
	// Find all users in the database
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode each document
	var users []User
	for cursor.Next(context.Background()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	// Check for errors after iterating through the cursor
	if err := cursor.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Return the list of users as JSON
	return c.JSON(users)
}

// deleteUser handles deleting a user by their ID
func deleteUser(c *fiber.Ctx) error {
	// Extract the user ID from the URL
	idParam := c.Params("id")

	// Convert the ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	// Delete the user from the database
	result, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Return a success message
	return c.SendString("User deleted successfully")
}

// updateUser handles updating a user by their ID
func updateUser(c *fiber.Ctx) error {
	// Extract the user ID from the URL
	idParam := c.Params("id")

	// Convert the ID to an ObjectID
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	// Create a new User instance to hold the updated data
	updatedUser := new(User)

	// Parse the request body into the User instance
	if err := c.BodyParser(updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Create an update document
	update := bson.M{
		"$set": bson.M{
			"name": updatedUser.Name,
			"age":  updatedUser.Age,
		},
	}

	// Update the user in the database
	result, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Check if any document was matched and modified
	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).SendString("User not found")
	}

	// Set the updated ID to the user
	updatedUser.Id = id

	// Return the updated user as JSON
	return c.JSON(updatedUser)
}
