package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ParseToken(token string) (string, int64, error) {
	parsedToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("your-refresh-secret-key"), nil
	})

	if err != nil {
		return "", 0, fmt.Errorf("failed to parse token: %w", err)
	}

	// Проверяем валидность токена
	if !parsedToken.Valid {
		return "", 0, fmt.Errorf("token is invalid")
	}

	// Извлекаем claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, fmt.Errorf("invalid token claims")
	}

	// Проверяем и извлекаем userId
	userIdClaim, ok := claims["userId"]
	if !ok {
		return "", 0, fmt.Errorf("userId not found in token claims")
	}
	userIdFloat, ok := userIdClaim.(float64)
	if !ok {
		return "", 0, fmt.Errorf("userId is not a valid number")
	}
	userId := int64(userIdFloat)

	// Проверяем и извлекаем email
	emailClaim, ok := claims["email"]
	if !ok {
		return "", 0, fmt.Errorf("email not found in token claims")
	}
	email, ok := emailClaim.(string)
	if !ok {
		return "", 0, fmt.Errorf("email is not a valid string")
	}

	return email, userId, nil
}
