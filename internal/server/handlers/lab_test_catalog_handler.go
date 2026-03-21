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

type LabTestCatalogService interface {
	Create(ctx context.Context, labtestcatalog *models.LabTestCatalog) error
	List(ctx context.Context) ([]models.LabTestCatalog, error)
	ListPaginated(pagination repositories.Pagination[models.LabTestCatalog]) (*repositories.Pagination[models.LabTestCatalog], error)
	Get(ctx context.Context, id uint) (models.LabTestCatalog, error)
	Update(ctx context.Context, request domain.UpdateLabTestCatalogRequest) (*models.LabTestCatalog, error)
	Delete(ctx context.Context, request domain.DeleteLabTestCatalogRequest) error
}

type LabTestCatalogHandler struct {
	svc LabTestCatalogService
}

func NewLabTestCatalogHandlers(svc LabTestCatalogService) *LabTestCatalogHandler {
	return &LabTestCatalogHandler{svc: svc}
}

// CreateLabTestCatalog godoc
//
//	@Summary		Create labtestcatalog
//	@Description	Create labtestcatalog
//	@ID				labtestcatalogs-create
//	@Tags			LabTestCatalogs Actions
//	@Accept			json
//	@Produce		json
//	@Param			params	body		requests.CreateLabTestCatalogRequest	true	"LabTestCatalog name and age"
//	@Success		201		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/labTestCatalogs [post]
func (p *LabTestCatalogHandler) Create(c echo.Context) error {
	// authClaims, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	var createLabTestCatalogRequest requests.CreateLabTestCatalogRequest
	if err := c.Bind(&createLabTestCatalogRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := createLabTestCatalogRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	labtestcatalog := &models.LabTestCatalog{
		Name:         createLabTestCatalogRequest.Name,
		DefaultPrice: createLabTestCatalogRequest.DefaultPrice,
	}

	if err := p.svc.Create(c.Request().Context(), labtestcatalog); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to create labtestcatalog: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "LabTestCatalog successfully created")
}

// GetLabTestCatalogs godoc
//
//	@Summary		Get labtestcatalogs
//	@Description	Get the list of all labtestcatalogs
//	@ID				labtestcatalogs-get
//	@Tags			LabTestCatalogs Actions
//	@Produce		json
//	@Success		200	{array}	responses.LabTestCatalogResponse
//	@Security		ApiKeyAuth
//	@Router			/labTestCatalogs [get]
func (p *LabTestCatalogHandler) List(c echo.Context) error {
	labtestcatalogs, err := p.svc.List(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list lab test catalogs: "+err.Error())
	}

	response := responses.NewLabTestCatalogsResponse(labtestcatalogs)
	return responses.Response(c, http.StatusOK, response)
}

func (p *LabTestCatalogHandler) ListPaginated(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	sort := c.QueryParam("sort")

	pagination := repositories.Pagination[models.LabTestCatalog]{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

	paginatedResult, err := p.svc.ListPaginated(pagination)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list paginated lab test catalogs: "+err.Error())
	}

	return responses.Response(c, http.StatusOK, paginatedResult)
}

func (p *LabTestCatalogHandler) Get(c echo.Context) error {
	labtestcatalogID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	labtestcatalog, err := p.svc.Get(c.Request().Context(), uint(labtestcatalogID))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to get lab test catalog: "+err.Error())
	}

	response := responses.NewLabTestCatalogResponse(labtestcatalog)
	return responses.Response(c, http.StatusOK, response)
}

// UpdateLabTestCatalog godoc
//
//	@Summary		Update labtestcatalog
//	@Description	Update labtestcatalog
//	@ID				labtestcatalogs-update
//	@Tags			LabTestCatalogs Actions
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int								true	"LabTestCatalog ID"
//	@Param			params	body		requests.UpdateLabTestCatalogRequest	true	"LabTestCatalog name and age"
//	@Success		200		{object}	responses.Data
//	@Failure		400		{object}	responses.Error
//	@Failure		404		{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/labTestCatalogs/{id} [put]
func (p *LabTestCatalogHandler) Update(c echo.Context) error {
	// auth, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	parsedLabTestCatalogID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	labtestcatalogID, err := safecast.Convert[uint](parsedLabTestCatalogID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	var updateLabTestCatalogRequest requests.UpdateLabTestCatalogRequest
	if err := c.Bind(&updateLabTestCatalogRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := updateLabTestCatalogRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	_, err = p.svc.Update(c.Request().Context(), domain.UpdateLabTestCatalogRequest{
		LabTestCatalogID: labtestcatalogID,
		Name:             updateLabTestCatalogRequest.Name,
		DefaultPrice:     updateLabTestCatalogRequest.DefaultPrice,
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

// DeleteLabTestCatalog godoc
//
//	@Summary		Delete labtestcatalog
//	@Description	Delete labtestcatalog
//	@ID				labtestcatalogs-delete
//	@Tags			LabTestCatalogs Actions
//	@Param			id	path		int	true	"LabTestCatalog ID"
//	@Success		204	{object}	responses.Data
//	@Failure		404	{object}	responses.Error
//	@Security		ApiKeyAuth
//	@Router			/labTestCatalogs/{id} [delete]
func (p *LabTestCatalogHandler) Delete(c echo.Context) error {
	// auth, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	LabTestCatalogID, err := safecast.Convert[uint](parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	err = p.svc.Delete(c.Request().Context(), domain.DeleteLabTestCatalogRequest{
		LabTestCatalogID: LabTestCatalogID,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPostNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "LabTestCatalog not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "Forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete labtestcatalog: "+err.Error())
		}
	}

	return responses.MessageResponse(c, http.StatusNoContent, "LabTestCatalog deleted successfully")
}
