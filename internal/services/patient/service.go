package patient

import (
	"context"
	"fmt"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
)

//go:generate go tool mockgen -source=$GOFILE -destination=service_mock_test.go -package=${GOPACKAGE}_test -typed=true

type patientRepository interface {
	Create(ctx context.Context, patient *models.Patient) error
	GetPatients(ctx context.Context) ([]models.Patient, error)
	GetPatient(ctx context.Context, id uint) (models.Patient, error)
	Update(ctx context.Context, patient *models.Patient) error
	Delete(ctx context.Context, patient *models.Patient) error
}

type Service struct {
	patientRepository patientRepository
}

func NewService(patientRepository patientRepository) *Service {
	return &Service{patientRepository: patientRepository}
}

func (s *Service) Create(ctx context.Context, patient *models.Patient) error {
	if err := s.patientRepository.Create(ctx, patient); err != nil {
		return fmt.Errorf("create patient in repository: %w", err)
	}

	return nil
}

func (s *Service) GetPatients(ctx context.Context) ([]models.Patient, error) {
	patients, err := s.patientRepository.GetPatients(ctx)
	if err != nil {
		return nil, fmt.Errorf("get patients from repository: %w", err)
	}

	return patients, nil
}

func (s *Service) GetPatient(ctx context.Context, id uint) (models.Patient, error) {
	patient, err := s.patientRepository.GetPatient(ctx, id)
	if err != nil {
		return models.Patient{}, fmt.Errorf("get patient from repository: %w", err)
	}

	return patient, nil
}

func (s *Service) UpdatePatient(ctx context.Context, request domain.UpdatePatientRequest) (*models.Patient, error) {
	patient, err := s.patientRepository.GetPatient(ctx, request.PatientID)
	if err != nil {
		return nil, fmt.Errorf("get stored patient from repository: %w", err)
	}

	patient.Name = request.Name
	patient.Age = request.Age
	patient.Gender = request.Gender

	if err := s.patientRepository.Update(ctx, &patient); err != nil {
		return nil, fmt.Errorf("update patient in repository: %w", err)
	}

	return &patient, nil
}

func (s *Service) DeletePatient(ctx context.Context, request domain.DeletePatientRequest) error {
	patient, err := s.patientRepository.GetPatient(ctx, request.PatientID)
	if err != nil {
		return fmt.Errorf("get stored patient from repository: %w", err)
	}

	if err := s.patientRepository.Delete(ctx, &patient); err != nil {
		return fmt.Errorf("delete patient in repository: %w", err)
	}

	return nil
}
