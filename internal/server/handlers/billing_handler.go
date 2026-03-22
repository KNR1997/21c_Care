package handlers

import (
	"context"
	"go-echo-starter/internal/models"
	"go-echo-starter/internal/repositories"
	"go-echo-starter/internal/requests"
	"go-echo-starter/internal/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type billingService interface {
	// List(ctx context.Context) ([]models.Billing, error)
	ListPaginated(pagination repositories.Pagination[models.Billing]) (*repositories.Pagination[models.Billing], error)
	Create(ctx context.Context, visitID int64) (*models.Billing, error)
	Get(ctx context.Context, id uint) (models.Billing, error)
}

type BillingHandler struct {
	svc billingService
}

func NewBillingHandlers(billingService billingService) *BillingHandler {
	return &BillingHandler{svc: billingService}
}

// func (p *BillingHandler) List(c echo.Context) error {
// 	billings, err := p.svc.List(c.Request().Context())
// 	if err != nil {
// 		return responses.ErrorResponse(c, http.StatusNotFound, "failed to list billings: "+err.Error())
// 	}

// 	response := responses.NewBillingsResponse(billings)
// 	return responses.Response(c, http.StatusOK, response)
// }

func (p *BillingHandler) ListPaginated(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	sort := c.QueryParam("sort")

	pagination := repositories.Pagination[models.Billing]{
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

func (b *BillingHandler) Create(c echo.Context) error {
	var createBillingRequest requests.CreateBillingRequest
	if err := c.Bind(&createBillingRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to bind request: "+err.Error())
	}

	if err := createBillingRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid request")
	}

	billing, err := b.svc.Create(c.Request().Context(), createBillingRequest.VisitID)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "failed to create billing: "+err.Error())
	}

	response := responses.NewBillingResponse(*billing)
	return responses.Response(c, http.StatusCreated, response)
}

func (p *BillingHandler) Get(c echo.Context) error {
	billingID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "invalid id")
	}

	billing, err := p.svc.Get(c.Request().Context(), uint(billingID))
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "failed to get billing: "+err.Error())
	}

	response := responses.NewBillingResponse(billing)
	return responses.Response(c, http.StatusOK, response)
}
