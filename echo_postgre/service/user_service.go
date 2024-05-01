package service

import (
	"echo_postgre/common"
	dtoRequest "echo_postgre/dto/req"
	dtoResponse "echo_postgre/dto/resp"
	"echo_postgre/model"
	"echo_postgre/repository"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	userRepository repository.IUserRepository
}
type IUserService interface {
	HandleSearchUsers(echo.Context, *common.Filter, *common.Paging) (*[]model.Users, error)
	HandleGetUserById(echo.Context, map[string]interface{}) (*dtoResponse.UserResponseDTO, error)
	HandleUpdateUsers(echo.Context, map[string]interface{}, dtoRequest.UpdateUserDTO) error
	HandleDeleteUsers(echo.Context, map[string]interface{}) error
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) HandleDeleteUsers(ctx echo.Context, cond map[string]interface{}) error {
	if err := u.userRepository.Delete(ctx.Request().Context(), cond); err != nil {
		return err
	}

	return nil
}

func (u *UserService) HandleGetUserById(ctx echo.Context, cond map[string]interface{}) (*dtoResponse.UserResponseDTO, error) {
	userData, err := u.userRepository.GetById(ctx.Request().Context(), cond)

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u *UserService) HandleSearchUsers(ctx echo.Context, filter *common.Filter, paging *common.Paging) (*[]model.Users, error) {
	data, err := u.userRepository.Search(ctx.Request().Context(), filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) HandleUpdateUsers(ctx echo.Context, cond map[string]interface{}, data dtoRequest.UpdateUserDTO) error {
	if err := u.userRepository.Update(ctx.Request().Context(), cond, data); err != nil {
		return err
	}

	return nil
}
