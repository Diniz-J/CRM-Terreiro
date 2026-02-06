package model

import "time"

type Member struct {
	ID            string `db:"id" json:"id"`                         //(UUID)
	NomeCompleto  string `db:"nome" json:"nome"`                     //(string, obrigatório, min 3, max 255)
	NomeReligioso string `db:"nome_religioso" json:"nome_religioso"` //(string, opcional, max 255)
	CPF           string `db:"cpf" json:"cpf"`                       //(string, único, validado)
	//rg (string, opcional) rg nao existe mais
	DataNascimento string `db:"data_nascimento" json:"data_nascimento"` //(date)
	Sexo           string `db:"gênero" json:"gênero"`                   // (enum: masculino, feminino, outro)
	Telefone       string `db:"número" json:"número"`                   // (string, obrigatório, validado)
	Email          string `db:"email" json:"email"`                     // (string, único, validado)

	/////////endereco (objeto)
	//- rua (string)
	//- numero (string)
	//- complemento (string, opcional)
	//- bairro (string)
	//- cidade (string)
	//- estado (string, 2 chars)
	//- cep (string, validado)

	Cargo         string    `db:"cargo" json:"cargo"`             //enum: membro, iniciado, ogã, ekeji, sacerdote)
	Status        string    `db:"status" json:"status"`           //(enum: ativo, inativo, afastado)
	DataIniciacao string    `db:"odun" json:"odun"`               // (date, opcional)
	Observacoes   string    `db:"observacoes" json:"observaçoes"` //(text, opcional)
	CreatedAt     time.Time `db:"created_at" json:"created_at"`   //(timestamp)
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`   //(timestamp)
	DeletedAt     time.Time `db:"deleted_at" json:"deleted_at"`   //(timestamp, nullable - soft delete)
}

const (
	CargoOgan      = "Ogan"
	CargoEkeji     = "Ekeji"
	CargoMembro    = "Membro"
	CargoIniciado  = "Iniciado"
	CargoSacerdote = "Sacerdote"
	CargoPP        = "Pai Pequeno"
	CargoMP        = "Mãe Pequena"
)

const (
	StatusAtivo    = "Ativo"
	StatusInativo  = "Inativo"
	StatusAfastado = "Afastado"
)

const (
	SexoFem   = "Feminino"
	SexoMas   = "Masculino"
	SexoOutro = "Outro"
)
