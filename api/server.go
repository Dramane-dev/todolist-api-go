package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Dramane-dev/todolist-api/api/config/mysql"
	"github.com/Dramane-dev/todolist-api/api/controllers"
	"github.com/Dramane-dev/todolist-api/api/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	databaseConnection := mysql.New(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	userService := controllers.New(databaseConnection)

	router := gin.Default()
	router.POST("/api/user/signup", userService.Signup)
	router.POST("/api/user/signin", userService.Signin)
	router.Use(middlewares.VerifyToken())
	router.GET("/api/health-check", controllers.HealthCheckController)
	router.GET("/api/users", userService.GetAllUsers)
	router.GET("/api/user/:userId", userService.GetUserById)
	router.PATCH("/api/user/:userId", userService.UpdateUser)
	router.DELETE("/api/user/:userId", userService.DeleteUser)
	router.Run()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env not found... %v", err)
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error when getting .env file, %v", err)
	} else {
		fmt.Println("[ .env file getted successfully ✅ ]")
	}
}
