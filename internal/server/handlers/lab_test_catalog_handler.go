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

type labtestcatalogService interface {
	Create(ctx context.Context, labtestcatalog *models.LabTestCatalog) error
	GetLabTestCatalogs(ctx context.Context) ([]models.LabTestCatalog, error)
	// GetLabTestCatalog(ctx context.Context, id uint) (models.LabTestCatalog, error)
	UpdateLabTestCatalog(ctx context.Context, request domain.UpdateLabTestCatalogRequest) (*models.LabTestCatalog, error)
	DeleteLabTestCatalog(ctx context.Context, request domain.DeleteLabTestCatalogRequest) error
}

type LabTestCatalogHandlers struct {
	labtestcatalogService labtestcatalogService
}

func NewLabTestCatalogHandlers(labtestcatalogService labtestcatalogService) *LabTestCatalogHandlers {
	return &LabTestCatalogHandlers{labtestcatalogService: labtestcatalogService}
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
//	@Router			/labtestcatalogs [post]
func (p *LabTestCatalogHandlers) CreateLabTestCatalog(c echo.Context) error {
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

	if err := p.labtestcatalogService.Create(c.Request().Context(), labtestcatalog); err != nil {
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
//	@Router			/labtestcatalogs [get]
func (p *LabTestCatalogHandlers) GetLabTestCatalogs(c echo.Context) error {
	labtestcatalogs, err := p.labtestcatalogService.GetLabTestCatalogs(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Failed to get all labtestcatalogs: "+err.Error())
	}

	response := responses.NewLabTestCatalogResponse(labtestcatalogs)
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
//	@Router			/labtestcatalogs/{id} [put]
func (p *LabTestCatalogHandlers) UpdateLabTestCatalog(c echo.Context) error {
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

	_, err = p.labtestcatalogService.UpdateLabTestCatalog(c.Request().Context(), domain.UpdateLabTestCatalogRequest{
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
//	@Router			/labtestcatalogs/{id} [delete]
func (p *LabTestCatalogHandlers) DeleteLabTestCatalog(c echo.Context) error {
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

	err = p.labtestcatalogService.DeleteLabTestCatalog(c.Request().Context(), domain.DeleteLabTestCatalogRequest{
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
