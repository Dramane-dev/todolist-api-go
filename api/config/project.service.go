package config

import "github.com/Dramane-dev/todolist-api/api/models"

type ProjectService interface {
	GetAllProjects() ([]*models.Project, error)
	GetAllProjectsByUserId(userId string) ([]*models.Project, error)
	GetProjectById(projectId string) (*models.Project, error)
	CreateProject(project *models.Project) (*models.Project, error)
	UpdateProject(projectId string) (*models.Project, error)
	DeleteProject(projectId string) (*string, error)
}
