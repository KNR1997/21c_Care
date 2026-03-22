package billing

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
	"log/slog"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, visit *models.Billing) error
	Get(ctx context.Context, id uint) (models.Billing, error)
	ListPaginated(pagination repositories.Pagination[models.Billing]) (*repositories.Pagination[models.Billing], error)
	GetByVisitID(ctx context.Context, visitID int64) (*models.Billing, error)
}

type Service struct {
	repo        Repository
	drugRepo    *repositories.PrescribedDrugRepository
	labTestRepo *repositories.LabTestRepository
}

func NewService(
	repo Repository,
	drugRepo *repositories.PrescribedDrugRepository,
	labTestRepo *repositories.LabTestRepository,
) *Service {
	return &Service{
		repo:        repo,
		drugRepo:    drugRepo,
		labTestRepo: labTestRepo,
	}
}

func (s *Service) ListPaginated(pagination repositories.Pagination[models.Billing]) (*repositories.Pagination[models.Billing], error) {
	paginatedResult, err := s.repo.ListPaginated(pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated billings from repository: %w", err)
	}
	return paginatedResult, nil
}

func (s *Service) Create(ctx context.Context, visitID int64) (*models.Billing, error) {
	// Check if billing already exists for this visit
	existingBilling, err := s.repo.GetByVisitID(ctx, visitID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("check existing billing: %w", err)
	}

	slog.Info("Creating billing for visit", "visitID", visitID)
	slog.Debug("Existing billing check", "existingBilling", existingBilling)

	if existingBilling != nil {
		return nil, fmt.Errorf("billing already exists for visit ID: %d", visitID)
	}

	drugs, err := s.drugRepo.GetByVisitID(ctx, visitID)
	if err != nil {
		return nil, fmt.Errorf("get drugs: %w", err)
	}

	tests, err := s.labTestRepo.GetByVisitID(ctx, visitID)
	if err != nil {
		return nil, fmt.Errorf("get lab tests: %w", err)
	}

	var drugsTotal float64
	for _, d := range drugs {
		drugsTotal += d.Price
	}

	var testsTotal float64
	for _, t := range tests {
		testsTotal += t.Price
	}

	consultationFee := 500.0

	grandTotal := drugsTotal + testsTotal + consultationFee

	bill := &models.Billing{
		VisitID:         visitID,
		ConsultationFee: consultationFee,
		DrugsTotal:      drugsTotal,
		LabTestsTotal:   testsTotal,
		GrandTotal:      grandTotal,
	}

	if err := s.repo.Create(ctx, bill); err != nil {
		return nil, fmt.Errorf("save billing: %w", err)
	}

	return bill, nil
}

func (s *Service) Get(ctx context.Context, id uint) (models.Billing, error) {
	billing, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.Billing{}, fmt.Errorf("get billing from repository: %w", err)
	}

	return billing, nil
}
