package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/model"
)

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{}
}

// Scan
func scanMember(rows *sql.Rows) (*model.Member, error) {
	var member model.Member

	err := rows.Scan(
		&member.ID,
		&member.NomeCompleto,
		&member.NomeReligioso,
		&member.CPF,
		&member.RG,
		&member.DataNascimento,
		&member.Sexo,
		&member.Telefone,
		&member.Email,

		&member.Endereco.Rua,
		&member.Endereco.Numero,
		&member.Endereco.Complemento,
		&member.Endereco.Bairro,
		&member.Endereco.Cidade,
		&member.Endereco.Estado,
		&member.Endereco.CEP,

		&member.Cargo,
		&member.Status,
		&member.DataIniciacao,
		&member.Observacoes,
		&member.CreatedAt,
		&member.UpdatedAt,
		&member.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

func scanMemberRow(row *sql.Row) (*model.Member, error) {
	var member model.Member

	member.Endereco = model.Endereco{} //Garantia de não ser nil

	err := row.Scan(
		&member.ID,
		&member.NomeCompleto,
		&member.NomeReligioso,
		&member.CPF,
		&member.RG,
		&member.DataNascimento,
		&member.Sexo,
		&member.Telefone,
		&member.Email,

		&member.Endereco.Rua,
		&member.Endereco.Numero,
		&member.Endereco.Complemento,
		&member.Endereco.Bairro,
		&member.Endereco.Cidade,
		&member.Endereco.Estado,
		&member.Endereco.CEP,

		&member.Cargo,
		&member.Status,
		&member.DataIniciacao,
		&member.Observacoes,
		&member.CreatedAt,
		&member.UpdatedAt,
		&member.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &member, nil
}

// Save new member
func (r *MemberRepository) Save(ctx context.Context, member *model.Member) error {
	query := `
		INSERT INTO members (id, nome, nome_religioso, cpf, rg, data_nascimento, sexo, telefone, email,
		endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado,
		endereco_cep, cargo, status, odun, observacoes, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`

	_, err := r.db.ExecContext(ctx, query,
		member.ID,
		member.NomeCompleto,
		member.NomeReligioso,
		member.CPF,
		member.RG,
		member.DataNascimento,
		member.Sexo,
		member.Telefone,
		member.Email,

		member.Endereco.Rua,
		member.Endereco.Numero,
		member.Endereco.Complemento,
		member.Endereco.Bairro,
		member.Endereco.Cidade,
		member.Endereco.Estado,
		member.Endereco.CEP,

		member.Cargo,
		member.Status,
		member.DataIniciacao,
		member.Observacoes,
		member.CreatedAt,
		member.UpdatedAt)

	if err != nil {
		return fmt.Errorf("Failed to save member: %w", err)
	}

	return nil
}

func (r *MemberRepository) FindByID(ctx context.Context, id string) (*model.Member, error) {
	query := `
		SELECT id, nome, nome_religioso, cpf, rg, data_nascimento, sexo, telefone, email,
		endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado,
		endereco_cep, cargo, status, odun, observacoes, created_at, updated_at, deleted_at   FROM members 
		WHERE id = ? AND deleted_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, id)

	member, err := scanMemberRow(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Failed to find member by id(%s): %w", id, err)
	}
	return member, nil
}
