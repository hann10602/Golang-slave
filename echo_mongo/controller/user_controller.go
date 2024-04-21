package controller

import (
	"echo_mongo/common"
	dto "echo_mongo/dto/user"
	"echo_mongo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.IUserService
}

func (ctl *UserController) GetUsers(c echo.Context) error {
	users, err := ctl.UserService.HandleGetUsers(c)

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
		Message:    "Searched data successfully",
		Data:       users,
	})
}

func (ctl *UserController) GetUserById(c echo.Context) error {
	userId := c.Param("id")
	user, err := ctl.UserService.HandleGetUserById(c, userId)

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
		Message:    "Searched data successfully",
		Data:       user,
	})
}

func (ctl *UserController) UpdateUser(c echo.Context) error {
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

	objectId, err := ctl.UserService.HandleUpdateUser(c, user, userId)

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
		Message:    "Updated data successfully",
		Data:       objectId,
	})
}

func (ctl *UserController) DeleteUser(c echo.Context) error {
	userId := c.Param("id")

	err := ctl.UserService.HandleDeleteUser(c, userId)

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
		Message:    "Deleted data successfully",
		Data:       true,
	})
}
