package handlers

import (
	"context"
	"go-echo-starter/internal/domain"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
	"go-echo-starter/internal/requests"
	"go-echo-starter/internal/responses"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type visitService interface {
	Create(ctx context.Context, visit *models.Visit, aiResult *domain.AIResponse) error
	List(ctx context.Context) ([]models.Visit, error)
	ListPaginated(pagination repositories.Pagination[models.Visit]) (*repositories.Pagination[models.Visit], error)
	Get(ctx context.Context, id uint) (models.Visit, error)
	Preview(ctx context.Context, rawInput string) (*domain.AIResponse, error)
}

type VisitHandlers struct {
	svc visitService
}

func NewVisitHandlers(visitService visitService) *VisitHandlers {
	return &VisitHandlers{svc: visitService}
}

// CreateVisit godoc
//
//	@Summary		Create visit
//	@Description	Create visit
//	@ID				visits-create
//	@Tags			Visits Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreateVisitRequest	true	"Visit title and content"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/visits [post]
func (v *VisitHandlers) Create(c echo.Context) error {
	var createVisitRequest requests.CreateVisitRequest
	if err := c.Bind(&createVisitRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to bind request: "+err.Error())
	}

	if err := createVisitRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	visit := &models.Visit{
		PatientID: createVisitRequest.PatientID,
		RawInput:  createVisitRequest.RawInput,
	}

	if err := v.svc.Create(c.Request().Context(), visit, createVisitRequest.AIResult); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to create visit: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "visit created")
}

// GetVisits godoc
//
//	@Summary		Get visits
//	@Description	Get the list of all visits
//	@ID				visits-get
//	@Tags			Visits Actions
//	@Produce		json
//	@Success		200	{array}	responses.VisitResponse
//	@Security		ApiKeyAuth
//	@Router			/visits [get]
// func (p *VisitHandlers) List(c echo.Context) error {
// 	visits, err := p.svc.List(c.Request().Context())
// 	if err != nil {
// 		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list visits: "+err.Error())
// 	}

// 	response := responses.NewVisitsResponse(visits)
// 	return responses.Response(c, http.StatusOK, response)
// }

// ListPaginated godoc
//
//	@Summary		List visits with pagination
//	@Description	Get the list of visits with pagination, sorting, and filtering options
//	@ID				visits-list-paginated
//	@Tags			Visits Actions
//	@Produce		json
//	@Param			limit	query		int		false	"Number of items per page"	default(10)
//	@Param			page	query		int		false	"Page number"				default(1)
//	@Param			sort	query		string	false	"Sort field (e.g., created_at, updated_at)"
//	@Success		200		{object}	responses.VisitPaginationResponse	"Paginated list of visits"
//	@Failure		404		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/visits [get]
func (p *VisitHandlers) ListPaginated(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	sort := c.QueryParam("sort")

	pagination := repositories.Pagination[models.Visit]{
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

// GetVisit godoc
//
//	@Summary		Get visit by ID
//	@Description	Get a single visit by its unique identifier
//	@ID				visits-get-by-id
//	@Tags			Visits Actions
//	@Produce		json
//	@Param			id	path		int	true	"Visit ID"
//	@Success		200	{object}	responses.VisitResponse
//	@Failure		400	{object}	responses.Error
//	@Failure		404	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/visits/{id} [get]
func (p *VisitHandlers) Get(c echo.Context) error {
	patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	visit, err := p.svc.Get(c.Request().Context(), uint(patientID))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to get patient: "+err.Error())
	}

	response := responses.NewVisitResponse(visit)
	return responses.Response(c, http.StatusOK, response)
}

// PreviewVisit godoc
//
//	@Summary		Classify visit prompt
//	@Description	Preview a visit's prompt details by classifying using AI
//	@ID				visits-preview
//	@Tags			Visits Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body	requests.PreviewVisitRequest	true	"Visit title and content"
//	@Success		200	{object}	responses.VisitResponse
//	@Failure		400	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/visits/preview [post]
func (h *VisitHandlers) Preview(c echo.Context) error {
	var requestData requests.PreviewVisitRequest

	if err := c.Bind(&requestData); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to bind request: "+err.Error())
	}

	if err := requestData.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	result, err := h.svc.Preview(c.Request().Context(), requestData.RawInput)
	slog.Debug("AI preview result", "result", result, "error", err)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
