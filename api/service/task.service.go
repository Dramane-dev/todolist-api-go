package service

import "github.com/Dramane-dev/todolist-api/api/models"

type TaskService interface {
	GetAllTasks() ([]*models.Task, error)
	GetAllTasksByProjectId(projectId string) ([]*models.Task, error)
	GetTaskById(taskId string) (*models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	UpdateTask(taskId string, task map[string]interface{}) (*models.Task, error)
	DeleteTask(taskId string) (*string, error)
}
