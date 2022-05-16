package controllers

import (
	"github.com/Dramane-dev/todolist-api/api/config"
)

type Controller struct {
	database config.UserService
}

func New(database config.UserService) *Controller {
	return &Controller{
		database: database,
	}
}
