package repository

import (
	"context"
	model "echo_mongo/model/user"
)

type IUserRepository interface {
	GetUserById(context context.Context, id string) (*model.User, error)
}
