package controller

import (
	"context"
	"echo_postgre/service"
)

type IUserController interface {
	SearchUser(context.Context) error
	GetUserById(context.Context) error
	CreateUser(context.Context) error
	UpdateUser(context.Context) error
	DeleteUser(context.Context) error
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{
		userService: userService,
	}
}
