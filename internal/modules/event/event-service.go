package event

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrEventNotFound         = errors.New("event not found")
	ErrInvalidDate           = errors.New("invalid date format, expected YYYY-MM-DD")
	ErrMissingRequiredFields = errors.New("name, event_type and event_status are required")
)

type EventRepositoryInterface interface {
	CreateEvent(ctx context.Context, e *Event) error
	GetEventByID(ctx context.Context, id string) (*Event, error)
	ListEvents(ctx context.Context) ([]Event, error)
	UpdateEvent(ctx context.Context, e *Event) error
	DeleteEvent(ctx context.Context, id string) error
	GetEventsByDate(ctx context.Context, date time.Time) ([]Event, error)
}

type EventService struct {
	repo EventRepositoryInterface
}

func NewEventService(repo EventRepositoryInterface) *EventService {
	return &EventService{repo: repo}
}

type EventInput struct {
	Name        string  `json:"name"`
	Date        string  `json:"date"`
	Description *string `json:"description"`
	EventType   string  `json:"event_type"`
	EventStatus string  `json:"event_status"`
	Location    *string `json:"location"`
}

func (s *EventService) CreateEvent(ctx context.Context, input EventInput) (*Event, error) {
	if input.Name == "" || input.EventType == "" || input.EventStatus == "" {
		return nil, ErrMissingRequiredFields
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidDate, err)
	}

	now := time.Now()
	e := &Event{
		ID:          uuid.New().String(),
		Name:        input.Name,
		Date:        date,
		Description: input.Description,
		EventType:   input.EventType,
		EventStatus: input.EventStatus,
		Location:    input.Location,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := s.repo.CreateEvent(ctx, e); err != nil {
		return nil, fmt.Errorf("failed to create event: %w", err)
	}

	return e, nil
}

func (s *EventService) GetEventByID(ctx context.Context, id string) (*Event, error) {
	e, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get event by id: %w", err)
	}

	if e == nil {
		return nil, ErrEventNotFound
	}
	return e, nil
}

func (s *EventService) ListEvents(ctx context.Context) ([]Event, error) {
	events, err := s.repo.ListEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list events: %w", err)
	}
	return events, nil
}

func (s *EventService) UpdateEvent(ctx context.Context, id string, input EventInput) (*Event, error) {
	existing, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find event to update: %w", err)
	}

	if existing == nil {
		return nil, ErrEventNotFound
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidDate, err)
	}

	if input.Name == "" || input.EventType == "" || input.EventStatus == "" {
		return nil, ErrMissingRequiredFields
	}

	e := &Event{
		ID:          existing.ID,
		Name:        input.Name,
		Date:        date,
		Description: input.Description,
		EventType:   input.EventType,
		EventStatus: input.EventStatus,
		Location:    input.Location,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.UpdateEvent(ctx, e); err != nil {
		return nil, fmt.Errorf("failed to update event: %w", err)
	}

	return e, nil
}

func (s *EventService) DeleteEvent(ctx context.Context, id string) error {
	existing, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find event to delete: %w", err)
	}
	if existing == nil {
		return ErrEventNotFound
	}

	if err := s.repo.DeleteEvent(ctx, id); err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}
	return nil
}
