package handlers

import (
	"context"
	"errors"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/requests"
	"go-echo-starter/internal/responses"
	"net/http"
	"strconv"

	safecast "github.com/ccoveille/go-safecast"
	"github.com/labstack/echo/v4"
)

type patientService interface {
	Create(ctx context.Context, patient *models.Patient) error
	GetPatients(ctx context.Context) ([]models.Patient, error)
	// GetPatient(ctx context.Context, id uint) (models.Patient, error)
	UpdatePatient(ctx context.Context, request domain.UpdatePatientRequest) (*models.Patient, error)
	DeletePatient(ctx context.Context, request domain.DeletePatientRequest) error
}

type PatientHandlers struct {
	patientService patientService
}

func NewPatientHandlers(patientService patientService) *PatientHandlers {
	return &PatientHandlers{patientService: patientService}
}

// CreatePatient godoc
//
//	@Summary		Create patient
//	@Description	Create patient
//	@ID				patients-create
//	@Tags			Patients Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreatePatientRequest	true	"Patient name and age"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/patients [post]
func (p *PatientHandlers) CreatePatient(c echo.Context) error {
	// authClaims, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	var createPatientRequest requests.CreatePatientRequest
	if err := c.Bind(&createPatientRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := createPatientRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	patient := &models.Patient{
		Name:   createPatientRequest.Name,
		Age:    createPatientRequest.Age,
		Gender: createPatientRequest.Gender,
	}

	if err := p.patientService.Create(c.Request().Context(), patient); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to create patient: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "Patient successfully created")
}

// GetPatients godoc
//
//	@Summary		Get patients
//	@Description	Get the list of all patients
//	@ID				patients-get
//	@Tags			Patients Actions
//	@Produce		json
//	@Success		200	{array}	responses.PatientResponse
//	@Security		ApiKeyAuth
//	@Router			/patients [get]
func (p *PatientHandlers) GetPatients(c echo.Context) error {
	patients, err := p.patientService.GetPatients(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Failed to get all patients: "+err.Error())
	}

	response := responses.NewPatientResponse(patients)
	return responses.Response(c, http.StatusOK, response)
}

// UpdatePatient godoc
//
//	@Summary		Update patient
//	@Description	Update patient
//	@ID				patients-update
//	@Tags			Patients Actions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"Patient ID"
//	@Param			params	body		requests.UpdatePatientRequest	true	"Patient name and age"
//	@Success		200		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Failure		404		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/patients/{id} [put]
func (p *PatientHandlers) UpdatePatient(c echo.Context) error {
	// auth, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	parsedPatientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	patientID, err := safecast.Convert[uint](parsedPatientID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	var updatePatientRequest requests.UpdatePatientRequest
	if err := c.Bind(&updatePatientRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := updatePatientRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	_, err = p.patientService.UpdatePatient(c.Request().Context(), domain.UpdatePatientRequest{
		PatientID: patientID,
		Name:      updatePatientRequest.Name,
		Age:       updatePatientRequest.Age,
		Gender:    updatePatientRequest.Gender,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPostNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "Post not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "Forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to update post: "+err.Error())
		}
	}

	return responses.MessageResponse(c, http.StatusOK, "Post successfully updated")
}

// DeletePatient godoc
//
//	@Summary		Delete patient
//	@Description	Delete patient
//	@ID				patients-delete
//	@Tags			Patients Actions
//	@Param			id	path		int	true	"Patient ID"
//	@Success		204	{object}	responses.Data
//	@Failure		404	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/patients/{id} [delete]
func (p *PatientHandlers) DeletePatient(c echo.Context) error {
	// auth, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	PatientID, err := safecast.Convert[uint](parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	err = p.patientService.DeletePatient(c.Request().Context(), domain.DeletePatientRequest{
		PatientID: PatientID,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPostNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "Patient not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "Forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete patient: "+err.Error())
		}
	}

	return responses.MessageResponse(c, http.StatusNoContent, "Patient deleted successfully")
}
