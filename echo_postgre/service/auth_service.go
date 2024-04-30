package service

import (
	dtoRequest "echo_postgre/dto/req"
	dto "echo_postgre/dto/resp"
	"echo_postgre/initializer"
	"echo_postgre/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	userRepository     repository.IUserRepository
	settingsRepository repository.ISettingsRepository
}

type IAuthService interface {
	HandleLogin(echo.Context, dtoRequest.LoginDto) (string, error)
	HandleRegister(echo.Context, dtoRequest.CreateUserDTO) error
}

func NewAuthService(userRepository repository.IUserRepository, settingsRepository repository.ISettingsRepository) IAuthService {
	return &AuthService{
		userRepository:     userRepository,
		settingsRepository: settingsRepository,
	}
}

func (a *AuthService) HandleLogin(ctx echo.Context, data dtoRequest.LoginDto) (string, error) {
	config, err := initializer.LoadConfig(".")

	if err != nil {
		return "", err
	}

	user, err := a.userRepository.GetByUsernameAndPassword(ctx.Request().Context(), map[string]interface{}{
		"username": data.Username,
		"password": data.Password,
	})

	if err != nil {
		return "", err
	}

	claims := &dto.JwtCustomClaims{
		Email: user.Username,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(config.JWTSecret))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (a *AuthService) HandleRegister(ctx echo.Context, data dtoRequest.CreateUserDTO) error {
	userId, err := a.userRepository.Create(ctx.Request().Context(), data)

	if err != nil && userId == 0 {
		return err
	}

	if err := a.settingsRepository.Create(ctx.Request().Context(), userId); err != nil {
		return err
	}

	return nil
}
