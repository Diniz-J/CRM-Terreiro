# teiunecc-admin

API REST para gestão de terreiro de Candomblé/Umbanda, desenvolvida em Go com arquitetura em camadas.

## Pré-requisitos

- Go 1.21+
- Docker e Docker Compose (para rodar com Docker)
- MySQL 8.0 (para rodar sem Docker)

## Tecnologias

- Go 1.21+
- Fiber v2
- MySQL 8.0
- Docker

## Módulos implementados

| Módulo | Descrição |
|---|---|
| Auth | Registro e autenticação via JWT |
| Members | CRUD completo de membros do terreiro |
| Events | CRUD completo de eventos |
| Attendance | Registro e gestão de presença em eventos |

## Estrutura do projeto

```
teiunecc-admin/
├── cmd/
│   └── main.go
├── internal/
│   ├── modules/
│   │   ├── model/
│   │   ├── repository/
│   │   ├── service/
│   │   ├── handler/
│   │   └── routes/
│   └── shared/
│       ├── config/
│       ├── database/
│       ├── middleware/
│       └── validator/
└── migrations/
```

## Configuração

Copie o `.env.example` para `.env` e preencha as variáveis antes de rodar.

| Variável | Descrição |
|---|---|
| DB_HOST | Host do banco de dados |
| DB_PORT | Porta do banco (padrão: 3306) |
| DB_USER | Usuário do banco |
| DB_PASSWORD | Senha do banco |
| DB_NAME | Nome do banco |
| SERVER_PORT | Porta da aplicação (padrão: 8080) |
| SERVER_ENV | Ambiente (development / production) |
| JWT_SECRET | Chave secreta para assinatura dos tokens JWT |

## Rodando com Docker

Suba o banco de dados:

```bash
docker-compose up -d
```

Rode a aplicação:

```bash
go run cmd/main.go
```

As migrations são aplicadas automaticamente na inicialização.

## Rodando sem Docker

Certifique-se de ter um MySQL 8.0 rodando localmente. Configure o `.env` com as credenciais do seu banco e rode:

```bash
go run cmd/main.go
```

As migrations são aplicadas automaticamente na inicialização.

## Autenticação

As rotas de `/members`, `/events` e `/attendances` são protegidas e exigem token JWT no header:

```
Authorization: Bearer <token>
```

O token é obtido via `POST /auth/login`.

## Rotas

### Auth

> Rotas públicas — não exigem token.

#### `POST /auth/register`
Cria credenciais de acesso para um membro já cadastrado.

```json
{
  "cpf": "123.456.789-00",
  "password": "suasenha"
}
```

#### `POST /auth/login`
Autentica o membro e retorna um token JWT válido por 24h.

```json
{
  "cpf": "123.456.789-00",
  "password": "suasenha"
}
```

Resposta:
```json
{
  "token": "<jwt>",
  "nome": "Maria Silva",
  "nome_religioso": null,
  "cargo": "Ekeji"
}
```

---

### Members

> Rotas protegidas — exigem `Authorization: Bearer <token>`.

#### `GET /members`
Lista todos os membros ativos.

#### `POST /members`
Cria um novo membro.

```json
{
  "nome": "Maria Silva",
  "cpf": "123.456.789-00",
  "data_nascimento": "1990-01-01T00:00:00Z",
  "sexo": "Feminino",
  "telefone": "(11) 99999-9999",
  "email": "maria@exemplo.com",
  "cargo": "Ekeji",
  "status": "ativo"
}
```

Campos opcionais: `nome_religioso`, `rg`, `endereco_rua`, `endereco_numero`, `endereco_complemento`, `endereco_bairro`, `endereco_cidade`, `endereco_estado`, `endereco_cep`, `observacoes`.

Valores válidos para `cargo`: `Ogan`, `Ekeji`, `Membro`, `Iniciado`, `Sacerdote`, `Pai de Santo`, `Mae de Santo`.

Valores válidos para `sexo`: `Masculino`, `Feminino`, `Outro`.

#### `GET /members/:id`
Busca membro pelo ID.

#### `PUT /members/:id`
Atualiza os dados de um membro. Aceita os mesmos campos do `POST`.

#### `DELETE /members/:id`
Remove um membro (soft delete).

---

### Events

> Rotas protegidas — exigem `Authorization: Bearer <token>`.

#### `GET /events`
Lista todos os eventos.

#### `POST /events`
Cria um novo evento.

```json
{
  "name": "Festa de Xango",
  "event_type": "Gira",
  "event_status": "Agendado",
  "date": "2026-05-01T20:00:00Z"
}
```

Campos opcionais: `description`, `location`.

Valores válidos para `event_type`: `Gira`, `Funcao`.

Valores válidos para `event_status`: `Agendado`, `Cancelado`, `Concluido`.

#### `GET /events/:id`
Busca evento pelo ID.

#### `PUT /events/:id`
Atualiza um evento. Aceita os mesmos campos do `POST`.

#### `DELETE /events/:id`
Remove um evento (soft delete).

---

### Attendance

> Rotas protegidas — exigem `Authorization: Bearer <token>`.

#### `POST /attendances`
Registra a presença de um membro em um evento.

```json
{
  "event_id": "<uuid do evento>",
  "member_id": "<uuid do membro>",
  "status": "Presente",
  "marked_by": "Nome de quem registrou"
}
```

Campo opcional: `notes`.

Valores válidos para `status`: `Presente`, `Ausente`, `Justificado`, `Pendente`.

#### `GET /attendances/:id`
Busca registro de presença pelo ID.

#### `PUT /attendances/:id`
Atualiza um registro de presença. Aceita os mesmos campos do `POST`.

#### `DELETE /attendances/:id`
Remove um registro de presença (soft delete).

#### `GET /events/:event_id/attendances`
Lista todas as presenças de um evento.

#### `GET /members/:member_id/attendances`
Lista todas as presenças de um membro.

## Desenvolvedor

Rodrigo Junior (Diniz-J)
