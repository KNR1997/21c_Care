package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"log/slog"

	"gorm.io/gorm"
)

type VisitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) *VisitRepository {
	return &VisitRepository{db: db}
}

func (r *VisitRepository) Create(ctx context.Context, visit *models.Visit) error {
	if err := r.db.WithContext(ctx).Create(visit).Error; err != nil {
		return fmt.Errorf("execute insert visit query: %w", err)
	}

	return nil
}

// In repositories/visit.go
func (r *VisitRepository) List(ctx context.Context) ([]models.Visit, error) {
	var visits []models.Visit

	// Log the query being executed
	slog.Info("Executing GetVisits query")

	result := r.db.WithContext(ctx).Find(&visits)

	// Log the result
	slog.Info("GetVisits result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select visits query: %w", result.Error)
	}

	// Log the actual visits found
	slog.Info("Visits found", "count", len(visits))

	return visits, nil
}

func (r *VisitRepository) ListPaginated(pagination Pagination[models.Visit]) (*Pagination[models.Visit], error) {
	var visits []models.Visit

	r.db.
		Preload("Patient").
		Scopes(paginate(models.Visit{}, &pagination, r.db)).
		Find(&visits)

	pagination.Data = visits

	return &pagination, nil
}

func (r *VisitRepository) Get(ctx context.Context, id uint) (models.Visit, error) {
	var visit models.Visit

	err := r.db.WithContext(ctx).
		Preload("Patient").
		Preload("LabTests").
		Preload("PrescribedDrugs").
		Preload("ClinicalNotes").
		Where("id = ?", id).
		Take(&visit).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Visit{}, errors.Join(models.ErrVisitNotFound, err)
	} else if err != nil {
		return models.Visit{}, fmt.Errorf("execute select visit by id query: %w", err)
	}

	return visit, nil
}

func (r *VisitRepository) Update(ctx context.Context, visit *models.Visit) error {
	if err := r.db.WithContext(ctx).Save(visit).Error; err != nil {
		return fmt.Errorf("execute update visit query: %w", err)
	}

	return nil
}

func (r *VisitRepository) Delete(ctx context.Context, visit *models.Visit) error {
	if err := r.db.WithContext(ctx).Delete(visit).Error; err != nil {
		return fmt.Errorf("execute delete visit query: %w", err)
	}

	return nil
}
