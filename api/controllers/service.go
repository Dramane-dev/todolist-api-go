package controllers

import (
	"github.com/Dramane-dev/todolist-api/api/config"
	"github.com/Dramane-dev/todolist-api/api/middlewares"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	database config.UserService
}

type ProjectController struct {
	database config.ProjectService
}

// router.Use(middlewares.VerifyToken())

var jwtMiddleware = middlewares.VerifyToken()

func NewUserDatabaseInstance(router *gin.Engine, database config.UserService) error {
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

func NewProjectDatabaseInstance(router *gin.Engine, database config.ProjectService) error {
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
