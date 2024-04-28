package repository

import (
	"context"
	"echo_postgre/common"
	dto "echo_postgre/dto/req/users"
	"echo_postgre/model"
)

type IUserRepository interface {
	Search(context.Context, common.Filter, common.Paging) (*[]model.Users, error)
	GetById(context.Context, map[string]interface{}) (*model.Users, error)
	Create(context.Context, dto.CreateUserDTO) (uint, error)
	Update(context.Context, map[string]interface{}, dto.UpdateUserDTO) error
	Delete(context.Context, map[string]interface{}) error
}
