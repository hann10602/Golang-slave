package dto

import "time"

type CreateSettingsDTO struct {
	Id        uint       `json:"id" gorm:"column:id"`
	Username  string     `json:"username" gorm:"column:username"`
	Password  string     `json:"password" gorm:"column:password"`
	Role      string     `json:"role,omitempty" gorm:"default:'USER';column:role"`
	Status    string     `json:"status,omitempty" gorm:"default:'ACTIVE';column:status"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`
}
