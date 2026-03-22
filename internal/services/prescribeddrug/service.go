package prescribeddrug

import (
	"context"
	"fmt"
	"go-echo-starter/internal/models"
)

type Repository interface {
	Create(ctx context.Context, drug *models.PrescribedDrug) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, drug *models.PrescribedDrug) error {
	if err := s.repo.Create(ctx, drug); err != nil {
		return fmt.Errorf("create prescribed drug: %w", err)
	}
	return nil
}
