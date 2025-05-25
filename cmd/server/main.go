package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang-migrate/migrate/v4"
	"github.com/passawutwannadee/tb-it03/config"
	"github.com/passawutwannadee/tb-it03/internal/app"
	httphandler "github.com/passawutwannadee/tb-it03/internal/handler/http"
	"github.com/passawutwannadee/tb-it03/pkg/postgres"

	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := postgres.Connect(ctx, config.C.Database)
	if err != nil {
		log.Fatalf("Failed to initialize datyabase: %v", err)
	}

	r := chi.NewRouter()

	// Middleware for logging
	r.Use(middleware.Logger)
	// r.Use(cors.Handler(cors.Options{
	// 	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	// 	AllowedOrigins: []string{"https://*", "http://*"},
	// 	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	// }))

	// Routers
	httphandler.Routes(r, &app.AppConfig{
		Database: db,
	})

	server := &http.Server{
		Handler:      r,
		Addr:         ":" + config.C.App.HTTPPort,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Starting server on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Graceful shutdown handler
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown

	gracefulCtx, graceCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer graceCancel()

	cancel()

	if err := server.Shutdown(gracefulCtx); err != nil {
		log.Printf("HTTP server Shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func init() {

	config.Init("./")

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable&TimeZone=Asia/Bangkok",
		config.C.Database.Username, config.C.Database.Password, config.C.Database.Host, config.C.Database.Port, config.C.Database.Name)

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}
