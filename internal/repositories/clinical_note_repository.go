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

func (r *ClinicalNoteRepository) Create(ctx context.Context, note *models.ClinicalNote) error {
	if err := r.db.WithContext(ctx).Create(note).Error; err != nil {
		return fmt.Errorf("execute insert clinicalnote query: %w", err)
	}

	return nil
}

// In repositories/clinicalnote.go
func (r *ClinicalNoteRepository) GetClinicalNotes(ctx context.Context) ([]models.ClinicalNote, error) {
	var notes []models.ClinicalNote

	// Log the query being executed
	slog.Info("Executing GetClinicalNotes query")

	result := r.db.WithContext(ctx).Find(notes)

	// Log the result
	slog.Info("GetClinicalNotes result",
		"count", result.RowsAffected,
		"error", result.Error)

	if result.Error != nil {
		return nil, fmt.Errorf("execute select clinicalnotes query: %w", result.Error)
	}

	// Log the actual clinicalnotes found
	slog.Info("ClinicalNotes found", "count", len(notes))

	return notes, nil
}

func (r *ClinicalNoteRepository) GetClinicalNote(ctx context.Context, id uint) (models.ClinicalNote, error) {
	var note models.ClinicalNote
	err := r.db.WithContext(ctx).Where("id = ?", id).Take(&note).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.ClinicalNote{}, errors.Join(models.ErrClinicalNoteNotFound, err)
	} else if err != nil {
		return models.ClinicalNote{}, fmt.Errorf("execute select clinicalnote by id query: %w", err)
	}

	return note, nil
}

func (r *ClinicalNoteRepository) Update(ctx context.Context, note *models.ClinicalNote) error {
	if err := r.db.WithContext(ctx).Save(note).Error; err != nil {
		return fmt.Errorf("execute update clinicalnote query: %w", err)
	}

	return nil
}

func (r *ClinicalNoteRepository) Delete(ctx context.Context, note *models.ClinicalNote) error {
	if err := r.db.WithContext(ctx).Delete(note).Error; err != nil {
		return fmt.Errorf("execute delete clinicalnote query: %w", err)
	}

	return nil
}
