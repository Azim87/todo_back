package service

import (
	"time"
	"todo/utils"
)

func CreateAccessToken(email string, userId int64) (string, time.Time, error) {
	accessToken, accessExp, err := utils.GenerateAccessToken(email, userId)
	if err != nil {
		return "", time.Time{}, err
	}
	return accessToken, accessExp, nil
}

func CreateRefreshToken(email string, userId int64) (string, time.Time, error) {
	refreshToken, refreshExp, err := utils.GenerateRefreshToken(email, userId)
	if err != nil {
		return "", time.Time{}, err
	}
	return refreshToken, refreshExp, nil
}

func VerifyToken(token string) (int64, error) {
	userId, err := utils.VerifyToken(token)

	if err != nil {
		return 0, err
	}

	return userId, nil
}
