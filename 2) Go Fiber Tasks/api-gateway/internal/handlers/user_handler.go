package handlers

import (
	"api-gateway/internal/services"
	"io/ioutil"

	"github.com/gofiber/fiber/v2"
)

// HandleUserRequest processes the incoming user-related API requests
func HandleUserRequest(c *fiber.Ctx) error {
	// Extract the necessary details from the request
	requestPath := c.Params("*")

	// Forward the request to the user service
	response, err := services.ForwardUserRequest(c.Method(), requestPath, c.Body())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer response.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Set the response headers
	for key, values := range response.Header {
		for _, value := range values {
			c.Set(key, value)
		}
	}

	// Send back the response to the client
	return c.Status(response.StatusCode).Send(responseBody)
}
