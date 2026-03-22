package repositories

import (
	"context"
	"fmt"
	"go-echo-starter/internal/models"

	"gorm.io/gorm"
)

type PrescribedDrugRepository struct {
	db *gorm.DB
}

func NewPrescribedDrugRepository(db *gorm.DB) *PrescribedDrugRepository {
	return &PrescribedDrugRepository{db: db}
}

func (r *PrescribedDrugRepository) Create(ctx context.Context, prescribedDrug *models.PrescribedDrug) error {
	if err := r.db.WithContext(ctx).Create(prescribedDrug).Error; err != nil {
		return fmt.Errorf("execute insert PrescribedDrug query: %w", err)
	}

	return nil
}

func (r *PrescribedDrugRepository) GetByVisitID(ctx context.Context, visitID int64) ([]models.PrescribedDrug, error) {
	var prescribedDrugs []models.PrescribedDrug
	if err := r.db.WithContext(ctx).Where("visit_id = ?", visitID).Find(&prescribedDrugs).Error; err != nil {
		return nil, fmt.Errorf("execute select PrescribedDrug query: %w", err)
	}

	return prescribedDrugs, nil
}
