package domain

type UpdateDrugCatalogRequest struct {
	// DrugCatalogID is the drugcatalog to update.
	DrugCatalogID uint

	Name         string
	DefaultPrice float64
}

type DeleteDrugCatalogRequest struct {
	// DrugCatalogID is the drugcatalog to update.
	DrugCatalogID uint
}
