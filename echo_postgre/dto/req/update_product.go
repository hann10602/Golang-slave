package dto

import "time"

type UpdateProductDTO struct {
	Name        uint       `json:"name" gorm:"column:name"`
	Inventory   uint       `json:"inventory" gorm:"column:inventory"`
	Price       uint       `json:"price" gorm:"column:price"`
	Description string     `json:"description" gorm:"column:description"`
	Thumbnail   string     `json:"thumbnail" gorm:"column:thumbnail"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}
