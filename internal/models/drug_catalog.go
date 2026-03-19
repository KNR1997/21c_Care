package models

type DrugCatalog struct {
	ID           int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string  `json:"name" gorm:"unique;not null"`
	DefaultPrice float64 `json:"default_price" gorm:"type:numeric(10,2);default:0"`
}
