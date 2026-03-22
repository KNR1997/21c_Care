package responses

import "go-echo-starter/internal/models"

type PrescribedDrugResponse struct {
	ID        int64   `json:"id" example:"1"`
	VisitID   int64   `json:"visit_id" example:"1"`
	DrugName  string  `json:"drug_name" example:"ECG"`
	Dosage    string  `json:"dosage" example:"7500"`
	Frequency string  `json:"frequency" example:"Once a day"`
	Duration  string  `json:"duration" example:"5 days"`
	Price     float64 `json:"price" example:"7500"`
}

func NewPrescribedDrugsResponse(prescribeddrugs []models.PrescribedDrug) *[]PrescribedDrugResponse {
	prescribeddrugResponse := make([]PrescribedDrugResponse, 0)

	for i := range prescribeddrugs {
		prescribeddrugResponse = append(prescribeddrugResponse, PrescribedDrugResponse{
			ID:        prescribeddrugs[i].ID,
			VisitID:   prescribeddrugs[i].VisitID,
			DrugName:  prescribeddrugs[i].DrugName,
			Dosage:    prescribeddrugs[i].Dosage,
			Frequency: prescribeddrugs[i].Frequency,
			Duration:  prescribeddrugs[i].Duration,
			Price:     prescribeddrugs[i].Price,
		})
	}

	return &prescribeddrugResponse
}

func NewPrescribedDrugResponse(prescribeddrug models.PrescribedDrug) *PrescribedDrugResponse {
	return &PrescribedDrugResponse{
		ID:        prescribeddrug.ID,
		VisitID:   prescribeddrug.VisitID,
		DrugName:  prescribeddrug.DrugName,
		Dosage:    prescribeddrug.Dosage,
		Frequency: prescribeddrug.Frequency,
		Duration:  prescribeddrug.Duration,
		Price:     prescribeddrug.Price,
	}
}
