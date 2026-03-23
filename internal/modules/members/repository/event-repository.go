package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/model"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *model.Event) error {
	query := `
		INSERT INTO events (id, name, date, description, location, event_type, event_status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`
	_, err := r.db.ExecContext(ctx, query,
		event.ID,
		event.Name,
		event.Description,
		event.EventType,
		event.EventStatus,
		event.Date,
		event.Location)

	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}
	return nil
}
