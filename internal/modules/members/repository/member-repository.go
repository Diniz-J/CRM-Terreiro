package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/model"
)

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

type scannable interface {
	Scan(dest ...any) error
}

const memberSelectColumns = `
	id, nome, nome_religioso, cpf, rg, data_nascimento, sexo, telefone, email,
		endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado,
		endereco_cep, cargo, status, odun, observacoes, created_at, updated_at, deleted_at
`

// Scan
func scanMember(s scannable) (*model.Member, error) {
	var member model.Member

	member.Endereco = model.Endereco{} //Garantia de não ser nil

	err := s.Scan(
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
		return fmt.Errorf("insert member: %w", err)
	}

	return nil
}

func (r *MemberRepository) FindByID(ctx context.Context, id string) (*model.Member, error) {
	query := `
		SELECT ` + memberSelectColumns + ` FROM members 
		WHERE id = ? AND deleted_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, id)

	member, err := scanMember(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("failed to found id(%s): %w", id, err)
		}
		return nil, fmt.Errorf("failed to scan id(%s): %w", id, err)
	}
	return member, nil
}

func (r *MemberRepository) SearchByName(ctx context.Context, nome string) ([]*model.Member, error) {
	query := `
		SELECT ` + memberSelectColumns + ` FROM members 
		WHERE nome LIKE ? AND deleted_at IS NULL`

	search := "%" + nome + "%"

	rows, err := r.db.QueryContext(ctx, query, search)
	if err != nil {
		return nil, fmt.Errorf("failed to search: %w", err)
	}
	defer rows.Close()

	var members []*model.Member

	for rows.Next() {
		member, err := scanMember(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}
		members = append(members, member)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)

	}
	return members, nil
}

func (r *MemberRepository) Update(ctx context.Context, member *model.Member) error {
	query := `
		UPDATE members
		SET nome = ?, nome_religioso = ?, cpf = ?, rg = ?, data_nascimento = ?, sexo = ?, telefone = ?,
		email = ?, endereco_rua = ?, endereco_numero = ?, endereco_complemento = ?, endereco_bairro = ?,
		endereco_cidade = ?, endereco_estado = ?, endereco_cep = ?, cargo = ?, status = ?, odun = ?, observacoes = ?
		WHERE id = ?
`
	result, err := r.db.ExecContext(ctx, query,
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
		member.UpdatedAt,

		member.ID)
	if err != nil {
		return fmt.Errorf("update member: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return nil
	}

	return nil
}

func (r *MemberRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM members
		WHERE id = ? AND deleted_at IS NULL`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete member: %w", err)
	}
	return nil
}
