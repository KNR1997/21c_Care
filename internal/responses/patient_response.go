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
