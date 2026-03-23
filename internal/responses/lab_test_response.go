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

type LabTestPaginationResponse struct {
	Limit      int               `json:"limit" example:"10"`
	Page       int               `json:"page" example:"1"`
	Sort       string            `json:"sort" example:"1"`
	TotalRows  int64             `json:"total_rows" example:"8"`
	TotalPages int               `json:"total_pages" example:"2"`
	Data       []LabTestResponse `json:"data"`
}
