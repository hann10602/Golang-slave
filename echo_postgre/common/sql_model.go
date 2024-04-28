package common

import (
	"time"

	"gorm.io/gorm"
)

type SQLModel struct {
	gorm.Model
	Id        uint       `json:"id"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}
