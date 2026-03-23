package model

import "time"

type Attendance struct {
	ID        string     `db:"id" json:"id"`
	EventID   string     `db:"event_id" json:"event_id"`
	MemberID  string     `db:"member_id" json:"member_id"`
	Notes     *string    `db:"notes" json:"notes,omitempty"`
	Status    string     `db:"status" json:"status"`
	MarkedAt  *time.Time `db:"marked_at" json:"marked_at,omitempty"`
	MarkedBy  *string    `db:"marked_by" json:"marked_by,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

type MarkAttendanceRequest struct {
	EventID  string  `json:"event_id" validate:"required"`
	MemberID string  `json:"member_id" validate:"required"`
	Status   string  `json:"status" validate:"required"`
	MarkedBy string  `json:"marked_by" validate:"required"`
	Notes    *string `json:"notes,omitempty"`
}

type AttendanceResponse struct {
	ID       string    `json:"id"`
	EventID  string    `json:"event_id"`
	MemberID string    `json:"member_id"`
	Status   string    `json:"status"`
	MarkedAt time.Time `json:"marked_at"`
	Notes    *string   `json:"notes,omitempty"`
	Message  string    `json:"message"` // Presença marcada com sucesso
}

const (
	AttendanceStatusPresent = "Presente"
	AttendanceStatusAbsent  = "Ausente"
	AttendanceStatusExcused = "Justificado"
	AttendanceStatusPending = "Pendente"
)
