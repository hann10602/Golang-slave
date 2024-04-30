package route

import (
	"echo_postgre/controller"
	"echo_postgre/initializer"
	customMiddleware "echo_postgre/middleware"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo               *echo.Echo
	UserController     controller.UserController
	SettingsController controller.SettingsController
	AuthController     controller.AuthController
}

func (API *API) SetUpRouter() {
	config, err := initializer.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variable")
	}

	API.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.ClientOrigin},
		AllowCredentials: true,
	}))

	apiV1 := API.Echo.Group("api/v1")

	userGroup := apiV1.Group("/users")
	userGroup.Use(customMiddleware.VerifyToken)
	userGroup.GET("", API.UserController.SearchUser)
	userGroup.GET("/:id", API.UserController.GetUserById)
	userGroup.PUT("/:id", API.UserController.UpdateUser)
	userGroup.DELETE("/:id", API.UserController.DeleteUser)

	settingsGroup := apiV1.Group("/settings")
	settingsGroup.Use(customMiddleware.VerifyToken)
	settingsGroup.PUT("/:id", API.SettingsController.UpdateSettings)

	authGroup := apiV1.Group("/auth")
	authGroup.POST("/login", API.AuthController.Login)
	authGroup.POST("/register", API.AuthController.Register)
}
