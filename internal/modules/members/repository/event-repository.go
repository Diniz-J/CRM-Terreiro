package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/model"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

const eventSelectColumns = `
	id, name, date, description, location, event_type, event_status, created_at, updated_at, deleted_at
`

// Scan
func scanEvent(s scannable) (*model.Event, error) {
	var event model.Event

	err := s.Scan(
		&event.ID,
		&event.Name,
		&event.Date,
		&event.Description,
		&event.Location,
		&event.EventType,
		&event.EventStatus,
		&event.CreatedAt,
		&event.UpdatedAt,
		&event.DeletedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("event not found")
		}
		return nil, err
	}

	return &event, nil
}
func (r *EventRepository) CreateEvent(ctx context.Context, event *model.Event) error {
	query := `
		INSERT INTO events (id, name, date, description, location, event_type, event_status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
	`
	_, err := r.db.ExecContext(ctx, query,
		event.ID,
		event.Name,
		event.Date,
		event.Description,
		event.EventType,
		event.EventStatus,
		event.Location)

	if err != nil {
		return fmt.Errorf("create event: %w", err)
	}
	return nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
	query := `
		SELECT ` + eventSelectColumns + ` FROM events
		WHERE id = ? AND deleted_at IS NULL
	`
	row := r.db.QueryRowContext(ctx, query, id)

	event := &model.Event{}

	event, err := scanEvent(row)
	if err != nil {
		return nil, fmt.Errorf("failed to scan id(%s): %w", id, err)
	}
	return event, nil
}
