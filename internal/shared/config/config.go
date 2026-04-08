package config

import (
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	Auth     AuthConfig
}

type AuthConfig struct {
	JWTSecret string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
	Env  string
}

func Load() *Config {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET não configurado — a aplicação não pode subir sem um secret definido")
	}

	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			Username: getEnv("DB_USER", "terreiro_user"),
			Password: getEnv("DB_PASSWORD", "terreiro_pass"),
			Name:     getEnv("DB_NAME", "terreiro_crm"),
		},

		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env:  getEnv("SERVER_ENV", "development"),
		},
		Auth: AuthConfig{
			JWTSecret: jwtSecret,
		},
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	log.Printf("ENV %s not set, using default", key)
	return fallback
}
