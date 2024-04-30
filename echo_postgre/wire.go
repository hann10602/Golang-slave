//go:build wireinject
// +build wireinject

package main

import (
	"echo_postgre/controller"
	"echo_postgre/repository/implement"
	"echo_postgre/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeUserController(db *gorm.DB) controller.UserController {
	wire.Build(implement.NewUserImplement, service.NewUserService, controller.NewUserController)
	return controller.UserController{}
}

func InitializeSettingsController(db *gorm.DB) controller.SettingsController {
	wire.Build(implement.NewSettingsImplement, service.NewSettingsService, controller.NewSettingsController)
	return controller.SettingsController{}
}

func InitializeAuthController(db *gorm.DB) controller.AuthController {
	wire.Build(implement.NewUserImplement, implement.NewSettingsImplement, service.NewAuthService, controller.NewAuthController)
	return controller.AuthController{}
}
