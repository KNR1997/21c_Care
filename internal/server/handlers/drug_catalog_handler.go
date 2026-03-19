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

type drugcatalogService interface {
	Create(ctx context.Context, drugcatalog *models.DrugCatalog) error
	GetDrugCatalogs(ctx context.Context) ([]models.DrugCatalog, error)
	// GetDrugCatalog(ctx context.Context, id uint) (models.DrugCatalog, error)
	UpdateDrugCatalog(ctx context.Context, request domain.UpdateDrugCatalogRequest) (*models.DrugCatalog, error)
	DeleteDrugCatalog(ctx context.Context, request domain.DeleteDrugCatalogRequest) error
}

type DrugCatalogHandlers struct {
	drugcatalogService drugcatalogService
}

func NewDrugCatalogHandlers(drugcatalogService drugcatalogService) *DrugCatalogHandlers {
	return &DrugCatalogHandlers{drugcatalogService: drugcatalogService}
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
func (p *DrugCatalogHandlers) CreateDrugCatalog(c echo.Context) error {
	// authClaims, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	var createDrugCatalogRequest requests.CreateDrugCatalogRequest
	if err := c.Bind(&createDrugCatalogRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := createDrugCatalogRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	drugcatalog := &models.DrugCatalog{
		Name:         createDrugCatalogRequest.Name,
		DefaultPrice: createDrugCatalogRequest.DefaultPrice,
	}

	if err := p.drugcatalogService.Create(c.Request().Context(), drugcatalog); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to create drugcatalog: "+err.Error())
	}

	return responses.MessageResponse(c, http.StatusCreated, "DrugCatalog successfully created")
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
func (p *DrugCatalogHandlers) GetDrugCatalogs(c echo.Context) error {
	drugcatalogs, err := p.drugcatalogService.GetDrugCatalogs(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Failed to get all drugcatalogs: "+err.Error())
	}

	response := responses.NewDrugCatalogResponse(drugcatalogs)
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
func (p *DrugCatalogHandlers) UpdateDrugCatalog(c echo.Context) error {
	// auth, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	parsedDrugCatalogID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	drugcatalogID, err := safecast.Convert[uint](parsedDrugCatalogID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	var updateDrugCatalogRequest requests.UpdateDrugCatalogRequest
	if err := c.Bind(&updateDrugCatalogRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to bind request: "+err.Error())
	}

	if err := updateDrugCatalogRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	_, err = p.drugcatalogService.UpdateDrugCatalog(c.Request().Context(), domain.UpdateDrugCatalogRequest{
		DrugCatalogID: drugcatalogID,
		Name:          updateDrugCatalogRequest.Name,
		DefaultPrice:  updateDrugCatalogRequest.DefaultPrice,
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
func (p *DrugCatalogHandlers) DeleteDrugCatalog(c echo.Context) error {
	// auth, err := getAuthClaims(c)
	// if err != nil {
	// 	return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
	// }

	parsedID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	DrugCatalogID, err := safecast.Convert[uint](parsedID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Failed to parse post id: "+err.Error())
	}

	err = p.drugcatalogService.DeleteDrugCatalog(c.Request().Context(), domain.DeleteDrugCatalogRequest{
		DrugCatalogID: DrugCatalogID,
	})
	if err != nil {
		switch {
		case errors.Is(err, models.ErrPostNotFound):
			return responses.ErrorResponse(c, http.StatusNotFound, "DrugCatalog not found")
		case errors.Is(err, models.ErrForbidden):
			return responses.ErrorResponse(c, http.StatusForbidden, "Forbidden")
		default:
			return responses.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete drugcatalog: "+err.Error())
		}
	}

	return responses.MessageResponse(c, http.StatusNoContent, "DrugCatalog deleted successfully")
}
