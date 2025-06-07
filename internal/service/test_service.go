package service

import (
	"context"

	"github.com/mohit838/inventory-managements-golang/internal/models"
	"github.com/mohit838/inventory-managements-golang/internal/repository"
)

type TestService interface {
	FetchAllData(ctx context.Context) ([]models.Test, error)
}

type testService struct {
	repo repository.TestRepository
}

func NewTestService(repo repository.TestRepository) TestService {
	return &testService{repo: repo}
}

func (s *testService) FetchAllData(ctx context.Context) ([]models.Test, error) {
	return s.repo.GetAll(ctx)
}
