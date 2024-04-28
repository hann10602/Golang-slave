package service

import (
	dto "echo_postgre/dto/req"
	"echo_postgre/repository"

	"github.com/labstack/echo/v4"
)

type SettingsService struct {
	settingsRepository repository.ISettingsRepository
}

type ISettingsService interface {
	HandleUpdateSettings(echo.Context, map[string]interface{}, dto.UpdateSettingsDTO) error
}

func NewSettingsService(settingsRepository repository.ISettingsRepository) ISettingsService {
	return &SettingsService{
		settingsRepository: settingsRepository,
	}
}

func (s *SettingsService) HandleUpdateSettings(ctx echo.Context, cond map[string]interface{}, data dto.UpdateSettingsDTO) error {
	if err := s.settingsRepository.Update(ctx.Request().Context(), cond, data); err != nil {
		return err
	}

	return nil
}
