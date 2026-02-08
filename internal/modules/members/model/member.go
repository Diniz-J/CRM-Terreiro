package model

import "time"

type Member struct {
	ID             string     `db:"id" json:"id"`                                   //(UUID)
	NomeCompleto   string     `db:"nome" json:"nome"`                               //(string, obrigatório, min 3, max 255)
	NomeReligioso  *string    `db:"nome_religioso" json:"nome_religioso,omitempty"` //(string, opcional, max 255)
	CPF            string     `db:"cpf" json:"cpf"`                                 //(string, único, validado)
	RG             *string    `db:"rg" json:"rg,omitempty"`
	DataNascimento time.Time  `db:"data_nascimento" json:"data_nascimento"`
	Sexo           string     `db:"sexo" json:"sexo"`
	Telefone       string     `db:"telefone" json:"telefone"` // (string, obrigatório, validado)
	Email          string     `db:"email" json:"email"`       // (string, único, validado)
	Endereco       Endereco   `db:"endereco" json:"endereco"`
	Cargo          string     `db:"cargo" json:"cargo"`
	Status         string     `db:"status" json:"status"`
	DataIniciacao  *time.Time `db:"odun" json:"odun,omitempty"`
	Observacoes    *string    `db:"observacoes" json:"observacoes,omitempty"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at,omitempty"` //(timestamp, nullable - soft delete)
}

type Endereco struct {
	Rua         *string `db:"rua" json:"rua,omitempty"`
	Numero      *string `db:"numero" json:"numero,omitempty"`
	Complemento *string `db:"complemento" json:"complemento,omitempty"`
	Bairro      *string `db:"bairro" json:"bairro,omitempty"`
	Cidade      *string `db:"cidade" json:"cidade,omitempty"`
	Estado      *string `db:"estado" json:"estado,omitempty"`
	CEP         *string `db:"cep" json:"cep,omitempty"`
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
