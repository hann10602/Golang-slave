package dto

import "time"

type CreateOrderDTO struct {
	Name      uint       `json:"name" gorm:"column:name"`
	Inventory uint       `json:"inventory" gorm:"column:inventory"`
	Price     uint       `json:"price" gorm:"column:price"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
}
