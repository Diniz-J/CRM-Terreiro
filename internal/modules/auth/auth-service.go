package auth

import (
	"context"
	"fmt"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/member"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryInterface interface {
	CreateCredentials(ctx context.Context, c *Credentials) error
	GetCredentialByCPF(ctx context.Context, cpf string) (*Credentials, error)
	GetMemberIDByCPF(ctx context.Context, cpf string) (string, error)
}

type AuthService struct {
	repo  AuthRepositoryInterface
	mrepo member.MemberRepositoryInterface
}

func NewAuthService(repo AuthRepositoryInterface, mrepo member.MemberRepositoryInterface) *AuthService {
	return &AuthService{repo: repo,
		mrepo: mrepo}
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) error {
	memberID, err := s.repo.GetMemberIDByCPF(ctx, req.CPF)
	if err != nil {
		return fmt.Errorf("register: %w", err)
	}

	// gera o hash da senha antes de salvar
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("register: erro ao gerar hash: %w", err)
	}

	credential := &Credentials{
		ID:           uuid.NewString(),
		MemberID:     memberID,
		PasswordHash: string(hash),
		IsActive:     true,
	}

	return s.repo.CreateCredentials(ctx, credential)
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	credential, err := s.repo.GetCredentialByCPF(ctx, req.CPF)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(credential.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("senha invalida")
	}

	// TODO: gerar JWT e montar LoginResponse
	return nil, nil
}
