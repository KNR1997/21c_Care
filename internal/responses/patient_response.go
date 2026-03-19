package responses

import "go-echo-starter/internal/models"

type PatientResponse struct {
	Name   string `json:"name" example:"Saman Perera"`
	Age    int    `json:"age" example:"34"`
	Gender string `json:"gender" example:"Male"`
}

func NewPatientResponse(patients []models.Patient) *[]PatientResponse {
	patientResponse := make([]PatientResponse, 0)

	for i := range patients {
		patientResponse = append(patientResponse, PatientResponse{
			Name:   patients[i].Name,
			Age:    patients[i].Age,
			Gender: patients[i].Gender,
		})
	}

	return &patientResponse
}
