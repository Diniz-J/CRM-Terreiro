package handler

import (
	"errors"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/service"
	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	service *service.MemberService
}

func NewMemberHandler(service *service.MemberService) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) handleServiceError(c *fiber.Ctx, err error) error {
	if errors.Is(err, service.ErrMemberNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "NOT_FOUND",
				"message": err.Error(),
			},
		})
	}
	if errors.Is(err, service.ErrInvalidCPF) ||
		errors.Is(err, service.ErrInvalidEmail) ||
		errors.Is(err, service.ErrInvalidPhone) ||
		errors.Is(err, service.ErrDuplicateCPF) ||
		errors.Is(err, service.ErrDuplicateEmail) {
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

func (h *MemberHandler) CreateMember(c *fiber.Ctx) error {
	var input service.MemberInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	member, err := h.service.CreateMember(c.Context(), input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(member)
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

	var input service.MemberInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.Map{
				"code":    "BAD_REQUEST",
				"message": "invalid body",
			},
		})
	}

	member, err := h.service.UpdateMember(c.Context(), id, input)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(member)
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

	member, err := h.service.GetMember(c.Context(), id)
	if err != nil {
		return h.handleServiceError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(member)
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
