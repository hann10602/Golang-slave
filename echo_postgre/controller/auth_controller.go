package controller

import (
	"echo_postgre/common"
	dto "echo_postgre/dto/req"
	"echo_postgre/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

func (a AuthController) Register(ctx echo.Context) error {
	var user dto.CreateUserDTO

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	if err := a.authService.HandleRegister(ctx, user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Created user successfully",
		Data:       true,
	})
}

func (a AuthController) Login(ctx echo.Context) error {
	var user dto.LoginDto

	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	token, err := a.authService.HandleLogin(ctx, user)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, &common.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		})
	}

	return ctx.JSON(http.StatusOK, &common.Response{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       token,
	})
}
