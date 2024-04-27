package dto

type UpdateSettingsDTO struct {
	IsNotification   *bool `json:"isNotification" gorm:"notnull;default:false;column:isNotification"`
	IsReceiveMessage *bool `json:"isReceiveMessage" gorm:"notnull;default:false;column:isReceiveMessage"`
	Language         *bool `json:"language" gorm:"notnull;default:en;column:language"`
}
