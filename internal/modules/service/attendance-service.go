package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
	"github.com/google/uuid"
)

var (
	ErrAttendanceNotFound = errors.New("attendance not found")
	ErrMissingRequirement = errors.New("event_id, member_id and status are required")
)

type AttendanceRepository interface {
	MarkAttendance(ctx context.Context, attendance *model.Attendance) error
	GetAttendanceByID(ctx context.Context, id string) (*model.Attendance, error)
	ListAttendancesByEvent(ctx context.Context, eventID string) ([]model.Attendance, error)
	ListAttendancesByMember(ctx context.Context, memberID string) ([]model.Attendance, error)
	UpdateAttendance(ctx context.Context, attendance *model.Attendance) error
	DeleteAttendance(ctx context.Context, id string) error
}

type AttendanceService struct {
	repo AttendanceRepository
}

func NewAttendanceService(repo AttendanceRepository) *AttendanceService {
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

func (s *AttendanceService) MarkAttendance(ctx context.Context, input AttendanceInput) (*model.Attendance, error) {
	if input.EventID == "" || input.MemberID == "" || input.Status == "" {
		return nil, ErrMissingRequirement
	}

	now := time.Now()
	markedAt := now
	if input.MarkedAt != nil {
		markedAt = *input.MarkedAt
	}
	attendance := &model.Attendance{
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

	err := s.repo.MarkAttendance(ctx, attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to mark attendance: %w", err)
	}

	return attendance, nil
}
