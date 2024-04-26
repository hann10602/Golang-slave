package main

import (
	"echo_postgre/db"
	"echo_postgre/initializer"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		SSLMode:  config.DBSslmode,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{config.ClientOrigin},
		AllowCredentials: true,
	}))

	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
