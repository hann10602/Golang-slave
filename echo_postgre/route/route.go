package route

import "github.com/labstack/echo/v4"

func Init() *echo.Echo {
	e := echo.New()

	userGroup := e.Group("api/v1/users")
}
