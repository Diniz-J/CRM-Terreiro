package model

import "time"

type Event struct {
	ID          string     `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description *string    `db:"description" json:"description,omitempty"`
	EventType   string     `db:"event_type" json:"event_type"`
	EventStatus string     `db:"event_status" json:"event_status"`
	Date        time.Time  `db:"date" json:"date"`
	Location    *string    `db:"location" json:"location,omitempty"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

const (
	EventTypeGira   = "Gira"
	EventTypeFuncao = "Função"

	EventStatusScheduled = "Agendado"
	EventStatusCancelled = "Cancelado"
	EventStatusCompleted = "Concluído"
)
