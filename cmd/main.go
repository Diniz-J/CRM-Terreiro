package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Diniz-J/teiunecc-admin/internal/shared/config"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/database"
)

func main() {

	//carregar config
	cfg := config.Load()

	//connect DB
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to database")

	// TODO: Inicializar módulos (members, payments, etc)

	// TODO: Configurar rotas

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
