package mysql

import (
	"errors"
	"strings"

	"github.com/Dramane-dev/todolist-api/api/functions"
	"github.com/Dramane-dev/todolist-api/api/models"

	"github.com/google/uuid"
)

type User struct {
	UserId        string `gorm:"column:userId;primaryKey" json:"userId"`
	Lastname      string `gorm:"column:lastname" json:"lastname"`
	FirstName     string `gorm:"column:firstname" json:"firstname"`
	Email         string `gorm:"column:email" json:"email"`
	MailConfirmed bool   `gorm:"column:mailConfirmed" json:"mailConfirmed"`
}

func (db *MySQLDatabase) GetAllUsers() ([]*models.UserInformations, error) {
	var users []*models.UserInformations

	errWhenGetAllUsers := db.connection.Preload("Projects.Tasks").Preload("Projects.Attachments").Model(&User{}).Find(&users).Error

	if errWhenGetAllUsers != nil {
		return nil, errWhenGetAllUsers
	}

	return users, nil
}

func (db *MySQLDatabase) GetUserById(userId string) (*models.UserInformations, error) {
	var user models.UserInformations
	errWhenGetUserById := db.connection.Preload("Projects.Tasks").Preload("Projects.Attachments").Model(&User{}).First(&user, "userId = ?", userId).Error

	if errWhenGetUserById != nil {
		return nil, errWhenGetUserById
	}
	return &user, nil
}

func (db *MySQLDatabase) GetUserByEmail(userEmail string) (*models.User, error) {
	var user models.User
	userNotFound := db.connection.Model(&User{}).Find(&user, "email = ?", userEmail).Error

	if userNotFound != nil {
		return nil, userNotFound
	}
	return &user, nil
}

func (db *MySQLDatabase) Signup(user *models.User) (*models.UserInformations, error) {
	user.UserId = "USR" + uuid.NewString()
	passwordHashed, errWhenHashingPassword := functions.HashPassword(user.Password)

	if errWhenHashingPassword != nil {
		return nil, errors.New(errWhenHashingPassword.Error())
	}

	user.Password = passwordHashed
	errWhenCreateUser := db.connection.Create(user).Error

	if errWhenCreateUser != nil {
		if strings.Contains(errWhenCreateUser.Error(), "Duplicate entry") {
			return nil, errors.New("user already exist in database ❌")
		}

		return nil, errWhenCreateUser
	}
	return db.GetUserById(user.UserId)
}

func (db *MySQLDatabase) Signin(user *models.UserCredentials) (*string, error) {
	usr, userNotFound := db.GetUserByEmail(user.Email)

	if userNotFound != nil {
		return nil, errors.New("User not found...❌")
	}

	samePassword := functions.CheckPasswordHash(user.Password, usr.Password)

	if !samePassword {
		return nil, errors.New("User credentials incorrect ❌")
	}

	token, errWhenGenerateToken := functions.GenerateNewToken(usr.UserId)

	if errWhenGenerateToken != nil {
		return nil, errWhenGenerateToken
	}

	return &token, nil
}

func (db *MySQLDatabase) UpdateUser(userId string, user map[string]interface{}) (*models.UserInformations, error) {
	errWhenUpdateUser := db.connection.Model(&models.User{}).Where("userId = ?", userId).Updates(user).Error

	if errWhenUpdateUser != nil {
		return nil, errWhenUpdateUser
	} else {
		return db.GetUserById(userId)
	}
}

func (db *MySQLDatabase) DeleteUser(userId string) (string, error) {
	var user models.User
	errWhenDeleteUser := db.connection.Delete(&user, "userId = ?", userId).Error

	if errWhenDeleteUser != nil {
		return "", errWhenDeleteUser
	}
	return "User deleted successfully ✅", errWhenDeleteUser
}
