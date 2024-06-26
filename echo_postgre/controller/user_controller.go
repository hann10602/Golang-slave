package controller

import (
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	"echo_postgre/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (u UserController) DeleteUser(ctx echo.Context) error {
	userId := ctx.Param("id")

	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received user id",
			Data:       false,
		})
	}

	if err := u.userService.HandleDeleteUser(ctx, map[string]interface{}{
		"id": userId,
	}); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Deleted user successfully",
		Data:       true,
	})
}

func (u UserController) GetUserById(ctx echo.Context) error {
	userId, err := strconv.ParseUint(ctx.Param("id"), 10, 32)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received user id",
			Data:       false,
		})
	}

	user, err := u.userService.HandleGetUserById(ctx, uint(userId))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Get user successfully",
		Data:       user,
	})
}

func (u UserController) SearchUser(ctx echo.Context) error {
	var paging common.Paging
	var filter common.Filter

	err := ctx.Bind(&paging)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	err = ctx.Bind(&filter)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	paging.Process()

	data, err := u.userService.HandleSearchUsers(ctx, &filter, &paging)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Search users successfully",
		Data:       common.HandleResponseWithPagination(data, paging),
	})
}

func (u UserController) UpdateUser(ctx echo.Context) error {
	var user dtoReq.UpdateUserDTO

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	userId := ctx.Param("id")

	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received user id",
			Data:       false,
		})
	}

	err := u.userService.HandleUpdateUser(ctx, map[string]interface{}{
		"id": userId,
	}, user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated user successfully",
		Data:       true,
	})
}
