package event

import (
	"errors"

	"github.com/Diniz-J/CRM-Terreiro/internal/shared/pagination"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	service *EventService
}

func NewEventHandler(service *EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (h *EventHandler) handleServiceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, ErrEventNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "NOT_FOUND",
				"message": err.Error(),
			},
		})
	}
	if errors.Is(err, ErrInvalidDate) ||
		errors.Is(err, ErrMissingRequiredFields) {
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
			"message": "internal server error",
		},
	})
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var input EventInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	e, err := h.service.CreateEvent(c.Context(), input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(e)
}

func (h *EventHandler) GetEventByID(c *fiber.Ctx) error {
	id := c.Params("id")

	e, err := h.service.GetEventByID(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(e)
}

func (h *EventHandler) ListEvents(c *fiber.Ctx) error {
	p := pagination.Normalize(c.QueryInt("page"), c.QueryInt("page_size"))

	result, err := h.service.ListEvents(c.Context(), p)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *EventHandler) UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")

	var input EventInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	e, err := h.service.UpdateEvent(c.Context(), id, input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(e)
}

func (h *EventHandler) DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.DeleteEvent(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
