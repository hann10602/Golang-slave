package main

import (
	"echo_postgre/db"
	"echo_postgre/initializer"
	"echo_postgre/route"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config, err := initializer.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variable")
	}

	db, err := db.NewConnection(&db.Config{
		Host:     config.DBHost,
		User:     config.DBUser,
		Password: config.DBPassword,
		Port:     config.DBPort,
		DBName:   config.DBDBname,
		SSLMode:  config.DBSSLMode,
	})

	if err != nil {
		log.Fatal(err)
	}

	userController := InitializeUserController(db)
	settingsController := InitializeSettingsController(db)
	authController := InitializeAuthController(db)

	e := echo.New()

	api := &route.API{
		Echo:               e,
		UserController:     userController,
		SettingsController: settingsController,
		AuthController:     authController,
	}

	api.SetUpRouter()

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
