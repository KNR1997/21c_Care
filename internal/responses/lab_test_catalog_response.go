package responses

import "go-echo-starter/internal/models"

type LabTestCatalogResponse struct {
	ID           int64   `json:"id" example:"1"`
	Name         string  `json:"name" example:"Blood Test"`
	DefaultPrice float64 `json:"default_price" example:"7500"`
}

func NewLabTestCatalogsResponse(labtestcatalogs []models.LabTestCatalog) *[]LabTestCatalogResponse {
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

func NewLabTestCatalogResponse(labtestcatalog models.LabTestCatalog) *LabTestCatalogResponse {
	return &LabTestCatalogResponse{
		ID:           labtestcatalog.ID,
		Name:         labtestcatalog.Name,
		DefaultPrice: labtestcatalog.DefaultPrice,
	}
}

type LabTestCatalogPaginationResponse struct {
	Limit      int                      `json:"limit" example:"10"`
	Page       int                      `json:"page" example:"1"`
	Sort       string                   `json:"sort" example:"1"`
	TotalRows  int64                    `json:"total_rows" example:"8"`
	TotalPages int                      `json:"total_pages" example:"2"`
	Data       []LabTestCatalogResponse `json:"data"`
}
