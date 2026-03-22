package handlers

import (
	"context"
	"go-echo-starter/internal/services/report"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type reportService interface {
	GenerateVisitPDF(ctx context.Context, visitID uint) ([]byte, error)
}

type ReportHandler struct {
	svc reportService
}

func NewReportHandler(reportService *report.Service) *ReportHandler {
	return &ReportHandler{svc: reportService}
}

func (h *ReportHandler) GeneratePDF(c echo.Context) error {

	idParam := c.Param("id")

	visitID, _ := strconv.Atoi(idParam)

	visitIDUint := uint(visitID)

	pdfBytes, err := h.svc.GenerateVisitPDF(c.Request().Context(), visitIDUint)
	if err != nil {
		return err
	}

	c.Response().Header().Set(
		"Content-Disposition",
		"attachment; filename=invoice.pdf",
	)

	return c.Blob(
		http.StatusOK,
		"application/pdf",
		pdfBytes,
	)
}
