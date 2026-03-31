package auth

import "time"

type Credentials struct {
	ID           string    `db:"id" json:"id"`
	MemberID     string    `db:"member_id" json:"member_id"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
