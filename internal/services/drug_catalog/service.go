package drugCatalog

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
)

type drugCatalogRepository interface {
	Create(ctx context.Context, drugcatalog *models.DrugCatalog) error
	GetDrugCatalogs(ctx context.Context) ([]models.DrugCatalog, error)
	GetDrugCatalog(ctx context.Context, id uint) (models.DrugCatalog, error)
	Update(ctx context.Context, drugcatalog *models.DrugCatalog) error
	Delete(ctx context.Context, drugcatalog *models.DrugCatalog) error
}

type Service struct {
	drugCatalogRepository drugCatalogRepository
}

func NewService(drugCatalogRepository drugCatalogRepository) *Service {
	return &Service{drugCatalogRepository: drugCatalogRepository}
}

func (s *Service) Create(ctx context.Context, drugcatalog *models.DrugCatalog) error {
	if err := s.drugCatalogRepository.Create(ctx, drugcatalog); err != nil {
		return fmt.Errorf("create drugcatalog in repository: %w", err)
	}

	return nil
}

func (s *Service) GetDrugCatalogs(ctx context.Context) ([]models.DrugCatalog, error) {
	drugcatalogs, err := s.drugCatalogRepository.GetDrugCatalogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("get drugcatalogs from repository: %w", err)
	}

	return drugcatalogs, nil
}

func (s *Service) GetDrugCatalog(ctx context.Context, id uint) (models.DrugCatalog, error) {
	drugcatalog, err := s.drugCatalogRepository.GetDrugCatalog(ctx, id)
	if err != nil {
		return models.DrugCatalog{}, fmt.Errorf("get drugcatalog from repository: %w", err)
	}

	return drugcatalog, nil
}

func (s *Service) UpdateDrugCatalog(ctx context.Context, request domain.UpdateDrugCatalogRequest) (*models.DrugCatalog, error) {
	drugcatalog, err := s.drugCatalogRepository.GetDrugCatalog(ctx, request.DrugCatalogID)
	if err != nil {
		return nil, fmt.Errorf("get stored drugcatalog from repository: %w", err)
	}

	drugcatalog.Name = request.Name
	drugcatalog.DefaultPrice = request.DefaultPrice

	if err := s.drugCatalogRepository.Update(ctx, &drugcatalog); err != nil {
		return nil, fmt.Errorf("update drugcatalog in repository: %w", err)
	}

	return &drugcatalog, nil
}

func (s *Service) DeleteDrugCatalog(ctx context.Context, request domain.DeleteDrugCatalogRequest) error {
	drugcatalog, err := s.drugCatalogRepository.GetDrugCatalog(ctx, request.DrugCatalogID)
	if err != nil {
		return fmt.Errorf("get stored drugcatalog from repository: %w", err)
	}

	if err := s.drugCatalogRepository.Delete(ctx, &drugcatalog); err != nil {
		return fmt.Errorf("delete drugcatalog in repository: %w", err)
	}

	return nil
}
