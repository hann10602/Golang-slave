package repository

import (
	"context"
	dto "echo_mongo/dto/user"
	model "echo_mongo/model/user"
)

type IUserRepository interface {
	GetUsers(context.Context) ([]*model.User, error)
	GetUserById(context.Context, string) (*model.User, error)
	UpdateUser(context.Context, dto.UpdateUserDto, string) (*model.User, error)
	DeleteUser(context.Context, string) error
}
