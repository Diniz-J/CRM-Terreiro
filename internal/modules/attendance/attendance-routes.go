package attendance

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App, h *AttendanceHandler) {
	app.Get("/events/:event_id/attendances", h.ListAttendancesByEvent)
	app.Get("/members/:member_id/attendances", h.ListAttendancesByMember)
	app.Post("/attendances", h.MarkAttendance)
	app.Get("/attendances/:id", h.GetAttendanceByID)
	app.Put("/attendances/:id", h.UpdateAttendance)
	app.Delete("/attendances/:id", h.DeleteAttendance)
}
