package model

import "time"

type Event struct {
	ID          string
	Name        string
	Description string
	EventType   string
	EventStatus string
	Date        time.Time
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

const (
	EventTypeGira   = "Gira"
	EventTypeFuncao = "Função"

	EventStatusScheduled = "Agendado"
	EventStatusCancelled = "Cancelado"
	EventStatusCompleted = "Concluído"
)
