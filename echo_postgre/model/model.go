package model

import (
	"echo_postgre/common"
)

type (
	Users struct {
		common.SQLModel
		Username string      `json:"username" gorm:"size:255;unique;notnull"`
		Password string      `json:"password" gorm:"size:255;notnull"`
		Role     string      `json:"role" gorm:"notnull"`
		Status   string      `json:"status" gorm:"notnull"`
		Settings Settings    `json:"settings" gorm:"foreignKey:UserId;references:Id"`
		Products []*Products `gorm:"many2many:orders"`
	}

	Settings struct {
		common.SQLModel
		IsNotification   bool   `json:"isNotification" gorm:"notnull;default:false"`
		IsReceiveMessage bool   `json:"isReceiveMessage" gorm:"notnull;default:false"`
		Language         string `json:"language" gorm:"notnull;default:'en'"`
		UserId           uint   `json:"-"`
	}

	Products struct {
		common.SQLModel
		Name        string   `json:"name" gorm:"notnull"`
		Description string   `json:"description" gorm:"notnull"`
		Inventory   uint     `json:"inventory" gorm:"notnull"`
		Price       uint     `json:"price" gorm:"notnull"`
		Thumbnail   string   `json:"thumbnail"`
		Users       []*Users `gorm:"many2many:orders"`
	}

	Orders struct {
		common.SQLModel
		Quantity   uint `json:"quantity" gorm:"notnull"`
		UsersId    uint `json:"userId" gorm:"notnull"`
		ProductsId uint `json:"productId" gorm:"notnull"`
	}
)
