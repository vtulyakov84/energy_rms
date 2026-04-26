package database

import (
    "context"
    "fmt"
    "log"
    "os"
    
    "github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
    Pool *pgxpool.Pool
}

func NewDB() (*DB, error) {
    // Получаем строку подключения из переменной окружения
    connString := os.Getenv("DATABASE_URL")
    if connString == "" {
        connString = "postgres://postgres:password@localhost:5432/users_db?sslmode=disable"
    }
    
    config, err := pgxpool.ParseConfig(connString)
    if err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, fmt.Errorf("failed to create connection pool: %w", err)
    }
    
    // Проверяем подключение
    if err := pool.Ping(context.Background()); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }
    
    log.Println("Database connected successfully")
    return &DB{Pool: pool}, nil
}

func (db *DB) Close() {
    if db.Pool != nil {
        db.Pool.Close()
    }
}