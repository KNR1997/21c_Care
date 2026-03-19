package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"log/slog"

	"gorm.io/gorm"
)

type DrugCatalogRepository struct {
	db *gorm.DB
}

func NewDrugCatalogRepository(db *gorm.DB) *DrugCatalogRepository {
	return &DrugCatalogRepository{db: db}
}

func (r *DrugCatalogRepository) Create(ctx context.Context, drugcatalog *models.DrugCatalog) error {
	if err := r.db.WithContext(ctx).Create(drugcatalog).Error; err != nil {
		return fmt.Errorf("execute insert drugcatalog query: %w", err)
	}

	return nil
}

// In repositories/drugcatalog.go
func (r *DrugCatalogRepository) GetDrugCatalogs(ctx context.Context) ([]models.DrugCatalog, error) {
	var drugcatalogs []models.DrugCatalog

	// Log the query being executed
	slog.Info("Executing GetDrugCatalogs query")

	result := r.db.WithContext(ctx).Find(&drugcatalogs)

	// Log the result
	slog.Info("GetDrugCatalogs result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select drugcatalogs query: %w", result.Error)
	}

	// Log the actual drugcatalogs found
	slog.Info("DrugCatalogs found", "count", len(drugcatalogs))

	return drugcatalogs, nil
}

func (r *DrugCatalogRepository) GetDrugCatalog(ctx context.Context, id uint) (models.DrugCatalog, error) {
	var drugcatalog models.DrugCatalog
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&drugcatalog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.DrugCatalog{}, errors.Join(models.ErrDrugCatalogNotFound, err)
	} else if err != nil {
		return models.DrugCatalog{}, fmt.Errorf("execute select drugcatalog by id query: %w", err)
	}

	return drugcatalog, nil
}

func (r *DrugCatalogRepository) Update(ctx context.Context, drugcatalog *models.DrugCatalog) error {
	if err := r.db.WithContext(ctx).Save(drugcatalog).Error; err != nil {
		return fmt.Errorf("execute update drugcatalog query: %w", err)
	}

	return nil
}

func (r *DrugCatalogRepository) Delete(ctx context.Context, drugcatalog *models.DrugCatalog) error {
	if err := r.db.WithContext(ctx).Delete(drugcatalog).Error; err != nil {
		return fmt.Errorf("execute delete drugcatalog query: %w", err)
	}

	return nil
}
