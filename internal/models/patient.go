package models

import "time"

type Patient struct {
	// gorm.Model
	ID        uint   `gorm:"primarykey"`
	Name      string `json:"name" gorm:"type:varchar(255);not null"`
	Age       int    `json:"age"`
	Gender    string `json:"gender" gorm:"type:varchar(10)"`
	CreatedAt time.Time
}
