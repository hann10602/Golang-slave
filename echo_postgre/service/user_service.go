package service

import (
	"echo_postgre/common"
	dtoReq "echo_postgre/dto/req"
	dtoResp "echo_postgre/dto/resp"
	"echo_postgre/repository"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	userRepository repository.IUserRepository
}
type IUserService interface {
	HandleSearchUsers(echo.Context, *common.Filter, *common.Paging) (*[]dtoResp.UserResponseDTO, error)
	HandleGetUserById(echo.Context, uint) (*dtoResp.UserResponseDTO, error)
	HandleUpdateUser(echo.Context, map[string]interface{}, dtoReq.UpdateUserDTO) error
	HandleDeleteUser(echo.Context, map[string]interface{}) error
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) HandleDeleteUser(ctx echo.Context, cond map[string]interface{}) error {
	if err := u.userRepository.Delete(ctx.Request().Context(), cond); err != nil {
		return err
	}

	return nil
}

func (u *UserService) HandleGetUserById(ctx echo.Context, id uint) (*dtoResp.UserResponseDTO, error) {
	userData, err := u.userRepository.GetById(ctx.Request().Context(), id)

	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u *UserService) HandleSearchUsers(ctx echo.Context, filter *common.Filter, paging *common.Paging) (*[]dtoResp.UserResponseDTO, error) {
	data, err := u.userRepository.Search(ctx.Request().Context(), filter, paging)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (u *UserService) HandleUpdateUser(ctx echo.Context, cond map[string]interface{}, data dtoReq.UpdateUserDTO) error {
	if err := u.userRepository.Update(ctx.Request().Context(), cond, data); err != nil {
		return err
	}

	return nil
}
