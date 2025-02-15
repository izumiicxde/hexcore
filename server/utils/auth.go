package utils

import (
	"fmt"
	"hexcore/config"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Passwords
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func GenerateVerificationCode() string {
	code := rand.Intn(900000) + 100000
	return strconv.Itoa(code)
}

// JWT
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

func ParseJWT(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token is signed with HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Envs.JWT_SECRET), nil
	})

	if err != nil {
		return nil, nil, err
	}

	// Extract claims if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, claims, nil
	}
	return nil, nil, fmt.Errorf("invalid token")
}
