package controllers

import (
	"log"
	"net/http"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/gin-gonic/gin"
)

func (projectService *ProjectController) GetAllProjectsByUserId(ctx *gin.Context) {
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect"})
		return
	}

	projects, errWhenGetAllProjects := projectService.database.GetAllProjectsByUserId(userId)

	if errWhenGetAllProjects != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, errWhenGetAllProjects)
	}

	data := map[string]interface{}{
		"projects": projects,
	}

	ctx.JSON(http.StatusOK, data)
}

func (projectService *ProjectController) GetAllProjects(ctx *gin.Context) {
	projects, errWhenGetAllProjectByUserId := projectService.database.GetAllProjects()

	if errWhenGetAllProjectByUserId != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenGetAllProjectByUserId})
		return
	}

	data := map[string]interface{}{
		"projects": projects,
	}

	ctx.JSON(http.StatusOK, data)
}

func (projectService *ProjectController) GetProjectById(ctx *gin.Context) {
	projectID, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect"})
		return
	}

	project, errWhenGetProjectById := projectService.database.GetProjectById(projectID)

	if errWhenGetProjectById != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenGetProjectById})
		return
	}

	data := map[string]interface{}{
		"project": project,
	}

	ctx.JSON(http.StatusOK, data)
}

func (projectService *ProjectController) CreateProject(ctx *gin.Context) {
	var project *models.Project
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect"})
		return
	}

	err := ctx.BindJSON(&project)
	project.UserId = userId

	if err != nil {
		log.Println("Error when create a project ", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	project, errWhenCreateProject := projectService.database.CreateProject(project)

	if errWhenCreateProject != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errWhenCreateProject})
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func (projectService *ProjectController) UpdateProject(ctx *gin.Context) {
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect"})
		return
	}

	project, errWhenUpdateProject := projectService.database.UpdateProject(projectId)

	if errWhenUpdateProject != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenUpdateProject})
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func (projectService *ProjectController) DeleteProject(ctx *gin.Context) {
	projectId, ok := ctx.Params.Get("projectId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "projectId not provided or incorrect"})
		return
	}

	project, errWhenDeleteProject := projectService.database.DeleteProject(projectId)

	if errWhenDeleteProject != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": errWhenDeleteProject})
		return
	}

	ctx.JSON(http.StatusOK, project)
}
