package member

import "time"

type Member struct {
	ID             string     `db:"id" json:"id"`                                   //(UUID)
	NomeCompleto   string     `db:"nome" json:"name"`                               //(string, obrigatório, min 3, max 255)
	NomeReligioso  *string    `db:"nome_religioso" json:"religious_name,omitempty"` //(string, opcional, max 255)
	CPF            string     `db:"cpf" json:"cpf"`                                 //(string, único, validado)
	RG             *string    `db:"rg" json:"rg,omitempty"`
	DataNascimento time.Time  `db:"data_nascimento" json:"birth_date"`
	Sexo           string     `db:"sexo" json:"gender"`
	Telefone       string     `db:"telefone" json:"phone"` // (string, obrigatório, validado)
	Email          string     `db:"email" json:"email"`    // (string, único, validado)
	Endereco       Endereco   `db:"endereco" json:"address"`
	Cargo          string     `db:"cargo" json:"role"`
	Status         string     `db:"status" json:"status"`
	Odun           *time.Time `db:"odun" json:"odun,omitempty"`
	Observacoes    *string    `db:"observacoes" json:"notes,omitempty"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at,omitempty"` //(timestamp, nullable - soft delete)
}

type Endereco struct {
	Rua         *string `db:"rua" json:"street,omitempty"`
	Numero      *string `db:"numero" json:"number,omitempty"`
	Complemento *string `db:"complemento" json:"complement,omitempty"`
	Bairro      *string `db:"bairro" json:"neighborhood,omitempty"`
	Cidade      *string `db:"cidade" json:"city,omitempty"`
	Estado      *string `db:"estado" json:"state,omitempty"`
	CEP         *string `db:"cep" json:"zip_code,omitempty"`
}

const (
	CargoOgan      = "Ogan"
	CargoEkeji     = "Ekeji"
	CargoMembro    = "Membro"
	CargoIniciado  = "Iniciado"
	CargoSacerdote = "Sacerdote"
	CargoPP        = "Pai Pequeno"
	CargoMP        = "Mãe Pequena"

	StatusAtivo    = "Ativo"
	StatusInativo  = "Inativo"
	StatusAfastado = "Afastado"

	SexoFem   = "Feminino"
	SexoMas   = "Masculino"
	SexoOutro = "Outro"
)
