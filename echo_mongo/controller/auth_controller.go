package controller

import (
	"echo_mongo/common"
	dto "echo_mongo/dto/auth"
	model "echo_mongo/model/user"
	"echo_mongo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService service.IAuthService
}

func (ctl *AuthController) Login(c echo.Context) error {
	var data dto.LoginDto

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	token, err := ctl.AuthService.HandleLogin(c, data)

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
		Message:    "Login successfully",
		Data:       token,
	})
}

func (ctl *AuthController) Register(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})

		return nil
	}

	data, err := ctl.AuthService.HandleRegister(c, user)

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
		Message:    "Register successfully",
		Data:       data,
	})
}
