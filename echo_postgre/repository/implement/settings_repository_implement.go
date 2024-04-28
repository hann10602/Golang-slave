package implement

import (
	"context"
	dto "echo_postgre/dto/req"
	"echo_postgre/enum"
	"echo_postgre/model"
	"echo_postgre/repository"

	"gorm.io/gorm"
)

type SettingsImplement struct {
	db *gorm.DB
}

func NewSettingsImplement(db *gorm.DB) repository.ISettingsRepository {
	return &SettingsImplement{
		db: db,
	}
}

// func (u *SettingsImplement) Delete(ctx context.Context, cond map[string]interface{}) error {
// 	if err := u.db.Table(enum.SETTINGS_TABLE).Where(cond).Updates(map[string]interface{}{
// 		"status": "DELETED",
// 	}).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func (u *SettingsImplement) GetByUserId(ctx context.Context, userId uint) (*model.Settings, error) {
	var settings model.Settings

	if err := u.db.Table(enum.SETTINGS_TABLE).Where("user_id = ?", userId).Create(&settings).Error; err != nil {
		return nil, err
	}

	return &settings, nil
}

func (u *SettingsImplement) Create(ctx context.Context, userId uint) error {
	settings := &model.Settings{
		UserId: userId,
	}

	if err := u.db.Table(enum.SETTINGS_TABLE).Create(&settings).Error; err != nil {
		return err
	}

	return nil
}

func (u *SettingsImplement) Update(ctx context.Context, cond map[string]interface{}, data dto.UpdateSettingsDTO) error {
	if err := u.db.Table(enum.SETTINGS_TABLE).Where(cond).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
