package service

import (
	"context"
	dto "echo_postgre/dto/req/settings"
	"echo_postgre/repository"
)

type SettingsService struct {
	settingsRepository repository.ISettingsRepository
}

type ISettingsService interface {
	HandleUpdateSettings(context.Context, map[string]interface{}, *dto.UpdateSettingsDTO) error
}

func NewSettingsService(settingsRepository repository.ISettingsRepository) ISettingsService {
	return &SettingsService{
		settingsRepository: settingsRepository,
	}
}

func (s *SettingsService) HandleUpdateSettings(ctx context.Context, cond map[string]interface{}, data *dto.UpdateSettingsDTO) error {
	if err := s.settingsRepository.Update(ctx, cond, data); err != nil {
		return err
	}

	return nil
}
