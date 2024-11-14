package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const accessSecret = "your-access-secret-key"
const refreshSecret = "your-refresh-secret-key"

func GenerateAccessToken(email string, userId int64) (string, error) {
	// Set token expiration time (e.g., 15 minutes)
	expirationTime := time.Now().Add(15 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"ext":    expirationTime,
	})

	return token.SignedString([]byte(accessSecret))
}

func GenerateRefreshToken(email string, userId int64) (string, error) {
	// Set refresh token expiration time (e.g., 7 days)
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"ext":    expirationTime,
	})

	return token.SignedString([]byte(refreshSecret))
}

func VerifyToken(token string, secretKey []byte) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New(`token is invalid`)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")

	}
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
