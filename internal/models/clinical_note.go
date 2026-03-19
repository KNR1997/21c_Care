package models

import "time"

type ClinicalNote struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	VisitID   int64     `json:"visit_id" gorm:"not null"`
	Note      string    `json:"note" gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
}
