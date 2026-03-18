-- Migration: Criar tabela members
-- Autor: Rodrigo Diniz
-- Data: 2026-02-08

CREATE TABLE IF NOT EXISTS members (
    id VARCHAR(36) PRIMARY KEY COMMENT 'uuid do membro',

    -- Dados pessoais
    nome VARCHAR(255) NOT NULL COMMENT 'Nome civil completo',
    nome_religioso VARCHAR(255) DEFAULT NULL COMMENT 'Nome religioso (opcional)',
    cpf VARCHAR(14) UNIQUE NOT NULL COMMENT 'CPF (XXX.XXX.XXX-XX)',
    rg VARCHAR(20) DEFAULT NULL COMMENT 'RG (opcional)',
    data_nascimento DATE NOT NULL COMMENT 'Data de nascimento',
    sexo ENUM('Feminino', 'Masculino', 'Outro') NOT NULL COMMENT 'Sexo',

    -- Contato
    telefone VARCHAR(20) NOT NULL COMMENT '(XX) XXXXX-XXXX',
    email VARCHAR(255) UNIQUE NOT NULL COMMENT 'Email',

    -- Endereço
    endereco_rua VARCHAR(255) DEFAULT NULL,
    endereco_numero VARCHAR(20) DEFAULT NULL,
    endereco_complemento VARCHAR(100) DEFAULT NULL,
    endereco_bairro VARCHAR(100) DEFAULT NULL,
    endereco_cidade VARCHAR(100) DEFAULT NULL,
    endereco_estado CHAR(2) DEFAULT NULL COMMENT 'UF (2 caracteres)',
    endereco_cep VARCHAR(10) DEFAULT NULL COMMENT 'CEP XXXXX-XXX',

    -- INFORMAÇÕES RELIGIOSAS
    cargo ENUM('Membro', 'Iniciado', 'Ogan', 'Ekeji', 'Sacerdote', 'Pai Pequeno', 'Mãe Pequena') DEFAULT 'membro' COMMENT 'Cargo no terreiro',
    status ENUM('ativo', 'inativo', 'afastado') DEFAULT 'ativo' COMMENT 'Status do membro',
    odun DATE DEFAULT NULL COMMENT 'Data de iniciação (se aplicável)',
    observacoes TEXT DEFAULT NULL COMMENT 'Observações gerais',

    -- Auditoria
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Data de criação',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Data de atualização',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT 'Data de exclusão (soft delete)',

    -- Índice para performance
    INDEX idx_cpf (cpf),
    INDEX idx_email (email),
    INDEX idx_nome (nome),
    INDEX idx_status (status),
    INDEX idx_cargo (cargo),
    INDEX idx_deleted_at (deleted_at)

)   ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Tabelas de membros do terreiro'