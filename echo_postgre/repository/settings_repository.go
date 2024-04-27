package repository

import (
	"context"
	dto "echo_postgre/dto/req/settings"
)

type ISettingsRepository interface {
	Update(context.Context, map[string]interface{}, *dto.UpdateSettingsDTO) error
	// Delete(context.Context, map[string]interface{}) error
}
