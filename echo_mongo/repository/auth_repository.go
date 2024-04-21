package repository

import (
	"context"
	dto "echo_mongo/dto/auth"
	model "echo_mongo/model/user"
)

type IAuthRepository interface {
	Login(context.Context, dto.LoginDto) (*model.User, error)
	Register(context.Context, model.User) (*model.User, error)
}
