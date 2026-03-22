package models

import "time"

type LabTest struct {
	ID        uint      `gorm:"primarykey"`
	VisitID   int64     `json:"visit_id" gorm:"not null"`
	TestName  string    `json:"test_name" gorm:"type:varchar(255);not null"`
	Price     float64   `json:"price" gorm:"type:numeric(10,2);default:0"`
	CreatedAt time.Time `json:"created_at"`
}
