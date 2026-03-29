package event

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App, h *EventHandler) {
	app.Get("/events", h.ListEvents)
	app.Post("/events", h.CreateEvent)
	app.Get("/events/:id", h.GetEventByID)
	app.Put("/events/:id", h.UpdateEvent)
	app.Delete("/events/:id", h.DeleteEvent)
}
