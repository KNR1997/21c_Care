package labtestcatalog

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
)

type Repository interface {
	Create(ctx context.Context, catalog *models.LabTestCatalog) error
	List(ctx context.Context) ([]models.LabTestCatalog, error)
	ListPaginated(pagination repositories.Pagination[models.LabTestCatalog]) (*repositories.Pagination[models.LabTestCatalog], error)
	Get(ctx context.Context, id uint) (models.LabTestCatalog, error)
	Update(ctx context.Context, catalog *models.LabTestCatalog) error
	Delete(ctx context.Context, catalog *models.LabTestCatalog) error
	IsExisting(ctx context.Context, name string) (bool, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, catalog *models.LabTestCatalog) error {
	exists, err := s.repo.IsExisting(ctx, catalog.Name)
	if err != nil {
		return fmt.Errorf("check existing lab test: %w", err)
	}

	if exists {
		return fmt.Errorf("lab test with name '%s' already exists", catalog.Name)
	}
	if err := s.repo.Create(ctx, catalog); err != nil {
		return fmt.Errorf("create lab test catalog in repository: %w", err)
	}

	return nil
}

func (s *Service) List(ctx context.Context) ([]models.LabTestCatalog, error) {
	catalog, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("get lab test catalogs from repository: %w", err)
	}

	return catalog, nil
}

func (s *Service) ListPaginated(pagination repositories.Pagination[models.LabTestCatalog]) (*repositories.Pagination[models.LabTestCatalog], error) {
	paginatedResult, err := s.repo.ListPaginated(pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated lab test catalogs from repository: %w", err)
	}
	return paginatedResult, nil
}

func (s *Service) Get(ctx context.Context, id uint) (models.LabTestCatalog, error) {
	catalog, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.LabTestCatalog{}, fmt.Errorf("get lab test catalog from repository: %w", err)
	}

	return catalog, nil
}

func (s *Service) Update(ctx context.Context, request domain.UpdateLabTestCatalogRequest) (*models.LabTestCatalog, error) {
	catalog, err := s.repo.Get(ctx, request.LabTestCatalogID)
	if err != nil {
		return nil, fmt.Errorf("get stored lab test catalog from repository: %w", err)
	}

	catalog.Name = request.Name
	catalog.DefaultPrice = request.DefaultPrice

	if err := s.repo.Update(ctx, &catalog); err != nil {
		return nil, fmt.Errorf("update lab test catalog in repository: %w", err)
	}

	return &catalog, nil
}

func (s *Service) Delete(ctx context.Context, request domain.DeleteLabTestCatalogRequest) error {
	catalog, err := s.repo.Get(ctx, request.LabTestCatalogID)
	if err != nil {
		return fmt.Errorf("get stored lab test catalog from repository: %w", err)
	}

	if err := s.repo.Delete(ctx, &catalog); err != nil {
		return fmt.Errorf("delete lab test catalog in repository: %w", err)
	}

	return nil
}
