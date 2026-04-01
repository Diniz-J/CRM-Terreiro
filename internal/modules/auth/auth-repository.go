package auth

import (
	"context"
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateCredentials(ctx context.Context, c *Credentials) error {
	query := `
		INSERT INTO credentials (id, member_id, password_hash, is_active, created_at, updated_at) 
		VALUES (?, ?, ?, ?, NOW(), NOW())
		`

	_, err := r.db.ExecContext(ctx, query,
		c.ID,
		c.MemberID,
		c.PasswordHash,
		c.IsActive)
	if err != nil {
		return fmt.Errorf("create credentials: %w", err)
	}
	return nil
}

func (r *AuthRepository) GetCredentialByCPF(ctx context.Context, cpf string) (*Credentials, error) {
	query := `
		SELECT c.password_hash, c.member_id
		FROM credentials c
		JOIN members m ON c.member_id = m.id
		WHERE m.cpf = ?
	`

	c := &Credentials{}
	err := r.db.QueryRowContext(ctx, query, cpf).Scan(&c.PasswordHash, &c.MemberID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("credenciais nao encontradas para o CPF informado")
	}
	if err != nil {
		return nil, fmt.Errorf("get credential by cpf: %w", err)
	}
	return c, nil
}
