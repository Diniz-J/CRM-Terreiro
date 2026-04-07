package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrSenhaInvalida = errors.New("senha invalida")

type JwtClaims struct {
	MemberID string `json:"member_id"`
	jwt.RegisteredClaims
}

type Credentials struct {
	ID           string    `db:"id" json:"id"`
	MemberID     string    `db:"member_id" json:"member_id"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type LoginRequest struct {
	CPF      string `json:"cpf" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token         string  `json:"token"`
	Nome          string  `json:"nome"`
	NomeReligioso *string `json:"nome_religioso"`
	Cargo         string  `json:"cargo"`
}

type RegisterRequest struct {
	CPF      string `json:"cpf" validate:"required"`
	Password string `json:"password" validate:"required"`
}
