package controller

import (
	"echo_postgre/common"
	dto "echo_postgre/dto/req/settings"
	"echo_postgre/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type SettingsController struct {
	settingsService service.ISettingsService
}

func NewSettingsController(settingsService service.ISettingsService) SettingsController {
	return SettingsController{
		settingsService: settingsService,
	}
}

func (s SettingsController) UpdateSettings(ctx echo.Context) error {
	var settings dto.UpdateSettingsDTO

	if err := ctx.Bind(&settings); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	id := ctx.Param("id")

	if id == "" {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received id",
			Data:       false,
		})
	}

	if err := s.settingsService.HandleUpdateSettings(ctx, map[string]interface{}{
		"id": id,
	}, &settings); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated settings successfully",
		Data:       true,
	})
}
