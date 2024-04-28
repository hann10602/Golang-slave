package model

import (
	"echo_postgre/common"
)

type (
	Users struct {
		common.SQLModel
		Username string   `json:"username" gorm:"size:255;unique;notnull"`
		Password string   `json:"password" gorm:"size:255;notnull"`
		Role     string   `json:"role" gorm:"notnull"`
		Status   string   `json:"status" gorm:"notnull"`
		Settings Settings `json:"settings" gorm:"foreignKey:UserId;references:Id"`
		Orders   []Orders `json:"orders" gorm:"foreignKey:UserId"`
	}

	Settings struct {
		common.SQLModel
		IsNotification   bool   `json:"isNotification" gorm:"notnull;default:false"`
		IsReceiveMessage bool   `json:"isReceiveMessage" gorm:"notnull;default:false"`
		Language         string `json:"language" gorm:"notnull;default:'en'"`
		UserId           uint   `json:"-"`
	}

	Orders struct {
		common.SQLModel
		Quantity bool        `json:"quantity" gorm:"notnull"`
		Products []*Products `gorm:"many2many"`
		UserId   uint        `json:"-"`
	}

	Products struct {
		common.SQLModel
		Name      uint      `json:"name" gorm:"notnull"`
		Inventory uint      `json:"inventory" gorm:"notnull"`
		Price     uint      `json:"price" gorm:"notnull"`
		Orders    []*Orders `gorm:"many2many"`
		UserId    uint      `json:"-"`
	}
)
