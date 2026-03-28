package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
)

type AttendanceRepository struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

const attendanceSelectColumns = `
	id, event_id, member_id, status, notes, marked_at, marked_by, created_at, updated_at, deleted_at
`

func scanAttendance(s scannable) (*model.Attendance, error) {
	var attendance model.Attendance

	err := s.Scan(
		&attendance.ID,
		&attendance.EventID,
		&attendance.MemberID,
		&attendance.Status,
		&attendance.Notes,
		&attendance.MarkedAt,
		&attendance.MarkedBy,
		&attendance.CreatedAt,
		&attendance.UpdatedAt,
		&attendance.DeletedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("attendance not found")
		}
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepository) MarkAttendance(ctx context.Context, attendance *model.Attendance) error {
	query := `
		INSERT INTO attendances (id, event_id, member_id, status, notes, marked_at, marked_by, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := r.db.ExecContext(ctx, query,
		attendance.ID,
		attendance.EventID,
		attendance.MemberID,
		attendance.Status,
		attendance.Notes,
		attendance.MarkedAt,
		attendance.MarkedBy)
	if err != nil {
		return fmt.Errorf("failed to mark attendance: %w", err)
	}
	return nil
}

func (r *AttendanceRepository) GetAttendanceByID(ctx context.Context, id string) (*model.Attendance, error) {
	query := `
		SELECT ` + attendanceSelectColumns + ` FROM attendances
		WHERE id = ? AND deleted_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, id)

	attendance := &model.Attendance{}

	attendance, err := scanAttendance(row)
	if err != nil {
		return nil, fmt.Errorf("failed to scan attendance: %w", err)
	}

	return attendance, nil
}
