package attendance

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	ErrAttendanceNotFound = errors.New("attendance not found")
	ErrMissingRequirement = errors.New("event_id, member_id and status are required")
)

type AttendanceRepositoryInterface interface {
	MarkAttendance(ctx context.Context, a *Attendance) error
	GetAttendanceByID(ctx context.Context, id string) (*Attendance, error)
	ListAttendancesByEvent(ctx context.Context, eventID string) ([]Attendance, error)
	ListAttendancesByMember(ctx context.Context, memberID string) ([]Attendance, error)
	UpdateAttendance(ctx context.Context, a *Attendance) error
	DeleteAttendance(ctx context.Context, id string) error
}

type AttendanceService struct {
	repo AttendanceRepositoryInterface
}

func NewAttendanceService(repo AttendanceRepositoryInterface) *AttendanceService {
	return &AttendanceService{repo: repo}
}

type AttendanceInput struct {
	EventID  string     `json:"event_id"`
	MemberID string     `json:"member_id"`
	Status   string     `json:"status"`
	Notes    *string    `json:"notes"`
	MarkedAt *time.Time `json:"marked_at"`
	MarkedBy *string    `json:"marked_by"`
}

func (s *AttendanceService) MarkAttendance(ctx context.Context, input AttendanceInput) (*Attendance, error) {
	if input.EventID == "" || input.MemberID == "" || input.Status == "" {
		return nil, ErrMissingRequirement
	}

	now := time.Now()
	markedAt := now
	if input.MarkedAt != nil {
		markedAt = *input.MarkedAt
	}
	a := &Attendance{
		ID:        uuid.New().String(),
		EventID:   input.EventID,
		MemberID:  input.MemberID,
		Status:    input.Status,
		Notes:     input.Notes,
		MarkedAt:  markedAt,
		MarkedBy:  input.MarkedBy,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.MarkAttendance(ctx, a)
	if err != nil {
		return nil, fmt.Errorf("failed to mark attendance: %w", err)
	}

	return a, nil
}

func (s *AttendanceService) GetAttendanceByID(ctx context.Context, id string) (*Attendance, error) {
	a, err := s.repo.GetAttendanceByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get attendance: %w", err)
	}
	if a == nil {
		return nil, ErrAttendanceNotFound
	}
	return a, nil
}

func (s *AttendanceService) ListAttendancesByEvent(ctx context.Context, eventID string) ([]Attendance, error) {
	attendances, err := s.repo.ListAttendancesByEvent(ctx, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendances: %w", err)
	}
	return attendances, nil
}

func (s *AttendanceService) ListAttendancesByMember(ctx context.Context, memberID string) ([]Attendance, error) {
	attendances, err := s.repo.ListAttendancesByMember(ctx, memberID)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendances: %w", err)
	}
	return attendances, nil
}

func (s *AttendanceService) UpdateAttendance(ctx context.Context, id string, input AttendanceInput) (*Attendance, error) {
	existing, err := s.repo.GetAttendanceByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find attendance to update: %w", err)
	}

	if existing == nil {
		return nil, ErrAttendanceNotFound
	}
	if input.EventID == "" || input.MemberID == "" || input.Status == "" {
		return nil, ErrMissingRequirement
	}

	now := time.Now()
	markedAt := now
	if input.MarkedAt != nil {
		markedAt = *input.MarkedAt
	}
	a := &Attendance{
		ID:        existing.ID,
		EventID:   input.EventID,
		MemberID:  input.MemberID,
		Status:    input.Status,
		Notes:     input.Notes,
		MarkedAt:  markedAt,
		MarkedBy:  input.MarkedBy,
		CreatedAt: existing.CreatedAt,
		UpdatedAt: now,
	}

	if err := s.repo.UpdateAttendance(ctx, a); err != nil {
		return nil, fmt.Errorf("failed to update attendance: %w", err)
	}

	return a, nil
}

func (s *AttendanceService) DeleteAttendance(ctx context.Context, id string) error {
	existing, err := s.repo.GetAttendanceByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to find attendance to delete: %w", err)
	}
	if existing == nil {
		return ErrAttendanceNotFound
	}

	if err := s.repo.DeleteAttendance(ctx, id); err != nil {
		return fmt.Errorf("failed to delete attendance: %w", err)
	}
	return nil
}
