package main

import (
	"gorm_demo/db"
	"gorm_demo/route"
)

func main() {
	db.DbConnection()

	server := route.Init()

	server.Logger.Fatal(server.Start(":8081"))
}