package model

import "time"

type Attendance struct {
	ID        string     `db:"id" json:"id"`
	EventID   string     `db:"event_id" json:"event_id"`
	MemberID  string     `db:"member_id" json:"member_id"`
	Date      time.Time  `db:"date" json:"date"`
	Status    string     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
