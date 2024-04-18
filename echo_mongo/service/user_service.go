package service

import (
	dto "echo_mongo/dto/user"
	model "echo_mongo/model/user"
	"echo_mongo/repository"

	"github.com/labstack/echo/v4"
)

type UserService struct {
	UserRepository repository.IUserRepository
}

type IUserService interface {
	HandleGetUsers(echo.Context) ([]*model.User, error)
	HandleGetUserById(echo.Context, string) (*model.User, error)
	HandleCreateUser(echo.Context, model.User) (interface{}, error)
	HandleUpdateUser(echo.Context, dto.UpdateUserDto, string) (*model.User, error)
	HandleDeleteUser(echo.Context, string) error
}

func NewUserService(userRepository repository.IUserRepository) IUserService {
	return UserService{
		UserRepository: userRepository,
	}
}

func (u UserService) HandleGetUsers(ctx echo.Context) ([]*model.User, error) {
	users, err := u.UserRepository.GetUsers(ctx.Request().Context())

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserService) HandleGetUserById(ctx echo.Context, id string) (*model.User, error) {
	user, err := u.UserRepository.GetUserById(ctx.Request().Context(), id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) HandleCreateUser(ctx echo.Context, dto model.User) (interface{}, error) {
	objectId, err := u.UserRepository.CreateUser(ctx.Request().Context(), dto)

	if err != nil {
		return nil, err
	}

	return objectId, nil
}

func (u UserService) HandleUpdateUser(ctx echo.Context, dto dto.UpdateUserDto, id string) (*model.User, error) {
	user, err := u.UserRepository.UpdateUser(ctx.Request().Context(), dto, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) HandleDeleteUser(ctx echo.Context, id string) error {
	err := u.UserRepository.DeleteUser(ctx.Request().Context(), id)

	if err != nil {
		return err
	}

	return nil
}
