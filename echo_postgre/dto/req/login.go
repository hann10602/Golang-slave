package dto

type LoginDto struct {
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}