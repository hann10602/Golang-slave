package implement

import (
	"context"
	"echo_postgre/common"
	dto "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
	"echo_postgre/enum"
	"echo_postgre/model"
	"echo_postgre/repository"

	"gorm.io/gorm"
)

type UserImplement struct {
	db *gorm.DB
}

func NewUserImplement(db *gorm.DB) repository.IUserRepository {
	return &UserImplement{
		db: db,
	}
}

func (u *UserImplement) Create(ctx context.Context, data dto.CreateUserDTO) (uint, error) {
	if err := u.db.Table(enum.USER_TABLE).Create(&data).Error; err != nil {
		return data.Id, err
	}

	return data.Id, nil
}

func (u *UserImplement) GetByUsernameAndPassword(ctx context.Context, cond map[string]interface{}) (*model.Users, error) {
	var user model.Users

	if err := u.db.Table(enum.USER_TABLE).Where(cond).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserImplement) Delete(ctx context.Context, cond map[string]interface{}) error {
	if err := u.db.Table(enum.USER_TABLE).Where(cond).Updates(map[string]interface{}{
		"status": "DELETED",
	}).Error; err != nil {
		return err
	}

	return nil
}

func (u *UserImplement) GetById(ctx context.Context, cond map[string]interface{}) (*dtoResp.UserResponseDTO, error) {
	var user model.Users

	if err := u.db.Table(enum.USER_TABLE).Preload("Settings").Where(cond).First(&user).Error; err != nil {
		return nil, err
	}

	dto := dtoResp.UserResponseDTO{
		Id:        user.Id,
		Username:  user.Username,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Settings: dtoResp.SettingsResponseDTO{
			IsNotification:   user.Settings.IsNotification,
			IsReceiveMessage: user.Settings.IsReceiveMessage,
			Language:         user.Settings.Language,
		},
	}

	return &dto, nil
}

func (u *UserImplement) Search(ctx context.Context, filter common.Filter, paging common.Paging) (*[]model.Users, error) {
	var users []model.Users

	if err := u.db.Table(enum.USER_TABLE).Preload("Settings").Where("status <> ?", "DELETED").Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserImplement) Update(ctx context.Context, cond map[string]interface{}, data dto.UpdateUserDTO) error {
	if err := u.db.Table(enum.USER_TABLE).Where(cond).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
