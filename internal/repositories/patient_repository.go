package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"log/slog"

	"gorm.io/gorm"
)

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) Create(ctx context.Context, patient *models.Patient) error {
	if err := r.db.WithContext(ctx).Create(patient).Error; err != nil {
		return fmt.Errorf("execute insert patient query: %w", err)
	}

	return nil
}

// In repositories/patient.go
func (r *PatientRepository) GetPatients(ctx context.Context) ([]models.Patient, error) {
	var patients []models.Patient

	// Log the query being executed
	slog.Info("Executing GetPatients query")

	result := r.db.WithContext(ctx).Find(&patients)

	// Log the result
	slog.Info("GetPatients result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select patients query: %w", result.Error)
	}

	// Log the actual patients found
	slog.Info("Patients found", "count", len(patients))

	return patients, nil
}

func (r *PatientRepository) GetPatient(ctx context.Context, id uint) (models.Patient, error) {
	var patient models.Patient
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&patient).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Patient{}, errors.Join(models.ErrPatientNotFound, err)
	} else if err != nil {
		return models.Patient{}, fmt.Errorf("execute select patient by id query: %w", err)
	}

	return patient, nil
}

func (r *PatientRepository) Update(ctx context.Context, patient *models.Patient) error {
	if err := r.db.WithContext(ctx).Save(patient).Error; err != nil {
		return fmt.Errorf("execute update patient query: %w", err)
	}

	return nil
}

func (r *PatientRepository) Delete(ctx context.Context, patient *models.Patient) error {
	if err := r.db.WithContext(ctx).Delete(patient).Error; err != nil {
		return fmt.Errorf("execute delete patient query: %w", err)
	}

	return nil
}
