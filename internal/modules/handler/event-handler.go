package handler

import (
	"errors"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/service"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(service *service.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) handleServiceError (c *fiber.Ctx, err error) error {
	if errors.Is(err, service.ErrEventNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "NOT_FOUND",
				"message": err.Error(),
			},
		})
	}
	if errors.Is(err, service.ErrInvalidDate) ||
		errors.Is(err, service.ErrMissingRequiredFields) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": err.Error(),
			},
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error": fiber.Map{
			"code":    "INTERNAL_ERROR",
			"message": err.Error(),
		},
	})
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var input service.EventInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	member, err := h.service.CreateEvent(c.Context(), input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(member)
}

func (h *EventHandler)

