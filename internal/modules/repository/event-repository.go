package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
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
		event.Location,
		event.EventType,
		event.EventStatus)

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

func (r *EventRepository) ListEvents(ctx context.Context) ([]model.Event, error) {
	query := `
		SELECT ` + eventSelectColumns + ` FROM events
		WHERE deleted_at IS NULL
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list events: %w", err)
	}
	defer rows.Close()

	var events []model.Event

	for rows.Next() {
		event, err := scanEvent(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}
		if event == nil {
			continue
		}

		events = append(events, *event)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate events: %w", err)
	}
	return events, nil
}

func (r *EventRepository) UpdateEvent(ctx context.Context, event *model.Event) error {
	query := `
		UPDATE events
		SET name = ?, date = ?, description = ?, location = ?, event_type = ?, event_status = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
	`

	result, err := r.db.ExecContext(ctx, query,
		event.Name,
		event.Date,
		event.Description,
		event.Location,
		event.EventType,
		event.EventStatus,
		event.ID)
	if err != nil {
		return fmt.Errorf("failed to update event: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("event not found")
	}

	return nil
}

func (r *EventRepository) DeleteEvent(ctx context.Context, id string) error {
	query := `
		UPDATE events
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL
		`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("event not found")
	}
	return nil
}

func (r *EventRepository) GetEventsByDate(ctx context.Context, date time.Time) ([]model.Event, error) {
	start := date.Truncate(24 * time.Hour)
	end := date.Add(24 * time.Hour)

	query := `
		SELECT ` + eventSelectColumns + ` FROM events
		WHERE date >= ? AND date < ? AND deleted_at IS NULL
	`
	rows, err := r.db.QueryContext(ctx, query, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to get events by date: %w", err)
	}
	defer rows.Close()

	var events []model.Event

	for rows.Next() {
		event, err := scanEvent(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan events: %w", err)
		}

		if event == nil {
			continue
		}

		events = append(events, *event)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate events: %w", err)
	}

	return events, nil

}
