package repository

import (
	"context"
	"echo_postgre/common"
	dtoRequest "echo_postgre/dto/req"
	dtoResponse "echo_postgre/dto/resp"
	"echo_postgre/model"
)

type IUserRepository interface {
	Search(context.Context, common.Filter, common.Paging) (*[]model.Users, error)
	GetById(context.Context, map[string]interface{}) (*dtoResponse.UserResponseDTO, error)
	Create(context.Context, dtoRequest.CreateUserDTO) (uint, error)
	Update(context.Context, map[string]interface{}, dtoRequest.UpdateUserDTO) error
	Delete(context.Context, map[string]interface{}) error
}
