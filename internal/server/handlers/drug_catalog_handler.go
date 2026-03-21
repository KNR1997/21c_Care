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

type DrugCatalogService interface {
	Create(ctx context.Context, catalog *models.DrugCatalog) error
	List(ctx context.Context) ([]models.DrugCatalog, error)
	ListPaginated(pagination repositories.Pagination[models.DrugCatalog]) (*repositories.Pagination[models.DrugCatalog], error)
	Get(ctx context.Context, id uint) (models.DrugCatalog, error)
	Update(ctx context.Context, request domain.UpdateDrugCatalogRequest) (*models.DrugCatalog, error)
	Delete(ctx context.Context, request domain.DeleteDrugCatalogRequest) error
}

type DrugCatalogHandler struct {
	svc DrugCatalogService
}

func NewDrugCatalogHandlers(svc DrugCatalogService) *DrugCatalogHandler {
	return &DrugCatalogHandler{svc: svc}
}

// CreateDrugCatalog godoc
//
//	@Summary		Create drugcatalog
//	@Description	Create drugcatalog
//	@ID				drugcatalogs-create
//	@Tags			DrugCatalogs Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreateDrugCatalogRequest	true	"DrugCatalog name and age"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/drugCatalogs [post]
func (h *DrugCatalogHandler) Create(c echo.Context) error {
	var req requests.CreateDrugCatalogRequest
	if err := c.Bind(&req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to bind request: "+err.Error())
	}

	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	catalog := &models.DrugCatalog{
		Name:         req.Name,
		DefaultPrice: req.DefaultPrice,
	}

	if err := h.svc.Create(c.Request().Context(), catalog); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to create drug catalog: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "drug catalog created")
}

// GetDrugCatalogs godoc
//
//	@Summary		Get drugcatalogs
//	@Description	Get the list of all drugcatalogs
//	@ID				drugcatalogs-get
//	@Tags			DrugCatalogs Actions
//	@Produce		json
//	@Success		200	{array}	responses.DrugCatalogResponse
//	@Security		ApiKeyAuth
//	@Router			/drugCatalogs [get]
func (h *DrugCatalogHandler) List(c echo.Context) error {
	catalogs, err := h.svc.List(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "failed to list drug catalogs: "+err.Error())
	}

	return responses.Response(c, http.StatusOK, responses.NewDrugCatalogsResponse(catalogs))
}

func (p *DrugCatalogHandler) ListPaginated(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	sort := c.QueryParam("sort")

	pagination := repositories.Pagination[models.DrugCatalog]{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	paginatedResult, err := p.svc.ListPaginated(pagination)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list paginated drug catalogs: "+err.Error())
	}

	return responses.Response(c, http.StatusOK, paginatedResult)
}

func (p *DrugCatalogHandler) Get(c echo.Context) error {
	drugID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	drug, err := p.svc.Get(c.Request().Context(), uint(drugID))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to get drug: "+err.Error())
	}

	response := responses.NewDrugCatalogResponse(drug)
	return responses.Response(c, http.StatusOK, response)
}

// UpdateDrugCatalog godoc
//
//	@Summary		Update drugcatalog
//	@Description	Update drugcatalog
//	@ID				drugcatalogs-update
//	@Tags			DrugCatalogs Actions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"DrugCatalog ID"
//	@Param			params	body		requests.UpdateDrugCatalogRequest	true	"DrugCatalog name and age"
//	@Success		200		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Failure		404		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/drugCatalogs/{id} [put]
func (h *DrugCatalogHandler) Update(c echo.Context) error {
	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	catalogID, err := safecast.Convert[uint](parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	var req requests.UpdateDrugCatalogRequest
	if err := c.Bind(&req); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to bind request: "+err.Error())
	}

	if err := req.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	_, err = h.svc.Update(c.Request().Context(), domain.UpdateDrugCatalogRequest{
		DrugCatalogID: catalogID,
		Name:          req.Name,
		DefaultPrice:  req.DefaultPrice,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDrugCatalogNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "drug catalog not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "failed to update drug catalog")
		}
	}

	return responses.MessageResponse(c, http.StatusOK, "drug catalog updated")
}

// DeleteDrugCatalog godoc
//
//	@Summary		Delete drugcatalog
//	@Description	Delete drugcatalog
//	@ID				drugcatalogs-delete
//	@Tags			DrugCatalogs Actions
//	@Param			id	path		int	true	"DrugCatalog ID"
//	@Success		204	{object}	responses.Data
//	@Failure		404	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/drugCatalogs/{id} [delete]
func (h *DrugCatalogHandler) Delete(c echo.Context) error {
	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	catalogID, err := safecast.Convert[uint](parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	if err := h.svc.Delete(c.Request().Context(), domain.DeleteDrugCatalogRequest{
		DrugCatalogID: catalogID,
	}); err != nil {
		switch {
		case errors.Is(err, models.ErrDrugCatalogNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "drug catalog not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "failed to delete drug catalog")
		}
	}

	return responses.MessageResponse(c, http.StatusNoContent, "drug catalog deleted")
}
