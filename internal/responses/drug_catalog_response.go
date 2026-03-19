package responses

import "go-echo-starter/internal/models"

type DrugCatalogResponse struct {
	Name         string  `json:"name" example:"Dupixent"`
	DefaultPrice float64 `json:"default_price" example:"1200"`
}

func NewDrugCatalogResponse(drugcatalogs []models.DrugCatalog) *[]DrugCatalogResponse {
	drugcatalogResponse := make([]DrugCatalogResponse, 0)

	for i := range drugcatalogs {
		drugcatalogResponse = append(drugcatalogResponse, DrugCatalogResponse{
			Name:         drugcatalogs[i].Name,
			DefaultPrice: drugcatalogs[i].DefaultPrice,
		})
	}

	return &drugcatalogResponse
}
