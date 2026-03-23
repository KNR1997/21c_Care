package responses

import "go-echo-starter/internal/models"

type VisitResponse struct {
	ID                     int64                     `json:"id" example:"1"`
	PatientID              int64                     `json:"patient_id" example:"1"`
	RawInput               string                    `json:"raw_input" example:"Echo"`
	PatientResponse        *PatientResponse          `json:"patient,omitempty"`
	LabTestResponse        *[]LabTestResponse        `json:"lab_tests,omitempty"`
	PrescribedDrugResponse *[]PrescribedDrugResponse `json:"prescribed_drugs,omitempty"`
	ClinicalNoteResponse   *[]ClinicalNoteResponse   `json:"clinical_notes,omitempty"`
}

func NewVisitsResponse(visits []models.Visit) *[]VisitResponse {
	visitResponse := make([]VisitResponse, 0)

	for i := range visits {
		visitResponse = append(visitResponse, VisitResponse{
			ID:        visits[i].ID,
			PatientID: visits[i].PatientID,
			RawInput:  visits[i].RawInput,
		})
	}

	return &visitResponse
}

func NewVisitResponse(visit models.Visit) *VisitResponse {
	return &VisitResponse{
		ID:                     visit.ID,
		PatientID:              visit.PatientID,
		RawInput:               visit.RawInput,
		PatientResponse:        NewPatientResponse(visit.Patient),
		LabTestResponse:        NewLabTestsResponse(visit.LabTests),
		PrescribedDrugResponse: NewPrescribedDrugsResponse(visit.PrescribedDrugs),
		ClinicalNoteResponse:   NewClinicalNotesResponse(visit.ClinicalNotes),
	}
}

type VisitPaginationResponse struct {
	Limit      int             `json:"limit" example:"10"`
	Page       int             `json:"page" example:"1"`
	Sort       string          `json:"sort" example:"1"`
	TotalRows  int64           `json:"total_rows" example:"8"`
	TotalPages int             `json:"total_pages" example:"2"`
	Data       []VisitResponse `json:"data"`
}
