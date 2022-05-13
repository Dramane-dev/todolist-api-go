package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

func HealthCheckController(ctx *gin.Context) {
	response := HealthCheckResponse{Message: "OK"}

	ctx.JSON(http.StatusOK, response)
}
