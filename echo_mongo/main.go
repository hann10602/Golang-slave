package main

import (
	"echo_mongo/controller"
	"echo_mongo/db"
	"echo_mongo/initializer"
	"echo_mongo/repository/implement"
	"echo_mongo/routes"
	"echo_mongo/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config, err := initializer.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	mongoDB := &db.MongoDB{
		DBName: config.DBName,
	}

	mongoDB.Connect()
	defer mongoDB.Close()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.ClientOrigin},
		AllowCredentials: true,
	}))

	userRepo := implement.NewUserImplement(mongoDB.Client.Database(config.DBName))
	userService := service.NewUserService(userRepo)
	userController := controller.UserController{
		UserService: userService,
	}

	authRepo := implement.NewAuthImplement(mongoDB.Client.Database(config.DBName))
	authService := service.NewAuthService(authRepo)
	authController := controller.AuthController{
		AuthService: authService,
	}

	api := routes.API{
		Echo:           e,
		UserController: userController,
		AuthController: authController,
	}

	api.SetUpRouter()

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
