package member

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App, h *MemberHandler) {
	app.Get("/members", h.ListMembers)
	app.Post("/members", h.CreateMember)
	app.Get("/members/:id", h.GetMember)
	app.Put("/members/:id", h.UpdateMember)
	app.Delete("/members/:id", h.DeleteMember)
}
