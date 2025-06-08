package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mohit838/inventory-managements-golang/models"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

// Get user by email
func (r *userRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.GetContext(ctx, &user, `SELECT * FROM users WHERE email = ? AND is_deleted = FALSE LIMIT 1`, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert new user
func (r *userRepo) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users (name, email, password, role_id, is_active, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, NOW(), NOW())`

	_, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		user.RoleID,
		user.IsActive,
	)

	return err
}
