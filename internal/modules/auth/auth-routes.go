package auth

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App, h *AuthHandler) {
	app.Post("/auth/register", h.Register)
	app.Post("/auth/login", h.Login)
}
