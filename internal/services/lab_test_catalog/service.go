package labtestcatalog

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
)

type labTestCatalogRepository interface {
	Create(ctx context.Context, labtestcatalog *models.LabTestCatalog) error
	GetLabTestCatalogs(ctx context.Context) ([]models.LabTestCatalog, error)
	GetLabTestCatalog(ctx context.Context, id uint) (models.LabTestCatalog, error)
	Update(ctx context.Context, labtestcatalog *models.LabTestCatalog) error
	Delete(ctx context.Context, labtestcatalog *models.LabTestCatalog) error
}

type Service struct {
	labTestCatalogRepository labTestCatalogRepository
}

func NewService(labTestCatalogRepository labTestCatalogRepository) *Service {
	return &Service{labTestCatalogRepository: labTestCatalogRepository}
}

func (s *Service) Create(ctx context.Context, labtestcatalog *models.LabTestCatalog) error {
	if err := s.labTestCatalogRepository.Create(ctx, labtestcatalog); err != nil {
		return fmt.Errorf("create labtestcatalog in repository: %w", err)
	}

	return nil
}

func (s *Service) GetLabTestCatalogs(ctx context.Context) ([]models.LabTestCatalog, error) {
	labtestcatalogs, err := s.labTestCatalogRepository.GetLabTestCatalogs(ctx)
	if err != nil {
		return nil, fmt.Errorf("get labtestcatalogs from repository: %w", err)
	}

	return labtestcatalogs, nil
}

func (s *Service) GetLabTestCatalog(ctx context.Context, id uint) (models.LabTestCatalog, error) {
	labtestcatalog, err := s.labTestCatalogRepository.GetLabTestCatalog(ctx, id)
	if err != nil {
		return models.LabTestCatalog{}, fmt.Errorf("get labtestcatalog from repository: %w", err)
	}

	return labtestcatalog, nil
}

func (s *Service) UpdateLabTestCatalog(ctx context.Context, request domain.UpdateLabTestCatalogRequest) (*models.LabTestCatalog, error) {
	labtestcatalog, err := s.labTestCatalogRepository.GetLabTestCatalog(ctx, request.LabTestCatalogID)
	if err != nil {
		return nil, fmt.Errorf("get stored labtestcatalog from repository: %w", err)
	}

	labtestcatalog.Name = request.Name
	labtestcatalog.DefaultPrice = request.DefaultPrice

	if err := s.labTestCatalogRepository.Update(ctx, &labtestcatalog); err != nil {
		return nil, fmt.Errorf("update labtestcatalog in repository: %w", err)
	}

	return &labtestcatalog, nil
}

func (s *Service) DeleteLabTestCatalog(ctx context.Context, request domain.DeleteLabTestCatalogRequest) error {
	labtestcatalog, err := s.labTestCatalogRepository.GetLabTestCatalog(ctx, request.LabTestCatalogID)
	if err != nil {
		return fmt.Errorf("get stored labtestcatalog from repository: %w", err)
	}

	if err := s.labTestCatalogRepository.Delete(ctx, &labtestcatalog); err != nil {
		return fmt.Errorf("delete labtestcatalog in repository: %w", err)
	}

	return nil
}
