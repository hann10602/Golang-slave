package db

import (
	"gorm_demo/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DbConnection() {
	dsn := "root:mny10602@tcp(localhost:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{}, &models.Note{}, &models.CreditCard{})
}

func DbManager() *gorm.DB {
	return db
}