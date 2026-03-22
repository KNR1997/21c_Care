package repositories

import (
	"context"
	"fmt"
	"go-echo-starter/internal/models"

	"gorm.io/gorm"
)

type LabTestRepository struct {
	db *gorm.DB
}

func NewLabTestRepository(db *gorm.DB) *LabTestRepository {
	return &LabTestRepository{db: db}
}

func (r *LabTestRepository) Create(ctx context.Context, test *models.LabTest) error {
	if err := r.db.WithContext(ctx).Create(test).Error; err != nil {
		return fmt.Errorf("execute insert LabTest query: %w", err)
	}

	return nil
}

func (r *LabTestRepository) GetByVisitID(ctx context.Context, visitID int64) ([]models.LabTest, error) {
	var labTests []models.LabTest
	if err := r.db.WithContext(ctx).Where("visit_id = ?", visitID).Find(&labTests).Error; err != nil {
		return nil, fmt.Errorf("execute select LabTest query: %w", err)
	}

	return labTests, nil
}
