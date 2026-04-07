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

| Método | Rota | Descrição |
|---|---|---|
| POST | /auth/register | Cria credenciais para um membro existente |
| POST | /auth/login | Autentica e retorna token JWT |

### Members
| Método | Rota | Descrição |
|---|---|---|
| GET | /members | Lista todos os membros |
| POST | /members | Cria um membro |
| GET | /members/:id | Busca membro por ID |
| PUT | /members/:id | Atualiza um membro |
| DELETE | /members/:id | Remove um membro |

### Events
| Método | Rota | Descrição |
|---|---|---|
| GET | /events | Lista todos os eventos |
| POST | /events | Cria um evento |
| GET | /events/:id | Busca evento por ID |
| PUT | /events/:id | Atualiza um evento |
| DELETE | /events/:id | Remove um evento |

### Attendance
| Método | Rota | Descrição |
|---|---|---|
| POST | /attendances | Registra presença |
| GET | /attendances/:id | Busca presença por ID |
| PUT | /attendances/:id | Atualiza presença |
| DELETE | /attendances/:id | Remove presença |
| GET | /events/:event_id/attendances | Lista presenças de um evento |
| GET | /members/:member_id/attendances | Lista presenças de um membro |

## Desenvolvedor

Rodrigo Junior (Diniz-J)
