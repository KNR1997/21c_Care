package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type BasicPatient struct {
	Name   string `json:"name" validate:"required" example:"Saman Perera"`
	Age    int    `json:"age" validate:"required" example:"34"`
	Gender string `json:"gender" validate:"required" example:"Male"`
}

func (bp BasicPatient) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.Name, validation.Required),
		validation.Field(&bp.Age, validation.Required),
		validation.Field(&bp.Gender, validation.Required),
	)
}

type CreatePatientRequest struct {
	BasicPatient
}

type UpdatePatientRequest struct {
	BasicPatient
}
