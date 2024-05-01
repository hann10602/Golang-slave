package dto

import "time"

type UserResponseDTO struct {
	Id               uint       `json:"id"`
	Username         string     `json:"username"`
	Password         string     `json:"password"`
	Role             string     `json:"role"`
	Status           string     `json:"status"`
	CreatedAt        *time.Time `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
	IsNotification   bool       `json:"isNotification"`
	IsReceiveMessage bool       `json:"isReceiveMessage"`
	Language         string     `json:"language"`
}
