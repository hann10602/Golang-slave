package db

import (
	"echo_postgre/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	if err := db.SetupJoinTable(&model.Products{}, "Orders", &model.OrderProducts{}); err != nil {
		log.Fatal(err)

		return nil, err
	}

	if err = db.AutoMigrate(&model.Users{}, &model.Settings{}, &model.Orders{}, &model.Products{}, &model.OrderProducts{}); err != nil {
		log.Fatal("Migrate entity failed...")

		return nil, err
	}

	fmt.Println("Connect database successfully...")
	return db, nil
}
