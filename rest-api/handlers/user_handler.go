package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
    
    "user-service/models"
    "user-service/repositories"
)

type UserHandler struct {
    repo *repositories.UserRepository
}

func NewUserHandler(repo *repositories.UserRepository) *UserHandler {
    return &UserHandler{repo: repo}
}

// GetAllUsers - GET /api/users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
    // Проверяем, что это GET запрос и нет ID в пути
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Если путь содержит что-то кроме /api/users, обрабатываем как GetByID
    if r.URL.Path != "/api/users" {
        h.GetUserByID(w, r)
        return
    }
    
    users, err := h.repo.GetAll(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    if users == nil {
        users = []models.User{}
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}

// GetUserByID - GET /api/users/{id}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    // Извлекаем ID из пути
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) != 4 {
        http.Error(w, "Invalid path", http.StatusBadRequest)
        return
    }
    
    id, err := strconv.Atoi(pathParts[3])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }
    
    user, err := h.repo.GetByID(r.Context(), id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    if user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// CreateUser - POST /api/users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    
    var req models.CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    // Простая валидация
    if req.Username == "" {
        http.Error(w, "Username is required", http.StatusBadRequest)
        return
    }
    
    if req.Email == "" {
        http.Error(w, "Email is required", http.StatusBadRequest)
        return
    }
    
    user, err := h.repo.Create(r.Context(), &req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}