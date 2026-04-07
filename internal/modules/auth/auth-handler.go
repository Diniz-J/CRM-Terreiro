package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) handleServiceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, ErrCredentialNotFound) || errors.Is(err, ErrSenhaInvalida) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "UNAUTHORIZED",
				"message": "credenciais invalidas",
			},
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    "INTERNAL_ERROR",
			"message": "internal server error",
		},
	})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid request body",
			},
		})
	}

	if err := h.service.Register(c.Context(), &req); err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    "CREATED",
		"message": "user successfully created",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid request body",
			},
		})
	}

	resp, err := h.service.Login(c.Context(), &req)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
