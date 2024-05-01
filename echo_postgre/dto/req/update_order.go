package dto

import (
	"echo_postgre/model"
	"time"
)

type UpdateOrderDTO struct {
	Quantity  bool              `json:"quantity" gorm:"notnull"`
	Products  []*model.Products `gorm:"many2many"`
	UserId    uint              `json:"-"`
	UpdatedAt *time.Time        `json:"updatedAt,omitempty"`
}
