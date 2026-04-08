package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Diniz-J/CRM-Terreiro/internal/modules/attendance"
	"github.com/Diniz-J/CRM-Terreiro/internal/modules/auth"
	"github.com/Diniz-J/CRM-Terreiro/internal/modules/event"
	"github.com/Diniz-J/CRM-Terreiro/internal/modules/member"
	"github.com/Diniz-J/CRM-Terreiro/internal/shared/config"
	"github.com/Diniz-J/CRM-Terreiro/internal/shared/database"
	"github.com/Diniz-J/CRM-Terreiro/internal/shared/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env nao encontrado, usando variaveis de ambiente do sistema")
	}

	//carregar config
	cfg := config.Load()

	//connect DB
	db, err := database.Connect(cfg.Database)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to database")

	if err := database.RunMigrations(db, "migrations"); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Migrations applied")

	memberRepo := member.NewMemberRepository(db)
	memberService := member.NewMemberService(memberRepo)
	memberHandler := member.NewMemberHandler(memberService)

	eventRepo := event.NewEventRepository(db)
	eventService := event.NewEventService(eventRepo)
	eventHandler := event.NewEventHandler(eventService)

	attendanceRepo := attendance.NewAttendanceRepository(db)
	attendanceService := attendance.NewAttendanceService(attendanceRepo)
	attendanceHandler := attendance.NewAttendanceHandler(attendanceService)

	authRepo := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepo, memberRepo, cfg.Auth.JWTSecret)
	authHandler := auth.NewAuthHandler(authService)

	app := fiber.New()

	app.Use(middleware.Logger)
	app.Use(middleware.CorsMiddleware)

	// rotas publicas — nao exigem token
	auth.Routes(app, authHandler)

	// rotas protegidas — exigem token JWT valido
	app.Use(middleware.NewAuthMiddleware(cfg.Auth.JWTSecret))
	member.Routes(app, memberHandler)
	event.Routes(app, eventHandler)
	attendance.Routes(app, attendanceHandler)

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
