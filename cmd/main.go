package main

import (
	"log"
	"net/http"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/handler"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/repository"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/routes"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/members/service"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/config"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/database"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/middleware"
	"github.com/gorilla/mux"
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
	memberRepo := repository.NewMemberRepository(db)
	memberService := service.NewMemberService(memberRepo)
	memberHandler := handler.NewMemberHandler(memberService)

	// TODO: Configurar rotas
	r := mux.NewRouter()
	routes.MemberRoutes(r, memberHandler)

	//Middlewares
	chain := middleware.Chain(middleware.Logger, middleware.CorsMiddleware)
	http.Handle("/", chain(r))

	log.Printf("Servidor rodando na porta %s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, nil))
}
