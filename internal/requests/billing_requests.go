package requests

import validation "github.com/go-ozzo/ozzo-validation/v4"

type CreateBillingRequest struct {
	VisitID int64 `json:"visit_id" validate:"required" example:1`
}

func (cbr CreateBillingRequest) Validate() error {
	return validation.ValidateStruct(&cbr,
		validation.Field(&cbr.VisitID, validation.Required),
	)
}
