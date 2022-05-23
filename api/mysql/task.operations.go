package mysql

import (
	"errors"
	"strings"
	"time"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/google/uuid"
)

func (db *MySQLDatabase) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	errWhenGettingTask := db.connection.Find(&tasks).Error

	if errWhenGettingTask != nil {
		return nil, errWhenGettingTask
	}

	return tasks, nil
}

func (db *MySQLDatabase) GetAllTasksByProjectId(projectId string) ([]*models.Task, error) {
	var tasks []*models.Task

	errWhenGettingTask := db.connection.Model(&models.Task{}).Where("projectId = ?", projectId).Find(&tasks).Error

	if errWhenGettingTask != nil {
		return nil, errWhenGettingTask
	}

	return tasks, nil
}

func (db *MySQLDatabase) GetTaskById(taskId string) (*models.Task, error) {
	var task *models.Task

	errWhenGetTaskById := db.connection.Model(&models.Task{}).Where("taskId = ?", taskId).Find(&task).Error

	if errWhenGetTaskById != nil {
		return nil, errWhenGetTaskById
	}

	return task, nil
}

func (db *MySQLDatabase) CreateTask(task *models.Task) (*models.Task, error) {
	task.TaskId = uuid.NewString()
	currentDate := time.Now()
	task.CreatedAt = strings.Split(currentDate.String(), ".")[0]
	errWhenCreateTask := db.connection.Model(&models.Task{}).Create(task).Error

	if errWhenCreateTask != nil {
		return nil, errWhenCreateTask
	}

	return db.GetTaskById(task.TaskId)
}

func (db *MySQLDatabase) UpdateTask(taskId string, task map[string]interface{}) (*models.Task, error) {
	errWhenUpdateTask := db.connection.Model(&models.Task{}).Where("taskId = ?", taskId).Updates(task).Error

	if errWhenUpdateTask != nil {
		return nil, errWhenUpdateTask
	}
	return db.GetTaskById(taskId)
}

func (db *MySQLDatabase) DeleteTask(taskId string) (*string, error) {
	task, _ := db.GetTaskById(taskId)

	if !(len(task.TaskId) > 0) {
		return nil, errors.New("project not found ...❌")
	}

	errWhenDeleteTask := db.connection.Model(&models.Task{}).Where("taskId = ?", taskId).Delete(taskId).Error

	if errWhenDeleteTask != nil {
		return nil, errWhenDeleteTask
	}

	var strResponse string = "Project deleted successfully ✅"

	return &strResponse, errWhenDeleteTask
}
