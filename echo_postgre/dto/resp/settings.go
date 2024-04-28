package dto

type SettingsResponseDTO struct {
	IsNotification   bool   `json:"isNotification" gorm:"column:is_notification"`
	IsReceiveMessage bool   `json:"isReceiveMessage" gorm:"column:is_receive_message"`
	Language         string `json:"language" gorm:"column:language"`
}
