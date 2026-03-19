package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"log/slog"

	"gorm.io/gorm"
)

type LabTestCatalogRepository struct {
	db *gorm.DB
}

func NewLabTestCatalogRepository(db *gorm.DB) *LabTestCatalogRepository {
	return &LabTestCatalogRepository{db: db}
}

func (r *LabTestCatalogRepository) Create(ctx context.Context, labtestcatalog *models.LabTestCatalog) error {
	if err := r.db.WithContext(ctx).Create(labtestcatalog).Error; err != nil {
		return fmt.Errorf("execute insert labtestcatalog query: %w", err)
	}

	return nil
}

// In repositories/labtestcatalog.go
func (r *LabTestCatalogRepository) GetLabTestCatalogs(ctx context.Context) ([]models.LabTestCatalog, error) {
	var labtestcatalogs []models.LabTestCatalog

	// Log the query being executed
	slog.Info("Executing GetLabTestCatalogs query")

	result := r.db.WithContext(ctx).Find(&labtestcatalogs)

	// Log the result
	slog.Info("GetLabTestCatalogs result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select labtestcatalogs query: %w", result.Error)
	}

	// Log the actual labtestcatalogs found
	slog.Info("LabTestCatalogs found", "count", len(labtestcatalogs))

	return labtestcatalogs, nil
}

func (r *LabTestCatalogRepository) GetLabTestCatalog(ctx context.Context, id uint) (models.LabTestCatalog, error) {
	var labtestcatalog models.LabTestCatalog
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&labtestcatalog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.LabTestCatalog{}, errors.Join(models.ErrLabTestCatalogNotFound, err)
	} else if err != nil {
		return models.LabTestCatalog{}, fmt.Errorf("execute select labtestcatalog by id query: %w", err)
	}

	return labtestcatalog, nil
}

func (r *LabTestCatalogRepository) Update(ctx context.Context, labtestcatalog *models.LabTestCatalog) error {
	if err := r.db.WithContext(ctx).Save(labtestcatalog).Error; err != nil {
		return fmt.Errorf("execute update labtestcatalog query: %w", err)
	}

	return nil
}

func (r *LabTestCatalogRepository) Delete(ctx context.Context, labtestcatalog *models.LabTestCatalog) error {
	if err := r.db.WithContext(ctx).Delete(labtestcatalog).Error; err != nil {
		return fmt.Errorf("execute delete labtestcatalog query: %w", err)
	}

	return nil
}
