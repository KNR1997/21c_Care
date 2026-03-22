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

func (r *DrugCatalogRepository) Create(ctx context.Context, catalog *models.DrugCatalog) error {
	if err := r.db.WithContext(ctx).Create(catalog).Error; err != nil {
		return fmt.Errorf("execute insert drugcatalog query: %w", err)
	}

	return nil
}

// In repositories/drugcatalog.go
func (r *DrugCatalogRepository) List(ctx context.Context) ([]models.DrugCatalog, error) {
	var catalogs []models.DrugCatalog

	// Log the query being executed
	slog.Info("Executing GetDrugCatalogs query")

	result := r.db.WithContext(ctx).Find(&catalogs)

	// Log the result
	slog.Info("GetDrugCatalogs result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select drugcatalogs query: %w", result.Error)
	}

	// Log the actual drugcatalogs found
	slog.Info("DrugCatalogs found", "count", len(catalogs))

	return catalogs, nil
}

func (r *DrugCatalogRepository) ListPaginated(pagination Pagination[models.DrugCatalog]) (*Pagination[models.DrugCatalog], error) {
	var catalogs []models.DrugCatalog

	r.db.Scopes(paginate(catalogs, &pagination, r.db)).Find(&catalogs)
	pagination.Data = catalogs

	return &pagination, nil
}

func (r *DrugCatalogRepository) Get(ctx context.Context, id uint) (models.DrugCatalog, error) {
	var catalog models.DrugCatalog
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&catalog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.DrugCatalog{}, errors.Join(models.ErrDrugCatalogNotFound, err)
	} else if err != nil {
		return models.DrugCatalog{}, fmt.Errorf("execute select drugcatalog by id query: %w", err)
	}

	return catalog, nil
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

func (r *DrugCatalogRepository) GetByName(ctx context.Context, name string) (models.DrugCatalog, error) {
	var catalog models.DrugCatalog
	err := r.db.WithContext(ctx).Where("name = ?", name).Take(&catalog).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.DrugCatalog{}, errors.Join(models.ErrDrugCatalogNotFound, err)
	} else if err != nil {
		return models.DrugCatalog{}, fmt.Errorf("execute select drugcatalog by name query: %w", err)
	}
	return catalog, nil
}
