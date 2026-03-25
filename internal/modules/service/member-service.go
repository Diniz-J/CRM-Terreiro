package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/model"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/validator"
	"github.com/google/uuid"
)

var (
	ErrMemberNotFound = errors.New("member not found")
	ErrInvalidCPF     = errors.New("invalid CPF")
	ErrInvalidEmail   = errors.New("invalid email")
	ErrInvalidPhone   = errors.New("invalid phone")
)

type MemberRepository interface {
	Save(ctx context.Context, member *model.Member) error
	FindByID(ctx context.Context, id string) (*model.Member, error)
	Update(ctx context.Context, member *model.Member) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]model.Member, error)
	SearchByName(ctx context.Context, nome string) ([]*model.Member, error)
}

type MemberService struct {
	repo MemberRepository
}

func NewMemberService(repo MemberRepository) *MemberService {
	return &MemberService{repo: repo}
}

type MemberInput struct {
	NomeCompleto   string
	NomeReligioso  *string
	CPF            string
	RG             *string
	DataNascimento time.Time
	Sexo           string
	Telefone       string
	Email          string
	Cargo          string
	Status         string
	Odun           *time.Time
	Observacoes    *string
	Rua            *string
	Numero         *string
	Complemento    *string
	Bairro         *string
	Cidade         *string
	Estado         *string
	CEP            *string
}

func (s *MemberService) CreateMember(ctx context.Context, input MemberInput) (*model.Member, error) {
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
	member := &model.Member{
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
		Endereco: model.Endereco{
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
	if err := s.repo.Save(ctx, member); err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	return member, nil

}

func (s *MemberService) GetMember(ctx context.Context, id string) (*model.Member, error) {
	member, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get member: %w", err)
	}

	if member == nil {
		return nil, ErrMemberNotFound
	}

	return member, nil
}

func (s *MemberService) UpdateMember(ctx context.Context, id string, input MemberInput) (*model.Member, error) {
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

	member := &model.Member{
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
		Endereco: model.Endereco{
			Rua:         input.Rua,
			Numero:      input.Numero,
			Complemento: input.Complemento,
			Bairro:      input.Bairro,
			Cidade:      input.Cidade,
			Estado:      input.Estado,
			CEP:         input.CEP,
		},
	}
	if err := s.repo.Update(ctx, member); err != nil {
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

func (s *MemberService) ListMembers(ctx context.Context) ([]model.Member, error) {
	member, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list all members: %w", err)
	}
	return member, nil
}

func (s *MemberService) SearchByName(ctx context.Context, nome string) ([]*model.Member, error) {
	member, err := s.repo.SearchByName(ctx, nome)
	if err != nil {
		return nil, fmt.Errorf("failed to search name: %w", err)
	}

	if member == nil {
		return nil, ErrMemberNotFound
	}
	return member, nil
}
