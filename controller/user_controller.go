package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/models"
	"todo/service"
)

func Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data!"})
		return
	}

	err = user.SaveUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "success": false})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User saved successfully!", "success": true})
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse requested data!"})
		return
	}

	err = user.GetUserByEmail()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login user!"})
		return
	}

	accessToken, err := service.CreateAccessToken(user.Email, user.Id)
	refreshToken, err := service.CreateRefreshToken(user.Email, user.Id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user!"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "User logged in!", "success": true, "accessToken": accessToken, "refreshToken": refreshToken})
}
