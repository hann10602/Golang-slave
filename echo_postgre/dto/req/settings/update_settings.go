package dto

import "time"

type UpdateSettingsDTO struct {
	IsNotification   *bool      `json:"isNotification" gorm:"notnull;column:is_notification"`
	IsReceiveMessage *bool      `json:"isReceiveMessage" gorm:"notnull;column:is_receive_message"`
	Language         *string    `json:"language" gorm:"notnull;default:en;column:language"`
	UpdatedAt        *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}
