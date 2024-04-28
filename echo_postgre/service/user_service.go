package service

import (
	"echo_postgre/common"
	dto "echo_postgre/dto/req/users"
	"echo_postgre/model"
	"echo_postgre/repository"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	userRepository     repository.IUserRepository
	settingsRepository repository.ISettingsRepository
}
type IUserService interface {
	HandleSearchUsers(echo.Context, common.Filter, common.Paging) (*[]model.Users, error)
	HandleGetUserById(echo.Context, map[string]interface{}) (*model.Users, error)
	HandleCreateUsers(echo.Context, dto.CreateUserDTO) error
	HandleUpdateUsers(echo.Context, map[string]interface{}, dto.UpdateUserDTO) error
	HandleDeleteUsers(echo.Context, map[string]interface{}) error
}

func NewUserService(userRepository repository.IUserRepository, settingsRepository repository.ISettingsRepository) IUserService {
	return &UserService{
		userRepository:     userRepository,
		settingsRepository: settingsRepository,
	}
}

func (u *UserService) HandleCreateUsers(ctx echo.Context, data dto.CreateUserDTO) error {
	userId, err := u.userRepository.Create(ctx.Request().Context(), data)

	if err != nil && userId == 0 {
		return err
	}

	if err := u.settingsRepository.Create(ctx.Request().Context(), userId); err != nil {
		return err
	}

	return nil
}

func (u *UserService) HandleDeleteUsers(ctx echo.Context, cond map[string]interface{}) error {
	if err := u.userRepository.Delete(ctx.Request().Context(), cond); err != nil {
		return err
	}

	return nil
}

func (u *UserService) HandleGetUserById(ctx echo.Context, cond map[string]interface{}) (*model.Users, error) {
	data, err := u.userRepository.GetById(ctx.Request().Context(), cond)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) HandleSearchUsers(ctx echo.Context, filter common.Filter, paging common.Paging) (*[]model.Users, error) {
	data, err := u.userRepository.Search(ctx.Request().Context(), filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) HandleUpdateUsers(ctx echo.Context, cond map[string]interface{}, data dto.UpdateUserDTO) error {
	if err := u.userRepository.Update(ctx.Request().Context(), cond, data); err != nil {
		return err
	}

	return nil
}
