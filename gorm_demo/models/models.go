package models

import "gorm.io/gorm"

type (
	User struct {
		gorm.Model
		Username    string `json:"username", gorm:"size:64"`
		Password    string `json:"password", gorm:"size:255"`
		Notes []Note

	}

	Note struct {
		gorm.Model
		Name string `gorm:"size:255"`
		Content string `gorm:"type:text"`
		UserId uint64 `gorm:"index"`
		User User
	}

	CreditCard struct {
		gorm.Model
		Number string
		UserId uint64
	}
)