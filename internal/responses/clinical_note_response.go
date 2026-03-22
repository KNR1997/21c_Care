package responses

import "go-echo-starter/internal/models"

type ClinicalNoteResponse struct {
	ID      int64  `json:"id"`
	VisitID int64  `json:"visit_id" example:"1"`
	Note    string `json:"note" example:"some notes"`
}

func NewClinicalNotesResponse(clinicalnotes []models.ClinicalNote) *[]ClinicalNoteResponse {
	clinicalnoteResponse := make([]ClinicalNoteResponse, 0)

	for i := range clinicalnotes {
		clinicalnoteResponse = append(clinicalnoteResponse, ClinicalNoteResponse{
			ID:      clinicalnotes[i].ID,
			VisitID: clinicalnotes[i].VisitID,
			Note:    clinicalnotes[i].Note,
		})
	}

	return &clinicalnoteResponse
}

func NewClinicalNoteResponse(clinicalnote models.ClinicalNote) *ClinicalNoteResponse {
	return &ClinicalNoteResponse{
		ID:      clinicalnote.ID,
		VisitID: clinicalnote.VisitID,
		Note:    clinicalnote.Note,
	}
}
