package domain

type Drug struct {
	Name      string `json:"name"`
	Dosage    string `json:"dosage"`
	Frequency string `json:"frequency"`
	Duration  string `json:"duration"`
}

type AIResponse struct {
	Drugs    []Drug   `json:"drugs"`
	LabTests []string `json:"lab_tests"`
	Notes    string   `json:"notes"`
}
