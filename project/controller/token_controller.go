package controller

import (
	"net/http"
	"todo/models"
	"todo/service"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

func RefreshTokens(context *gin.Context) {
	var requestData struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Invalid request body"})
		return
	}

	token := requestData.RefreshToken

	if token == "" {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Refresh token is required"})
		return
	}

	email, userId, _ := utils.ParseToken(token)

	if email != "" {
		newAccessToken, accessExpiry, err := service.CreateAccessToken(email, int64(userId))
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "failed to generate access token"})
			return
		}

		newRefreshToken, refreshExpiry, err := service.CreateRefreshToken(email, int64(userId))
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "failed to generate refresh token"})
			return
		}

		if err := models.UpdateToken(userId, newAccessToken, newRefreshToken, accessExpiry, refreshExpiry); err != nil {
			context.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Could not save tokens!"})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"accessToken":  newAccessToken,
			"refreshToken": newRefreshToken,
			"success":      true,
		})
	} else {
		context.JSON(
			666,
			gin.H{"error": "Refresh token has expired!"})
		return
	}

}
