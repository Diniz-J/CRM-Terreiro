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

func EventRoutes(app *fiber.App, h *handler.EventHandler) {
	app.Get("/events", h.ListEvents)
	app.Post("/events", h.CreateEvent)
	app.Get("/events/:id", h.GetEventByID)
	app.Put("/events/:id", h.UpdateEvent)
	app.Delete("/events/:id", h.DeleteEvent)
}

func AttendanceRoutes(app *fiber.App, h *handler.AttendanceHandler) {
	app.Get("/events/:event_id/attendances", h.ListAttendancesByEvent)
	app.Get("/members/:member_id/attendances", h.ListAttendancesByMember)
	app.Post("/attendances", h.MarkAttendance)
	app.Get("/attendances/:id", h.GetAttendanceByID)
	app.Put("/attendances/:id", h.UpdateAttendance)
	app.Delete("/attendances/:id", h.DeleteAttendance)
}
