package member

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	service *MemberService
}

func NewMemberHandler(service *MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) handleServiceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, ErrMemberNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "NOT_FOUND",
				"message": err.Error(),
			},
		})
	}
	if errors.Is(err, ErrInvalidCPF) ||
		errors.Is(err, ErrInvalidEmail) ||
		errors.Is(err, ErrInvalidPhone) ||
		errors.Is(err, ErrDuplicateCPF) ||
		errors.Is(err, ErrDuplicateEmail) {
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

func (h *MemberHandler) CreateMember(c *fiber.Ctx) error {
	var input MemberInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	m, err := h.service.CreateMember(c.Context(), input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(m)
}

func (h *MemberHandler) UpdateMember(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "missing id",
			},
		})
	}

	var input MemberInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	m, err := h.service.UpdateMember(c.Context(), id, input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(m)
}

func (h *MemberHandler) DeleteMember(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "missing id",
			},
		})
	}

	err := h.service.DeleteMember(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *MemberHandler) GetMember(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "missing id",
			},
		})
	}

	m, err := h.service.GetMember(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(m)
}

func (h *MemberHandler) ListMembers(c *fiber.Ctx) error {
	nome := c.Query("nome")

	if nome != "" {
		members, err := h.service.SearchByName(c.Context(), nome)
		if err != nil {
			return h.handleServiceError(c, err)
		}
		return c.Status(fiber.StatusOK).JSON(members)
	}
	members, err := h.service.ListMembers(c.Context())
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(members)
}
