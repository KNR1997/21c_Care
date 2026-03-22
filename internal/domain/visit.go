package domain

import "go-echo-starter/internal/models"

type VisitDetails struct {
	PatientName string
	Drugs       []models.PrescribedDrug
	LabTests    []models.LabTest
	Notes       []models.ClinicalNote
	Bill        models.Billing
}
