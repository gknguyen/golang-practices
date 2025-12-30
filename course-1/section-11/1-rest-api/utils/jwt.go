package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "your_secret_key"

func GenerateToken(userID int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"email":  email,
		"exp":    jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
	})
	return token.SignedString([]byte(secretKey))
}
