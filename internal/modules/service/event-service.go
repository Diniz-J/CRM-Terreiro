package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
	"github.com/google/uuid"
)

type EventRepository interface {
	CreateEvent(ctx context.Context, event *model.Event) error
	GetEventByID(ctx context.Context, id string) (*model.Event, error)
	ListEvents(ctx context.Context) ([]model.Event, error)
	UpdateEvent(ctx context.Context, event *model.Event) error
	DeleteEvent(ctx context.Context, id string) error
	GetEventsByDate(ctx context.Context, date time.Time) ([]model.Event, error)
}
type EventService struct {
	repo EventRepository
}

func NewEventService(repo EventRepository) *EventService {
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

func (s *EventService) CreateEvent(ctx context.Context, input EventInput) (*model.Event, error) {
	if input.Name == "" || input.EventType == "" || input.EventStatus == "" {
		return nil, fmt.Errorf("name, event_type and event_status are required")
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}

	now := time.Now()
	event := &model.Event{
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
	if err := s.repo.CreateEvent(ctx, event); err != nil {
		return nil, fmt.Errorf("failed to create event: %w", err)
	}

	return event, nil
}

func (s *EventService) GetEventByID(ctx context.Context, id string) (*model.Event, error) {
	event, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get event by id: %w", err)
	}

	if event == nil {
		return nil, fmt.Errorf("event not found")
	}
	return event, nil
}

func (s *EventService) ListEvents(ctx context.Context) ([]model.Event, error) {
	events, err := s.repo.ListEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list events: %w", err)
	}
	return events, nil
}

func (s *EventService) UpdateEvent(ctx context.Context, id string, input EventInput) (*model.Event, error) {
	existing, err := s.repo.GetEventByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find event to update: %w", err)
	}

	if existing == nil {
		return nil, fmt.Errorf("event not found")
	}

	date, err := time.Parse("2006-01-02", input.Date)
	if err != nil {
		return nil, fmt.Errorf("invalid date format, expected YYYY-MM-DD: %w", err)
	}

	if input.Name == "" || input.EventType == "" || input.EventStatus == "" {
		return nil, fmt.Errorf("name, event_type and event_status are required")
	}

	event := &model.Event{
		ID:          existing.ID,
		Name:        input.Name,
		Date:        date,
		Description: input.Description,
		EventType:   input.EventType,
		EventStatus: input.EventStatus,
		Location:    input.Location,
		UpdatedAt:   time.Now(),
	}

	if err := s.repo.UpdateEvent(ctx, event); err != nil {
		return nil, fmt.Errorf("failed to update event: %w", err)
	}

	return event, nil
}
