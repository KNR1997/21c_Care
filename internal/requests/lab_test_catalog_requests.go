package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BasicLabTestCatalog struct {
	Name         string  `json:"name" validate:"required" example:"Blood Test"`
	DefaultPrice float64 `json:"default_price" validate:"required" example:"7500"`
}

func (bp BasicLabTestCatalog) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.Name, validation.Required),
		validation.Field(&bp.DefaultPrice, validation.Required),
	)
}

type CreateLabTestCatalogRequest struct {
	BasicLabTestCatalog
}

type UpdateLabTestCatalogRequest struct {
	BasicLabTestCatalog
}
