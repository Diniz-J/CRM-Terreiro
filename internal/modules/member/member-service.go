package member

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Diniz-J/CRM-Terreiro/internal/shared/validator"
	"github.com/google/uuid"
)

var (
	ErrMemberNotFound = errors.New("member not found")
	ErrInvalidCPF     = errors.New("invalid CPF")
	ErrInvalidEmail   = errors.New("invalid email")
	ErrInvalidPhone   = errors.New("invalid phone")
	ErrDuplicateCPF   = errors.New("CPF already registered")
	ErrDuplicateEmail = errors.New("email already registered")
)

type MemberRepositoryInterface interface {
	Save(ctx context.Context, m *Member) error
	FindByID(ctx context.Context, id string) (*Member, error)
	Update(ctx context.Context, m *Member) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]Member, error)
	SearchByName(ctx context.Context, nome string) ([]*Member, error)
}

type MemberService struct {
	repo MemberRepositoryInterface
}

func NewMemberService(repo MemberRepositoryInterface) *MemberService {
	return &MemberService{repo: repo}
}

type MemberInput struct {
	NomeCompleto   string     `json:"name"`
	NomeReligioso  *string    `json:"religious_name"`
	CPF            string     `json:"cpf"`
	RG             *string    `json:"rg"`
	DataNascimento time.Time  `json:"birth_date"`
	Sexo           string     `json:"gender"`
	Telefone       string     `json:"phone"`
	Email          string     `json:"email"`
	Cargo          string     `json:"role"`
	Status         string     `json:"status"`
	Odun           *time.Time `json:"odun"`
	Observacoes    *string    `json:"notes"`
	Rua            *string    `json:"address_street"`
	Numero         *string    `json:"address_number"`
	Complemento    *string    `json:"address_complement"`
	Bairro         *string    `json:"address_neighborhood"`
	Cidade         *string    `json:"address_city"`
	Estado         *string    `json:"address_state"`
	CEP            *string    `json:"address_zip_code"`
}

func (s *MemberService) CreateMember(ctx context.Context, input MemberInput) (*Member, error) {
	if !validator.CPF(input.CPF) {
		return nil, ErrInvalidCPF
	}

	if !validator.Email(input.Email) {
		return nil, ErrInvalidEmail
	}

	if !validator.Phone(input.Telefone) {
		return nil, ErrInvalidPhone
	}

	now := time.Now()
	m := &Member{
		ID:             uuid.New().String(),
		NomeCompleto:   input.NomeCompleto,
		NomeReligioso:  input.NomeReligioso,
		CPF:            input.CPF,
		RG:             input.RG,
		DataNascimento: input.DataNascimento,
		Sexo:           input.Sexo,
		Telefone:       input.Telefone,
		Email:          input.Email,
		Cargo:          input.Cargo,
		Status:         input.Status,
		Odun:           input.Odun,
		Observacoes:    input.Observacoes,
		Endereco: Endereco{
			Rua:         input.Rua,
			Numero:      input.Numero,
			Complemento: input.Complemento,
			Bairro:      input.Bairro,
			Cidade:      input.Cidade,
			Estado:      input.Estado,
			CEP:         input.CEP,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.repo.Save(ctx, m); err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	return m, nil
}

func (s *MemberService) GetMember(ctx context.Context, id string) (*Member, error) {
	m, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}

	if m == nil {
		return nil, ErrMemberNotFound
	}

	return m, nil
}

func (s *MemberService) UpdateMember(ctx context.Context, id string, input MemberInput) (*Member, error) {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find member to update: %w", err)
	}
	if existing == nil {
		return nil, ErrMemberNotFound
	}

	if !validator.CPF(input.CPF) {
		return nil, ErrInvalidCPF
	}

	if !validator.Email(input.Email) {
		return nil, ErrInvalidEmail
	}

	if !validator.Phone(input.Telefone) {
		return nil, ErrInvalidPhone
	}

	m := &Member{
		ID:             id,
		NomeCompleto:   input.NomeCompleto,
		NomeReligioso:  input.NomeReligioso,
		CPF:            input.CPF,
		RG:             input.RG,
		DataNascimento: input.DataNascimento,
		Sexo:           input.Sexo,
		Telefone:       input.Telefone,
		Email:          input.Email,
		Cargo:          input.Cargo,
		Status:         input.Status,
		Odun:           input.Odun,
		Observacoes:    input.Observacoes,
		Endereco: Endereco{
			Rua:         input.Rua,
			Numero:      input.Numero,
			Complemento: input.Complemento,
			Bairro:      input.Bairro,
			Cidade:      input.Cidade,
			Estado:      input.Estado,
			CEP:         input.CEP,
		},
	}
	if err := s.repo.Update(ctx, m); err != nil {
		return nil, fmt.Errorf("failed to update member: %w", err)
	}

	return s.repo.FindByID(ctx, id)
}

func (s *MemberService) DeleteMember(ctx context.Context, id string) error {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to search member: %w", err)
	}
	if existing == nil {
		return ErrMemberNotFound
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete member: %w", err)
	}
	return nil
}

func (s *MemberService) ListMembers(ctx context.Context) ([]Member, error) {
	members, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list all members: %w", err)
	}
	return members, nil
}

func (s *MemberService) SearchByName(ctx context.Context, nome string) ([]*Member, error) {
	members, err := s.repo.SearchByName(ctx, nome)
	if err != nil {
		return nil, fmt.Errorf("failed to search name: %w", err)
	}

	if members == nil {
		return nil, ErrMemberNotFound
	}
	return members, nil
}
