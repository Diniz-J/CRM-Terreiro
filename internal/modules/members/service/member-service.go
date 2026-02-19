package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/model"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/repository"
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
		return nil, fmt.Errorf("cpf inválido")
	}

	if !shared.Email(input.Email) {
		return nil, fmt.Errorf("email inválido")
	}

	if !shared.Phone(input.Telefone) {
		return nil, fmt.Errorf("telefone inválido")
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
