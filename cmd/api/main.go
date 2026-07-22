package main

import (
	"WhatShouldICook/internal/config"
	"context"
	"fmt"
	"log"

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

	pool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}
	defer pool.Close()

	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Failed to acquire connection: %v", err)
	}
	defer conn.Release()

	fmt.Println("Connected to PSQL")

}
