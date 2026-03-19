package domain

type UpdatePatientRequest struct {
	// PatientID is the patient to update.
	PatientID uint

	Name   string
	Age    int
	Gender string
}

type DeletePatientRequest struct {
	// PatientID is the patient to update.
	PatientID uint
}
