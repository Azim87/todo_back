package service

import "todo/utils"

func CreateAccessToken(email string, userId int64) (string, error) {
	token, err := utils.GenerateAccessToken(email, userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateRefreshToken(email string, userId int64) (string, error) {
	token, err := utils.GenerateRefreshToken(email, userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ValidateJwtToken(accessToken string) (int64, error) {
	token, err := utils.VerifyToken(accessToken, []byte("your-access-secret-key"))
	if err != nil {
		return 0, err
	}
	return token, nil
}
