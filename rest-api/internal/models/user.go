package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// DTO для создания пользователя
type CreateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// DTO для обновления пользователя
type UpdateUserRequest struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

// Ответ с ошибкой
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}
