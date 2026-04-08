package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Diniz-J/CRM-Terreiro/internal/modules/member"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepositoryInterface interface {
	CreateCredentials(ctx context.Context, c *Credentials) error
	GetCredentialByCPF(ctx context.Context, cpf string) (*Credentials, error)
	GetMemberIDByCPF(ctx context.Context, cpf string) (string, error)
}

type AuthService struct {
	repo   AuthRepositoryInterface
	mrepo  member.MemberRepositoryInterface
	secret string
}

func NewAuthService(repo AuthRepositoryInterface, mrepo member.MemberRepositoryInterface, secret string) *AuthService {
	return &AuthService{
		repo:   repo,
		mrepo:  mrepo,
		secret: secret,
	}
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

func gerarToken(memberID string, secret string) (string, error) {
	claims := JwtClaims{
		MemberID: memberID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	credential, err := s.repo.GetCredentialByCPF(ctx, req.CPF)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(credential.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, ErrSenhaInvalida
	}

	token, err := gerarToken(credential.MemberID, s.secret)
	if err != nil {
		return nil, fmt.Errorf("login: erro ao gerar token: %w", err)
	}

	membro, err := s.mrepo.FindByID(ctx, credential.MemberID)
	if err != nil {
		return nil, fmt.Errorf("login: erro ao buscar membro: %w", err)
	}

	return &LoginResponse{
		Token:         token,
		Nome:          membro.NomeCompleto,
		NomeReligioso: membro.NomeReligioso,
		Cargo:         membro.Cargo,
	}, nil
}
