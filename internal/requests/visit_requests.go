package requests

import (
	"go-echo-starter/internal/domain"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type BasicVisit struct {
	PatientID int64              `json:"patient_id" validate:"required" example:"1"`
	RawInput  string             `json:"raw_input" validate:"required" example:"Patient has fever for 3 days. Prescribe Paracetamol 500mg twice daily. Order CBC test."`
	AIResult  *domain.AIResponse `json:"ai_result"  validate:"required"`
}

func (bp BasicVisit) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.PatientID, validation.Required),
		validation.Field(&bp.RawInput, validation.Required),
	)
}

type CreateVisitRequest struct {
	BasicVisit
}

type UpdateVisitRequest struct {
	BasicVisit
}

type PreviewVisitRequest struct {
	RawInput string `json:"raw_input" validate:"required" example:"Patient has fever for 3 days. Prescribe Paracetamol 500mg twice daily. Order CBC test."`
}

func (request PreviewVisitRequest) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.RawInput, validation.Required),
	)
}
