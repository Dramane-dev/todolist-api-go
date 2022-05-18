package mysql

import (
	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/google/uuid"
)

func (db *MySQLDatabase) GetAllProjects() ([]*models.Project, error) {
	var projects []*models.Project
	// var user *models.User

	errWhenGettingProject := db.connection.Find(&projects).Error

	// for i := range projects {
	// 	errWhenGettingUser := db.connection.Model(&user).Where("userId = ?", projects[i].UserId).Find(&user).Error

	// 	if errWhenGettingUser == nil {
	// 		projects[i].UserLastname = user.Lastname
	// 		projects[i].UserFirstname = user.FirstName
	// 		projects[i].UserEmail = user.Email
	// 	}
	// }

	if errWhenGettingProject != nil {
		return nil, errWhenGettingProject
	}

	return projects, nil
}

func (db *MySQLDatabase) GetAllProjectsByUserId(userId string) ([]*models.Project, error) {
	var projects []*models.Project

	errWhenGettingProject := db.connection.Model(&models.Project{}).Find(&projects).Error

	if errWhenGettingProject != nil {
		return nil, errWhenGettingProject
	}

	return projects, nil
}

func (db *MySQLDatabase) GetProjectById(projectId string) (*models.Project, error) {
	var projects *models.Project
	return projects, nil
}

func (db *MySQLDatabase) CreateProject(project *models.Project) (*models.Project, error) {
	project.ProjectId = uuid.NewString()
	errWhenCreateProject := db.connection.Model(&models.Project{}).Create(project).Error

	if errWhenCreateProject != nil {
		return nil, errWhenCreateProject
	}

	return project, nil
}

func (db *MySQLDatabase) UpdateProject(projectId string) (*models.Project, error) {
	var projects *models.Project
	return projects, nil
}

func (db *MySQLDatabase) DeleteProject(projectId string) (*string, error) {
	var projects string = ""
	return &projects, nil
}
