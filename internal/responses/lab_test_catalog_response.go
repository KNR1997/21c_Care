package responses

import "go-echo-starter/internal/models"

type LabTestCatalogResponse struct {
	ID           int64   `json:"id" example:"1"`
	Name         string  `json:"name" example:"Blood Test"`
	DefaultPrice float64 `json:"default_price" example:"7500"`
}

func NewLabTestCatalogResponse(labtestcatalogs []models.LabTestCatalog) *[]LabTestCatalogResponse {
	labtestcatalogResponse := make([]LabTestCatalogResponse, 0)

	for i := range labtestcatalogs {
		labtestcatalogResponse = append(labtestcatalogResponse, LabTestCatalogResponse{
			ID:           labtestcatalogs[i].ID,
			Name:         labtestcatalogs[i].Name,
			DefaultPrice: labtestcatalogs[i].DefaultPrice,
		})
	}

	return &labtestcatalogResponse
}
