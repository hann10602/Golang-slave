package routes

import (
	"echo_mongo/controller"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo           *echo.Echo
	UserController controller.UserController
}
