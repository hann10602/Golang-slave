package repository

import (
	"context"
	"echo_postgre/model"
)

type IAuthRepository interface {
	Login(context.Context) (*model.Users, error)
}