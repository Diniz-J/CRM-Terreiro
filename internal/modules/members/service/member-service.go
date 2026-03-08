package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/model"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/repository"
	shared "github.com/Diniz-J/teiunecc-admin/internal/shared/validator"
	"github.com/google/uuid"
)

type MemberService struct {
	repo *repository.MemberRepository
}

func NewMemberService(repo *repository.MemberRepository) *MemberService {
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
	if !shared.CPF(input.CPF) {
		return nil, fmt.Errorf("invalid CPF")
	}

	if !shared.Email(input.Email) {
		return nil, fmt.Errorf("invalid email")
	}

	if !shared.Phone(input.Telefone) {
		return nil, fmt.Errorf("invalid phone")
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
		return nil, err
	}

	if member == nil {
		return nil, errors.New("member not found")
	}

	return member, nil
}

func (s *MemberService) UpdateMember(ctx context.Context, id string, input MemberInput) (*model.Member, error) {
	if !shared.CPF(input.CPF) {
		return nil, fmt.Errorf("invalid CPF")
	}

	if !shared.Email(input.Email) {
		return nil, fmt.Errorf("invalid email")
	}

	if !shared.Phone(input.Telefone) {
		return nil, fmt.Errorf("invalid phone")
	}

	now := time.Now()
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
		UpdatedAt: now,
	}
	if err := s.repo.Update(ctx, member); err != nil {
		return nil, fmt.Errorf("failed to update member: %w", err)
	}

	return s.repo.FindByID(ctx, id)
}
