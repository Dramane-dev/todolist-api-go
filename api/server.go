package api

import (
	"github.com/Dramane-dev/todolist-api/api/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/api/health-check", controllers.HealthCheckController)
	router.GET("/api/users", controllers.GetAllUsers)
	router.GET("/api/user/:userId", controllers.GetUserById)
	router.POST("/api/user", controllers.CreateUser)
	router.PATCH("/api/user/:userId", controllers.UpdateUser)
	router.DELETE("/api/user/:userId", controllers.DeleteUser)
	router.Run()
}
