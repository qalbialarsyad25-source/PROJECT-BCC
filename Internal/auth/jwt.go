package controller

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, email string) (string, error) {
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := tokenJWT.SignedString([]byte("secret-key"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
