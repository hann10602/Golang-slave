package dto

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Email string `json:"email" bson:"email"`
	Role  string `json:"role" bson:"role"`
	jwt.RegisteredClaims
}