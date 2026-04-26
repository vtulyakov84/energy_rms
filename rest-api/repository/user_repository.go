package repositories

import (
    "context"
    "fmt"
    
    "user-service/models"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
    db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
    query := `SELECT id, username, email, age, created_at, updated_at 
              FROM users 
              ORDER BY id`
    
    rows, err := r.db.Query(ctx, query)
    if err != nil {
        return nil, fmt.Errorf("failed to query users: %w", err)
    }
    defer rows.Close()
    
    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(
            &user.ID,
            &user.Username,
            &user.Email,
            &user.Age,
            &user.CreatedAt,
            &user.UpdatedAt,
        )
        if err != nil {
            return nil, fmt.Errorf("failed to scan user: %w", err)
        }
        users = append(users, user)
    }
    
    if err = rows.Err(); err != nil {
        return nil, fmt.Errorf("rows iteration error: %w", err)
    }
    
    return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*models.User, error) {
    query := `SELECT id, username, email, age, created_at, updated_at 
              FROM users 
              WHERE id = $1`
    
    var user models.User
    err := r.db.QueryRow(ctx, query, id).Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.Age,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return nil, nil // Пользователь не найден
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    
    return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, req *models.CreateUserRequest) (*models.User, error) {
    query := `INSERT INTO users (username, email, age, created_at, updated_at) 
              VALUES ($1, $2, $3, NOW(), NOW()) 
              RETURNING id, created_at, updated_at`
    
    var user models.User
    user.Username = req.Username
    user.Email = req.Email
    user.Age = req.Age
    
    err := r.db.QueryRow(ctx, query, req.Username, req.Email, req.Age).Scan(
        &user.ID,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    
    if err != nil {
        return nil, fmt.Errorf("failed to create user: %w", err)
    }
    
    return &user, nil
}