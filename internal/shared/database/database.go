package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Diniz-J/teiunecc-admin/internal/shared/config"
)

func Connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	//u:p@tcp(h:p)/d
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Error opening database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	return db, nil
}
