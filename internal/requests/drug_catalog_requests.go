package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BasicDrugCatalog struct {
	Name         string  `json:"name" validate:"required" example:"Dupixent"`
	DefaultPrice float64 `json:"default_price" validate:"required" example:"1200"`
}

func (bp BasicDrugCatalog) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.Name, validation.Required),
		validation.Field(&bp.DefaultPrice, validation.Required),
	)
}

type CreateDrugCatalogRequest struct {
	BasicDrugCatalog
}

type UpdateDrugCatalogRequest struct {
	BasicDrugCatalog
}
