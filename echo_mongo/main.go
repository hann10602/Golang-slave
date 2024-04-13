package main

import (
	"echo_mongo/db"
	"echo_mongo/initializer"
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
	e.Use(middleware.CORSWithConfig{
		AllowOrigins:	  []string{config.ClientOrigin},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept}
	})

	userRepo
}