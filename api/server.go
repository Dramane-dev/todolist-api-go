package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Dramane-dev/todolist-api/api/controllers"
	"github.com/Dramane-dev/todolist-api/api/mysql"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func Run() {
	databaseConnection := mysql.New(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST_DEV"), os.Getenv("DB_NAME"))
	router := gin.Default()
	router.MaxMultipartMemory = 500 << 20
	userServiceErr := controllers.NewUserDatabaseInstance(router, databaseConnection)
	projectServiceErr := controllers.NewProjectDatabaseInstance(router, databaseConnection)
	taskServiceErr := controllers.NewTaskDatabaseInstance(router, databaseConnection)
	attachmentServiceErr := controllers.NewAttachmentDatabaseInstance(router, databaseConnection)
	paymentServiceErr := controllers.NewPaymentDatabaseInstance(router, databaseConnection)

	if userServiceErr != nil {
		panic(userServiceErr)
	}

	if projectServiceErr != nil {
		panic(projectServiceErr)
	}

	if taskServiceErr != nil {
		panic(taskServiceErr)
	}

	if attachmentServiceErr != nil {
		panic(attachmentServiceErr)
	}

	if paymentServiceErr != nil {
		panic(paymentServiceErr)
	}

	router.Run()
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf(".env not found... %v", err)
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error when getting .env file, %v", err)
	} else {
		fmt.Println("[ .env file getted successfully ✅ ]")
	}
}
