package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "context"
    "syscall"
    "time"
    
    "user-service/database"
    "user-service/handlers"
    "user-service/repositories"
)

func main() {
    // Подключение к базе данных
    db, err := database.NewDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()
    
    // Инициализация репозитория и обработчиков
    userRepo := repositories.NewUserRepository(db.Pool)
    userHandler := handlers.NewUserHandler(userRepo)
    
    // Настройка маршрутов
    mux := http.NewServeMux()
    
    // REST маршруты для пользователей
    mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            userHandler.GetAllUsers(w, r)
        case http.MethodPost:
            userHandler.CreateUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    
    mux.HandleFunc("/api/users/", userHandler.GetUserByID)
    
    // Health check endpoint
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"healthy"}`))
    })
    
    // Настройка сервера
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
    server := &http.Server{
        Addr:         ":" + port,
        Handler:      mux,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    // Graceful shutdown
    go func() {
        log.Printf("Server starting on port %s", port)
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed to start: %v", err)
        }
    }()
    
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    
    log.Println("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        log.Fatalf("Server forced to shutdown: %v", err)
    }
    
    log.Println("Server exited properly")
}