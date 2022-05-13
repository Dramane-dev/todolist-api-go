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
	userId, ok := ctx.Params.Get("userId")
	var user entities.User

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect..."})
		return
	}

	errWhenBindingUser := ctx.BindJSON(&user)

	if errWhenBindingUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": errWhenBindingUser.Error()})
		return
	}

	userUpdated, errWhenUpdateUser := models.Update(userId, &user)

	if errWhenUpdateUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": errWhenUpdateUser.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userUpdated)
}

func DeleteUser(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect..."})
		return
	}

	userDeleted, errWhenDeleteUser := models.Delete(userId)

	if errWhenDeleteUser != nil {
		ctx.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"error": errWhenDeleteUser})
		return
	}

	ctx.JSON(http.StatusOK, userDeleted)
}
