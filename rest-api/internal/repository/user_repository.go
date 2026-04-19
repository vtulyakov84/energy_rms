package repository

import (
	"context"
	"fmt"

	"energy_rms/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetAll(ctx context.Context) ([]models.User, error)
	Create(ctx context.Context, user *models.CreateUserRequest) (*models.User, error)
	Update(ctx context.Context, id int, user *models.UpdateUserRequest) (*models.User, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) (bool, error)
}

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, name, age, created_at FROM users WHERE id = $1`

	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.CreatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, nil // Пользователь не найден
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `SELECT id, name, age, created_at FROM users ORDER BY id`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации: %w", err)
	}

	return users, nil
}

func (r *PostgresUserRepository) Create(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id, name, age, created_at`

	var user models.User
	err := r.db.QueryRow(ctx, query, req.Name, req.Age).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return &user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, id int, req *models.UpdateUserRequest) (*models.User, error) {
	// Проверяем существование пользователя
	exists, err := r.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, nil
	}

	// Динамическое обновление только переданных полей
	query := `UPDATE users SET name = COALESCE($1, name), age = COALESCE($2, age) 
              WHERE id = $3 RETURNING id, name, age, created_at`

	var user models.User
	err = r.db.QueryRow(ctx, query, req.Name, req.Age, id).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка обновления пользователя: %w", err)
	}

	return &user, nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`

	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("ошибка удаления пользователя: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return nil // Пользователь не найден, но это не ошибка
	}

	return nil
}

func (r *PostgresUserRepository) Exists(ctx context.Context, id int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)`

	var exists bool
	err := r.db.QueryRow(ctx, query, id).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка проверки существования: %w", err)
	}

	return exists, nil
}
