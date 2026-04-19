package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"energy_rms/internal/handlers"
	"energy_rms/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	db_url := os.Getenv("DB_URL")

	if db_url == "" {
		db_host := os.Getenv("DB_HOST")
		db_port := os.Getenv("DB_PORT")
		db_user := os.Getenv("DB_USER")
		db_pasw := os.Getenv("DB_PASSWORD")
		db_name := os.Getenv("DB_NAME")
		db_sslmode := os.Getenv("DB_SSLMODE")

		db_url = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s",
			db_user, db_pasw, db_host, db_port, db_name, db_sslmode,
		)
	}

	// Подключение к БД с пулом соединений
	config, err := pgxpool.ParseConfig(db_url)
	if err != nil {
		log.Fatal("Ошибка парсинга конфига БД:", err)
	}

	config.MaxConns = 10
	config.MinConns = 2

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer db.Close()

	// Проверка подключения
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("БД не отвечает:", err)
	}
	log.Println("✅ Подключено к PostgreSQL")

	// Инициализация репозитория и хендлеров
	userRepo := repository.NewPostgresUserRepository(db)
	userHandlers := handlers.NewUserHandlers(userRepo)

	// Настройка роутов
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", userHandlers.HealthCheck)
	mux.HandleFunc("GET /user/{id}", userHandlers.GetUser)
	mux.HandleFunc("GET /users", userHandlers.GetUsers)
	mux.HandleFunc("POST /user", userHandlers.CreateUser)
	mux.HandleFunc("PUT /user/{id}", userHandlers.UpdateUser)
	mux.HandleFunc("DELETE /user/{id}", userHandlers.DeleteUser)

	// HTTP сервер
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		log.Println("🚀 Сервер запущен на :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Ошибка сервера:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("🛑 Завершение работы...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Ошибка при остановке сервера:", err)
	}

	log.Println("👋 Сервер остановлен")
}
