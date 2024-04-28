package repository

import (
	"context"
	dto "echo_postgre/dto/req/settings"
)

type ISettingsRepository interface {
	Create(context.Context, uint) error
	Update(context.Context, map[string]interface{}, *dto.UpdateSettingsDTO) error
}
