package utils

import (
	"hexcore/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userId uint, role string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":  userId,
		"role":    role,
		"expires": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	})

	tokenstr, err := token.SignedString([]byte(config.Envs.JWT_SECRET))
	if err != nil {
		panic(err)
	}

	return tokenstr
}
