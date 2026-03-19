package models

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrInvalidAuthToken = errors.New("invalid authorization jwt token")

	ErrPostNotFound           = errors.New("post not found")
	ErrPatientNotFound        = errors.New("patient not found")
	ErrDrugCatalogNotFound    = errors.New("drug catalog not found")
	ErrLabTestCatalogNotFound = errors.New("lab test catalog not found")
	ErrClinicalNoteNotFound   = errors.New("clinical note not found")
	ErrBillingNotFound        = errors.New("billing not found")
	ErrVisitNotFound          = errors.New("visit not found")

	ErrForbidden = errors.New("operation forbidden")
)
