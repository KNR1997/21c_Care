package models

import "time"

type PrescribedDrug struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	VisitID   int64     `json:"visit_id" gorm:"not null"`
	DrugName  string    `json:"drug_name" gorm:"type:varchar(255);not null"`
	Dosage    string    `json:"dosage"`
	Frequency string    `json:"frequency"`
	Duration  string    `json:"duration"`
	Price     float64   `json:"price" gorm:"type:numeric(10,2);default:0"`
	CreatedAt time.Time `json:"created_at"`
}
