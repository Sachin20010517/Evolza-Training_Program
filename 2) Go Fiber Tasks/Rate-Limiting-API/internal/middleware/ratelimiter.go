package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	// Create a rate limiter
	limiter := limiter.New(limiter.Config{
		Max:        10, // Maximum requests per minute
		Expiration: 1 * time.Minute,
	})

	// Return the rate limiting middleware
	return limiter
}
