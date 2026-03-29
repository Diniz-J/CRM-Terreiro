package attendance

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type AttendanceRepository struct {
	db *sql.DB
}

func NewAttendanceRepository(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

type scannable interface {
	Scan(dest ...any) error
}

const attendanceSelectColumns = `
	id, event_id, member_id, status, notes, marked_at, marked_by, created_at, updated_at, deleted_at
`

func scanAttendance(s scannable) (*Attendance, error) {
	var a Attendance

	err := s.Scan(
		&a.ID,
		&a.EventID,
		&a.MemberID,
		&a.Status,
		&a.Notes,
		&a.MarkedAt,
		&a.MarkedBy,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *AttendanceRepository) MarkAttendance(ctx context.Context, a *Attendance) error {
	query := `
		INSERT INTO attendances (id, event_id, member_id, status, notes, marked_at, marked_by, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := r.db.ExecContext(ctx, query,
		a.ID,
		a.EventID,
		a.MemberID,
		a.Status,
		a.Notes,
		a.MarkedAt,
		a.MarkedBy)
	if err != nil {
		return fmt.Errorf("failed to mark attendance: %w", err)
	}
	return nil
}

func (r *AttendanceRepository) GetAttendanceByID(ctx context.Context, id string) (*Attendance, error) {
	query := `
		SELECT ` + attendanceSelectColumns + ` FROM attendances
		WHERE id = ? AND deleted_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, id)

	a, err := scanAttendance(row)
	if err != nil {
		return nil, fmt.Errorf("failed to scan attendance: %w", err)
	}

	return a, nil
}

func (r *AttendanceRepository) ListAttendancesByEvent(ctx context.Context, eventID string) ([]Attendance, error) {
	query := `
		SELECT ` + attendanceSelectColumns + ` FROM attendances
		WHERE event_id = ? AND deleted_at IS NULL`

	rows, err := r.db.QueryContext(ctx, query, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendances: %w", err)
	}
	defer rows.Close()

	var attendances []Attendance

	for rows.Next() {
		a, err := scanAttendance(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attendances: %w", err)
		}
		if a == nil {
			continue
		}

		attendances = append(attendances, *a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate attendances: %w", err)
	}
	return attendances, nil
}

func (r *AttendanceRepository) ListAttendancesByMember(ctx context.Context, memberID string) ([]Attendance, error) {
	query := `
		SELECT ` + attendanceSelectColumns + ` FROM attendances
		WHERE member_id = ? AND deleted_at IS NULL`

	rows, err := r.db.QueryContext(ctx, query, memberID)
	if err != nil {
		return nil, fmt.Errorf("failed to list attendances: %w", err)
	}
	defer rows.Close()

	var attendances []Attendance

	for rows.Next() {
		a, err := scanAttendance(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan attendances: %w", err)
		}
		if a == nil {
			continue
		}

		attendances = append(attendances, *a)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate attendances: %w", err)
	}
	return attendances, nil
}

func (r *AttendanceRepository) UpdateAttendance(ctx context.Context, a *Attendance) error {
	query := `
		UPDATE attendances
		SET status = ?, notes = ?, marked_at = ?, marked_by = ?, updated_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	result, err := r.db.ExecContext(ctx, query,
		a.Status,
		a.Notes,
		a.MarkedAt,
		a.MarkedBy,
		a.ID)

	if err != nil {
		return fmt.Errorf("failed to update attendance: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("attendance not found")
	}

	return nil
}

func (r *AttendanceRepository) DeleteAttendance(ctx context.Context, attendanceID string) error {
	query := `
		UPDATE attendances
		SET deleted_at = NOW()
		WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, attendanceID)
	if err != nil {
		return fmt.Errorf("failed to delete attendance: %w", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("attendance not found")
	}

	return nil
}
