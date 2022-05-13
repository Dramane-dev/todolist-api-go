package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Dramane-dev/todolist-api/api/config"
	"github.com/Dramane-dev/todolist-api/api/entities"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type UserModel struct {
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

	config.GetDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
}

func (*UserModel) GetAll() ([]entities.User, error) {
	db, err := config.GetDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("SELECT * FROM users")

		if err2 != nil {
			return nil, err2
		} else {
			var users []entities.User

			for rows.Next() {
				var user entities.User
				rows.Scan(&user.UserId, &user.Lastname, &user.FirstName, &user.Email, &user.Password, &user.MailConfirmed)
				users = append(users, user)
			}
			return users, nil
		}
	}
}

func (*UserModel) GetById(userId string) (*entities.User, error) {
	db, err := config.GetDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	if err != nil {
		return nil, err
	} else {
		result, err2 := db.Query("SELECT * FROM users WHERE userId = ?", userId)

		if err2 != nil {
			return nil, err2
		} else {
			var user entities.User
			for result.Next() {
				result.Scan(&user.UserId, &user.Lastname, &user.FirstName, &user.Email, &user.Password, &user.MailConfirmed)
			}
			return &user, nil
		}
	}
}

func Create(user *entities.User) (*entities.User, error) {
	db, err := config.GetDB(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	if err != nil {
		return nil, err
	} else {
		user.UserId = uuid.NewString()
		result, err2 := db.Exec("INSERT INTO users(userId, lastname, firstname, email, password) VALUES(?, ?, ?, ?, ?)", user.UserId, user.Lastname, user.FirstName, user.Email, user.Password)

		if err2 != nil {
			if strings.Contains(err2.Error(), "Duplicate entry") {
				return nil, errors.New("user already exist in database ❌")
			}
			return nil, err2
		} else {
			rowsAffected, errorWhenCreateUser := result.RowsAffected()

			if rowsAffected > 0 {
				return user, err2
			} else {
				return nil, errorWhenCreateUser
			}
		}
	}
}
