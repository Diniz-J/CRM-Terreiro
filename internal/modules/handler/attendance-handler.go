package handler

import (
	"errors"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/service"
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	service *service.AttendanceService
}

func NewAttendanceHandler(service *service.AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{service: service}
}

func (h *AttendanceHandler) handleAttendanceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, service.ErrAttendanceNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "NOT_FOUND",
				"message": err.Error(),
			},
		})
	}
	if errors.Is(err, service.ErrMissingRequirement) {
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
	var input service.AttendanceInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid request body",
			},
		})
	}

	attendance, err := h.service.MarkAttendance(c.Context(), input)
	if err != nil {
		return h.handleAttendanceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(attendance)
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

	attendance, err := h.service.GetAttendanceByID(c.Context(), id)
	if err != nil {
		return h.handleAttendanceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(attendance)
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
		return h.handleAttendanceError(c, err)
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
		return h.handleAttendanceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(attendances)
}
