package models

type Visit struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	PatientID int64  `json:"patient_id" gorm:"not null"`
	RawInput  string `json:"raw_input" gorm:"type:text;not null"`
	// CreatedAt time.Time `json:"created_at"`

	// Relations
	// Patient         Patient          `json:"patient" gorm:"foreignKey:PatientID"`
	// PrescribedDrugs []PrescribedDrug `json:"prescribed_drugs"`
	// LabTests        []LabTest        `json:"lab_tests"`
	// ClinicalNotes   []ClinicalNote   `json:"clinical_notes"`
	// Billing         Billing          `json:"billing"`
}
