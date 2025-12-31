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

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	isValidToken := parsedToken.Valid
	if !isValidToken {
		return 0, jwt.ErrTokenNotValidYet
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
