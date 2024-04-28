package repository

import (
	"context"
	dto "echo_postgre/dto/req"
	"echo_postgre/model"
)

type ISettingsRepository interface {
	GetByUserId(context.Context, uint) (*model.Settings, error)
	Create(context.Context, uint) error
	Update(context.Context, map[string]interface{}, dto.UpdateSettingsDTO) error
}
