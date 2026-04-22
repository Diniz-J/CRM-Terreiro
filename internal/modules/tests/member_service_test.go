package tests

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Diniz-J/CRM-Terreiro/internal/modules/member"
)

// mock manual da interface MemberRepositoryInterface
type mockMemberRepo struct {
	saveFn         func(ctx context.Context, m *member.Member) error
	findByIDFn     func(ctx context.Context, id string) (*member.Member, error)
	updateFn       func(ctx context.Context, m *member.Member) error
	deleteFn       func(ctx context.Context, id string) error
	findAllFn      func(ctx context.Context) ([]member.Member, error)
	searchByNameFn func(ctx context.Context, nome string) ([]*member.Member, error)
}

func (r *mockMemberRepo) Save(ctx context.Context, m *member.Member) error {
	return r.saveFn(ctx, m)
}
func (r *mockMemberRepo) FindByID(ctx context.Context, id string) (*member.Member, error) {
	return r.findByIDFn(ctx, id)
}
func (r *mockMemberRepo) Update(ctx context.Context, m *member.Member) error {
	return r.updateFn(ctx, m)
}
func (r *mockMemberRepo) Delete(ctx context.Context, id string) error {
	return r.deleteFn(ctx, id)
}
func (r *mockMemberRepo) FindAll(ctx context.Context) ([]member.Member, error) {
	return r.findAllFn(ctx)
}
func (r *mockMemberRepo) SearchByName(ctx context.Context, nome string) ([]*member.Member, error) {
	return r.searchByNameFn(ctx, nome)
}

// input valido reutilizavel nos testes
func validMemberInput() member.MemberInput {
	return member.MemberInput{
		NomeCompleto:   "Maria das Dores",
		CPF:            "529.982.247-25",
		DataNascimento: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Sexo:           member.SexoFem,
		Telefone:       "(11) 91234-5678",
		Email:          "maria@example.com",
		Cargo:          member.CargoMembro,
		Status:         member.StatusAtivo,
	}
}

func TestCreateMember_CPFInvalido(t *testing.T) {
	svc := member.NewMemberService(&mockMemberRepo{})

	input := validMemberInput()
	input.CPF = "000.000.000-00"

	_, err := svc.CreateMember(context.Background(), input)
	if !errors.Is(err, member.ErrInvalidCPF) {
		t.Errorf("esperado ErrInvalidCPF, obteve: %v", err)
	}
}

func TestCreateMember_EmailInvalido(t *testing.T) {
	svc := member.NewMemberService(&mockMemberRepo{})

	input := validMemberInput()
	input.Email = "nao-eh-um-email"

	_, err := svc.CreateMember(context.Background(), input)
	if !errors.Is(err, member.ErrInvalidEmail) {
		t.Errorf("esperado ErrInvalidEmail, obteve: %v", err)
	}
}

func TestCreateMember_TelefoneInvalido(t *testing.T) {
	svc := member.NewMemberService(&mockMemberRepo{})

	input := validMemberInput()
	input.Telefone = "1234"

	_, err := svc.CreateMember(context.Background(), input)
	if !errors.Is(err, member.ErrInvalidPhone) {
		t.Errorf("esperado ErrInvalidPhone, obteve: %v", err)
	}
}

func TestCreateMember_Sucesso(t *testing.T) {
	repo := &mockMemberRepo{
		saveFn: func(ctx context.Context, m *member.Member) error {
			return nil
		},
	}
	svc := member.NewMemberService(repo)

	m, err := svc.CreateMember(context.Background(), validMemberInput())
	if err != nil {
		t.Fatalf("esperado sucesso, obteve erro: %v", err)
	}
	if m == nil {
		t.Fatal("esperado membro retornado, obteve nil")
	}
	if m.Email != "maria@example.com" {
		t.Errorf("email incorreto: %s", m.Email)
	}
}

func TestCreateMember_CPFDuplicado(t *testing.T) {
	repo := &mockMemberRepo{
		saveFn: func(ctx context.Context, m *member.Member) error {
			return member.ErrDuplicateCPF
		},
	}
	svc := member.NewMemberService(repo)

	_, err := svc.CreateMember(context.Background(), validMemberInput())
	if !errors.Is(err, member.ErrDuplicateCPF) {
		t.Errorf("esperado ErrDuplicateCPF, obteve: %v", err)
	}
}

func TestGetMember_NaoEncontrado(t *testing.T) {
	repo := &mockMemberRepo{
		findByIDFn: func(ctx context.Context, id string) (*member.Member, error) {
			return nil, nil
		},
	}
	svc := member.NewMemberService(repo)

	_, err := svc.GetMember(context.Background(), "id-inexistente")
	if !errors.Is(err, member.ErrMemberNotFound) {
		t.Errorf("esperado ErrMemberNotFound, obteve: %v", err)
	}
}

func TestGetMember_Sucesso(t *testing.T) {
	esperado := &member.Member{ID: "abc-123", NomeCompleto: "Joao Silva"}
	repo := &mockMemberRepo{
		findByIDFn: func(ctx context.Context, id string) (*member.Member, error) {
			return esperado, nil
		},
	}
	svc := member.NewMemberService(repo)

	m, err := svc.GetMember(context.Background(), "abc-123")
	if err != nil {
		t.Fatalf("esperado sucesso, obteve erro: %v", err)
	}
	if m.ID != esperado.ID {
		t.Errorf("ID incorreto: %s", m.ID)
	}
}

func TestDeleteMember_NaoEncontrado(t *testing.T) {
	repo := &mockMemberRepo{
		findByIDFn: func(ctx context.Context, id string) (*member.Member, error) {
			return nil, nil
		},
	}
	svc := member.NewMemberService(repo)

	err := svc.DeleteMember(context.Background(), "id-inexistente")
	if !errors.Is(err, member.ErrMemberNotFound) {
		t.Errorf("esperado ErrMemberNotFound, obteve: %v", err)
	}
}

func TestDeleteMember_Sucesso(t *testing.T) {
	repo := &mockMemberRepo{
		findByIDFn: func(ctx context.Context, id string) (*member.Member, error) {
			return &member.Member{ID: "abc-123"}, nil
		},
		deleteFn: func(ctx context.Context, id string) error {
			return nil
		},
	}
	svc := member.NewMemberService(repo)

	err := svc.DeleteMember(context.Background(), "abc-123")
	if err != nil {
		t.Errorf("esperado sucesso, obteve erro: %v", err)
	}
}

func TestUpdateMember_CPFInvalido(t *testing.T) {
	repo := &mockMemberRepo{
		findByIDFn: func(ctx context.Context, id string) (*member.Member, error) {
			return &member.Member{ID: "abc-123"}, nil
		},
	}
	svc := member.NewMemberService(repo)

	input := validMemberInput()
	input.CPF = "000.000.000-00"

	_, err := svc.UpdateMember(context.Background(), "abc-123", input)
	if !errors.Is(err, member.ErrInvalidCPF) {
		t.Errorf("esperado ErrInvalidCPF, obteve: %v", err)
	}
}

func TestUpdateMember_MembroNaoEncontrado(t *testing.T) {
	repo := &mockMemberRepo{
		findByIDFn: func(ctx context.Context, id string) (*member.Member, error) {
			return nil, nil
		},
	}
	svc := member.NewMemberService(repo)

	_, err := svc.UpdateMember(context.Background(), "id-inexistente", validMemberInput())
	if !errors.Is(err, member.ErrMemberNotFound) {
		t.Errorf("esperado ErrMemberNotFound, obteve: %v", err)
	}
}

func TestListMembers_RetornaVazio(t *testing.T) {
	repo := &mockMemberRepo{
		findAllFn: func(ctx context.Context) ([]member.Member, error) {
			return []member.Member{}, nil
		},
	}
	svc := member.NewMemberService(repo)

	members, err := svc.ListMembers(context.Background())
	if err != nil {
		t.Fatalf("esperado sucesso, obteve erro: %v", err)
	}
	if len(members) != 0 {
		t.Errorf("esperado slice vazio, obteve %d membros", len(members))
	}
}

func TestListMembers_ErroNoRepo(t *testing.T) {
	erroEsperado := errors.New("falha no banco")
	repo := &mockMemberRepo{
		findAllFn: func(ctx context.Context) ([]member.Member, error) {
			return nil, erroEsperado
		},
	}
	svc := member.NewMemberService(repo)

	_, err := svc.ListMembers(context.Background())
	if err == nil {
		t.Fatal("esperado erro, obteve nil")
	}
}
