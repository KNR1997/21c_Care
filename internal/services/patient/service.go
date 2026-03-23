package patient

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
)

//go:generate go tool mockgen -source=$GOFILE -destination=service_mock_test.go -package=${GOPACKAGE}_test -typed=true

type patientRepository interface {
	Create(ctx context.Context, patient *models.Patient) error
	List(ctx context.Context) ([]models.Patient, error)
	ListPaginated(pagination repositories.Pagination[models.Patient]) (*repositories.Pagination[models.Patient], error)
	Get(ctx context.Context, id uint) (models.Patient, error)
	Update(ctx context.Context, patient *models.Patient) error
	Delete(ctx context.Context, patient *models.Patient) error
	IsExisting(ctx context.Context, name string) (bool, error)
}

type Service struct {
	repo patientRepository
}

func NewService(repo patientRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, patient *models.Patient) error {
	// Check if patient already exists
	exists, err := s.repo.IsExisting(ctx, patient.Name)
	if err != nil {
		return fmt.Errorf("check existing patient: %w", err)
	}

	if exists {
		return fmt.Errorf("patient with name '%s' already exists", patient.Name)
	}

	// Create the new patient
	if err := s.repo.Create(ctx, patient); err != nil {
		return fmt.Errorf("create patient in repository: %w", err)
	}

	return nil
}

func (s *Service) List(ctx context.Context) ([]models.Patient, error) {
	patients, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("get patients from repository: %w", err)
	}

	return patients, nil
}

func (s *Service) ListPaginated(pagination repositories.Pagination[models.Patient]) (*repositories.Pagination[models.Patient], error) {
	paginatedResult, err := s.repo.ListPaginated(pagination)
	if err != nil {
		return nil, fmt.Errorf("get paginated patients from repository: %w", err)
	}
	return paginatedResult, nil
}

func (s *Service) Get(ctx context.Context, id uint) (models.Patient, error) {
	patient, err := s.repo.Get(ctx, id)
	if err != nil {
		return models.Patient{}, fmt.Errorf("get patient from repository: %w", err)
	}

	return patient, nil
}

func (s *Service) Update(ctx context.Context, request domain.UpdatePatientRequest) (*models.Patient, error) {
	patient, err := s.repo.Get(ctx, request.PatientID)
	if err != nil {
		return nil, fmt.Errorf("get stored patient from repository: %w", err)
	}

	patient.Name = request.Name
	patient.Age = request.Age
	patient.Gender = request.Gender

	if err := s.repo.Update(ctx, &patient); err != nil {
		return nil, fmt.Errorf("update patient in repository: %w", err)
	}

	return &patient, nil
}

func (s *Service) Delete(ctx context.Context, request domain.DeletePatientRequest) error {
	patient, err := s.repo.Get(ctx, request.PatientID)
	if err != nil {
		return fmt.Errorf("get stored patient from repository: %w", err)
	}

	if err := s.repo.Delete(ctx, &patient); err != nil {
		return fmt.Errorf("delete patient in repository: %w", err)
	}

	return nil
}
