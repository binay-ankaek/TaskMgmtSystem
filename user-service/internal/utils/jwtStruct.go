package utils

import (
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}
