package domain

type UpdateLabTestCatalogRequest struct {
	// LabTestCatalogID is the labtestcatalog to update.
	LabTestCatalogID uint

	Name         string
	DefaultPrice float64
}

type DeleteLabTestCatalogRequest struct {
	// LabTestCatalogID is the labtestcatalog to update.
	LabTestCatalogID uint
}
