package controllers

import (
	"log"
	"net/http"

	"github.com/Dramane-dev/todolist-api/api/models"

	"github.com/gin-gonic/gin"
)

func (userService *UserController) GetAllUsers(ctx *gin.Context) {
	users, _ := userService.database.GetAllUsers()
	data := map[string]interface{}{
		"users": users,
	}

	ctx.JSON(http.StatusOK, data)
}

func (userService *UserController) GetUserById(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect ...❌"})
		return
	}

	user, errWhenGetUserById := userService.database.GetUserById(userId)

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

func (userService *UserController) Signup(ctx *gin.Context) {
	var user models.User

	err := ctx.BindJSON(&user)

	if err != nil {
		log.Println("Error when create a user ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	usr, err2 := userService.database.Signup(&user)

	if err2 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, usr)
	}
}

func (userService *UserController) Signin(ctx *gin.Context) {
	var userCredentials models.UserCredentials
	err := ctx.BindJSON(&userCredentials)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, errWhenUserSigning := userService.database.Signin(&userCredentials)

	if errWhenUserSigning != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errWhenUserSigning.Error()})
		return
	} else {
		userCredentials.Token = *token
		ctx.JSON(http.StatusOK, gin.H{"user": userCredentials})
		return
	}
}

func (userService *UserController) UpdateUser(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")
	user := make(map[string]interface{})

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect...❌"})
		return
	}

	errWhenBindingUser := ctx.BindJSON(&user)

	if errWhenBindingUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": errWhenBindingUser.Error()})
		return
	}

	userUpdated, errWhenUpdateUser := userService.database.UpdateUser(userId, user)

	if errWhenUpdateUser != nil {
		ctx.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": errWhenUpdateUser.Error()})
		return
	}

	ctx.JSON(http.StatusOK, userUpdated)
}

func (userService *UserController) DeleteUser(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect...❌"})
		return
	}

	userDeleted, errWhenDeleteUser := userService.database.DeleteUser(userId)

	if errWhenDeleteUser != nil {
		ctx.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"error": errWhenDeleteUser})
		return
	}

	ctx.JSON(http.StatusOK, userDeleted)
}
