package responses

import "go-echo-starter/internal/models"

type DrugCatalogResponse struct {
	ID           int64   `json:"id" example:"1"`
	Name         string  `json:"name" example:"Dupixent"`
	DefaultPrice float64 `json:"default_price" example:"1200"`
}

func NewDrugCatalogsResponse(drugcatalogs []models.DrugCatalog) *[]DrugCatalogResponse {
	drugcatalogResponse := make([]DrugCatalogResponse, 0)

	for i := range drugcatalogs {
		drugcatalogResponse = append(drugcatalogResponse, DrugCatalogResponse{
			ID:           drugcatalogs[i].ID,
			Name:         drugcatalogs[i].Name,
			DefaultPrice: drugcatalogs[i].DefaultPrice,
		})
	}

	return &drugcatalogResponse
}

func NewDrugCatalogResponse(drugcatalog models.DrugCatalog) *DrugCatalogResponse {
	return &DrugCatalogResponse{
		ID:           drugcatalog.ID,
		Name:         drugcatalog.Name,
		DefaultPrice: drugcatalog.DefaultPrice,
	}
}
