package middleware

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Diniz-J/CRM-Terreiro/internal/modules/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("[%s] %s %s", c.Method(), c.Path(), c.IP())
	err := c.Next()
	log.Printf("response time: %v", time.Since(start))
	return err
}

func NewAuthMiddleware(secret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token ausente"})
		}
		if !strings.HasPrefix(header, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "formato de token invalido"})
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		claims := &auth.JwtClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			// verifica se o algoritmo e o esperado
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritmo inesperado: %v", t.Header["alg"])
			}
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token invalido"})
		}

		// salva o member_id no contexto pra handlers usarem
		c.Locals("member_id", claims.MemberID)
		return c.Next()
	}
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
