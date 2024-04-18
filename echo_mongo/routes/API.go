package routes

import (
	"echo_mongo/controller"
	"github.com/labstack/echo/v4"
)

type API struct {
	Echo           *echo.Echo
	UserController controller.UserController
}


func (API *API) SetUpRouter() {
	// config, err := initializer.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Could not load environment variables", err)
	// }

	apiV1 := API.Echo.Group("api/v1")
	user := apiV1.Group("/users")
	user.GET("", API.UserController.GetUsers)
	user.GET("/:id", API.UserController.GetUserById)
	user.POST("", API.UserController.CreateUser)
	user.PUT("/:id", API.UserController.GetUsers)
	user.DELETE("/:id", API.UserController.GetUsers)
}