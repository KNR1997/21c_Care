package handlers

import (
	"context"
	"errors"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
	"go-echo-starter/internal/requests"
	"go-echo-starter/internal/responses"
	"net/http"
	"strconv"

	safecast "github.com/ccoveille/go-safecast"
	"github.com/labstack/echo/v4"
)

type PatientService interface {
	Create(ctx context.Context, patient *models.Patient) error
	List(ctx context.Context) ([]models.Patient, error)
	ListPaginated(pagination repositories.Pagination[models.Patient]) (*repositories.Pagination[models.Patient], error)
	Get(ctx context.Context, id uint) (models.Patient, error)
	Update(ctx context.Context, request domain.UpdatePatientRequest) (*models.Patient, error)
	Delete(ctx context.Context, request domain.DeletePatientRequest) error
}

type PatientHandler struct {
	svc PatientService
}

func NewPatientHandler(svc PatientService) *PatientHandler {
	return &PatientHandler{svc: svc}
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
func (p *PatientHandler) Create(c echo.Context) error {
	var createPatientRequest requests.CreatePatientRequest
	if err := c.Bind(&createPatientRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := createPatientRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	patient := &models.Patient{
		Name:   createPatientRequest.Name,
		Age:    createPatientRequest.Age,
		Gender: createPatientRequest.Gender,
	}

	if err := p.svc.Create(c.Request().Context(), patient); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to create patient: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "patient successfully created")
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
func (p *PatientHandler) List(c echo.Context) error {
	patients, err := p.svc.List(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list patients: "+err.Error())
	}

	response := responses.NewPatientsResponse(patients)
	return responses.Response(c, http.StatusOK, response)
}

func (p *PatientHandler) ListPaginated(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	sort := c.QueryParam("sort")

	pagination := repositories.Pagination[models.Patient]{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	paginatedResult, err := p.svc.ListPaginated(pagination)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list paginated patients: "+err.Error())
	}

	return responses.Response(c, http.StatusOK, paginatedResult)
}

func (p *PatientHandler) Get(c echo.Context) error {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	patient, err := p.svc.Get(c.Request().Context(), uint(patientID))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to get patient: "+err.Error())
	}

	response := responses.NewPatientResponse(patient)
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
func (p *PatientHandler) Update(c echo.Context) error {
	parsedPatientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	patientID, err := safecast.Convert[uint](parsedPatientID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	var updatePatientRequest requests.UpdatePatientRequest
	if err := c.Bind(&updatePatientRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to bind request: "+err.Error())
	}

	if err := updatePatientRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	patient, err := p.svc.Update(c.Request().Context(), domain.UpdatePatientRequest{
		PatientID: patientID,
		Name:      updatePatientRequest.Name,
		Age:       updatePatientRequest.Age,
		Gender:    updatePatientRequest.Gender,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPostNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "patient not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "failed to update post: "+err.Error())
		}
	}

	response := responses.NewPatientResponse(*patient)
	return responses.Response(c, http.StatusOK, response)
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
func (p *PatientHandler) Delete(c echo.Context) error {
	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	PatientID, err := safecast.Convert[uint](parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	err = p.svc.Delete(c.Request().Context(), domain.DeletePatientRequest{
		PatientID: PatientID,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPostNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "patient not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "failed to delete patient: "+err.Error())
		}
	}

	return responses.MessageResponse(c, http.StatusNoContent, "patien deleted")
}
