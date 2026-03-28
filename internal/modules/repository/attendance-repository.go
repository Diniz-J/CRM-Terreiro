package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
)

type AttendanceRepository struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
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
