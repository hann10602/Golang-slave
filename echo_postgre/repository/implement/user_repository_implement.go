package implement

import (
	"context"
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
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

func (u *UserImplement) Create(ctx context.Context, data dtoReq.CreateUserDTO) (uint, error) {
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

func (u *UserImplement) GetById(ctx context.Context, id uint) (*dtoResp.UserResponseDTO, error) {
	dto := dtoResp.UserResponseDTO{}

	if err := u.db.Table(enum.USER_TABLE).Select("users.id, users.username, users.password, users.role, users.status, users.created_at, users.updated_at, settings.is_notification, settings.is_receive_message, settings.language").Joins("LEFT JOIN settings on users.id = settings.user_id").Where("users.id = ?", id).Scan(&dto).Error; err != nil {
		return nil, err
	}

	return &dto, nil
}

func (u *UserImplement) Search(ctx context.Context, filter *common.Filter, paging *common.Paging) (*[]dtoResp.UserResponseDTO, error) {
	var dto []dtoResp.UserResponseDTO

	u.db.Table(enum.USER_TABLE).Preload("Settings").Where("status <> ?", "DELETED")

	if filter != nil {
		if filter.Role != "" {
			u.db = u.db.Where("users.role = ?", filter.Role)
		}
		if filter.Status != "" {
			u.db = u.db.Where("users.status = ?", filter.Status)
		}
		if filter.Username != "" {
			u.db = u.db.Where("users.username = ?", filter.Username)
		}
	}

	if err := u.db.Table(enum.USER_TABLE).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := u.db.Table(enum.USER_TABLE).Select("users.id, users.username, users.password, users.role, users.status, users.created_at, users.updated_at, settings.is_notification, settings.is_receive_message, settings.language").Joins("LEFT JOIN settings on users.id = settings.user_id").Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Scan(&dto).Error; err != nil {
		return nil, err
	}

	return &dto, nil
}

func (u *UserImplement) Update(ctx context.Context, cond map[string]interface{}, data dtoReq.UpdateUserDTO) error {
	if err := u.db.Table(enum.USER_TABLE).Where(cond).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
