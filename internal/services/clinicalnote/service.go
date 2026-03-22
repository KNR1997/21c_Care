package clinicalnote

import (
	"context"
	"fmt"
	"go-echo-starter/internal/models"
)

type Repository interface {
	Create(ctx context.Context, note *models.ClinicalNote) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, note *models.ClinicalNote) error {
	if err := s.repo.Create(ctx, note); err != nil {
		return fmt.Errorf("create clinical note: %w", err)
	}
	return nil
}
