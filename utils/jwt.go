package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const accessSecret = "your-access-secret-key"
const refreshSecret = "your-refresh-secret-key"

func GenerateAccessToken(email string, userId int64) (string, time.Time, error) {
	expirationAccessTime := time.Now().Add(time.Minute * 360)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    expirationAccessTime.Unix(),
		})
	signedAccessToken, err := token.SignedString([]byte(accessSecret))

	if err != nil {
		return "", time.Time{}, err
	}

	return signedAccessToken, expirationAccessTime, nil
}

func GenerateRefreshToken(email string, userId int64) (string, time.Time, error) {
	expirationRefreshTime := time.Now().Add(time.Hour * 24 * 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    expirationRefreshTime.Unix(),
		})

	signedRefreshToken, err := token.SignedString([]byte(refreshSecret))

	if err != nil {
		return "", time.Time{}, err
	}

	return signedRefreshToken, expirationRefreshTime, nil
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})

	if err != nil {
		return 0, err
	}

	validToken := !parsedToken.Valid

	if !validToken {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, err
	}

	userId := int64(claims["userId"].(float64))

	return userId, nil
}
