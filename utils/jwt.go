package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
)

type UserClaims struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	jwt.RegisteredClaims
}

func init() {
	err := os.Setenv("TOKEN_SECRET", "TOKEN_SECRET")
	if err != nil {
		log.Fatal("fail to set token in the os env")
	}
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
}
