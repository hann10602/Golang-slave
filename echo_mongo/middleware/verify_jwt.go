package middleware

import (
	"echo_mongo/common"
	"echo_mongo/constant"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := strings.Split(c.Request().Header.Get("Authorization"), " ")

		verifySecretKey := func(token *jwt.Token) (interface{}, error) {
			return []byte(constant.JWT_SECRET), nil
		}

		_, err := jwt.Parse(tokenString[1], verifySecretKey)

		if tokenString[0] != "Bearer" || err != nil {
			return c.JSON(http.StatusUnauthorized, common.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token is invalid",
				Data:       nil,
			})
		}

		return next(c)
	}
}
