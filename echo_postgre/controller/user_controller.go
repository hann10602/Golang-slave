package controller

import (
	"echo_postgre/common"
	dto "echo_postgre/dto/req"
	"echo_postgre/service"
	"net/http"

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

	if err := u.userService.HandleDeleteUsers(ctx, map[string]interface{}{
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
	userId := ctx.Param("id")

	if userId == "" {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot received user id",
			Data:       false,
		})
	}

	user, err := u.userService.HandleGetUserById(ctx, map[string]interface{}{
		"id": userId,
	})

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
	// filter := ctx.Bind("filter")
	// paging := ctx.FormValue("paging")

	filter := common.Filter{
		Status: "ACTIVE",
	}

	paging := common.Paging{
		Page: 1,
	}

	data, err := u.userService.HandleSearchUsers(ctx, filter, paging)

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
		Data:       data,
	})
}

func (u UserController) UpdateUser(ctx echo.Context) error {
	var user dto.UpdateUserDTO

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

	err := u.userService.HandleUpdateUsers(ctx, map[string]interface{}{
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
