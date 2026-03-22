package labtest

import (
	"context"
	"fmt"
	"go-echo-starter/internal/models"
)

type Repository interface {
	Create(ctx context.Context, test *models.LabTest) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, test *models.LabTest) error {
	if err := s.repo.Create(ctx, test); err != nil {
		return fmt.Errorf("create lab test in repository: %w", err)
	}

	return nil
}
