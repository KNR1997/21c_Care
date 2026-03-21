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

func (r *LabTestCatalogRepository) Create(ctx context.Context, catalog *models.LabTestCatalog) error {
	if err := r.db.WithContext(ctx).Create(catalog).Error; err != nil {
		return fmt.Errorf("execute insert labtestcatalog query: %w", err)
	}

	return nil
}

// In repositories/labtestcatalog.go
func (r *LabTestCatalogRepository) List(ctx context.Context) ([]models.LabTestCatalog, error) {
	var catalogs []models.LabTestCatalog

	// Log the query being executed
	slog.Info("Executing GetLabTestCatalogs query")

	result := r.db.WithContext(ctx).Find(&catalogs)

	// Log the result
	slog.Info("GetLabTestCatalogs result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select labtestcatalogs query: %w", result.Error)
	}

	// Log the actual labtestcatalogs found
	slog.Info("LabTestCatalogs found", "count", len(catalogs))

	return catalogs, nil
}

func (r *LabTestCatalogRepository) ListPaginated(pagination Pagination[models.LabTestCatalog]) (*Pagination[models.LabTestCatalog], error) {
	var catalogs []models.LabTestCatalog

	r.db.Scopes(paginate(catalogs, &pagination, r.db)).Find(&catalogs)
	pagination.Data = catalogs

	return &pagination, nil
}

func (r *LabTestCatalogRepository) Get(ctx context.Context, id uint) (models.LabTestCatalog, error) {
	var catalog models.LabTestCatalog
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&catalog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.LabTestCatalog{}, errors.Join(models.ErrLabTestCatalogNotFound, err)
	} else if err != nil {
		return models.LabTestCatalog{}, fmt.Errorf("execute select labtestcatalog by id query: %w", err)
	}

	return catalog, nil
}

func (r *LabTestCatalogRepository) Update(ctx context.Context, catalog *models.LabTestCatalog) error {
	if err := r.db.WithContext(ctx).Save(catalog).Error; err != nil {
		return fmt.Errorf("execute update labtestcatalog query: %w", err)
	}

	return nil
}

func (r *LabTestCatalogRepository) Delete(ctx context.Context, catalog *models.LabTestCatalog) error {
	if err := r.db.WithContext(ctx).Delete(catalog).Error; err != nil {
		return fmt.Errorf("execute delete labtestcatalog query: %w", err)
	}

	return nil
}
