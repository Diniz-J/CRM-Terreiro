-- Migration: Criar tabela credentials
-- Autor: Rodrigo Diniz
-- Data: 2026-02-08

CREATE TABLE IF NOT EXISTS credentials (
    id VARCHAR(36) PRIMARY KEY COMMENT 'uuid das credenciais',
    member_id VARCHAR(36) NOT NULL COMMENT 'uuid do membro',
    password_hash VARCHAR(255) NOT NULL COMMENT 'hash da senha',
    is_active BOOLEAN NOT NULL DEFAULT TRUE COMMENT 'indica se a credencial está ativa',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'data de criação',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'data de atualização',

    FOREIGN KEY (member_id) REFERENCES members(id) ON DELETE RESTRICT,

    UNIQUE KEY unique_member_id (member_id)
)   ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Tabela de credenciais'