package member

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
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
func scanMember(s scannable) (*Member, error) {
	var m Member

	m.Endereco = Endereco{} //Garantia de não ser nil

	err := s.Scan(
		&m.ID,
		&m.NomeCompleto,
		&m.NomeReligioso,
		&m.CPF,
		&m.RG,
		&m.DataNascimento,
		&m.Sexo,
		&m.Telefone,
		&m.Email,

		&m.Endereco.Rua,
		&m.Endereco.Numero,
		&m.Endereco.Complemento,
		&m.Endereco.Bairro,
		&m.Endereco.Cidade,
		&m.Endereco.Estado,
		&m.Endereco.CEP,

		&m.Cargo,
		&m.Status,
		&m.Odun,
		&m.Observacoes,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.DeletedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

// Save salva um novo membro
func (r *MemberRepository) Save(ctx context.Context, m *Member) error {
	query := `
		INSERT INTO members (id, nome, nome_religioso, cpf, rg, data_nascimento, sexo, telefone, email,
		endereco_rua, endereco_numero, endereco_complemento, endereco_bairro, endereco_cidade, endereco_estado,
		endereco_cep, cargo, status, odun, observacoes, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW())
		`

	_, err := r.db.ExecContext(ctx, query,
		m.ID,
		m.NomeCompleto,
		m.NomeReligioso,
		m.CPF,
		m.RG,
		m.DataNascimento,
		m.Sexo,
		m.Telefone,
		m.Email,

		m.Endereco.Rua,
		m.Endereco.Numero,
		m.Endereco.Complemento,
		m.Endereco.Bairro,
		m.Endereco.Cidade,
		m.Endereco.Estado,
		m.Endereco.CEP,

		m.Cargo,
		m.Status,
		m.Odun,
		m.Observacoes)

	if err != nil {
		if isDuplicateEntry(err, "cpf") {
			return ErrDuplicateCPF
		}
		if isDuplicateEntry(err, "email") {
			return ErrDuplicateEmail
		}
		return fmt.Errorf("insert member: %w", err)
	}

	return nil
}

func isDuplicateEntry(err error, field string) bool {
	return strings.Contains(err.Error(), "Error 1062") && strings.Contains(err.Error(), field)
}

func (r *MemberRepository) Update(ctx context.Context, m *Member) error {
	query := `
		UPDATE members
		SET nome = ?, nome_religioso = ?, cpf = ?, rg = ?, data_nascimento = ?, sexo = ?, telefone = ?,
		email = ?, endereco_rua = ?, endereco_numero = ?, endereco_complemento = ?, endereco_bairro = ?,
		endereco_cidade = ?, endereco_estado = ?, endereco_cep = ?, cargo = ?, status = ?, odun = ?,
		observacoes = ?, updated_at = NOW()
		WHERE id = ?
`
	result, err := r.db.ExecContext(ctx, query,
		m.NomeCompleto,
		m.NomeReligioso,
		m.CPF,
		m.RG,
		m.DataNascimento,
		m.Sexo,
		m.Telefone,
		m.Email,

		m.Endereco.Rua,
		m.Endereco.Numero,
		m.Endereco.Complemento,
		m.Endereco.Bairro,
		m.Endereco.Cidade,
		m.Endereco.Estado,
		m.Endereco.CEP,

		m.Cargo,
		m.Status,
		m.Odun,
		m.Observacoes,

		m.ID)
	if err != nil {
		return fmt.Errorf("update member: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *MemberRepository) Delete(ctx context.Context, id string) error {
	query := `
		UPDATE members
		SET deleted_at = NOW()
		WHERE id = ? AND deleted_at IS NULL`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete member: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *MemberRepository) FindByID(ctx context.Context, id string) (*Member, error) {
	query := `
		SELECT ` + memberSelectColumns + ` FROM members
		WHERE id = ? AND deleted_at IS NULL`

	row := r.db.QueryRowContext(ctx, query, id)

	m, err := scanMember(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to scan id(%s): %w", id, err)
	}
	return m, nil
}

func (r *MemberRepository) Count(ctx context.Context) (int, error) {
	var total int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM members WHERE deleted_at IS NULL`).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("count members: %w", err)
	}
	return total, nil
}

func (r *MemberRepository) FindAll(ctx context.Context, limit, offset int) ([]Member, error) {
	query := `
		SELECT ` + memberSelectColumns + ` FROM members
		WHERE deleted_at IS NULL
		ORDER BY nome ASC
		LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to find all members: %w", err)
	}
	defer rows.Close()

	members := make([]Member, 0)

	for rows.Next() {
		m, err := scanMember(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan member: %w", err)
		}
		members = append(members, *m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to iterate all members: %w", err)
	}
	return members, nil
}

func (r *MemberRepository) SearchByName(ctx context.Context, nome string) ([]*Member, error) {
	query := `
		SELECT ` + memberSelectColumns + ` FROM members
		WHERE nome LIKE ? AND deleted_at IS NULL`

	search := "%" + nome + "%"

	rows, err := r.db.QueryContext(ctx, query, search)
	if err != nil {
		return nil, fmt.Errorf("failed to search: %w", err)
	}
	defer rows.Close()

	var members []*Member

	for rows.Next() {
		m, err := scanMember(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan: %w", err)
		}
		members = append(members, m)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}
	return members, nil
}
