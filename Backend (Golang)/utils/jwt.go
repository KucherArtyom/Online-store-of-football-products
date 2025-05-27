package utils

import (
	"Go-Kurs/models"
	"errors"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

func ParseToken(tokenString, secret string) (*models.Claims, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	claims := &models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
