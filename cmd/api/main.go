package main

import (
	"WhatShouldICook/internal/config"
	"WhatShouldICook/internal/handler"
	"WhatShouldICook/internal/repository/postgres"
	"WhatShouldICook/internal/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Загрузка конфига (config.Load()).
// Подключение к БД (db.Connect()).
// Создание репозиториев.
// Создание сервисов.
// Создание хендлеров.
// Настройка роутера (HTTP-маршрутов).
// Запуск HTTP-сервера.
// Graceful shutdown (ожидание сигналов Ctrl+C).

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		// log.Fatal - не выполняет differ
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.Database.URL)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()
	// проверка стабильности связи БД

	// conn, err := pool.Acquire(context.Background())
	// if err != nil {
	// 	log.Fatalf("Failed to acquire connection: %v", err)
	// }
	// defer conn.Release()
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping database %v", err)
	}

	cuisineRepo := postgres.NewCuisineRepo(pool)
	cuisineService := service.NewCuisineService(cuisineRepo)
	cuisineHandler := handler.NewCuisineHandler(cuisineService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1/cuisines", func(r chi.Router) {
		r.Get("/", cuisineHandler.List)
		r.Post("/", cuisineHandler.Create)
		r.Get("/{id}", cuisineHandler.GetByID)
		// r.Get("/{name}", cuisineHandler.GetByName) доделать, если нужно
		r.Put("/{id}", cuisineHandler.Update)
		r.Delete("/{id}", cuisineHandler.Delete)
	})

	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Starting server on addr: %v", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Failed to start server:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server shutdown failed: %v.\nIt will work forever(", err)
	}

	log.Println("Server stopped")

}
