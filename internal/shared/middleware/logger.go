package middleware

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("[%s] %s %s", c.Method(), c.Path(), c.IP())
	err := c.Next()
	log.Printf("response time: %v", time.Since(start))
	return err
}

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
	}
	if !strings.HasPrefix(token, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
	}
	return c.Next()
}

func CorsMiddleware(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Method() == fiber.MethodOptions {
		return c.SendStatus(fiber.StatusOK)
	}

	return c.Next()
}

func MetricsMiddleware(c *fiber.Ctx) error {
	err := c.Next()
	log.Printf("[METRICS] %s %s - Status: %d", c.Method(), c.Path(), c.Response().StatusCode())
	return err
}

// NewRateLimiter retorna um middleware com mutex correto
type rateLimiter struct {
	mu     sync.Mutex
	counts map[string]int
}

func NewRateLimiter(maxRequests int, windowSeconds int) fiber.Handler {
	rl := &rateLimiter{counts: make(map[string]int)}
	return func(c *fiber.Ctx) error {
		ip := c.IP()

		rl.mu.Lock()
		rl.counts[ip]++
		count := rl.counts[ip]
		rl.mu.Unlock()

		if count > maxRequests {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{"error": "Rate limit exceeded"})
		}

		go func() {
			time.Sleep(time.Duration(windowSeconds) * time.Second)
			rl.mu.Lock()
			rl.counts[ip] = 0
			rl.mu.Unlock()
		}()

		return c.Next()
	}
}
