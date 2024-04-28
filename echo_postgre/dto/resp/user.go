package dto

type UserResponseDTO struct {
	Id       uint                `json:"id" gorm:"column:id"`
	Username string              `json:"username" gorm:"column:username"`
	Password string              `json:"password" gorm:"column:password"`
	Role     string              `json:"role" gorm:"column:role"`
	Status   string              `json:"status" gorm:"column:status"`
	Settings SettingsResponseDTO `json:"settings"`
}
