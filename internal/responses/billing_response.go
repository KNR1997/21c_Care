package responses

import "go-echo-starter/internal/models"

type BillingResponse struct {
	ID                     int64                     `json:"id"`
	VisitID                int64                     `json:"visit_id" gorm:"unique;not null"`
	ConsultationFee        float64                   `json:"consultation_fee" gorm:"type:numeric(10,2);default:0"`
	DrugsTotal             float64                   `json:"drugs_total" gorm:"type:numeric(10,2);default:0"`
	LabTestsTotal          float64                   `json:"lab_tests_total" gorm:"type:numeric(10,2);default:0"`
	GrandTotal             float64                   `json:"grand_total" gorm:"type:numeric(10,2);default:0"`
	PatientResponse        *PatientResponse          `json:"patient,omitempty"`
	LabTestResponse        *[]LabTestResponse        `json:"lab_tests,omitempty"`
	PrescribedDrugResponse *[]PrescribedDrugResponse `json:"prescribed_drugs,omitempty"`
	ClinicalNoteResponse   *[]ClinicalNoteResponse   `json:"clinical_notes,omitempty"`
}

func NewBillingsResponse(billings []models.Billing) *[]BillingResponse {
	billingResponse := make([]BillingResponse, 0)

	for i := range billings {
		billingResponse = append(billingResponse, BillingResponse{
			ID:              billings[i].ID,
			VisitID:         billings[i].VisitID,
			ConsultationFee: billings[i].ConsultationFee,
			DrugsTotal:      billings[i].DrugsTotal,
			LabTestsTotal:   billings[i].LabTestsTotal,
			GrandTotal:      billings[i].GrandTotal,
			PatientResponse: NewPatientResponse(billings[i].Visit.Patient),
		})
	}

	return &billingResponse
}

func NewBillingResponse(billing models.Billing) *BillingResponse {
	return &BillingResponse{
		ID:                     billing.ID,
		VisitID:                billing.VisitID,
		ConsultationFee:        billing.ConsultationFee,
		DrugsTotal:             billing.DrugsTotal,
		LabTestsTotal:          billing.LabTestsTotal,
		GrandTotal:             billing.GrandTotal,
		PatientResponse:        NewPatientResponse(billing.Visit.Patient),
		LabTestResponse:        NewLabTestsResponse(billing.Visit.LabTests),
		PrescribedDrugResponse: NewPrescribedDrugsResponse(billing.Visit.PrescribedDrugs),
		ClinicalNoteResponse:   NewClinicalNotesResponse(billing.Visit.ClinicalNotes),
	}
}
