package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"energy_rms/internal/models"
	"energy_rms/internal/repository"
)

type UserHandlers struct {
	repo repository.UserRepository
}

func NewUserHandlers(repo repository.UserRepository) *UserHandlers {
	return &UserHandlers{
		repo: repo,
	}
}

// GET /user/{id}
func (h *UserHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r.URL.Path)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		log.Println("DB error:", err)
		h.sendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		h.sendError(w, "User not found", http.StatusNotFound)
		return
	}

	h.sendJSON(w, user, http.StatusOK)
}

// GET /users
func (h *UserHandlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := h.repo.GetAll(r.Context())
	if err != nil {
		log.Println("DB error:", err)
		h.sendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if users == nil {
		users = []models.User{}
	}

	h.sendJSON(w, users, http.StatusOK)
}

// POST /user
func (h *UserHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		h.sendError(w, "Name is required", http.StatusBadRequest)
		return
	}

	if req.Age <= 0 || req.Age > 150 {
		h.sendError(w, "Age must be between 1 and 150", http.StatusBadRequest)
		return
	}

	user, err := h.repo.Create(r.Context(), &req)
	if err != nil {
		log.Println("DB error:", err)
		h.sendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, user, http.StatusCreated)
}

// PUT /user/{id}
func (h *UserHandlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r.URL.Path)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.sendError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := h.repo.Update(r.Context(), id, &req)
	if err != nil {
		log.Println("DB error:", err)
		h.sendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if user == nil {
		h.sendError(w, "User not found", http.StatusNotFound)
		return
	}

	h.sendJSON(w, user, http.StatusOK)
}

// DELETE /user/{id}
func (h *UserHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := h.extractID(r.URL.Path)
	if err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		log.Println("DB error:", err)
		h.sendError(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.sendJSON(w, map[string]string{"message": "User deleted successfully"}, http.StatusOK)
}

// GET /health
func (h *UserHandlers) HealthCheck(w http.ResponseWriter, r *http.Request) {
	h.sendJSON(w, map[string]string{"status": "healthy"}, http.StatusOK)
}

// Вспомогательные методы
func (h *UserHandlers) extractID(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 || parts[1] == "" {
		return 0, http.ErrNoLocation
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil || id <= 0 {
		return 0, http.ErrNoLocation
	}

	return id, nil
}

func (h *UserHandlers) sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("JSON encode error:", err)
	}
}

func (h *UserHandlers) sendError(w http.ResponseWriter, message string, code int) {
	h.sendJSON(w, models.ErrorResponse{
		Error:   http.StatusText(code),
		Code:    code,
		Message: message,
	}, code)
}
