package routes

import (
	"echo_mongo/controller"
	"echo_mongo/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo           *echo.Echo
	UserController controller.UserController
	AuthController controller.AuthController
}

func (API *API) SetUpRouter() {
	// config, err := initializer.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Could not load environment variables", err)
	// }

	apiV1 := API.Echo.Group("api/v1")
	user := apiV1.Group("/user")
	user.Use(middleware.VerifyToken)
	user.GET("", API.UserController.GetUsers)
	user.GET("/:id", API.UserController.GetUserById)
	user.PUT("/:id", API.UserController.UpdateUser)
	user.DELETE("/:id", API.UserController.DeleteUser)

	auth := apiV1.Group("/auth")
	auth.POST("/login", API.AuthController.Login)
	auth.POST("/register", API.AuthController.Register)
}
