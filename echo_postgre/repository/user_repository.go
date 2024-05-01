package repository

import (
	"context"
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
	"echo_postgre/model"
)

type IUserRepository interface {
	Search(context.Context, *common.Filter, *common.Paging) (*[]dtoResp.UserResponseDTO, error)
	GetById(context.Context, uint) (*dtoResp.UserResponseDTO, error)
	GetByUsernameAndPassword(context.Context, map[string]interface{}) (*model.Users, error)
	Create(context.Context, dtoReq.CreateUserDTO) (uint, error)
	Update(context.Context, map[string]interface{}, dtoReq.UpdateUserDTO) error
	Delete(context.Context, map[string]interface{}) error
}
