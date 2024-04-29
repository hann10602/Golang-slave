package implement

import (
	"context"
	"echo_postgre/model"
	"echo_postgre/repository"

	"gorm.io/gorm"
)

type AuthImplement struct {
	db *gorm.DB
}

// Login implements repository.IAuthRepository.
func (a *AuthImplement) Login(context.Context) (*model.Users, error) {
	panic("unimplemented")
}

func NewAuthImplement(db *gorm.DB) repository.IAuthRepository {
	return &AuthImplement{
		db: db,
	}
}
