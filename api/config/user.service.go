package config

import (
	"github.com/Dramane-dev/todolist-api/api/models"
)

type UserService interface {
	GetAllUsers() ([]*models.UserInformations, error)

	GetUserById(userId string) (*models.UserInformations, error)

	Signup(user *models.User) (*models.UserInformations, error)

	UpdateUser(userId string, user map[string]interface{}) (*models.UserInformations, error)

	DeleteUser(userId string) (string, error)

	GetUserByEmail(userEmail string) (*models.User, error)

	Signin(user *models.UserCredentials) (*string, error)
}
