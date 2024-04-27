package controller

import (
	"echo_postgre/common"
	"echo_postgre/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SearchUser(echo.Context) error
	GetUserById(echo.Context) error
	CreateUser(echo.Context) error
	UpdateUser(echo.Context) error
	DeleteUser(echo.Context) error
}

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) CreateUser(ctx echo.Context) error {
}

func (u *UserController) DeleteUser(ctx echo.Context) error {
	panic("unimplemented")
}

func (u *UserController) GetUserById(ctx echo.Context) error {
	panic("unimplemented")
}

func (u *UserController) SearchUser(ctx echo.Context) error {
	filter := ctx.Bind("filter")
	paging := ctx.FormValue("paging")

	data, err := u.userService.HandleSearchUsers(ctx, &filter, &paging)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Search users successfully",
		Data:       data,
	})
}

func (u *UserController) UpdateUser(ctx echo.Context) error {
	panic("unimplemented")
}
