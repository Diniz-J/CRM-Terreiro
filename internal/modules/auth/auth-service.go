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
	// gera o hash da senha antes de salvar
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("register: erro ao gerar hash: %w", err)
	}

	credential := &Credentials{
		ID:           uuid.NewString(),
		MemberID:     req.CPF, // vamos ajustar isso logo
		PasswordHash: string(hash),
		IsActive:     true,
	}

	return s.repo.CreateCredentials(ctx, credential)
}

func (s *AuthService) Login(ctx context.Context, cpf string) (*LoginResponse, error) {
	return s.repo.GetCredentialByCPF(ctx, cpf)
}
