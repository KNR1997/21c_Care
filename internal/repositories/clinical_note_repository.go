package repositories

import (
	"context"
	"errors"
	"fmt"
	"go-echo-starter/internal/models"
	"log/slog"

	"gorm.io/gorm"
)

type ClinicalNoteRepository struct {
	db *gorm.DB
}

func NewClinicalNoteRepository(db *gorm.DB) *ClinicalNoteRepository {
	return &ClinicalNoteRepository{db: db}
}

func (r *ClinicalNoteRepository) Create(ctx context.Context, clinicalnote *models.ClinicalNote) error {
	if err := r.db.WithContext(ctx).Create(clinicalnote).Error; err != nil {
		return fmt.Errorf("execute insert clinicalnote query: %w", err)
	}

	return nil
}

// In repositories/clinicalnote.go
func (r *ClinicalNoteRepository) GetClinicalNotes(ctx context.Context) ([]models.ClinicalNote, error) {
	var clinicalnotes []models.ClinicalNote

	// Log the query being executed
	slog.Info("Executing GetClinicalNotes query")

	result := r.db.WithContext(ctx).Find(clinicalnotes)

	// Log the result
	slog.Info("GetClinicalNotes result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select clinicalnotes query: %w", result.Error)
	}

	// Log the actual clinicalnotes found
	slog.Info("ClinicalNotes found", "count", len(clinicalnotes))

	return clinicalnotes, nil
}

func (r *ClinicalNoteRepository) GetClinicalNote(ctx context.Context, id uint) (models.ClinicalNote, error) {
	var clinicalnote models.ClinicalNote
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&clinicalnote).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ClinicalNote{}, errors.Join(models.ErrClinicalNoteNotFound, err)
	} else if err != nil {
		return models.ClinicalNote{}, fmt.Errorf("execute select clinicalnote by id query: %w", err)
	}

	return clinicalnote, nil
}

func (r *ClinicalNoteRepository) Update(ctx context.Context, clinicalnote *models.ClinicalNote) error {
	if err := r.db.WithContext(ctx).Save(clinicalnote).Error; err != nil {
		return fmt.Errorf("execute update clinicalnote query: %w", err)
	}

	return nil
}

func (r *ClinicalNoteRepository) Delete(ctx context.Context, clinicalnote *models.ClinicalNote) error {
	if err := r.db.WithContext(ctx).Delete(clinicalnote).Error; err != nil {
		return fmt.Errorf("execute delete clinicalnote query: %w", err)
	}

	return nil
}
