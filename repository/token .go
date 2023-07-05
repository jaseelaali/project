package repository

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func Token(Email string) (*jwt.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	return token, nil

}
