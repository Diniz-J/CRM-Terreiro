-- Migration: Criar tabela events
-- Autor: Rodrigo Diniz
-- Data: 2026-02-08

CREATE TABLE IF NOT EXISTS events (
    id VARCHAR(36) PRIMARY KEY COMMENT 'uuid do evento',

    -- Dados do evento
    name VARCHAR(255) NOT NULL COMMENT 'Nome do evento',
    date DATE NOT NULL COMMENT 'Data do evento',
    description TEXT DEFAULT NULL COMMENT 'Descrição do evento',
    location VARCHAR(255) DEFAULT NULL COMMENT 'Local do evento',

    -- Classificação
    event_type ENUM('Gira', 'Função') NOT NULL COMMENT 'Tipo de evento',
    event_status ENUM('Agendado', 'Cancelado', 'Concluído') NOT NULL COMMENT 'Status do evento',

    -- Auditoria
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Data de criação',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Data de atualização',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT 'Data de exclusão (soft delete)',

    --
    INDEX idx_events_date (date),
    INDEX idx_events_type (event_type),
    INDEX idx_events_status (event_status),
    INDEX idx_events_deleted_at (deleted_at)

)   ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Tabelas de eventos do terreiro'