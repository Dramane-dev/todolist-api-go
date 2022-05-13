package controllers

import (
	"log"
	"net/http"

	"github.com/Dramane-dev/todolist-api/api/entities"
	"github.com/Dramane-dev/todolist-api/api/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	var userModel models.UserModel
	users, _ := userModel.GetAll()
	data := map[string]interface{}{
		"users": users,
	}

	ctx.JSON(http.StatusOK, data)
}

func GetUserById(ctx *gin.Context) {
	var userModel models.UserModel
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect"})
		return
	}

	user, errWhenGetUserById := userModel.GetById(userId)

	if errWhenGetUserById != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": errWhenGetUserById.Error(),
		})
		return
	}

	if user.UserId == "" && user.Email == "" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{})
		return
	}

	data := map[string]interface{}{
		"user": user,
	}
	ctx.JSON(http.StatusOK, data)
}

func CreateUser(ctx *gin.Context) {
	var user entities.User

	err := ctx.BindJSON(&user)

	if err != nil {
		log.Println("Error when create a user ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	usr, err2 := models.Create(&user)

	if err2 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, usr)
	}
}

func UpdateUser(ctx *gin.Context) {

}
