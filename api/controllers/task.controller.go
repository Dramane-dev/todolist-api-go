package controllers

import (
	"log"
	"net/http"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/gin-gonic/gin"
)

func (taskService *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, errWhenGetAllTaskByUserId := taskService.database.GetAllTasks()

	if errWhenGetAllTaskByUserId != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenGetAllTaskByUserId})
		return
	}

	data := map[string]interface{}{
		"tasks": tasks,
	}

	ctx.JSON(http.StatusOK, data)
}

func (taskService *TaskController) GetAllTasksByProjectId(ctx *gin.Context) {
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect...❌"})
		return
	}

	tasks, errWhenGetAllProjects := taskService.database.GetAllTasksByProjectId(projectId)

	if errWhenGetAllProjects != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errWhenGetAllProjects)
	}

	data := map[string]interface{}{
		"tasks": tasks,
	}

	ctx.JSON(http.StatusOK, data)
}

func (taskService *TaskController) GetTaskById(ctx *gin.Context) {
	taskId, ok := ctx.Params.Get("taskId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "taskId not provided or incorrect...❌"})
		return
	}

	task, errWhenGetTaskById := taskService.database.GetTaskById(taskId)

	if errWhenGetTaskById != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenGetTaskById})
		return
	}

	if len(task.TaskId) > 0 {
		data := map[string]interface{}{
			"task": task,
		}

		ctx.JSON(http.StatusOK, data)
	}

	ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Task not found...❌"})
}

func (taskService *TaskController) CreateTask(ctx *gin.Context) {
	var task *models.Task
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect"})
		return
	}

	err := ctx.BindJSON(&task)
	task.ProjectId = projectId

	if err != nil {
		log.Println("Error when create a task ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	task, errWhenCreatetask := taskService.database.CreateTask(task)

	if errWhenCreatetask != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errWhenCreatetask})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (taskService *TaskController) UpdateTask(ctx *gin.Context) {
	taskId, ok := ctx.Params.Get("taskId")
	project := make(map[string]interface{})

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "taskId not provided or incorrect...❌"})
		return
	}

	errWhenBindingTaskJson := ctx.BindJSON(&project)

	if errWhenBindingTaskJson != nil {
		ctx.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"error": errWhenBindingTaskJson.Error()})
	}

	taskUpdated, errWhenUpdateTask := taskService.database.UpdateTask(taskId, project)

	if errWhenUpdateTask != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenUpdateTask})
		return
	}

	ctx.JSON(http.StatusOK, taskUpdated)
}

func (taskService *TaskController) DeleteTask(ctx *gin.Context) {
	taskId, ok := ctx.Params.Get("taskId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "taskId not provided or incorrect...❌"})
		return
	}

	project, errWhenDeleteProject := taskService.database.DeleteTask(taskId)

	if errWhenDeleteProject != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenDeleteProject.Error()})
		return
	}

	ctx.JSON(http.StatusOK, project)
}
