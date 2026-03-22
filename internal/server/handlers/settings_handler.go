package handlers

import (
	"context"
	"go-echo-starter/internal/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type settingsService interface {
	Get(ctx context.Context) (map[string]interface{}, error)
}

type SettingsHandler struct {
	settingsService settingsService
}

func NewSettingsHandler(settingsService settingsService) *SettingsHandler {
	return &SettingsHandler{settingsService: settingsService}
}

func (h *SettingsHandler) Get(c echo.Context) error {
	settings, err := h.settingsService.Get(c.Request().Context())
	if err != nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Failed to get settings: "+err.Error())
	}

	return c.JSON(http.StatusOK, settings)
}
