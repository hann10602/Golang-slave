package service

import (
	dto "echo_mongo/dto/auth"
	"echo_mongo/initializer"
	model "echo_mongo/model/user"
	"echo_mongo/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	AuthRepository repository.IAuthRepository
}

type IAuthService interface {
	HandleLogin(echo.Context, dto.LoginDto) (string, error)
	HandleRegister(echo.Context, model.User) (*model.User, error)
}

func NewAuthService(authRepository repository.IAuthRepository) IAuthService {
	return AuthService{
		AuthRepository: authRepository,
	}
}

func (s AuthService) HandleLogin(ctx echo.Context, data dto.LoginDto) (string, error) {
	config, err := initializer.LoadConfig(".")

	if err != nil {
		return "", err
	}

	user, err := s.AuthRepository.Login(ctx.Request().Context(), data)

	if err != nil {
		return "", err
	}

	claims := &dto.JwtCustomClaims{
		Email: user.Email,
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

func (s AuthService) HandleRegister(ctx echo.Context, dto model.User) (*model.User, error) {
	user, err := s.AuthRepository.Register(ctx.Request().Context(), dto)

	if err != nil {
		return nil, err
	}

	return user, nil
}
