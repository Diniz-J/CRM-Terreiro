package handler

import (
	"errors"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/service"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(service *service.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) handleServiceError(c *fiber.Ctx, err error) error {
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

	event, err := h.service.CreateEvent(c.Context(), input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func (h *EventHandler) GetEventByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	event, err := h.service.GetEventByID(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func (h *EventHandler) ListEvents(c *fiber.Ctx) error {
	events, err := h.service.ListEvents(c.Context())
	if err != nil {
		return h.handleServiceError(c, err)
	}
	if events == nil {
		events = []model.Event{}
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h *EventHandler) UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	var input service.EventInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	event, err := h.service.UpdateEvent(c.Context(), id, input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func (h *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	err := h.service.DeleteEvent(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
