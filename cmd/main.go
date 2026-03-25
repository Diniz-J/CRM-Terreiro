package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Diniz-J/teiunecc-admin/internal/modules/handler"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/repository"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/routes"
	"github.com/Diniz-J/teiunecc-admin/internal/modules/service"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/config"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/database"
	"github.com/Diniz-J/teiunecc-admin/internal/shared/middleware"
	"github.com/gofiber/fiber/v2"
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
	app := fiber.New()

	app.Use(middleware.Logger)
	app.Use(middleware.CorsMiddleware)

	routes.MemberRoutes(app, memberHandler)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server...")
		_ = app.Shutdown()
	}()
	log.Printf("Servidor rodando na porta %s", cfg.Server.Port)
	log.Fatal(app.Listen(":" + cfg.Server.Port))
}
