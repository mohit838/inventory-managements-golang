package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/mohit838/inventory-managements-golang/internal/models"
)

type TestRepository interface {
	GetAll(ctx context.Context) ([]models.Test, error)
}

type testRepo struct {
	db *sqlx.DB
}

func NewTestRepository(db *sqlx.DB) TestRepository {
	return &testRepo{db: db}
}

func (r *testRepo) GetAll(ctx context.Context) ([]models.Test, error) {
	query := `SELECT id, name, email FROM users`
	var users []models.Test
	if err := r.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}
	return users, nil
}
