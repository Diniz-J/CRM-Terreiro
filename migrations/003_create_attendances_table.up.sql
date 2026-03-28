-- Migration: Criar tabela attendances
-- Autor: Rodrigo Diniz
-- Data: 2026-02-08

CREATE TABLE IF NOT EXISTS attendances (
    id VARCHAR(36) PRIMARY KEY COMMENT 'ID da inscrição',
    event_id VARCHAR(36) NOT NULL COMMENT 'ID do evento',
    member_id VARCHAR(36) NOT NULL COMMENT 'ID do usuário',
    status ENUM('Presente', 'Ausente', 'Justificado', 'Pendente') DEFAULT 'Pendente' COMMENT 'Status da inscrição',
    notes TEXT DEFAULT NULL COMMENT 'Anotações',
    marked_at TIMESTAMP DEFAULT NULL COMMENT 'Data da marcação',
    marked_by VARCHAR(255) DEFAULT NULL COMMENT 'Nome do usuário que marcou',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Data de criação',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Data de atualização',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT 'Data de exclusão',

    CONSTRAINT fk_attendance_event FOREIGN KEY (event_id) REFERENCES events(id),
    CONSTRAINT fk_attendance_member FOREIGN KEY (member_id) REFERENCES members(id),

    UNIQUE KEY uq_attendance_event_member (event_id, member_id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Tabela de inscrições'