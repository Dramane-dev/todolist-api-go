package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Dramane-dev/todolist-api/api/config/mysql"
	"github.com/Dramane-dev/todolist-api/api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Run() {
	databaseConnection := mysql.New(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	router := gin.Default()
	userServiceErr := controllers.NewUserDatabaseInstance(router, databaseConnection)
	projectServiceErr := controllers.NewProjectDatabaseInstance(router, databaseConnection)

	if userServiceErr != nil {
		panic(userServiceErr)
	}

	if projectServiceErr != nil {
		panic(projectServiceErr)
	}

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
