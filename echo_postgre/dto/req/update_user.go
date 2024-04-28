package dto

import "time"

type UpdateUserDTO struct {
	Username  *string    `json:"username,omitempty" gorm:"column:username"`
	Password  *string    `json:"password,omitempty" gorm:"column:password"`
	Role      *string    `json:"role,omitempty" gorm:"column:role"`
	Status    *string    `json:"status,omitempty" gorm:"column:status"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}
