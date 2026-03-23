package drugcatalog

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
)

type Repository interface {
	Create(ctx context.Context, catalog *models.DrugCatalog) error
	List(ctx context.Context) ([]models.DrugCatalog, error)
	ListPaginated(pagination repositories.Pagination[models.DrugCatalog]) (*repositories.Pagination[models.DrugCatalog], error)
	Get(ctx context.Context, id uint) (models.DrugCatalog, error)
	Update(ctx context.Context, catalog *models.DrugCatalog) error
	Delete(ctx context.Context, catalog *models.DrugCatalog) error
	IsExisting(ctx context.Context, name string) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, drugcatalog *models.DrugCatalog) error {
	exists, err := s.repo.IsExisting(ctx, drugcatalog.Name)
	if err != nil {
		return fmt.Errorf("check existing drug: %w", err)
	}

	if exists {
		return fmt.Errorf("drug with name '%s' already exists", drugcatalog.Name)
	}

	if err := s.repo.Create(ctx, drugcatalog); err != nil {
		return fmt.Errorf("create drugcatalog in repository: %w", err)
	}

	return nil
}

func (s *Service) List(ctx context.Context) ([]models.DrugCatalog, error) {
	catalogs, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("list drug catalogs: %w", err)
	}
	return catalogs, nil
}

func (s *Service) ListPaginated(pagination repositories.Pagination[models.DrugCatalog]) (*repositories.Pagination[models.DrugCatalog], error) {
	paginatedResult, err := s.repo.ListPaginated(pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated drug catalogs from repository: %w", err)
	}
	return paginatedResult, nil
}

func (s *Service) Get(ctx context.Context, id uint) (models.DrugCatalog, error) {
	catalog, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.DrugCatalog{}, fmt.Errorf("get drug catalog: %w", err)
	}
	return catalog, nil
}

func (s *Service) Update(ctx context.Context, request domain.UpdateDrugCatalogRequest) (*models.DrugCatalog, error) {
	catalog, err := s.repo.Get(ctx, request.DrugCatalogID)
	if err != nil {
		return nil, fmt.Errorf("get stored drug catalog: %w", err)
	}

	catalog.Name = request.Name
	catalog.DefaultPrice = request.DefaultPrice

	if err := s.repo.Update(ctx, &catalog); err != nil {
		return nil, fmt.Errorf("update drug catalog: %w", err)
	}

	return &catalog, nil
}

func (s *Service) Delete(ctx context.Context, request domain.DeleteDrugCatalogRequest) error {
	catalog, err := s.repo.Get(ctx, request.DrugCatalogID)
	if err != nil {
		return fmt.Errorf("get stored drug catalog: %w", err)
	}

	if err := s.repo.Delete(ctx, &catalog); err != nil {
		return fmt.Errorf("delete drug catalog: %w", err)
	}

	return nil
}
