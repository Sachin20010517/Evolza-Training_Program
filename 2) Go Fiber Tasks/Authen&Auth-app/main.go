package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// JWT secret key
var jwtSecret = []byte("secret")

// User model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT claims struct
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Dummy user for demonstration
var user = User{
	Username: "user1",
	Password: "password1",
}

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Public route: Login
	app.Post("/login", login)

	// Public route: Logout
	app.Get("/logout", logout)

	// JWT Middleware for protected routes
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtSecret,
		ContextKey: "user", // Ensure the context key is set for extracting user info
	}))

	// Protected route: Accessible only with valid JWT
	app.Get("/protected", protected)

	// Start server
	log.Fatal(app.Listen(":3000"))
}

// Login handler
func login(c *fiber.Ctx) error {
	var loginData User
	// Parse JSON body
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Check username and password
	if loginData.Username != user.Username || loginData.Password != user.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create JWT claims
	claims := Claims{
		Username: loginData.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not login"})
	}

	// Return token
	return c.JSON(fiber.Map{"token": tokenString})
}

// Protected handler
func protected(c *fiber.Ctx) error {
	// Get user info from JWT token
	userToken := c.Locals("user").(*jwt.Token)
	claims := userToken.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	return c.JSON(fiber.Map{"message": "Welcome " + username})
}

// Logout handler
func logout(c *fiber.Ctx) error {
	// Invalidate token logic can be added here (e.g., blacklisting tokens)
	return c.SendString("Logged out successfully")
}
