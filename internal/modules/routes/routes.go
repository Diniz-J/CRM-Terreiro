package routes

import (
	"github.com/Diniz-J/teiunecc-admin/internal/modules/handler"
	"github.com/gofiber/fiber/v2"
)

func MemberRoutes(app *fiber.App, h *handler.MemberHandler) {
	app.Get("/members", h.ListMembers)
	app.Post("/members", h.CreateMember)
	app.Get("/members/:id", h.GetMember)
	app.Put("/members/:id", h.UpdateMember)
	app.Delete("/members/:id", h.DeleteMember)
}
