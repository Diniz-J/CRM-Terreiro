package attendance

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	service *AttendanceService
}

func NewAttendanceHandler(service *AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: service}
}

func (h *AttendanceHandler) handleServiceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, ErrAttendanceNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "NOT_FOUND",
				"message": err.Error(),
			},
		})
	}
	if errors.Is(err, ErrMissingRequirement) {
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

func (h *AttendanceHandler) MarkAttendance(c *fiber.Ctx) error {
	var input AttendanceInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid request body",
			},
		})
	}

	a, err := h.service.MarkAttendance(c.Context(), input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(a)
}

func (h *AttendanceHandler) GetAttendanceByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "id is required",
			},
		})
	}

	a, err := h.service.GetAttendanceByID(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(a)
}

func (h *AttendanceHandler) ListAttendancesByEvent(c *fiber.Ctx) error {
	eventID := c.Params("event_id")
	if eventID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "event_id is required",
			},
		})
	}

	attendances, err := h.service.ListAttendancesByEvent(c.Context(), eventID)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (h *AttendanceHandler) ListAttendancesByMember(c *fiber.Ctx) error {
	memberID := c.Params("member_id")
	if memberID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "member_id is required",
			},
		})
	}

	attendances, err := h.service.ListAttendancesByMember(c.Context(), memberID)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(attendances)
}

func (h *AttendanceHandler) UpdateAttendance(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "missing id",
			},
		})
	}

	var input AttendanceInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	a, err := h.service.UpdateAttendance(c.Context(), id, input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(a)
}

func (h *AttendanceHandler) DeleteAttendance(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "missing id",
			},
		})
	}

	err := h.service.DeleteAttendance(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
