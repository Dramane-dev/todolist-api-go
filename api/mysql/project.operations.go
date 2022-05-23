package mysql

import (
	"errors"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/google/uuid"
)

func (db *MySQLDatabase) GetAllProjects() ([]*models.Project, error) {
	var projects []*models.Project

	errWhenGettingProject := db.connection.Preload("Tasks").Model(&models.Project{}).Find(&projects).Error

	if errWhenGettingProject != nil {
		return nil, errWhenGettingProject
	}

	return projects, nil
}

func (db *MySQLDatabase) GetAllProjectsByUserId(userId string) ([]*models.Project, error) {
	var projects []*models.Project

	errWhenGettingProject := db.connection.Preload("Tasks").Model(&models.Project{}).Where("userId = ?", userId).Find(&projects).Error

	if errWhenGettingProject != nil {
		return nil, errWhenGettingProject
	}

	return projects, nil
}

func (db *MySQLDatabase) GetProjectById(projectId string) (*models.Project, error) {
	var project *models.Project

	errWhenGetProjectById := db.connection.Preload("Tasks").Model(&models.Project{}).Where("projectId = ?", projectId).Find(&project).Error

	if errWhenGetProjectById != nil {
		return nil, errWhenGetProjectById
	}

	return project, nil
}

func (db *MySQLDatabase) CreateProject(project *models.Project) (*models.Project, error) {
	project.ProjectId = uuid.NewString()
	errWhenCreateProject := db.connection.Model(&models.Project{}).Create(project).Error

	if errWhenCreateProject != nil {
		return nil, errWhenCreateProject
	}

	return project, nil
}

func (db *MySQLDatabase) UpdateProject(projectId string, project map[string]interface{}) (*models.Project, error) {
	errWhenUpdateProject := db.connection.Model(&models.Project{}).Where("projectId = ?", projectId).Updates(project).Error

	if errWhenUpdateProject != nil {
		return nil, errWhenUpdateProject
	}
	return db.GetProjectById(projectId)
}

func (db *MySQLDatabase) DeleteProject(projectId string) (*string, error) {
	project, _ := db.GetProjectById(projectId)

	if !(len(project.ProjectId) > 0) {
		return nil, errors.New("project not found ...❌")
	}

	errWhenDeleteProject := db.connection.Model(&models.Project{}).Where("projectId = ?", projectId).Delete(projectId).Error

	if errWhenDeleteProject != nil {
		return nil, errWhenDeleteProject
	}

	var strResponse string = "Project deleted successfully ✅"

	return &strResponse, errWhenDeleteProject
}
