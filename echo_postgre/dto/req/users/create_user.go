package dto

import "echo_postgre/model"

type CreateUserDTO struct {
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
	Role     string `json:"role,omitempty" gorm:"column:role;defaul:USER"`
	Status   string `json:"status" gorm:"column:status;defaul:ACTIVE"`
	Settings *model.Settings
}
