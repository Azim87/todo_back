package controller

import (
	"net/http"
	"todo/models"
	"todo/service"

	"github.com/gin-gonic/gin"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse requested data!"})
		return
	}

	err = user.SaveUser()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": err.Error(), "success": false})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{"message": "User saved successfully!", "success": true})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Could not parse requested data!"})
		return
	}

	err = user.GetUserByEmail()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not login user!"})
		return
	}

	accessToken, accessExpiry, err := service.CreateAccessToken(user.Email, user.Id)
	refreshToken, refreshExpiry, err := service.CreateRefreshToken(user.Email, user.Id)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not authenticate user!"})
	}

	if err := models.SaveTokens(int(user.Id), accessToken, refreshToken, accessExpiry, refreshExpiry); err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "Could not save tokens!"})
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message":      "User logged in!",
			"success":      true,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
}
