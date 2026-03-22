package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"log/slog"

	"gorm.io/gorm"
)

type BillingRepository struct {
	db *gorm.DB
}

func NewBillingRepository(db *gorm.DB) *BillingRepository {
	return &BillingRepository{db: db}
}

func (r *BillingRepository) Create(ctx context.Context, billing *models.Billing) error {
	if err := r.db.WithContext(ctx).Create(billing).Error; err != nil {
		return fmt.Errorf("execute insert billing query: %w", err)
	}

	return nil
}

// In repositories/billing.go
func (r *BillingRepository) List(ctx context.Context) ([]models.Billing, error) {
	var billings []models.Billing

	// Log the query being executed
	slog.Info("Executing GetBillings query")

	result := r.db.WithContext(ctx).Find(billings)

	// Log the result
	slog.Info("GetBillings result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select billings query: %w", result.Error)
	}

	// Log the actual billings found
	slog.Info("Billings found", "count", len(billings))

	return billings, nil
}

func (r *BillingRepository) ListPaginated(pagination Pagination[models.Billing]) (*Pagination[models.Billing], error) {
	var billings []models.Billing

	r.db.
		Preload("Visit").
		Preload("Visit.Patient").
		Scopes(paginate(billings, &pagination, r.db)).Find(&billings)
	pagination.Data = billings

	return &pagination, nil
}

func (r *BillingRepository) Get(ctx context.Context, id uint) (models.Billing, error) {
	var billing models.Billing

	err := r.db.WithContext(ctx).
		Preload("Visit").
		Preload("Visit.Patient").
		Preload("Visit.LabTests").
		Preload("Visit.PrescribedDrugs").
		Preload("Visit.ClinicalNotes").
		Where("id = ?", id).Take(&billing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Billing{}, errors.Join(models.ErrBillingNotFound, err)
	} else if err != nil {
		return models.Billing{}, fmt.Errorf("execute select billing by id query: %w", err)
	}

	return billing, nil
}

func (r *BillingRepository) Update(ctx context.Context, billing *models.Billing) error {
	if err := r.db.WithContext(ctx).Save(billing).Error; err != nil {
		return fmt.Errorf("execute update billing query: %w", err)
	}

	return nil
}

func (r *BillingRepository) Delete(ctx context.Context, billing *models.Billing) error {
	if err := r.db.WithContext(ctx).Delete(billing).Error; err != nil {
		return fmt.Errorf("execute delete billing query: %w", err)
	}

	return nil
}

func (r *BillingRepository) GetByVisitID(ctx context.Context, visitID int64) (*models.Billing, error) {
	var billing models.Billing
	err := r.db.WithContext(ctx).Where("visit_id = ?", visitID).Take(&billing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // No billing found for this visit
	} else if err != nil {
		return nil, fmt.Errorf("execute select billing by visit_id query: %w", err)
	}
	return &billing, nil
}
