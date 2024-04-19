package controller

import (
	"echo_mongo/common"
	dto "echo_mongo/dto/user"
	model "echo_mongo/model/user"
	"echo_mongo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.IUserService
}

func (u *UserController) GetUsers(c echo.Context) error {
	users, err := u.UserService.HandleGetUsers(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	return c.JSON(http.StatusOK, common.Response{
		StatusCode: http.StatusOK,
		Message:    "Search data successfully",
		Data:       users,
	})
}

func (u *UserController) GetUserById(c echo.Context) error {
	userId := c.Param("id")
	user, err := u.UserService.HandleGetUserById(c, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	return c.JSON(http.StatusOK, common.Response{
		StatusCode: http.StatusOK,
		Message:    "Search data successfully",
		Data:       user,
	})
}

func (u *UserController) CreateUser(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	objectId, err := u.UserService.HandleCreateUser(c, user)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	return c.JSON(http.StatusOK, common.Response{
		StatusCode: http.StatusOK,
		Message:    "Search data successfully",
		Data:       objectId,
	})
}

func (u *UserController) UpdateUser(c echo.Context) error {
	var user dto.UpdateUserDto

	userId := c.Param("id")

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	objectId, err := u.UserService.HandleUpdateUser(c, user, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	return c.JSON(http.StatusOK, common.Response{
		StatusCode: http.StatusOK,
		Message:    "Update data successfully",
		Data:       objectId,
	})
}

func (u *UserController) DeleteUser(c echo.Context) error {
	userId := c.Param("id")

	err := u.UserService.HandleDeleteUser(c, userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	return c.JSON(http.StatusOK, common.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete data successfully",
		Data:       true,
	})
}
