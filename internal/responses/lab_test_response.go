package responses

import "go-echo-starter/internal/models"

type LabTestResponse struct {
	ID       uint    `json:"id" example:"1"`
	VisitID  int64   `json:"visit_id" example:"1"`
	TestName string  `json:"test_name" example:"ECG"`
	Price    float64 `json:"price" example:"7500"`
}

func NewLabTestsResponse(labtests []models.LabTest) *[]LabTestResponse {
	labtestResponse := make([]LabTestResponse, 0)

	for i := range labtests {
		labtestResponse = append(labtestResponse, LabTestResponse{
			ID:       labtests[i].ID,
			VisitID:  labtests[i].VisitID,
			TestName: labtests[i].TestName,
			Price:    labtests[i].Price,
		})
	}

	return &labtestResponse
}

func NewLabTestResponse(labtest models.LabTest) *LabTestResponse {
	return &LabTestResponse{
		ID:       labtest.ID,
		VisitID:  labtest.VisitID,
		TestName: labtest.TestName,
		Price:    labtest.Price,
	}
}
