package responses

import "go-echo-starter/internal/models"

type PatientResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name" example:"Saman Perera"`
	Age    int    `json:"age" example:"34"`
	Gender string `json:"gender" example:"Male"`
}

func NewPatientsResponse(patients []models.Patient) *[]PatientResponse {
	patientResponse := make([]PatientResponse, 0)

	for i := range patients {
		patientResponse = append(patientResponse, PatientResponse{
			ID:     patients[i].ID,
			Name:   patients[i].Name,
			Age:    patients[i].Age,
			Gender: patients[i].Gender,
		})
	}

	return &patientResponse
}

func NewPatientResponse(patient models.Patient) *PatientResponse {
	return &PatientResponse{
		ID:     patient.ID,
		Name:   patient.Name,
		Age:    patient.Age,
		Gender: patient.Gender,
	}
}

type PatientPaginationResponse struct {
	Limit      int               `json:"limit" example:"10"`
	Page       int               `json:"page" example:"1"`
	Sort       string            `json:"sort" example:"1"`
	TotalRows  int64             `json:"total_rows" example:"8"`
	TotalPages int               `json:"total_pages" example:"2"`
	Data       []PatientResponse `json:"data"`
}
