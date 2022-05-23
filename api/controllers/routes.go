package controllers

import (
	"github.com/Dramane-dev/todolist-api/api/middlewares"
	"github.com/Dramane-dev/todolist-api/api/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	database service.UserService
}

type ProjectController struct {
	database service.ProjectService
}

type TaskController struct {
	database service.TaskService
}

var jwtMiddleware = middlewares.VerifyToken()

func NewUserDatabaseInstance(router *gin.Engine, database service.UserService) error {
	userService := &UserController{
		database: database,
	}

	router.POST("/api/user/signup", userService.Signup)
	router.POST("/api/user/signin", userService.Signin)

	router.GET("/api/health-check", jwtMiddleware, HealthCheckController)
	router.GET("/api/users", jwtMiddleware, userService.GetAllUsers)
	router.GET("/api/user/:userId", jwtMiddleware, userService.GetUserById)
	router.PATCH("/api/user/:userId", jwtMiddleware, userService.UpdateUser)
	router.DELETE("/api/user/:userId", jwtMiddleware, userService.DeleteUser)

	return nil
}

func NewProjectDatabaseInstance(router *gin.Engine, database service.ProjectService) error {
	projectService := &ProjectController{
		database: database,
	}

	router.GET("/api/projects", jwtMiddleware, projectService.GetAllProjects)
	router.GET("/api/projects/:userId", jwtMiddleware, projectService.GetAllProjectsByUserId)
	router.GET("/api/project/:projectId", jwtMiddleware, projectService.GetProjectById)
	router.POST("/api/project/:userId", jwtMiddleware, projectService.CreateProject)
	router.PATCH("/api/project/:projectId", jwtMiddleware, projectService.UpdateProject)
	router.DELETE("/api/project/:projectId", jwtMiddleware, projectService.DeleteProject)

	return nil
}

func NewTaskDatabaseInstance(router *gin.Engine, database service.TaskService) error {
	taskService := &TaskController{
		database: database,
	}

	router.GET("/api/tasks", jwtMiddleware, taskService.GetAllTasks)
	router.GET("/api/tasks/:projectId", jwtMiddleware, taskService.GetAllTasksByProjectId)
	router.GET("/api/task/:taskId", jwtMiddleware, taskService.GetTaskById)
	router.POST("/api/task/:projectId", jwtMiddleware, taskService.CreateTask)
	router.PATCH("/api/task/:taskId", jwtMiddleware, taskService.UpdateTask)
	router.DELETE("/api/task/:taskId", jwtMiddleware, taskService.DeleteTask)

	return nil
}
