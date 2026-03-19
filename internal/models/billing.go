package models

import "time"

type Billing struct {
	ID              int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	VisitID         int64     `json:"visit_id" gorm:"unique;not null"`
	ConsultationFee float64   `json:"consultation_fee" gorm:"type:numeric(10,2);default:0"`
	DrugsTotal      float64   `json:"drugs_total" gorm:"type:numeric(10,2);default:0"`
	LabTestsTotal   float64   `json:"lab_tests_total" gorm:"type:numeric(10,2);default:0"`
	GrandTotal      float64   `json:"grand_total" gorm:"type:numeric(10,2);default:0"`
	CreatedAt       time.Time `json:"created_at"`

	// Relation
	Visit Visit `json:"visit" gorm:"foreignKey:VisitID"`
}
